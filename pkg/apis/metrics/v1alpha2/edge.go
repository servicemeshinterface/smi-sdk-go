package v1alpha2

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
)

// Direction defines whether an edge is sending traffic to a resource or
// receiving traffic from a resource.
type Direction string

// Side defines on which side of the edge the metrics are measured.
type Side string

const (
	// To is used when an edge is sending traffic to a resource
	To Direction = "to"
	// From is used when an edge is receiving traffic from a resource
	From Direction = "from"
	// Server indicates that the metrics are measured on the server side
	Server Side = "server"
	// Client indicated that the metrics are measured on the client side
	Client Side = "client"
)

// Edge describes the other resource that metrics are associated with
type Edge struct {
	Direction Direction           `json:"direction"`
	Side      Side                `json:"side"`
	Resource  *v1.ObjectReference `json:"resource"`
	Backend   *Backend            `json:"backend"`
}

// String returns a formatted string representation of this struct
func (e *Edge) String() string {
	return fmt.Sprintf("%#v", e)
}
