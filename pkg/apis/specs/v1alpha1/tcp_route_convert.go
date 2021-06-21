package v1alpha1

import (
	"github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha4"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// log is for logging in this package.
var tcproutelog = logf.Log.WithName("tcproute-resource")

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
// ConvertTo converts this TrafficTarget to the Hub version (v1alpha4).
func (src *TCPRoute) ConvertTo(dstRaw conversion.Hub) error {
	tcproutelog.Info("ConvertTo v1alpha4 from v1alpha1")

	dst := dstRaw.(*v1alpha4.TCPRoute)
	dst.ObjectMeta = src.ObjectMeta

	dst.TypeMeta = src.TypeMeta
	dst.APIVersion = "v1alpha4"

	return nil
}

/*
ConvertFrom is expected to modify its receiver to contain the converted object.
Most of the conversion is straightforward copying, except for converting our changed field.
*/

// ConvertFrom converts from the Hub version (v1alpha4) to this version.
func (dst *TCPRoute) ConvertFrom(srcRaw conversion.Hub) error {
	tcproutelog.Info("ConvertFrom v1alpha4 to v1alpha1")

	src := srcRaw.(*v1alpha4.TCPRoute)
	dst.ObjectMeta = src.ObjectMeta

	dst.TypeMeta = src.TypeMeta
	dst.APIVersion = "v1alpha1"

	return nil
}
