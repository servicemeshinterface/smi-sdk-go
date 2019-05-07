package v1beta1

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPRoutes is used to describe HTTP/1 and HTTP/2 traffic.
// It enumerates the routes that can be served by an application.
type HTTPRoutes struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Routes for inbound traffic
	Routes []HTTPRoute `json:"routes,omitempty" protobuf::bytes,2,opt,name=routes"`
}

// HTTPRoute defines an individual route for HTTP traffic
type HTTPRoute struct {
	// Methods for inbound traffic as defined in RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-4
	Methods []HTTPRouteMethod

	// PathRegex is a regular expression defining the route
	PathRegex string
}

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
	HTTPRouteMethodTrace HTTPRouteMethod = "PATCH"
)
