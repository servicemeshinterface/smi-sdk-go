package v1alpha2

import (
	"testing"

	"github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1"
	assert "github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestConvertToConvertsToAlpha1FromAlpha2(t *testing.T) {
	v1Test := &v1alpha1.TrafficSplit{}

	err := v2.ConvertTo(v1)
	assert.NoError(t, err)

	assert.Equal(t, v2.Spec.Service, v1Test.Spec.Service)

	assert.Len(t, v1Test.Spec.Backends, len(v2.Spec.Backends))

	for i, b := range v2.Spec.Backends {
		assert.Equal(t, b.Service, v1Test.Spec.Backends[i].Service)

		weight := resource.NewQuantity(int64(b.Weight), resource.DecimalSI)
		assert.Equal(t, weight, v1Test.Spec.Backends[i].Weight)
	}
}

func TestConvertToConvertsFromAlpha2ToAlpha1(t *testing.T) {
	v2Test := &TrafficSplit{}

	err := v2Test.ConvertFrom(v1)
	assert.NoError(t, err)

	assert.Equal(t, v1.Spec.Service, v2Test.Spec.Service)

	assert.Len(t, v2Test.Spec.Backends, len(v1.Spec.Backends))

	for i, b := range v1.Spec.Backends {
		assert.Equal(t, b.Service, v2Test.Spec.Backends[i].Service)

		weight, _ := b.Weight.AsInt64()
		assert.Equal(t, int(weight), v2Test.Spec.Backends[i].Weight)
	}
}

var v2 = &TrafficSplit{
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

var v1 = &v1alpha1.TrafficSplit{
	Spec: v1alpha1.TrafficSplitSpec{
		Service: "testservice",
		Backends: []v1alpha1.TrafficSplitBackend{
			v1alpha1.TrafficSplitBackend{
				Service: "v1",
				Weight:  resource.NewQuantity(int64(100), resource.DecimalSI),
			},
			v1alpha1.TrafficSplitBackend{
				Service: "v2",
				Weight:  resource.NewQuantity(int64(200), resource.DecimalSI),
			},
		},
	},
}
