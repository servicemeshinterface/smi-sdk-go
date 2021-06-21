package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPRouteGroup is used to describe HTTP/1 and HTTP/2 traffic.
// It enumerates the routes that can be served by an application.
type HTTPRouteGroup struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec HTTPRouteGroupSpec `json:"spec"`
}

// HTTPRouteGroupSpec is the specification for a HTTPRouteGroup
type HTTPRouteGroupSpec struct {
	// Routes for inbound traffic
	Matches []HTTPMatch `json:"matches,omitempty"`
}

// HTTPMatch defines an individual route for HTTP traffic
type HTTPMatch struct {
	// Name is the name of the match for referencing in a TrafficTarget
	Name string `json:"name,omitempty"`

	// Methods for inbound traffic as defined in RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-4
	Methods []string `json:"methods,omitempty"`

	// PathRegex is a regular expression defining the route
	PathRegex string `json:"pathRegex,omitempty"`

	// Headers is a list of headers used to match HTTP traffic
	Headers HTTPHeaders `json:"headers,omitempty"`
}

// HTTPHeaders is a map of key/value pairs which match HTTP header name and value
type HTTPHeaders map[string]string

// HTTPRouteMethod are methods allowed by the route
type HTTPRouteMethod string

const (
	// HTTPRouteMethodAll is a wildcard for all HTTP methods
	HTTPRouteMethodAll HTTPRouteMethod = "*"
	// HTTPRouteMethodGet HTTP GET method
	HTTPRouteMethodGet HTTPRouteMethod = "GET"
	// HTTPRouteMethodHead HTTP HEAD method
	HTTPRouteMethodHead HTTPRouteMethod = "HEAD"
	// HTTPRouteMethodPut HTTP PUT method
	HTTPRouteMethodPut HTTPRouteMethod = "PUT"
	// HTTPRouteMethodPost HTTP POST method
	HTTPRouteMethodPost HTTPRouteMethod = "POST"
	// HTTPRouteMethodDelete HTTP DELETE method
	HTTPRouteMethodDelete HTTPRouteMethod = "DELETE"
	// HTTPRouteMethodConnect HTTP CONNECT method
	HTTPRouteMethodConnect HTTPRouteMethod = "CONNECT"
	// HTTPRouteMethodOptions HTTP OPTIONS method
	HTTPRouteMethodOptions HTTPRouteMethod = "OPTIONS"
	// HTTPRouteMethodTrace HTTP TRACE method
	HTTPRouteMethodTrace HTTPRouteMethod = "TRACE"
	// HTTPRouteMethodPatch HTTP PATCH method
	HTTPRouteMethodPatch HTTPRouteMethod = "PATCH"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPRouteGroupList satisfy K8s code gen requirements
type HTTPRouteGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []HTTPRouteGroup `json:"items"`
}
