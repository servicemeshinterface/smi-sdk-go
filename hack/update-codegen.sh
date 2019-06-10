#!/bin/bash

set -eu

# ROOT_PACKAGE :: the package (relative to $GOPATH/src) that is the target for code generation
ROOT_PACKAGE="github.com/deislabs/smi-sdk-go"

# CUSTOM_RESOURCE_VERSION :: the version of the resource
CUSTOM_RESOURCE_VERSION="v1alpha1"

function generate_client() {
  
  SCRIPT_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
  ROOT_DIR="$( cd $SCRIPT_ROOT/.. && pwd )"
  GEN_VER=$( awk '/k8s.io\/code-generator/ { print $2 }' "$ROOT_DIR/go.mod" )
  CODEGEN_PKG=${GOPATH}/pkg/mod/k8s.io/code-generator@${GEN_VER}
  
  # CUSTOM_RESOURCE_NAME :: the name of the custom resource that we're generating client code for
  CUSTOM_RESOURCE_NAME=$1

  # make the generate script executable
  chmod +x ${CODEGEN_PKG}/generate-groups.sh

  # delete the generated code as this is additive, removed objects will not be cleaned
  rm -f ${ROOT_DIR}/pkg/apis/${CUSTOM_RESOURCE_NAME}/${CUSTOM_RESOURCE_VERSION}/zz_generated.deepcopy.go
  rm -rf ${ROOT_DIR}/pkg/gen/client/${CUSTOM_RESOURCE_NAME}
  
  find "${ROOT_DIR}/pkg/apis/${CUSTOM_RESOURCE_NAME}" -type f -exec sed -i 's/smi-spec.io/smispec.io/g' {} +
  
  # generate the code with:
  # --output-base    because this script should also be able to run inside the vendor dir of
  #                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
  #                  instead of the $GOPATH directly. For normal projects this can be dropped.
  GO111MODULE="on" "${CODEGEN_PKG}"/generate-groups.sh \
    "deepcopy,client,informer,lister" \
    "$ROOT_PACKAGE/pkg/gen/client/$CUSTOM_RESOURCE_NAME" \
    "$ROOT_PACKAGE/pkg/apis" \
    $CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION \
    --go-header-file "${SCRIPT_ROOT}"/boilerplate.go.txt
  
  # The generate-groups.sh script cannot handle group names with dashes, so we use
  # smispec.io as the group name, then replace it with smi-spec.io after code
  # generation.
  find "${ROOT_DIR}/pkg/apis/${CUSTOM_RESOURCE_NAME}" -type f -exec sed -i 's/smispec.io/smi-spec.io/g' {} +
  find "${ROOT_DIR}/pkg/gen/client/${CUSTOM_RESOURCE_NAME}" -type f -exec sed -i 's/smispec.io/smi-spec.io/g' {} +
}

echo "###### Generating Traffic Split Client ######"
generate_client "split"

echo ""
echo "##### Generating Traffic Access Client ######"
generate_client "specs"

echo ""
echo "##### Generating Traffic Spec Client ######"
generate_client "access"
