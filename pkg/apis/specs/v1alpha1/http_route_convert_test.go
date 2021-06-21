package v1alpha1

import (
	"testing"

	"github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha4"
	assert "github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestConvertHTTPRouteGroupToConvertsFromAlpha4ToAlpha2(t *testing.T) {
	v1Test := &HTTPRouteGroup{}

	err := v1Test.ConvertFrom(v4HTTPRoute)
	assert.NoError(t, err)

	assert.Equal(t, v4HTTPRoute.ObjectMeta, v1Test.ObjectMeta)
	assert.Equal(t, v4HTTPRoute.TypeMeta.Kind, v1Test.TypeMeta.Kind)
	assert.Equal(t, "v1alpha3", v1Test.TypeMeta.APIVersion)

	for i, m := range v4HTTPRoute.Spec.Matches {
		v1 := v1Test.Matches[i]

		assert.Equal(t, m.Name, v1.Name)

		// check the methods
		assert.Equal(t, m.Methods[0], v1.Methods[0])
		assert.Equal(t, m.Methods[1], v1.Methods[1])

		assert.Equal(t, m.PathRegex, v1.PathRegex)
	}
}

func TestConvertHTTPRouteGroupToConvertsFromAlpha3ToAlpha2(t *testing.T) {
	v4Test := &v1alpha4.HTTPRouteGroup{}

	err := v1HTTPRoute.ConvertTo(v4Test)
	assert.NoError(t, err)

	assert.Equal(t, v1HTTPRoute.ObjectMeta, v4Test.ObjectMeta)
	assert.Equal(t, v1HTTPRoute.TypeMeta.Kind, v4Test.TypeMeta.Kind)
	assert.Equal(t, "v1alpha4", v4Test.TypeMeta.APIVersion)

	for i, m := range v1HTTPRoute.Matches {
		v4 := v4Test.Spec.Matches[i]

		assert.Equal(t, m.Name, v4.Name)

		// check the methods
		assert.Equal(t, m.Methods[0], v4.Methods[0])
		assert.Equal(t, m.Methods[1], v4.Methods[1])

		assert.Equal(t, m.PathRegex, v4.PathRegex)

		// check the headers
		assert.Len(t, v4.Headers, 0)
	}
}

var v4HTTPRoute = &v1alpha4.HTTPRouteGroup{
	TypeMeta: v1.TypeMeta{
		Kind:       "HTTPRouteGroup",
		APIVersion: "v1alpha4",
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v4HTTPRouteGroup",
		Namespace: "default",
	},
	Spec: v1alpha4.HTTPRouteGroupSpec{
		Matches: []v1alpha4.HTTPMatch{
			v1alpha4.HTTPMatch{
				Name:      "testing1",
				Methods:   []string{"GET", "POST"},
				PathRegex: ".*",
				Headers:   map[string]string{"Foo": "Bar", "One": "Two"},
			},
			v1alpha4.HTTPMatch{
				Name:      "testing2",
				Methods:   []string{"DELETE", "POST"},
				PathRegex: "/post",
				Headers:   map[string]string{"abc": "123", "Mario": "Luigi"},
			},
		},
	},
}

var v1HTTPRoute = &HTTPRouteGroup{
	TypeMeta: v1.TypeMeta{
		Kind:       "HTTPRouteGroup",
		APIVersion: "v1alpha3",
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v3HTTPRouteGroup",
		Namespace: "default",
	},
	Matches: []HTTPMatch{
		HTTPMatch{
			Name:      "testing1",
			Methods:   []string{"GET", "POST"},
			PathRegex: ".*",
		},
		HTTPMatch{
			Name:      "testing2",
			Methods:   []string{"DELETE", "POST"},
			PathRegex: "/post",
		},
	},
}
