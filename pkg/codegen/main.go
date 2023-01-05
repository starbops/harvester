package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	loggingv1 "github.com/banzaicloud/logging-operator/pkg/sdk/logging/api/v1beta1"
	cniv1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	storagev1beta1 "github.com/kubernetes-csi/external-snapshotter/v2/pkg/apis/volumesnapshot/v1beta1"
	longhornv1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	monitoring "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	upgradev1 "github.com/rancher/system-upgrade-controller/pkg/apis/upgrade.cattle.io/v1"
	controllergen "github.com/rancher/wrangler/pkg/controller-gen"
	"github.com/rancher/wrangler/pkg/controller-gen/args"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	kubevirtv1 "kubevirt.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha4"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
)

func main() {
	os.Unsetenv("GOPATH")
	controllergen.Run(args.Options{
		OutputPackage: "github.com/harvester/harvester/pkg/generated",
		Boilerplate:   "scripts/boilerplate.go.txt",
		Groups: map[string]args.Group{
			"harvesterhci.io": {
				Types: []interface{}{
					harvesterv1.KeyPair{},
					harvesterv1.Preference{},
					harvesterv1.Setting{},
					harvesterv1.Upgrade{},
					harvesterv1.UpgradeLog{},
					harvesterv1.Version{},
					harvesterv1.VirtualMachineBackup{},
					harvesterv1.VirtualMachineRestore{},
					harvesterv1.VirtualMachineImage{},
					harvesterv1.VirtualMachineTemplate{},
					harvesterv1.VirtualMachineTemplateVersion{},
					harvesterv1.SupportBundle{},
					harvesterv1.Addon{},
				},
				GenerateTypes:   true,
				GenerateClients: true,
			},
			loggingv1.GroupVersion.Group: {
				Types: []interface{}{
					loggingv1.Logging{},
					loggingv1.ClusterFlow{},
					loggingv1.ClusterOutput{},
				},
				GenerateTypes:   false,
				GenerateClients: true,
			},
			kubevirtv1.SchemeGroupVersion.Group: {
				Types: []interface{}{
					kubevirtv1.VirtualMachine{},
					kubevirtv1.VirtualMachineInstance{},
					kubevirtv1.VirtualMachineInstanceMigration{},
				},
				GenerateTypes:   false,
				GenerateClients: true,
			},
			cniv1.SchemeGroupVersion.Group: {
				Types: []interface{}{
					cniv1.NetworkAttachmentDefinition{},
				},
				GenerateTypes:   false,
				GenerateClients: true,
			},
			networkingv1.SchemeGroupVersion.Group: {
				Types: []interface{}{
					networkingv1.Ingress{},
				},
				GenerateTypes:   false,
				GenerateClients: true,
			},
			storagev1beta1.SchemeGroupVersion.Group: {
				Types: []interface{}{
					storagev1beta1.VolumeSnapshotClass{},
					storagev1beta1.VolumeSnapshot{},
					storagev1beta1.VolumeSnapshotContent{},
				},
				GenerateTypes:   false,
				GenerateClients: true,
			},
			longhornv1.SchemeGroupVersion.Group: {
				Types: []interface{}{
					longhornv1.BackingImage{},
					longhornv1.BackingImageDataSource{},
					longhornv1.Volume{},
					longhornv1.Setting{},
					longhornv1.Backup{},
					longhornv1.Replica{},
				},
				GenerateClients: true,
			},
			upgradev1.SchemeGroupVersion.Group: {
				Types: []interface{}{
					upgradev1.Plan{},
				},
				GenerateTypes:   false,
				GenerateClients: true,
			},
			capi.GroupVersion.Group: {
				Types: []interface{}{
					capi.Cluster{},
					capi.Machine{},
				},
				GenerateTypes:   false,
				GenerateClients: true,
			},
			corev1.GroupName: {
				Types: []interface{}{
					corev1.PersistentVolume{},
				},
				InformersPackage: "k8s.io/client-go/informers",
				ClientSetPackage: "k8s.io/client-go/kubernetes",
				ListersPackage:   "k8s.io/client-go/listers",
			},
			monitoring.GroupName: {
				Types: []interface{}{
					monitoringv1.Prometheus{},
					monitoringv1.Alertmanager{},
				},
				GenerateClients: true,
			},
		},
	})
	nadControllerInterfaceRefactor()
	capiWorkaround()
	loggingWorkaround()
}

