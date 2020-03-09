package v1alpha2

// TrafficSplit describes which branch of a traffic split these metrics
// represent.
type TrafficSplit struct {
	apex   string `json:"apex"`
	leaf   string `json:"leaf"`
	weight int    `json:"weight"`
}
