package v1alpha1

import (
	"github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha4"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// log is for logging in this package.
var httproutelog = logf.Log.WithName("httproute-resource")

/*
Our "spoke" versions need to implement the
[`Convertible`](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/conversion?tab=doc#Convertible)
interface.  Namely, they'll need `ConvertTo` and `ConvertFrom` methods to convert to/from
the hub version.
*/

/*
ConvertTo is expected to modify its argument to contain the converted object.
Most of the conversion is straightforward copying, except for converting our changed field.
*/
// ConvertTo converts this TrafficTarget to the Hub version (v1alpha3).
func (src *HTTPRouteGroup) ConvertTo(dstRaw conversion.Hub) error {
	httproutelog.Info("ConvertTo v1alpha4 from v1alpha2")

	dst := dstRaw.(*v1alpha4.HTTPRouteGroup)
	dst.ObjectMeta = src.ObjectMeta

	dst.TypeMeta = src.TypeMeta
	dst.APIVersion = "v1alpha4"

	for _, m := range src.Matches {
		dst.Spec.Matches = append(dst.Spec.Matches, v1alpha4.HTTPMatch{
			Name:      m.Name,
			PathRegex: m.PathRegex,
			Methods:   m.Methods,
		})
	}

	return nil
}

/*
ConvertFrom is expected to modify its receiver to contain the converted object.
Most of the conversion is straightforward copying, except for converting our changed field.
*/

// ConvertFrom converts from the Hub version (v1alpha3) to this version.
func (dst *HTTPRouteGroup) ConvertFrom(srcRaw conversion.Hub) error {
	httproutelog.Info("ConvertFrom v1alpha4 to v1alpha2")

	src := srcRaw.(*v1alpha4.HTTPRouteGroup)
	dst.ObjectMeta = src.ObjectMeta

	dst.TypeMeta = src.TypeMeta
	dst.APIVersion = "v1alpha3"

	for _, m := range src.Spec.Matches {
		dst.Matches = append(dst.Matches, HTTPMatch{
			Name:      m.Name,
			PathRegex: m.PathRegex,
			Methods:   m.Methods,
		})
	}

	return nil
}
