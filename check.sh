#!/bin/bash

set -o errexit -o nounset

# check format

# add license header
find . -name \*.go -exec sh -c "if ! grep -q 'LICENSE-2.0' '{}';then echo no license header in '{}', run addlicense.sh to add;exit 1;fi" \;