package v1alpha2

import (
	"testing"

	"github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha4"
	assert "github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
)

func TestConvertToConvertsFromAlpha4ToAlpha2(t *testing.T) {
	v2Test := &TrafficSplit{}

	err := v2Test.ConvertFrom(v4Split)
	assert.NoError(t, err)

	assert.Equal(t, v4Split.Spec.Service, v2Test.Spec.Service)

	assert.Len(t, v4Split.Spec.Backends, len(v2Test.Spec.Backends))

	for i, b := range v4Split.Spec.Backends {
		assert.Equal(t, b.Service, v2Test.Spec.Backends[i].Service)
		assert.Equal(t, b.Weight, v2Test.Spec.Backends[i].Weight)
	}
}

func TestConvertToConvertsToAlpha4FromAlpha2(t *testing.T) {
	v4Test := &v1alpha4.TrafficSplit{}

	err := v2Split.ConvertTo(v4Test)
	assert.NoError(t, err)

	assert.Equal(t, v2Split.Spec.Service, v4Test.Spec.Service)

	assert.Len(t, v2Split.Spec.Backends, len(v4Test.Spec.Backends))

	for i, b := range v2Split.Spec.Backends {
		assert.Equal(t, b.Service, v4Test.Spec.Backends[i].Service)
		assert.Equal(t, b.Weight, v4Test.Spec.Backends[i].Weight)
	}
}

var v2Split = &TrafficSplit{
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
