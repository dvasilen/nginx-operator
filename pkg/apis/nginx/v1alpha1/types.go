package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Nginx struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              NginxSpec   `json:"spec"`
	Status            NginxStatus `json:"status,omitempty"`
}

type NginxSpec struct {
	// Number of desired pods. This is a pointer to distinguish between explicit
	// zero and not specified. Defaults to the default deployment replicas value.
	// +optional
	Replicas *int32 `json:"replicas"`
	// Docker image name. Defaults to "nginx:latest".
	// +optional
	Image string `json:"image"`
	// Reference to the nginx config object.
	Config *ConfigRef `json:"configRef"`
}

type NginxStatus struct {
	Pods []NginxPod `json:"pods"`
}

type NginxPod struct {
	Name  string `json:"name"`
	PodIP string `json:"podIP"`
}

// ConfigRef is a reference to a config object.
type ConfigRef struct {
	// Name of the config object.
	Name string `json:"name"`
	// Kind of the config object. Defaults to ConfigKindConfigMap.
	Kind ConfigKind `json:"kind"`
	// Optional value used by some ConfigKinds.
	Value string `json:"value"`
}

type ConfigKind string

const (
	// ConfigKindConfigMap is a Kind of configuration that points to a configmap
	ConfigKindConfigMap = ConfigKind("ConfigMap")
	// ConfigKindInline is a kinda of configuration that is setup as a annotation on the Pod
	// and is inject as a file on the container using the Downward API.
	ConfigKindInline = ConfigKind("Inline")
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NginxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Nginx `json:"items"`
}
