package upgradelog

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/rancher/apiserver/pkg/apierror"
	ctlbatchv1 "github.com/rancher/wrangler/pkg/generated/controllers/batch/v1"
	ctlcorev1 "github.com/rancher/wrangler/pkg/generated/controllers/core/v1"
	"github.com/rancher/wrangler/pkg/schemas/validation"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	ctlupgradelog "github.com/harvester/harvester/pkg/controller/master/upgradelog"
	ctlharvesterv1 "github.com/harvester/harvester/pkg/generated/controllers/harvesterhci.io/v1beta1"
	"github.com/harvester/harvester/pkg/util"
)

const (
	archiveSuffix = ".tar.gz"
)

type Handler struct {
	httpClient       *http.Client
	jobClient        ctlbatchv1.JobClient
	podCache         ctlcorev1.PodCache
	upgradeCache     ctlharvesterv1.UpgradeCache
	upgradeLogCache  ctlharvesterv1.UpgradeLogCache
	upgradeLogClient ctlharvesterv1.UpgradeLogClient
}

func (h Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if err := h.doAction(rw, req); err != nil {
		status := http.StatusInternalServerError
		if e, ok := err.(*apierror.APIError); ok {
			status = e.Code.Status
		}
		rw.WriteHeader(status)
		_, _ = rw.Write([]byte(err.Error()))
		return
	}
}

func (h Handler) doAction(rw http.ResponseWriter, req *http.Request) error {
	vars := mux.Vars(req)

	if req.Method == http.MethodGet {
		return h.doGet(vars["link"], rw, req)
	} else if req.Method == http.MethodPost {
		return h.doPost(vars["action"], rw, req)
	}

	return apierror.NewAPIError(validation.InvalidAction, fmt.Sprintf("Unsupported method %s", req.Method))
}

func (h Handler) doGet(link string, rw http.ResponseWriter, req *http.Request) error {
	switch link {
	case downloadArchiveLink:
		return h.downloadArchive(rw, req)
	default:
		return apierror.NewAPIError(validation.InvalidAction, fmt.Sprintf("Unsupported GET action %s", link))
	}
}

func (h Handler) doPost(action string, rw http.ResponseWriter, req *http.Request) error {
	switch action {
	case generateArchiveAction:
		return h.generateArchive(rw, req)
	default:
		return apierror.NewAPIError(validation.InvalidAction, fmt.Sprintf("Unsupported POST action %s", action))
	}
}

func (h Handler) downloadArchive(rw http.ResponseWriter, req *http.Request) error {
	vars := mux.Vars(req)
	upgradeLogName := vars["name"]
	upgradeLogNamespace := vars["namespace"]
	archiveName := req.URL.Query().Get("archiveName")

	logrus.Infof("Retrieve the archive (%s) for the UpgradeLog (%s/%s)", archiveName, upgradeLogNamespace, upgradeLogName)

	upgradeLog, err := h.upgradeLogCache.Get(upgradeLogNamespace, upgradeLogName)
	if err != nil {
		return fmt.Errorf("failed to get the upgradelog resource (%s/%s): %w", upgradeLogNamespace, upgradeLogName, err)
	}

	isDownloadReady := false
	for _, condition := range upgradeLog.Status.Conditions {
		if condition.Type == harvesterv1.DownloadReady && condition.Status == corev1.ConditionTrue {
			isDownloadReady = true
			break
		}
	}

	if !isDownloadReady {
		return fmt.Errorf("the archive (%s) of upgrade resource (%s/%s) is not ready yet", archiveName, upgradeLogNamespace, upgradeLogName)
	}

	downloaderPodIP, err := ctlupgradelog.GetDownloaderPodIP(h.podCache, upgradeLog)
	if err != nil {
		return fmt.Errorf("failed to get the downloader pod IP with upgradelog resource (%s/%s): %w", upgradeLogNamespace, upgradeLogName, err)
	}

	archiveFileName := fmt.Sprintf("%s%s", archiveName, archiveSuffix)
	downloadURL := fmt.Sprintf("http://%s/%s", downloaderPodIP, archiveFileName)
	downloadReq, err := http.NewRequestWithContext(req.Context(), http.MethodGet, downloadURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create the download request for the archive (%s): %w", archiveName, err)
	}

	downloadResp, err := h.httpClient.Do(downloadReq)
	if err != nil {
		return fmt.Errorf("failed to send the download request for the archive (%s): %w", archiveName, err)
	}
	defer downloadResp.Body.Close()

	if downloadResp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed with unexpected http status code %d", downloadResp.StatusCode)
	}

	// TODO: set Content-Length with archive size
	rw.Header().Set("Content-Disposition", "attachment; filename="+archiveFileName)
	contentType := downloadResp.Header.Get("Content-Type")
	if contentType != "" {
		rw.Header().Set("Content-Type", contentType)
	}

	if _, err := io.Copy(rw, downloadResp.Body); err != nil {
		return fmt.Errorf("failed to copy the downloaded content to the target (%s), err: %w", archiveFileName, err)
	}

	return nil
}

