#!/bin/sh

base=$(dirname $0)/..

go run $base/gen.go "$@"