// NB(GC), nadControllerInterfaceRefactor modify the generated resource name of NetworkAttachmentDefinition controller using a dash-separator,
// the original code is generated by https://github.com/rancher/wrangler/blob/e86bc912dfacbc81dc2d70171e4d103248162da6/pkg/controller-gen/generators/group_version_interface_go.go#L82-L97
// since the NAD crd uses a varietal plurals name(i.e network-attachment-definitions), and the default resource name generated by wrangler is
// `networkattachementdefinitions` that will raises crd not found exception of the NAD controller.
func nadControllerInterfaceRefactor() {
	absPath, _ := filepath.Abs("pkg/generated/controllers/k8s.cni.cncf.io/v1/interface.go")
	input, err := ioutil.ReadFile(absPath)
	if err != nil {
		logrus.Fatalf("failed to read the network-attachment-definition file: %v", err)
	}

	output := bytes.Replace(input, []byte("networkattachmentdefinitions"), []byte("network-attachment-definitions"), -1)

	if err = ioutil.WriteFile(absPath, output, 0644); err != nil {
		logrus.Fatalf("failed to update the network-attachment-definition file: %v", err)
	}
}

// capiWorkaround replaces the variable `SchemeGroupVersion` with `GroupVersion` in clusters.cluster.x-k8s.io client because
// `SchemeGroupVersion` is not declared in the vendor package but wrangler uses it.
// https://github.com/kubernetes-sigs/cluster-api/blob/56f9e9db7a9e9ca625ffe4bdc1e5e93a14d5e96c/api/v1alpha4/groupversion_info.go#L29
func capiWorkaround() {
	absPath, _ := filepath.Abs("pkg/generated/clientset/versioned/typed/cluster.x-k8s.io/v1alpha4/cluster.x-k8s.io_client.go")
	input, err := ioutil.ReadFile(absPath)
	if err != nil {
		logrus.Fatalf("failed to read the clusters.cluster.x-k8s.io client file: %v", err)
	}
	output := bytes.Replace(input, []byte("v1alpha4.SchemeGroupVersion"), []byte("v1alpha4.GroupVersion"), -1)

	if err = ioutil.WriteFile(absPath, output, 0644); err != nil {
		logrus.Fatalf("failed to update the clusters.cluster.x-k8s.io client file: %v", err)
	}
}

// loggingWorkaround replaces the variable `SchemeGroupVersion` with `GroupVersion` in logging.banzaicloud.io client because
// `SchemeGroupVersion` is not declared in the vendor package but wrangler uses it.
// https://github.com/banzaicloud/logging-operator/blob/e935c5d60604036a6f40cd4ab991420c6eaf096b/pkg/sdk/logging/api/v1beta1/groupversion_info.go#L27
func loggingWorkaround() {
	absPath, _ := filepath.Abs("pkg/generated/clientset/versioned/typed/logging.banzaicloud.io/v1beta1/logging.banzaicloud.io_client.go")
	input, err := ioutil.ReadFile(absPath)
	if err != nil {
		logrus.Fatalf("failed to read the logging.banzaicloud.io client file: %v", err)
	}
	output := bytes.Replace(input, []byte("v1beta1.SchemeGroupVersion"), []byte("v1beta1.GroupVersion"), -1)

	if err = ioutil.WriteFile(absPath, output, 0644); err != nil {
		logrus.Fatalf("failed to update the logging.banzaicloud.io client file: %v", err)
	}
}
