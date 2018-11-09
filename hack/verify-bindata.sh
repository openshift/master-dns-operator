#!/bin/sh

set -e

TMP_OUTPUT="$(mktemp)"
SRC_DIR="$(git rev-parse --show-toplevel)"

export EXISTING_FILE="${SRC_DIR}/pkg/operator/assets/bindata.go"
export OUTPUT_FILE="${TMP_OUTPUT}"

${SRC_DIR}/hack/update-bindata.sh

diff -Naup "${EXISTING_FILE}" "${TMP_OUTPUT}"
