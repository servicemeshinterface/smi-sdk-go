package v1alpha1

import (
	"testing"

	"github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha3"
	assert "github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestConvertToConvertsFromAlpha3ToAlpha1(t *testing.T) {
	v1Test := &TrafficTarget{}

	err := v1Test.ConvertFrom(v3Access)
	assert.NoError(t, err)

	assert.Equal(t, v3Access.ObjectMeta, v1Test.ObjectMeta)
	assert.Equal(t, v3Access.TypeMeta.Kind, v1Test.TypeMeta.Kind)
	assert.Equal(t, "v1alpha1", v1Test.TypeMeta.APIVersion)

	// test detination
	assert.Equal(t, v3Access.Spec.Destination.Kind, v1Test.Destination.Kind)
	assert.Equal(t, v3Access.Spec.Destination.Name, v1Test.Destination.Name)
	assert.Equal(t, v3Access.Spec.Destination.Namespace, v1Test.Destination.Namespace)

	// test sources
	assert.Len(t, v1Test.Sources, len(v3Access.Spec.Sources))
	for i, s := range v3Access.Spec.Sources {
		assert.Equal(t, s.Kind, v1Test.Sources[i].Kind)
		assert.Equal(t, s.Name, v1Test.Sources[i].Name)
		assert.Equal(t, s.Namespace, v1Test.Sources[i].Namespace)
	}

	// test rules
	assert.Len(t, v1Test.Specs, len(v3Access.Spec.Rules))
	for i, s := range v3Access.Spec.Rules {
		assert.Equal(t, s.Kind, v1Test.Specs[i].Kind)
		assert.Equal(t, s.Name, v1Test.Specs[i].Name)

		for n, m := range s.Matches {
			assert.Equal(t, m, v1Test.Specs[i].Matches[n])
		}
	}
}

func TestConvertToConvertsFromAlpha1ToAlpha3(t *testing.T) {
	v3Test := &v1alpha3.TrafficTarget{}

	err := v1Access.ConvertTo(v3Test)
	assert.NoError(t, err)

	assert.Equal(t, v1Access.ObjectMeta, v3Test.ObjectMeta)
	assert.Equal(t, v1Access.TypeMeta.Kind, v3Test.TypeMeta.Kind)
	assert.Equal(t, "v1alpha3", v3Test.TypeMeta.APIVersion)

	// test destination
	assert.Equal(t, v1Access.Destination.Kind, v3Test.Spec.Destination.Kind)
	assert.Equal(t, v1Access.Destination.Name, v3Test.Spec.Destination.Name)
	assert.Equal(t, v1Access.Destination.Namespace, v3Test.Spec.Destination.Namespace)

	// test sources
	assert.Len(t, v3Test.Spec.Sources, len(v1Access.Sources))
	for i, s := range v1Access.Sources {
		assert.Equal(t, s.Kind, v3Test.Spec.Sources[i].Kind)
		assert.Equal(t, s.Name, v3Test.Spec.Sources[i].Name)
		assert.Equal(t, s.Namespace, v3Test.Spec.Sources[i].Namespace)
	}

	// test rules
	assert.Len(t, v3Test.Spec.Rules, len(v1Access.Specs))
	for i, s := range v1Access.Specs {
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

var v1Access = &TrafficTarget{
	TypeMeta: v1.TypeMeta{
		Kind:       "TrafficTarget",
		APIVersion: "v1alpha1",
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v1Access",
		Namespace: "default",
	},
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
	Specs: []TrafficTargetSpec{
		TrafficTargetSpec{
			Kind:    "HTTPRouteGroup",
			Name:    "myname1",
			Matches: []string{"abc", "123"},
		},
		TrafficTargetSpec{
			Kind:    "HTTPRouteGroup",
			Name:    "myname2",
			Matches: []string{"123", "abc"},
		},
	},
}
