package metrics

import (
	v1 "k8s.io/api/core/v1"
	apiresource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Interval defines the time interval from which metrics were collected
type Interval struct {
	Timestamp metav1.Time     `json:"timestamp"`
	Window    metav1.Duration `json:"window"`
}

// Direction defines whether an edge is sending traffic to a resource or
// receiving traffic from a resource.
type Direction string

const (
	// To is used when an edge is sending traffic to a resource
	To Direction = "to"
	// From is used when an edge is receiving traffic from a resource
	From Direction = "from"
)

// Edge describes the other resource that metrics are associated with
type Edge struct {
	Direction Direction           `json:"direction"`
	Resource  *v1.ObjectReference `json:"resource"`
}

// Unit is associated with metrics and defines what unit the metric is using
type Unit string

// MilliSeconds is a time unit
const MilliSeconds Unit = "ms"

// Metric describes a name and value for specific metrics
type Metric struct {
	Name string `json:"name"`

	Unit  Unit                  `json:"unit,omitempty"`
	Value *apiresource.Quantity `json:"value"`
}

// TrafficMetrics provide the metrics for a specific resource
type TrafficMetrics struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	*Interval `json:",inline"`

	Resource *v1.ObjectReference `json:"resource"`
	Edge     *Edge               `json:"edge"`
	Metrics  []*Metric           `json:"metrics"`
}
