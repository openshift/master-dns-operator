#!/bin/sh

set -e

SRC_DIR="$(git rev-parse --show-toplevel)"

OUTPUT_DIR="$(mktemp -d)"

cd "${SRC_DIR}"
go build -o ./bin/controller-gen ./vendor/sigs.k8s.io/controller-tools/cmd/controller-gen

./bin/controller-gen \
	rbac \
	--name=master-dns-operator \
	--output-dir="${OUTPUT_DIR}"
    
cp ${OUTPUT_DIR}/rbac_role.yaml ${SRC_DIR}/manifests/02_role.yaml
sed -e 's/ default/ master-dns-operator/' -e 's/system/openshift-master-dns/' ${OUTPUT_DIR}/rbac_role_binding.yaml > ${SRC_DIR}/manifests/04_binding.yaml
