package metrics

import (
	"path"
)

const APIVersion = "metrics.smi-spec.io/v1beta1"

func baseURL() string {
	return path.Join("/apis", APIVersion)
}
