package v1alpha2

import (
	"path"
)

const APIVersion = "metrics.smi-spec.io/v1alpha2"

func baseURL() string {
	return path.Join("/apis", APIVersion)
}
