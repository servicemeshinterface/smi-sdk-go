# smi-sdk-go
SMI SDK for Golang

Shared code and example implementations

## Codegen

Client code for the trafficsplit CRD is generated.  The files:

* `gen/apis/trafficsplit/v1beta1/doc.go`
* `gen/apis/trafficsplit/v1beta1/register.go`
* `gen/apis/trafficsplit/v1beta1/trafficsplit.go`
* `gen/apis/trafficsplit/register.go`

are the original source files and can be manually edited.  All other files in
the `gen` directory are code generated and should not be edited directly.

Note that the code-generator does not support groupNames with hyphen characters.
Therefore, before running the code generator, you must rename all instances of
"smi-spec.io" in the project to "smispec.io".  The `update-codegen.sh` script
will rename these back to "smi-spec.io" after code generation is complete.

### Regenerating the Client Code

To regenerate the client code, follow these steps:

1. Install the code-generator tool: `go get k8s.io/code-generator`
1. Ensure the group name does not contain any hypen characters (see note above)
1. Run `hack/update-codegen.sh`