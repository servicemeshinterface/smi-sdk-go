
.PHONY: bootstrap
bootstrap:
	go mod tidy

.PHONY: test
test:
	go test -v ./...

.PHONY: verify
verify:
	./hack/verify-codegen.sh