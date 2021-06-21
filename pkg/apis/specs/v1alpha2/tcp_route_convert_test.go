package v1alpha2

import (
	"testing"

	"github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha4"
	assert "github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestConvertToConvertsFromAlpha4ToAlpha2(t *testing.T) {
	v2Test := &TCPRoute{}

	err := v2Test.ConvertFrom(v4TCPRoute)
	assert.NoError(t, err)

	assert.Equal(t, v4TCPRoute.ObjectMeta, v2Test.ObjectMeta)
	assert.Equal(t, v4TCPRoute.TypeMeta.Kind, v2Test.TypeMeta.Kind)
	assert.Equal(t, "v1alpha2", v2Test.TypeMeta.APIVersion)
}

func TestConvertToConvertsFromAlpha2ToAlpha4(t *testing.T) {
	v4Test := &v1alpha4.TCPRoute{}

	err := v2TCPRoute.ConvertTo(v4Test)
	assert.NoError(t, err)

	assert.Equal(t, v2TCPRoute.ObjectMeta, v4Test.ObjectMeta)
	assert.Equal(t, v2TCPRoute.TypeMeta.Kind, v4Test.TypeMeta.Kind)
	assert.Equal(t, "v1alpha4", v4Test.TypeMeta.APIVersion)

	// should have a blank matches as this does not exist in the v2 spec
	assert.Equal(t, v1alpha4.TCPRouteSpec{}, v4Test.Spec)
}

var v4TCPRoute = &v1alpha4.TCPRoute{
	TypeMeta: v1.TypeMeta{
		Kind:       "TCPRoute",
		APIVersion: "v1alpha4",
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v4Specs",
		Namespace: "default",
	},
	Spec: v1alpha4.TCPRouteSpec{
		Matches: v1alpha4.TCPMatch{
			Name:  "testing",
			Ports: []int{9090, 8080},
		},
	},
}

var v2TCPRoute = &TCPRoute{
	TypeMeta: v1.TypeMeta{
		Kind:       "TCPRoute",
		APIVersion: "v1alpha2",
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v1Specs",
		Namespace: "default",
	},
}
