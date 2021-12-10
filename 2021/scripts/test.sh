#!/bin/sh

set -e

base=$(dirname $0)/..

for d in $base/day*; do
    if [[ -d $d ]]; then
        go test $d/*.go
    fi
done