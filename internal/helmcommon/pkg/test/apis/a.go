package apis

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type A struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ASpec `json:"spec"`
}

type ASpec struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}
