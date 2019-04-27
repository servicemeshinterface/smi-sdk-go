package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TrafficRole does something
type TrafficRole struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Destination TrafficServiceReference `json:"destination,omitempty" protobuf:"bytes,2,opt,name=destination"`
	Rules       []TrafficRoleRules
}

// TrafficServiceReference is a reference to a service
type TrafficServiceReference struct {
	// Kind defines the type of service reference
	Kind ServiceKind
	// ValueFrom optional value to obtain the reference from
	ValueFrom TrafficServiceValueFrom
	Value     string
}

// TrafficServiceValueFrom allows a destination value to be obtained from another
// resource.
type TrafficServiceValueFrom struct {
	// Kind of Kubernetes object to retrieve the value from
	Kind ValueFromKind `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
	// Name of the object from which to obtain the value
	Name string `json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`
	// Namespace in which the object exists
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,3,opt,name=namespace"`
}

// ServiceKind defines a type of object for retrieving something
type ServiceKind string

// ServiceIdentity defines a service mesh identity
const ServiceIdentity ServiceKind = "ServiceIdentity"

// ValueFromKind defines a type of object from where values can be obtained
type ValueFromKind string

// ServiceAccount is a ValueFromKind which references Kubernetes service accounts
const ServiceAccount ValueFromKind = "ServiceAccount"

// TrafficRoleRules define the rules that apply to a TrafficRole
type TrafficRoleRules struct {
	// Access access level for the rule allow/deny
	Access TrafficRoleRuleAccess
	// Methods are HTTP verbs for access
	Methods []TrafficRoleRuleMethod
	// Paths are HTTP paths to allow access to
	// TODO what about grpc rules define L7 HTTP only
	Paths []string
}

// TrafficRoleRuleAccess defines the access level for a rule
type TrafficRoleRuleAccess string

// TrafficRoleRuleAccessAllow allows access to the service with this rule
const TrafficRoleRuleAccessAllow TrafficRoleRuleAccess = "allow"

// TrafficRoleRuleAccessDeny denies access to the service with this rule
const TrafficRoleRuleAccessDeny TrafficRoleRuleAccess = "deny"

// TrafficRoleRuleMethod are methods allowed by the rule
type TrafficRoleRuleMethod string

const (
	// TrafficRoleRuleMethodAll controls access to all HTTP methods
	TrafficRoleRuleMethodAll TrafficRoleRuleMethod = "*"
	// TrafficRoleRuleMethodGet controls access to the HTTP GET method
	TrafficRoleRuleMethodGet TrafficRoleRuleMethod = "GET"
	// TrafficRoleRuleMethodPut controls access to the HTTP PUT method
	TrafficRoleRuleMethodPut TrafficRoleRuleMethod = "PUT"
	// TrafficRoleRuleMethodPost controls access to the HTTP POST method
	TrafficRoleRuleMethodPost TrafficRoleRuleMethod = "POST"
	// TrafficRoleRuleMethodHead controls access to the HTTP HEAD method
	TrafficRoleRuleMethodHead TrafficRoleRuleMethod = "HEAD"
	// TrafficRoleRuleMethodDelete controls access to the HTTP DELETE method
	TrafficRoleRuleMethodDelete TrafficRoleRuleMethod = "DELETE"
)

// TrafficRoleBinding binding defines services to which a TrafficRole applies
type TrafficRoleBinding struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Subjects are services which the Role applies to
	Subjects []TrafficServiceReference
	// RoleRef is a reference to the Role to bind the subjects to
	RoleRef TrafficRoleRef
}

// TrafficRoleRef defines a role reference
type TrafficRoleRef struct {
	// Kind of object referenced
	Kind TrafficRoleRefKind
	// Name of the object referenced
	Name string
}

// TrafficRoleRefKind defines a base type for role references
type TrafficRoleRefKind string

// TrafficRoleRefRole is a reference to a TrafficRole object
const TrafficRoleRefRole TrafficRoleRefKind = "TrafficRole"
