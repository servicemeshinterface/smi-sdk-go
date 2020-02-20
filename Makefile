
.PHONY: bootstrap
bootstrap:
	go mod tidy

.PHONY: test
test:
	go test -v ./...
