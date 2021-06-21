package v1alpha3

import (
	"testing"

	"github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha4"
	assert "github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestConvertToConvertsFromAlpha4ToAlpha3(t *testing.T) {
	v3Test := &TCPRoute{}

	err := v3Test.ConvertFrom(v4TCPRoute)
	assert.NoError(t, err)

	assert.Equal(t, v4TCPRoute.ObjectMeta, v3Test.ObjectMeta)
	assert.Equal(t, v4TCPRoute.TypeMeta.Kind, v3Test.TypeMeta.Kind)
	assert.Equal(t, "v1alpha3", v3Test.TypeMeta.APIVersion)

	assert.Equal(t, TCPRouteSpec{}, v3Test.Spec)
}

func TestConvertToConvertsFromAlpha3ToAlpha4(t *testing.T) {
	v4Test := &v1alpha4.TCPRoute{}

	err := v3TCPRoute.ConvertTo(v4Test)
	assert.NoError(t, err)

	assert.Equal(t, v3TCPRoute.ObjectMeta, v4Test.ObjectMeta)
	assert.Equal(t, v3TCPRoute.TypeMeta.Kind, v4Test.TypeMeta.Kind)
	assert.Equal(t, "v1alpha4", v4Test.TypeMeta.APIVersion)

	// should have a blank matches as this does not exist in the v3 spec
	assert.Equal(t, v1alpha4.TCPRouteSpec{}, v4Test.Spec)
}

var v4TCPRoute = &v1alpha4.TCPRoute{
	TypeMeta: v1.TypeMeta{
		Kind:       "TCPRoute",
		APIVersion: "v1alpha4",
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v3Access",
		Namespace: "default",
	},
	Spec: v1alpha4.TCPRouteSpec{
		Matches: v1alpha4.TCPMatch{
			Name:  "testing",
			Ports: []int{9090, 8080},
		},
	},
}

var v3TCPRoute = &TCPRoute{
	TypeMeta: v1.TypeMeta{
		Kind:       "TrafficTarget",
		APIVersion: "v1alpha1",
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v1Access",
		Namespace: "default",
	},
	Spec: TCPRouteSpec{},
}
