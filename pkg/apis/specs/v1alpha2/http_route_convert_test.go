package v1alpha2

import (
	"testing"

	"github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha4"
	assert "github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestConvertHTTPRouteGroupToConvertsFromAlpha4ToAlpha2(t *testing.T) {
	v2Test := &HTTPRouteGroup{}

	err := v2Test.ConvertFrom(v4HTTPRoute)
	assert.NoError(t, err)

	assert.Equal(t, v4HTTPRoute.ObjectMeta, v2Test.ObjectMeta)
	assert.Equal(t, v4HTTPRoute.TypeMeta.Kind, v2Test.TypeMeta.Kind)
	assert.Equal(t, "v1alpha3", v2Test.TypeMeta.APIVersion)

	for i, m := range v4HTTPRoute.Spec.Matches {
		v2 := v2Test.Matches[i]

		assert.Equal(t, m.Name, v2.Name)

		// check the methods
		assert.Equal(t, m.Methods[0], v2.Methods[0])
		assert.Equal(t, m.Methods[1], v2.Methods[1])

		assert.Equal(t, m.PathRegex, v2.PathRegex)

		// check the headers
		for k, v := range m.Headers {
			assert.Equal(t, v2.Headers[k], v)
		}
	}
}

func TestConvertHTTPRouteGroupToConvertsFromAlpha3ToAlpha2(t *testing.T) {
	v4Test := &v1alpha4.HTTPRouteGroup{}

	err := v2HTTPRoute.ConvertTo(v4Test)
	assert.NoError(t, err)

	assert.Equal(t, v2HTTPRoute.ObjectMeta, v4Test.ObjectMeta)
	assert.Equal(t, v2HTTPRoute.TypeMeta.Kind, v4Test.TypeMeta.Kind)
	assert.Equal(t, "v1alpha4", v4Test.TypeMeta.APIVersion)

	for i, m := range v2HTTPRoute.Matches {
		v4 := v4Test.Spec.Matches[i]

		assert.Equal(t, m.Name, v4.Name)

		// check the methods
		assert.Equal(t, m.Methods[0], v4.Methods[0])
		assert.Equal(t, m.Methods[1], v4.Methods[1])

		assert.Equal(t, m.PathRegex, v4.PathRegex)

		// check the headers
		for k, v := range m.Headers {
			assert.Equal(t, v4.Headers[k], v)
		}
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

var v2HTTPRoute = &HTTPRouteGroup{
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
			Headers:   map[string]string{"Foo": "Bar", "One": "Two"},
		},
		HTTPMatch{
			Name:      "testing2",
			Methods:   []string{"DELETE", "POST"},
			PathRegex: "/post",
			Headers:   map[string]string{"abc": "123", "Mario": "Luigi"},
		},
	},
}
