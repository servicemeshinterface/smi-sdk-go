TARGETS ?= darwin/amd64 linux/amd64 windows/amd64
DIST_DIRS = find * -type d -exec

# go options
BINDIR := $(CURDIR)/bin
GO ?= go
GOFLAGS :=
TAGS :=
LDFLAGS := -w -s

.PHONY: build
build:
	GOBIN=$(BINDIR) $(GO) install $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' github.com/deislabs/smi-sdk-go/cmd/installer

.PHONY: build-cross
build-cross: LDFLAGS += -extldflags "-static"
build-cross:
	CGO_ENABLED=0 gox -parallel=3 \
		-output="_dist/{{.OS}}-{{.Arch}}/{{.Dir}}" \
		-osarch='$(TARGETS)' $(GOFLAGS) \
		$(if $(TAGS),-tags '$(TAGS)',) -ldflags '$(LDFLAGS)' \
		github.com/deislabs/smi-sdk-go/cmd/installer

.PHONY: checksum
checksum:
	for f in _dist/*.{gz,zip} ; do \
			shasum -a 256 "$${f}"  | awk '{print $$1}' > "$${f}.sha256" ; \
			echo -n "Checksum: " && cat $${f}.sha256 ; \
	done

.PHONY: dist
dist: build-cross
dist:
		cd _dist && \
		$(DIST_DIRS) cp ../LICENSE {} \; && \
		$(DIST_DIRS) cp ../README.md {} \; && \
		$(DIST_DIRS) tar -zcf installer-${VERSION}-{}.tar.gz {} \; && \
		$(DIST_DIRS) zip -r installer-${VERSION}-{}.zip {} \; \

.PHONY: bootstrap
bootstrap:
	go mod tidy

.PHONY: test
test:
	go test -v ./...
