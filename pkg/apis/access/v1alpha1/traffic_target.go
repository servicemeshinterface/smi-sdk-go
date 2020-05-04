package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TrafficTarget associates a set of traffic definitions (rules) with a service identity which is allocated to a group of pods.
// Access is controlled via referenced TrafficSpecs and by a list of source service identities.
// * If a pod which holds the referenced service identity makes a call to the destination on one of the defined routes then access
//   will be allowed
// * Any pod which attempts to connect and is not in the defined list of sources will be denied
// * Any pod which is in the defined list, but attempts to connect on a route which is not in the list of the
//   TrafficSpecs will be denied
type TrafficTarget struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Selector is the pod or group of pods to allow ingress traffic
	Destination IdentityBindingSubject `json:"destination"`

	// Sources are the pod or group of pods to allow ingress traffic
	Sources []IdentityBindingSubject `json:"sources"`

	// Rules are the traffic rules to allow (HTTPRoutes | TCPRoute),
	Specs []TrafficTargetSpec `json:"specs"`
}

// TrafficTargetSpec is the TrafficSpec to allow for a TrafficTarget
type TrafficTargetSpec struct {
	// Kind is the kind of TrafficSpec to allow
	Kind string `json:"kind"`
	// Name of the TrafficSpec to use
	Name string `json:"name"`
	// Matches is a list of TrafficSpec routes to allow traffic for
	Matches []string `json:"matches,omitempty"`
}

// IdentityBindingSubject is a Kubernetes objects which should be allowed access to the TrafficTarget
type IdentityBindingSubject struct {
	// Kind is the type of Subject to allow ingress (ServiceAccount | Group)
	Kind string `json:"kind"`
	// Name of the Subject, i.e. ServiceAccountName
	Name string `json:"name"`
	// Namespace where the Subject is deployed
	Namespace string `json:"namespace,omitempty"`
	// Port defines a TCP port to apply the TrafficTarget to
	Port int `json:"port,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TrafficTargetList satisfy K8s code gen requirements
type TrafficTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []TrafficTarget `json:"items"`
}
