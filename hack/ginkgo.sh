#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

ROOT=$(unset CDPATH && cd $(dirname "${BASH_SOURCE[0]}")/.. && pwd)
cd $ROOT

go run github.com/onsi/ginkgo/ginkgo "$@"
