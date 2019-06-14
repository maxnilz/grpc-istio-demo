#!/bin/bash
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

# This script rebuilds the generated code for the protocol buffers.
# To run this you will need protoc and goprotobuf installed;
# see https://github.com/golang/protobuf for instructions.
# You also need Go and Git installed.

set -e

BIN_DIR=${PWD}/.bin
GOOGLEAPIS_REPO=https://github.com/googleapis/googleapis

export GOBIN=${BIN_DIR}
export GO111MODULE=on

function die() {
  echo 1>&2 $*
  exit 1
}

# Check/Install protoc & related command line tools
function prepare_protoc(){
    dist=${BIN_DIR}/protoc
    # if the command exist. don`t fetch
    if [[ -f ${dist} ]] ; then
        return
    fi

    # Get protoc
    VERSION="3.7.1"

    # Use the go tool to determine OS.
    OS=$( go env GOOS )
    if [[ "$OS" = "darwin" ]]; then
      OS="osx"
    fi

    mkdir -p ${BIN_DIR}
    ZIP="protoc-${VERSION}-${OS}-x86_64.zip"
    URL="https://github.com/google/protobuf/releases/download/v${VERSION}/${ZIP}"

    wget ${URL} -O ${ZIP}

    # Unpack the protoc.
    unzip ${ZIP} -d /tmp/protoc
    mkdir -p third_party
    mv /tmp/protoc/include/* third_party/
    mv /tmp/protoc/bin/protoc ${BIN_DIR}/
    chmod +x ${BIN_DIR}/protoc
    rm -rf /tmp/protoc
    rm ${ZIP}

    # Install golang proto plugin
	go install -v github.com/golang/protobuf/protoc-gen-go
    go install -v github.com/gogo/protobuf/protoc-gen-gogoslick
	go install -v github.com/golang/mock/mockgen

	# Install protoc-gen-grpc-web
    VERSION="1.0.3"
	BINARY="protoc-gen-grpc-web-${VERSION}-${OS}-x86_64"
    URL="https://github.com/grpc/grpc-web/releases/download/${VERSION}/${BINARY}"
    wget ${URL}
    mv ${BINARY} ${BIN_DIR}/protoc-gen-grpc-web
    chmod +x ${BIN_DIR}/protoc-gen-grpc-web
}

# Sanity check that the right tools are accessible.
for tool in go git unzip; do
  q=$(which ${tool}) || die "didn't find ${tool}"
  echo 1>&2 "$tool: $q"
done

# Check/Install protoc & related command line tools
prepare_protoc

remove_dirs=
trap 'rm -rf ${remove_dirs}' EXIT

if [[ ! -d "third_party/googleapis" ]]; then
  apidir=$(mktemp -d -t regen-cds-api.XXXXXX)
  git clone ${GOOGLEAPIS_REPO} ${apidir}
  remove_dirs="$remove_dirs"
  mkdir -p third_party
  cp -rf ${apidir} third_party/googleapis
  rm -rf third_party/googleapis/.git
fi

wait

echo 1>&2 "All done!"

