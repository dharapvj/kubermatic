#!/usr/bin/env bash

# Copyright 2022 The Kubermatic Kubernetes Platform contributors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail

cd $(dirname $0)/..
source hack/lib.sh

KUBERMATIC_EDITION="${KUBERMATIC_EDITION:-ce}"

export CGO_ENABLED=1

go_test unit_tests \
  -tags "unit,${KUBERMATIC_EDITION}" -timeout 30m -race -v ./pkg/... ./cmd/... ./codegen/...

cd sdk
go_test unit_tests \
  -tags "unit,${KUBERMATIC_EDITION}" -timeout 30m -race -v ./...
