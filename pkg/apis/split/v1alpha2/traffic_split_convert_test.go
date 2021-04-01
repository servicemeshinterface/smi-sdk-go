package v1alpha2

import (
	"testing"

	"github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1"
	assert "github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/resource"
)

func testGetv1Split() *v1alpha1.TrafficSplit {
	weight1, _ := resource.ParseQuantity("100")
	weight2, _ := resource.ParseQuantity("200")

	v1Split := *v1
	v1Split.Spec.Backends[0].Weight = &weight1
	v1Split.Spec.Backends[1].Weight = &weight2

	return &v1Split
}

func TestConvertToConvertsToAlpha1FromAlpha2(t *testing.T) {
	v1Test := &v1alpha1.TrafficSplit{}

	err := v2.ConvertTo(v1Test)
	assert.NoError(t, err)

	assert.Equal(t, v2.Spec.Service, v1Test.Spec.Service)

	assert.Len(t, v1Test.Spec.Backends, len(v2.Spec.Backends))

	for i, b := range v2.Spec.Backends {
		assert.Equal(t, b.Service, v1Test.Spec.Backends[i].Service)

		weight := resource.NewQuantity(int64(b.Weight), resource.DecimalSI)
		assert.Equal(t, weight, v1Test.Spec.Backends[i].Weight)
	}
}

func TestConvertToConvertsFromAlpha1ToAlpha2(t *testing.T) {
	v2Test := &TrafficSplit{}
	v1Test := testGetv1Split()

	err := v2Test.ConvertFrom(v1Test)
	assert.NoError(t, err)

	assert.Equal(t, v1Test.Spec.Service, v2Test.Spec.Service)

	assert.Len(t, v2Test.Spec.Backends, len(v1Test.Spec.Backends))

	for i, b := range v1Test.Spec.Backends {
		assert.Equal(t, b.Service, v2Test.Spec.Backends[i].Service)

		weight := b.Weight.AsDec().UnscaledBig().Int64()
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
			},
			v1alpha1.TrafficSplitBackend{
				Service: "v2",
			},
		},
	},
}
