#!/bin/bash

set -o errexit -o nounset

# check format
#TODO


# check license header
function checkLicense() {
    file=$1
	if ! grep -q 'LICENSE-2.0' $file; then echo no license header in $file, run addlicense.sh to add; exit 1;fi
}
export -f checkLicense
find . -name \*.go | xargs sh -c "checkLicense"
