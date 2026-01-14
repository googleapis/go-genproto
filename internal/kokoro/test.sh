#!/bin/bash

# Fail on any error
set -eo pipefail

# Display commands being run
set -x

# cd to project dir on Kokoro instance
cd github/go-genproto
git config --global --add safe.directory "$(pwd)/./.git"

go version

# Set $GOPATH
export GOPATH="$HOME/go"
export GENPROTO_HOME=$GOPATH/src/google.golang.org/genproto
export PATH="$GOPATH/bin:$PATH"
export GO111MODULE=on
mkdir -p $GENPROTO_HOME

# Move code into $GOPATH and get dependencies
git clone . $GENPROTO_HOME
cd $GENPROTO_HOME

try3() { eval "$*" || eval "$*" || eval "$*"; }

# All packages, including +build tools, are fetched.
try3 go mod download
./internal/kokoro/vet.sh

go get github.com/jstemmer/go-junit-report

set +e

exit_code = 0
go_test_args=("-race" "-timeout" "10m")

gotestsum --packages="./..." \
    --junitfile sponge_log.xml \
    --format standard-verbose \
    -- "${go_test_args[@]}" 2>&1 | tee sponge_log.log
exit_code=$(($exit_code + $?))

# Send logs to Flaky Bot for continuous builds.
if [[ $KOKORO_BUILD_ARTIFACTS_SUBDIR = *"continuous"* ]]; then
  chmod +x $KOKORO_GFILE_DIR/linux_amd64/flakybot
  $KOKORO_GFILE_DIR/linux_amd64/flakybot \
    -logs_dir=$KOKORO_ARTIFACTS_DIR \
    -repo=googleapis/go-genproto
fi

exit $exit_code
