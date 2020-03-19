package v1alpha2

// Backend describes which branch of a traffic split these metrics
// represent.
type Backend struct {
	Apex   string `json:"apex"`
	Name   string `json:"name"`
	Weight int    `json:"weight"`
}
