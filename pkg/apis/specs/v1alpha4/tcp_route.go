package v1alpha4

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TCPRoute is used to describe TCP inbound connections
type TCPRoute struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TCPRouteSpec `json:"spec,omitempty"`
}

// TCPRouteSpec is the specification of a TCPRoute
type TCPRouteSpec struct {
	// Route match for inbound traffic
	Matches TCPMatch `json:"matches,omitempty"`
}

// TCPMatch defines an individual route for TCP traffic
type TCPMatch struct {
	// Name is the name of the match for referencing in a TrafficTarget
	Name string `json:"name,omitempty"`

	// Ports to allow inbound traffic on
	Ports []int `json:"ports,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TCPRouteList satisfy K8s code gen requirements
type TCPRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []TCPRoute `json:"items"`
}
