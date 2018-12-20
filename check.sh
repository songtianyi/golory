#!/bin/bash

set -o errexit -o nounset

# check format

format_all_go_files(){
    gofmt -w .
}

all_unformat_files(){
    gofmt -l .
}

lint(){
	which golint > /dev/null
	if [ $? -ne 0 ]; then
		echo '[N] golint not installed.' >&2
		exit 1
	fi
    count=$(golint ./... | wc -l)
    if [ $count -eq 0 ];
        then
            echo "[Y] golint passed."
    else
        echo "[N] You have $count lint errors."
		golint ./...
    fi
}

info(){
    count=$(gofmt -l . | wc -l)
    if [ $count -eq 0 ];
        then
            echo "[Y] All go files have been formatted."
    else
        echo ">[N] You have $count go file haven't been formatted:"
        all_unformat_files
    fi
    echo "--------------------------------------"

    all_source_code=$(find . -name "*[.sh|.go]" -type f | xargs cat | wc -l)
    echo "Total source code lines:$all_source_code."

    # TODO:can't recognize /** */
    annotation=$(find . -name "*[.sh|.go]" -type f | xargs cat |grep -e ^\s*\/\/.*$ -e ^\# |wc -l)
    echo "Total annotating code lines:$annotation."

    echo ">Tip:You can use 'gofmt -w .' to format all go files."
}

lint
info


# check license header
find . -name \*.go | xargs -n 1 -P 10 -I {} sh -c 'file="$@"; if ! grep -q 'LICENSE-2.0' $file; then echo no license header in $file, run addlicense.sh to add; exit 1;fi' _ {}
