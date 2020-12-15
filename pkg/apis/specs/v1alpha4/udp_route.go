package v1alpha4

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UDPRoute is used to describe UDP inbound connections
type UDPRoute struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec UDPRouteSpec `json:"spec,omitempty"`
}

// UDPRouteSpec is the specification of a UDPRoute
type UDPRouteSpec struct {
	// Route match for inbound traffic
	Matches UDPMatch `json:"matches,omitempty"`
}

// UDPMatch defines an individual route for UDP traffic
type UDPMatch struct {
	// Name is the name of the match for referencing in a TrafficTarget
	Name string `json:"name,omitempty"`

	// Ports to allow inbound traffic on
	Ports []int `json:"ports,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UDPRouteList satisfy K8s code gen requirements
type UDPRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []UDPRoute `json:"items"`
}
