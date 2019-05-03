package metrics

import (
	"testing"
	"time"

	assertLib "github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
)

func TestNewResource(t *testing.T) {
	assert := assertLib.New(t)

	testCases := []struct {
		obj *v1.ObjectReference
		has []string
		len int
	}{
		{
			&v1.ObjectReference{
				APIVersion: "apps/v1",
				Kind:       "Deployment",
				Name:       "foo",
				Namespace:  "bar",
			},
			[]string{"deployments", "bar"},
			0,
		},
		{
			&v1.ObjectReference{
				Kind: "Random",
			},
			[]string{"unsupported"},
			0,
		},
		{
			&v1.ObjectReference{
				APIVersion: "v1",
				Kind:       "Namespace",
				Name:       "foo",
			},
			[]string{"namespaces", "foo"},
			0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.obj.Kind, func(t *testing.T) {
			r := NewTrafficMetrics(tc.obj)

			assert.Equal("TrafficMetrics", r.TypeMeta.Kind)
			assert.Equal(APIVersion, r.TypeMeta.APIVersion)

			assert.WithinDuration(
				time.Now(),
				r.ObjectMeta.CreationTimestamp.Time,
				10*time.Millisecond)
			for _, item := range tc.has {
				assert.Contains(r.ObjectMeta.SelfLink, item)
			}
			assert.Equal(tc.obj.Name, r.ObjectMeta.Name)
			assert.Equal(tc.obj.Namespace, r.ObjectMeta.Namespace)

			assert.Equal(tc.obj, r.Resource)

			assert.Len(r.Metrics, 5)
			for _, item := range AvailableMetrics {
				assert.Equal(item, r.Get(item.Name))
			}

			assert.Nil(r.Get("foobar"))
		})
	}
}
