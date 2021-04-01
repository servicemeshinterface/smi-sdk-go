package v1alpha3

import (
	"testing"

	"github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha4"
	assert "github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
)

func TestConvertToConvertsFromAlpha4ToAlpha3(t *testing.T) {
	v3Test := &TrafficSplit{}

	err := v3Test.ConvertFrom(v4Split)
	assert.NoError(t, err)

	assert.Equal(t, v4Split.Spec.Service, v3Test.Spec.Service)

	assert.Len(t, v4Split.Spec.Backends, len(v3Test.Spec.Backends))

	for i, b := range v4Split.Spec.Backends {
		assert.Equal(t, b.Service, v3Test.Spec.Backends[i].Service)
		assert.Equal(t, b.Weight, v3Test.Spec.Backends[i].Weight)
	}

	for i, m := range v4Split.Spec.Matches {
		assert.Equal(t, m.Name, v3Test.Spec.Matches[i].Name)
		assert.Equal(t, m.Kind, v3Test.Spec.Matches[i].Kind)
		assert.Equal(t, m.APIGroup, v3Test.Spec.Matches[i].APIGroup)
	}
}

func TestConvertToConvertsToAlpha4FromAlpha3(t *testing.T) {
	v4Test := &v1alpha4.TrafficSplit{}

	err := v3Split.ConvertTo(v4Test)
	assert.NoError(t, err)

	assert.Equal(t, v3Split.Spec.Service, v4Test.Spec.Service)

	assert.Len(t, v3Split.Spec.Backends, len(v4Test.Spec.Backends))

	for i, b := range v3Split.Spec.Backends {
		assert.Equal(t, b.Service, v4Test.Spec.Backends[i].Service)
		assert.Equal(t, b.Weight, v4Test.Spec.Backends[i].Weight)
	}

	for i, m := range v3Split.Spec.Matches {
		assert.Equal(t, m.Name, v4Test.Spec.Matches[i].Name)
		assert.Equal(t, m.Kind, v4Test.Spec.Matches[i].Kind)
		assert.Equal(t, m.APIGroup, v4Test.Spec.Matches[i].APIGroup)
	}
}

var v3Split = &TrafficSplit{
	Spec: TrafficSplitSpec{
		Service: "testservice",
		Backends: []TrafficSplitBackend{
			TrafficSplitBackend{
				Service: "v1",
				Weight:  100,
			},
			TrafficSplitBackend{
				Service: "v2",
				Weight:  200,
			},
		},
		Matches: []corev1.TypedLocalObjectReference{
			corev1.TypedLocalObjectReference{
				Kind: "HTTPRouteGroup",
				Name: "route1",
			},
			corev1.TypedLocalObjectReference{
				Kind: "HTTPRouteGroup",
				Name: "route2",
			},
		},
	},
}

var v4Split = &v1alpha4.TrafficSplit{
	Spec: v1alpha4.TrafficSplitSpec{
		Service: "testservice",
		Backends: []v1alpha4.TrafficSplitBackend{
			v1alpha4.TrafficSplitBackend{
				Service: "v1",
				Weight:  100,
			},
			v1alpha4.TrafficSplitBackend{
				Service: "v2",
				Weight:  200,
			},
		},
		Matches: []corev1.TypedLocalObjectReference{
			corev1.TypedLocalObjectReference{
				Name: "match1",
				Kind: "HTTPRouteGroup",
			},
			corev1.TypedLocalObjectReference{
				Name: "match2",
				Kind: "HTTPRouteGroup",
			},
		},
	},
}
