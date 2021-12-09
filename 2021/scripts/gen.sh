#!/bin/sh

set -e

base=$(dirname $0)/../..

$base/scripts/gen.sh --year 2021 "$@"