func (h Handler) generateArchive(rw http.ResponseWriter, req *http.Request) error {
	vars := mux.Vars(req)
	upgradeLogName := vars["name"]
	upgradeLogNamespace := vars["namespace"]

	logrus.Infof("Generate an archive for the UpgradeLog (%s/%s)", upgradeLogNamespace, upgradeLogName)

	upgradeLog, err := h.upgradeLogCache.Get(upgradeLogNamespace, upgradeLogName)
	if err != nil {
		return fmt.Errorf("failed to get the upgradelog resource (%s/%s): %w", upgradeLogNamespace, upgradeLogName, err)
	}

	isUpgradeLogReady := false
	for _, condition := range upgradeLog.Status.Conditions {
		if condition.Type == harvesterv1.UpgradeLogReady && condition.Status == corev1.ConditionTrue {
			isUpgradeLogReady = true
			break
		}
	}

	if !isUpgradeLogReady {
		return fmt.Errorf("the logging infrastructure for the upgradelog resource (%s/%s) is not ready yet", upgradeLogNamespace, upgradeLogName)
	}

	// Get image version for log packager
	upgrade, err := h.upgradeCache.Get(upgradeLogNamespace, upgradeLog.Spec.Upgrade)
	if err != nil {
		return fmt.Errorf("failed to get the upgrade resource (%s/%s): %w", upgradeLogNamespace, upgradeLog.Spec.Upgrade, err)
	}
	imageVersion := upgrade.Status.PreviousVersion

	ts := time.Now().UTC()
	generatedTime := strings.Replace(ts.Format(time.RFC3339), ":", "-", -1)
	archiveName := fmt.Sprintf("%s-archive-%s", upgradeLog.Name, generatedTime)
	// TODO: update with the real size later
	archiveSize := int64(0)

	var component string
	if harvesterv1.UpgradeEnded.IsTrue(upgradeLog) {
		component = ctlupgradelog.DownloaderComponent
	} else {
		component = ctlupgradelog.AggregatorComponent
	}

	if _, err := h.jobClient.Create(ctlupgradelog.PrepareLogPackager(upgradeLog, imageVersion, archiveName, component)); err != nil {
		return fmt.Errorf("failed to create log packager job for the upgradelog resource (%s/%s): %w", upgradeLogNamespace, upgradeLogName, err)
	}
	toUpdate := upgradeLog.DeepCopy()
	ctlupgradelog.SetUpgradeLogArchive(toUpdate, archiveName, archiveSize, generatedTime, false)
	if _, err := h.upgradeLogClient.Update(toUpdate); err != nil {
		return fmt.Errorf("failed to update the upgradelog resource (%s/%s): %w", upgradeLogNamespace, upgradeLogName, err)
	}
	util.ResponseOKWithBody(rw, archiveName)

	return nil
}
