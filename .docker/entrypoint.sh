#!/bin/bash

set -euo pipefail

if [ ! -d bin ] ; then
	mkdir bin
fi


./build.sh ${1}

export GOGC=off
bin/day${1}
hyperfine --warmup 16 -N "bin/day${1}"
