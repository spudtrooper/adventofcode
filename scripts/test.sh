#!/bin/sh

set -e

base=$(dirname $0)/..

go test $base/common/*.go

rm -rf $base/8888/day99
$base/scripts/gen.sh -day 99 -year 8888 2>/dev/null
go test $base/8888/day99/lib/*.go
rm -rf $base/8888/day99