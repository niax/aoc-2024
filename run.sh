#!/bin/bash

set -euo pipefail

IMAGE=niax/aoc2024:1

docker image inspect "${IMAGE}" >/dev/null 2>&1 || docker build .docker -t "${IMAGE}"

exec docker run --rm -i -v "${PWD}:/code" "${IMAGE}" "${@}"
