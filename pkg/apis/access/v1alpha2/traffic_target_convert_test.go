package v1alpha2

import (
	"testing"

	"github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha3"
	assert "github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestConvertToConvertsFromAlpha3ToAlpha2(t *testing.T) {
	v2Test := &TrafficTarget{}

	err := v2Test.ConvertFrom(v3Access)
	assert.NoError(t, err)

	assert.Equal(t, v3Access.ObjectMeta, v2Test.ObjectMeta)
	assert.Equal(t, v3Access.TypeMeta.Kind, v2Test.TypeMeta.Kind)
	assert.Equal(t, "v1alpha2", v2Test.TypeMeta.APIVersion)

	// test destination
	assert.Equal(t, v3Access.Spec.Destination.Kind, v2Test.Spec.Destination.Kind)
	assert.Equal(t, v3Access.Spec.Destination.Name, v2Test.Spec.Destination.Name)
	assert.Equal(t, v3Access.Spec.Destination.Namespace, v2Test.Spec.Destination.Namespace)

	// test sources
	assert.Len(t, v2Test.Spec.Sources, len(v3Access.Spec.Sources))
	for i, s := range v3Access.Spec.Sources {
		assert.Equal(t, s.Kind, v2Test.Spec.Sources[i].Kind)
		assert.Equal(t, s.Name, v2Test.Spec.Sources[i].Name)
		assert.Equal(t, s.Namespace, v2Test.Spec.Sources[i].Namespace)
	}

	// test rules
	assert.Len(t, v2Test.Spec.Rules, len(v3Access.Spec.Rules))
	for i, s := range v3Access.Spec.Rules {
		assert.Equal(t, s.Kind, v2Test.Spec.Rules[i].Kind)
		assert.Equal(t, s.Name, v2Test.Spec.Rules[i].Name)

		for n, m := range s.Matches {
			assert.Equal(t, m, v2Test.Spec.Rules[i].Matches[n])
		}
	}
}

func TestConvertToConvertsFromAlpha1ToAlpha2(t *testing.T) {
	v3Test := &v1alpha3.TrafficTarget{}

	err := v2Access.ConvertTo(v3Test)
	assert.NoError(t, err)

	assert.Equal(t, v2Access.ObjectMeta, v3Test.ObjectMeta)
	assert.Equal(t, v2Access.TypeMeta.Kind, v3Test.TypeMeta.Kind)
	assert.Equal(t, "v1alpha3", v3Test.TypeMeta.APIVersion)

	// test destination
	assert.Equal(t, v2Access.Spec.Destination.Kind, v3Test.Spec.Destination.Kind)
	assert.Equal(t, v2Access.Spec.Destination.Name, v3Test.Spec.Destination.Name)
	assert.Equal(t, v2Access.Spec.Destination.Namespace, v3Test.Spec.Destination.Namespace)

	// test sources
	assert.Len(t, v3Test.Spec.Sources, len(v2Access.Spec.Sources))
	for i, s := range v2Access.Spec.Sources {
		assert.Equal(t, s.Kind, v3Test.Spec.Sources[i].Kind)
		assert.Equal(t, s.Name, v3Test.Spec.Sources[i].Name)
		assert.Equal(t, s.Namespace, v3Test.Spec.Sources[i].Namespace)
	}

	// test rules
	assert.Len(t, v3Test.Spec.Rules, len(v2Access.Spec.Rules))
	for i, s := range v2Access.Spec.Rules {
		assert.Equal(t, s.Kind, v3Test.Spec.Rules[i].Kind)
		assert.Equal(t, s.Name, v3Test.Spec.Rules[i].Name)

		for n, m := range s.Matches {
			assert.Equal(t, m, v3Test.Spec.Rules[i].Matches[n])
		}
	}
}

var v3Access = &v1alpha3.TrafficTarget{
	TypeMeta: v1.TypeMeta{
		Kind:       "TrafficTarget",
		APIVersion: "v1alpha3",
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v3Access",
		Namespace: "default",
	},
	Spec: v1alpha3.TrafficTargetSpec{
		Destination: v1alpha3.IdentityBindingSubject{
			Kind:      "ServiceAccount",
			Name:      "myservice",
			Namespace: "default",
		},
		Sources: []v1alpha3.IdentityBindingSubject{
			v1alpha3.IdentityBindingSubject{
				Kind:      "ServiceAccount",
				Name:      "mydestination1",
				Namespace: "default",
			},
			v1alpha3.IdentityBindingSubject{
				Kind:      "ServiceAccount",
				Name:      "mydestination2",
				Namespace: "default",
			},
		},
		Rules: []v1alpha3.TrafficTargetRule{
			v1alpha3.TrafficTargetRule{
				Kind:    "HTTPRouteGroup",
				Name:    "myname1",
				Matches: []string{"abc", "123"},
			},
			v1alpha3.TrafficTargetRule{
				Kind:    "HTTPRouteGroup",
				Name:    "myname2",
				Matches: []string{"123", "abc"},
			},
		},
	},
}

var v2Access = &TrafficTarget{
	TypeMeta: v1.TypeMeta{
		Kind:       "TrafficTarget",
		APIVersion: "v1alpha2",
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v1Access",
		Namespace: "default",
	},
	Spec: TrafficTargetSpec{
		Destination: IdentityBindingSubject{
			Kind:      "ServiceAccount",
			Name:      "myservice",
			Namespace: "default",
		},
		Sources: []IdentityBindingSubject{
			IdentityBindingSubject{
				Kind:      "ServiceAccount",
				Name:      "mydestination1",
				Namespace: "default",
			},
			IdentityBindingSubject{
				Kind:      "ServiceAccount",
				Name:      "mydestination2",
				Namespace: "default",
			},
		},
		Rules: []TrafficTargetRule{
			TrafficTargetRule{
				Kind:    "HTTPRouteGroup",
				Name:    "myname1",
				Matches: []string{"abc", "123"},
			},
			TrafficTargetRule{
				Kind:    "HTTPRouteGroup",
				Name:    "myname2",
				Matches: []string{"123", "abc"},
			},
		},
	},
}
