package v1alpha2

// Backend describes which branch of a traffic split these metrics
// represent.
type Backend struct {
	apex   string `json:"apex"`
	leaf   string `json:"leaf"`
	weight int    `json:"weight"`
}
