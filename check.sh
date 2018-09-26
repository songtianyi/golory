#!/bin/bash

set -o errexit -o nounset

# check format
#TODO

# check license header
find . -name \*.go | xargs -n 1 -P 10 -I {} sh -c 'file="$@"; if ! grep -q 'LICENSE-2.0' $file; then echo no license header in $file, run addlicense.sh to add; exit 1;fi' _ {}
