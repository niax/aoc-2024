#!/bin/bash

set -euo pipefail

if [ ! -d bin ] ; then
	mkdir bin
fi

MAIN_FILE="cmd/day${1}/day${1}.go"
if [ -e "${MAIN_FILE}" ]; then
	export GOAMD64=$(cpuid --json | jq '"v" + (.X64Level | tostring)' -r)
	go build -o bin/ ${MAIN_FILE}

	export GOGC=off
	time "bin/day${1}"
fi
