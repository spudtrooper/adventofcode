#!/bin/sh

set -e

base=$(dirname $0)/..

$base/scripts/test.sh
$base/2021/scripts/test.sh