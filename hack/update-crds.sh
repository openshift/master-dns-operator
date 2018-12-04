#!/bin/sh

set -e

SRC_DIR="$(git rev-parse --show-toplevel)"

OUTPUT_DIR="$(mktemp -d)"

cd "${SRC_DIR}"
go build -o ./bin/controller-gen ./vendor/sigs.k8s.io/controller-tools/cmd/controller-gen

./bin/controller-gen \
	crd \
	--output-dir="${OUTPUT_DIR}"
    
cp ${OUTPUT_DIR}/masterdns_v1alpha1_masterdnsoperatorconfig.yaml ${SRC_DIR}/manifests/01_config_crd.yaml
cp ${OUTPUT_DIR}/masterdns_v1alpha1_dnsendpoint.yaml ${SRC_DIR}/manifests/01_dns_crd.yaml
