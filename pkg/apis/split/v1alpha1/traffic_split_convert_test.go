package v1alpha1

import (
	"testing"

	"github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha4"
	assert "github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func testGetv1Split() *TrafficSplit {
	weight1, _ := resource.ParseQuantity("100")
	weight2, _ := resource.ParseQuantity("200")

	vs := *v1Split
	vs.Spec.Backends[0].Weight = &weight1
	vs.Spec.Backends[1].Weight = &weight2

	return &vs
}

func TestConvertToConvertsToAlpha1FromAlpha4(t *testing.T) {
	v1Test := TrafficSplit{}

	err := v1Test.ConvertFrom(v4Split)
	assert.NoError(t, err)

	assert.Equal(t, v4Split.Spec.Service, v1Test.Spec.Service)

	assert.Len(t, v1Test.Spec.Backends, len(v4Split.Spec.Backends))

	for i, b := range v4Split.Spec.Backends {
		assert.Equal(t, b.Service, v1Test.Spec.Backends[i].Service)

		weight := resource.NewQuantity(int64(b.Weight), resource.DecimalSI)
		assert.Equal(t, weight, v1Test.Spec.Backends[i].Weight)
	}
}

func TestConvertToConvertsFromAlpha1ToAlpha4(t *testing.T) {
	v4Test := *v4Split
	v1Test := testGetv1Split()

	err := v1Test.ConvertTo(&v4Test)
	assert.NoError(t, err)

	assert.Equal(t, v1Test.Spec.Service, v4Test.Spec.Service)

	assert.Len(t, v4Test.Spec.Backends, len(v1Test.Spec.Backends))

	for i, b := range v1Test.Spec.Backends {
		assert.Equal(t, b.Service, v4Test.Spec.Backends[i].Service)

		weight := b.Weight.AsDec().UnscaledBig().Int64()
		assert.Equal(t, int(weight), v4Test.Spec.Backends[i].Weight)
	}
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

var v1Split = &TrafficSplit{
	Spec: TrafficSplitSpec{
		Service: "testservice",
		Backends: []TrafficSplitBackend{
			TrafficSplitBackend{
				Service: "v1",
			},
			TrafficSplitBackend{
				Service: "v2",
			},
		},
	},
}
