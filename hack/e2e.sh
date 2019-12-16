#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

ROOT=$(unset CDPATH && cd $(dirname "${BASH_SOURCE[0]}")/.. && pwd)
cd $ROOT

go test -c -o output/bin/e2e.test ./test/e2e
go build -o output/bin/ginkgo github.com/onsi/ginkgo/ginkgo
./output/bin/ginkgo output/bin/e2e.test "$@"
