GO_PKGS = ./cmd/... ./pkg/...
IMAGE_BUILD_CMD = imagebuilder
IMAGE = "openshift/master-dns-operator:latest"

all: generate fmt vet lint build

.PHONY: lint
lint:
	golint --set_exit_status $(GO_PKGS)

.PHONY: fmt
fmt:
	gofmt -w -s ./pkg ./cmd

.PHONY: vet
vet:
	go vet $(GO_PKGS)

.PHONY: install
install:
	oc apply -f ./manifests

.PHONY: build
build: manager external-dns

.PHONY: manager
manager: generate
	go build -o bin/manager ./cmd/manager/

.PHONY: external-dns
external-dns:
	go build -o bin/external-dns ./vendor/github.com/kubernetes-incubator/external-dns/

.PHONY: test-unit
test-unit:
	go test ./pkg/... ./cmd/... -coverprofile cover.out

.PHONY: verify
verify: verify-bindata verify-gofmt lint vet

.PHONY: verify-gofmt
verify-gofmt:
	hack/verify-gofmt.sh

.PHONY: generate
generate: generate-go generate-crds generate-rbac generate-bindata

.PHONY: generate-go
generate-go:
	go generate $(GO_PKGS)

.PHONY: verify-bindata
verify-bindata:
	hack/verify-bindata.sh

.PHONY: generate-bindata
generate-bindata:
	hack/update-bindata.sh

.PHONY: generate-crds
generate-crds:
	hack/update-crds.sh

.PHONY: generate-rbac
generate-rbac:
	hack/update-rbac.sh

.PHONY: run
run:
	WATCH_NAMESPACE=openshift-master-dns IMAGE=quay.io/csrwng/master-dns-operator:latest ./bin/manager

.PHONY: image
image:  generate
	$(IMAGE_BUILD_CMD) -t $(IMAGE) .
