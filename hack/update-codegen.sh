#!/usr/bin/env bash

set -eu

ROOT_PACKAGE="github.com/servicemeshinterface/smi-sdk-go"
ROOT_DIR="$(git rev-parse --show-toplevel)"

# get code-generator version from go.sum
CODEGEN_VERSION="$(grep 'k8s.io/code-generator' go.sum | awk '{print $2}' | head -1)"
CODEGEN_PKG="$(echo `go env GOPATH`"/pkg/mod/k8s.io/code-generator@${CODEGEN_VERSION}")"

echo ">>> using codegen: ${CODEGEN_PKG}"
# ensure we can execute the codegen script
chmod +x ${CODEGEN_PKG}/generate-groups.sh

function generate_client() {
  TEMP_DIR=$(mktemp -d)
  CUSTOM_RESOURCE_NAME=$1
  CUSTOM_RESOURCE_VERSIONS=$2

  # delete the generated deepcopy
  for V in ${CUSTOM_RESOURCE_VERSIONS//,/ }; do
    rm -f ${ROOT_DIR}/pkg/apis/${CUSTOM_RESOURCE_NAME}/${V}/zz_generated.deepcopy.go
  done

  # delete the generated code as this is additive, removed objects will not be cleaned
  rm -rf ${ROOT_DIR}/pkg/gen/client/${CUSTOM_RESOURCE_NAME}

  # the generate-groups.sh script cannot handle group names with dashes, so we use smispec.io as the group name
  if [[ "$OSTYPE" == "darwin"* ]]; then
    find "${ROOT_DIR}/pkg/apis/${CUSTOM_RESOURCE_NAME}" -type f -exec sed -i '' 's/smi-spec.io/smispec.io/g' {} +
  else
    find "${ROOT_DIR}/pkg/apis/${CUSTOM_RESOURCE_NAME}" -type f -exec sed -i 's/smi-spec.io/smispec.io/g' {} +
  fi

  # code-generator makes assumptions about the project being located in `$GOPATH/src`.
  # To work around this we create a temporary directory, use it as output base and copy everything back once generated.
  "${CODEGEN_PKG}"/generate-groups.sh all \
    "$ROOT_PACKAGE/pkg/gen/client/$CUSTOM_RESOURCE_NAME" \
    "$ROOT_PACKAGE/pkg/apis" \
    $CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSIONS \
    --go-header-file "${ROOT_DIR}"/hack/boilerplate.go.txt \
    --output-base "${TEMP_DIR}"

  cp -r "${TEMP_DIR}/${ROOT_PACKAGE}/." "${ROOT_DIR}/"
  rm -rf ${TEMP_DIR}

  # replace smispec.io with smi-spec.io after code generation
  if [[ "$OSTYPE" == "darwin"* ]]; then
    find "${ROOT_DIR}/pkg/apis/${CUSTOM_RESOURCE_NAME}" -type f -exec sed -i '' 's/smispec.io/smi-spec.io/g' {} +
    find "${ROOT_DIR}/pkg/gen/client/${CUSTOM_RESOURCE_NAME}" -type f -exec sed -i '' 's/smispec.io/smi-spec.io/g' {} +
  else
    find "${ROOT_DIR}/pkg/apis/${CUSTOM_RESOURCE_NAME}" -type f -exec sed -i 's/smispec.io/smi-spec.io/g' {} +
    find "${ROOT_DIR}/pkg/gen/client/${CUSTOM_RESOURCE_NAME}" -type f -exec sed -i 's/smispec.io/smi-spec.io/g' {} +
  fi
}

echo "##### Generating specs client ######"
generate_client "specs" "v1alpha1,v1alpha2,v1alpha4"

echo ""
echo "###### Generating split client ######"
generate_client "split" "v1alpha1,v1alpha2,v1alpha3"

echo ""
echo "##### Generating access client ######"
generate_client "access" "v1alpha1"

echo ""
echo "##### Generating metrics client ######"
generate_client "metrics" "v1alpha1,v1alpha2"
