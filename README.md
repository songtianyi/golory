## golory

[![Build Status](https://travis-ci.org/1pb-club/golory.svg?branch=master)](https://travis-ci.org/1pb-club/golory)
[![Go Report Card](https://goreportcard.com/badge/github.com/1pb-club/golory)](https://goreportcard.com/report/github.com/1pb-club/golory)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

docs: http://golory.askuy.com

### Contribution guide
* use gofmt to format code before you issuing a pr
* use vendor to manage third-party dependencies
* support go version 1.9+
* steps to make a pull request
  ```
  # get golory source code
  go get github.com/1pb-club/golory
  # fork
  https://github.com/songtianyi/golory
  # set your fork as a remote
  cd $GOPATH/src/github.com/1pb-club/golory
  git remote -v
  git remote add fork https://github.com/songtianyi/golory.git
  git remote -v
  # push local commits to your forked repo
  git push fork master
  ```


