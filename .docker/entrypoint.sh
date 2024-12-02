#!/bin/bash

set -euo pipefail

if [ ! -d bin ] ; then
	mkdir bin
fi


./build.sh ${1}

export GOGC=off
time "bin/day${1}"
