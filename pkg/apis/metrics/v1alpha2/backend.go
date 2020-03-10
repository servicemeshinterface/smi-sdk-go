package v1alpha2

// Backend describes which branch of a traffic split these metrics
// represent.
type Backend struct {
	apex   string `json:"apex"`
	name   string `json:"name"`
	weight int    `json:"weight"`
}
