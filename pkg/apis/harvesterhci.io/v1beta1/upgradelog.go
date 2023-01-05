package v1beta1

import (
	"github.com/rancher/wrangler/pkg/condition"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	OperatorDeployed condition.Cond = "OperatorReady"
	InfraScaffolded  condition.Cond = "InfraReady"
	UpgradeLogReady  condition.Cond = "Ready"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName=lg;lgs,scope=Namespaced
// +kubebuilder:printcolumn:name="UPGRADE",type="string",JSONPath=`.spec.upgrade`

type UpgradeLog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UpgradeLogSpec   `json:"spec"`
	Status UpgradeLogStatus `json:"status,omitempty"`
}

type UpgradeLogSpec struct {
	// +kubebuilder:validation:Required
	Upgrade string `json:"upgrade,omitempty"`
}

type UpgradeLogStatus struct {
	// +optional
	Conditions []Condition `json:"conditions,omitempty"`
}
