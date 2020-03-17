
.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test:
	go test -v ./...

.PHONY: fmt
fmt:
	gofmt -l -s -w ./

.PHONY: verify-fmt
verify-fmt:
	gofmt -l -s ./ | grep ".*\.go"; if [ "$$?" = "0" ]; then exit 1; fi

.PHONY: codegen
codegen:
	./hack/update-codegen.sh

.PHONY: verify-codegen
verify-codegen:
	./hack/verify-codegen.sh
