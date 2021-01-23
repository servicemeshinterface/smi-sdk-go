package v1alpha2

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/conversion"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// log is for logging in this package.
var traffictargetlog = logf.Log.WithName("traffictarget-resource")

func (src *TrafficTarget) ConvertTo(dstRaw conversion.Hub) error {
	traffictargetlog.Info("ConvertTo", dstRaw)
	return fmt.Errorf("not implemented")
}

func (src *TrafficTarget) ConvertFrom(srcRaw conversion.Hub) error {
	traffictargetlog.Info("ConvertFrom", srcRaw)
	return fmt.Errorf("not implemented")
}
