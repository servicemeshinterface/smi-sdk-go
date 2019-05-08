package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IdentityBinding grants access for a specific identity to the rules in a TrafficTarget.
// It holds a list of subjects (service accounts for now) and a reference to the
// traffic target defining what has been granted.
type IdentityBinding struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Subjects are the pod or group of pods to allow ingress traffic
	Subjects []IdentityBindingSubjects `json:"subjects,omitempty" protobuf:"bytes,2,opt,name=subjects"`

	// TargetRef is the traffic target to which this binding applies
	TargetRef TrafficTargetRef `json:"targetRef,omitempty" protobuf:"bytes,3,opt,name=targetRef"`
}

// IdentityBindingSubjects are Kubernetes objects which should be allowed access to the TrafficTarget
type IdentityBindingSubjects struct {
	// Kind is the type of Subject to allow ingress (ServiceAccount | Group)
	Kind string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
	// Name of the Subject, i.e. ServiceAccountName
	Name string `json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`
	// Namespace where the Subject is deployed
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,3,opt,name=namespace"`
}

// TrafficTargetRef is a reference to a TrafficTarget object
type TrafficTargetRef struct {
	// Kind is the type of TrafficTarget reference (TrafficTarget)
	Kind string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
	// Name of the TrafficTarget
	Name string `json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`
	// Namespace where the TrafficTarget is deployed
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,3,opt,name=namespace"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IdentityBindingList satisfy K8s code gen requirements
type IdentityBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []IdentityBinding `json:"items"`
}
