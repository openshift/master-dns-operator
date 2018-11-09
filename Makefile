GO_PKGS = ./cmd/... ./pkg/...
IMAGE_BUILD_CMD = imagebuilder
IMAGE = "openshift/master-dns-operator:latest"

all: fmt vet lint build

.PHONY: lint
lint:
	golint $(GO_PKGS)

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

.PHONY: generate
generate: 
	go generate $(GO_PKGS)

.PHONY: image
image:  generate
	$(IMAGE_BUILD_CMD) -t $(IMAGE) .
