#!/bin/bash -e
#
# Copyright 2016 Google Inc.
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
#
# This script rebuilds the generated code for the protocol buffers.
# To run this you will need protoc and goprotobuf installed;
# see https://github.com/golang/protobuf for instructions.
# You also need Go and Git installed.

PKG=google.golang.org/genproto
PROTO_REPO=https://github.com/google/protobuf
PROTO_SUBDIR=src/google/protobuf
API_REPO=https://github.com/googleapis/googleapis

function die() {
  echo 1>&2 $*
  exit 1
}

# Sanity check that the right tools are accessible.
for tool in go git protoc protoc-gen-go; do
  q=$(which $tool) || die "didn't find $tool"
  echo 1>&2 "$tool: $q"
done

tmpdir=$(mktemp -d -t regen-cbt.XXXXXX)
trap 'rm -rf $tmpdir' EXIT
tmpapi=$(mktemp -d -t regen-cds.XXXXXX)
trap 'rm -rf $tmpapi' EXIT

echo -n 1>&2 "finding package dir... "
pkgdir=$(go list -f '{{.Dir}}' $PKG/protobuf)
echo 1>&2 $pkgdir
base=$(echo $pkgdir | sed "s,/$PKG/protobuf\$,,")
echo 1>&2 "base: $base"
cd $base

echo 1>&2 "fetching proto repos..."
git clone -q $PROTO_REPO $tmpdir &
git clone -q $API_REPO $tmpapi &
wait

declare -A filename_map

# Pass 1: copy protos from the google/protobuf repo.
for f in $(cd $PKG && find protobuf -name '*.proto'); do
  echo 1>&2 "finding latest version of $f... "
  up=google/protobuf/$(basename $f)
  cp $tmpdir/src/$up $PKG/$f
  filename_map[$up]=$f
done

# Pass 2: move the protos out of googleapis/google/api.
for f in $(cd $PKG && find googleapis/api -name '*.proto'); do
  echo 1>&2 "finding latest version of $f... "
  # Note: we use move here so that the next pass doesn't see them.
  up=google/api/$(basename $f)
  mv $tmpapi/$up $PKG/$f
  filename_map[$up]=$f
done

# Pass 3: copy the rest of googleapis/google
for f in $(cd "$tmpapi/google" && find * -name '*.proto'); do
  dst=$(dirname "$PKG/googleapis/$f")
  echo 1>&2 "finding latest version of $f... "
  mkdir -p $dst
  cp "$tmpapi/google/$f" $dst
  filename_map[google/$f]=googleapis/$f
done

# Mappings of well-known proto types.
declare -A known_types
known_types[google/protobuf/any.proto]=github.com/golang/protobuf/ptypes/any/any.proto
known_types[google/protobuf/duration.proto]=github.com/golang/protobuf/ptypes/duration/duration.proto
known_types[google/protobuf/empty.proto]=github.com/golang/protobuf/ptypes/empty/empty.proto
known_types[google/protobuf/struct.proto]=github.com/golang/protobuf/ptypes/struct/struct.proto
known_types[google/protobuf/timestamp.proto]=github.com/golang/protobuf/ptypes/timestamp/timestamp.proto
known_types[google/protobuf/wrappers.proto]=github.com/golang/protobuf/ptypes/wrappers/wrappers.proto

# Pass 4: fix the imports in each of the protos.
import_fixes=$tmpdir/fix_imports.sed
for up in "${!filename_map[@]}"; do
  dst=$PKG/${filename_map[$up]}
  echo "Renaming $up => $dst"
  echo >>$import_fixes "s,\"$up\";,\"$dst\"; // from $up,"
done | sort 1>&2

for up in "${!known_types[@]}"; do
  dst=${known_types[$up]}
  echo "Renaming $up => $dst"
  echo >>$import_fixes "s,\"$up\";,\"$dst\"; // from $up,"
done | sort 1>&2

sed -i -f $import_fixes $(find $PKG -name '*.proto')

types_map=""
for f in "${!known_types[@]}"; do
  pkg=${known_types[$f]}
  types_map="$types_map,M$f=$pkg"
done

# Run protoc once per package.
for dir in $(find $PKG -name '*.proto' | xargs dirname | sort -u); do
  echo 1>&2 "* $dir"
  protoc --go_out=plugins=grpc:. $dir/*.proto
done

# Sanity check the build.
echo 1>&2 "Checking that the libraries build..."
go build -v $PKG/...

echo 1>&2 "All done!"
