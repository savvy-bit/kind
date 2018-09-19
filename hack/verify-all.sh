#!/usr/bin/env bash
# Copyright 2018 The Kubernetes Authors.
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

set -o errexit
set -o nounset
set -o pipefail

# cd to repo root
REPO_ROOT=$(git rev-parse --show-toplevel)
cd "${REPO_ROOT}"

# exit code, if a script fails we'll set this to 1
res=0

# run all verify scripts
echo "verifying govet ..."
hack/verify-govet.sh || res=1
cd "${REPO_ROOT}"

echo "verifying gofmt ..."
hack/verify-gofmt.sh || res=1
cd "${REPO_ROOT}"

echo "verifying golint ..."
hack/verify-golint.sh || res=1
cd "${REPO_ROOT}"

echo "verifying generated ..."
hack/verify-generated.sh || res=1
cd "${REPO_ROOT}"

echo "verifying deps ..."
hack/verify-deps.sh || res=1
cd "${REPO_ROOT}"

# exit based on verify scripts
if [[ "${res}" = 0 ]]; then
  echo ""
  echo "All verify checks passed, congrats!"
else
  echo ""
  echo "One or more verify checks failed! See output above..."
fi
exit "${res}"

