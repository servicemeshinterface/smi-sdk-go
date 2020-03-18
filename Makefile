
.PHONY: bootstrap
bootstrap:
	go mod tidy

.PHONY: test
test:
	go test -v ./...

.PHONY: codegen
codegen:
	./hack/update-codegen.sh

.PHONY: verify
verify:
	./hack/verify-codegen.sh
