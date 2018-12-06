#!/bin/sh

set -e

TMP_OUTPUT="$(mktemp)"
SRC_DIR="$(git rev-parse --show-toplevel)"

gofmt -l -s ${SRC_DIR}/cmd ${SRC_DIR}/pkg > ${TMP_OUTPUT} 2>&1 || true
if [[  -s ${TMP_OUTPUT} ]]; then
	echo "*** Please run 'make fmt' in order to fix the following:"
	cat ${TMP_OUTPUT}
	exit 1
fi
