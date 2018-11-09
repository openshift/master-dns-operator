package v1alpha1

import (
	operatorsv1 "github.com/openshift/api/operator/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MasterDNSOperatorConfigSpec defines the desired state of MasterDNSOperatorConfig
type MasterDNSOperatorConfigSpec struct {
	operatorsv1.OperatorSpec `json:",inline"`

	// AutomaticUpdates if true indicates that the DNS entries should be automatically
	// updated based on IP address of Machine object
	AutomaticUpdates bool `json:"automaticUpdates"`
}

// MasterDNSOperatorConfigStatus defines the observed state of MasterDNSOperatorConfig
type MasterDNSOperatorConfigStatus struct {
	operatorsv1.OperatorStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MasterDNSOperatorConfig is the Schema for the masterdnsoperatorconfigs API
// +k8s:openapi-gen=true
type MasterDNSOperatorConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MasterDNSOperatorConfigSpec   `json:"spec,omitempty"`
	Status MasterDNSOperatorConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MasterDNSOperatorConfigList contains a list of MasterDNSOperatorConfig
type MasterDNSOperatorConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MasterDNSOperatorConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MasterDNSOperatorConfig{}, &MasterDNSOperatorConfigList{})
}
