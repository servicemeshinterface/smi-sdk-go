package metrics

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TrafficMetricsList provides a list of resources associated with a specific reference
type TrafficMetricsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Resource        *v1.ObjectReference `json:"resource"`

	Items []*TrafficMetrics `json:"items"`
}
