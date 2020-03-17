#!/usr/bin/env bash

set -eu

ROOT_PACKAGE="github.com/deislabs/smi-sdk-go"
ROOT_DIR="$(git rev-parse --show-toplevel)"
CODEGEN_VERSION="$(grep 'k8s.io/code-generator' go.sum | awk '{print $2}' | head -1)"
CODEGEN_PKG="~/go/pkg/mod/k8s.io/code-generator@${CODEGEN_VERSION}"

# make the generate script executable
chmod +x ${CODEGEN_PKG}/generate-groups.sh

echo ">> Using ${CODEGEN_PKG}"
echo ">> Project ${ROOT_DIR}"

function generate_client() {
  # CUSTOM_RESOURCE_NAME :: the name of the custom resource that we're generating client code for
  CUSTOM_RESOURCE_NAME=$1
  CUSTOM_RESOURCE_VERSIONS=$2

  # delete the generated code as this is additive, removed objects will not be cleaned

  # enumerate versions
  for V in ${CUSTOM_RESOURCE_VERSIONS//,/ }; do
    rm -f ${ROOT_DIR}/pkg/apis/${CUSTOM_RESOURCE_NAME}/${V}/zz_generated.deepcopy.go
  done
  rm -rf ${ROOT_DIR}/pkg/gen/client/${CUSTOM_RESOURCE_NAME}

  if [[ "$OSTYPE" == "darwin"* ]]; then
    find "${ROOT_DIR}/pkg/apis/${CUSTOM_RESOURCE_NAME}" -type f -exec sed -i '' 's/smi-spec.io/smispec.io/g' {} +
  else
    find "${ROOT_DIR}/pkg/apis/${CUSTOM_RESOURCE_NAME}" -type f -exec sed -i 's/smi-spec.io/smispec.io/g' {} +
  fi

  "${CODEGEN_PKG}"/generate-groups.sh \
    "deepcopy,client,informer,lister" \
    "$ROOT_PACKAGE/pkg/gen/client/$CUSTOM_RESOURCE_NAME" \
    "$ROOT_PACKAGE/pkg/apis" \
    $CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSIONS \
    --go-header-file "${ROOT_DIR}"/hack/boilerplate.go.txt

  # The generate-groups.sh script cannot handle group names with dashes, so we use
  # smispec.io as the group name, then replace it with smi-spec.io after code
  # generation.
  if [[ "$OSTYPE" == "darwin"* ]]; then
    find "${ROOT_DIR}/pkg/apis/${CUSTOM_RESOURCE_NAME}" -type f -exec sed -i '' 's/smispec.io/smi-spec.io/g' {} +
    find "${ROOT_DIR}/pkg/gen/client/${CUSTOM_RESOURCE_NAME}" -type f -exec sed -i '' 's/smispec.io/smi-spec.io/g' {} +
  else
    find "${ROOT_DIR}/pkg/apis/${CUSTOM_RESOURCE_NAME}" -type f -exec sed -i 's/smispec.io/smi-spec.io/g' {} +
    find "${ROOT_DIR}/pkg/gen/client/${CUSTOM_RESOURCE_NAME}" -type f -exec sed -i 's/smispec.io/smi-spec.io/g' {} +
  fi
}

echo "##### Generating specs client ######"
generate_client "specs" "v1alpha1,v1alpha2"

echo ""
echo "###### Generating split client ######"
generate_client "split" "v1alpha1,v1alpha2,v1alpha3"

echo ""
echo "##### Generating access client ######"
generate_client "access" "v1alpha1"

echo ""
echo "##### Generating metrics client ######"
generate_client "metrics" "v1alpha1,v1alpha2"