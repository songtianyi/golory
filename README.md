## golory

[![Build Status](https://travis-ci.org/1pb-club/golory.svg?branch=master)](https://travis-ci.org/1pb-club/golory)
[![Go Report Card](https://goreportcard.com/badge/github.com/1pb-club/golory)](https://goreportcard.com/report/github.com/1pb-club/golory)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

[docs](http://golory.askuy.com)

### Contribution guide
* use gofmt to format code before you issuing a pr
* use vendor to manage third-party dependencies
* support go version 1.9+
* add license header in your created file
* steps to make a pull request
  ```bash
  # get golory source code
  go get github.com/1pb-club/golory
  # fork repo
  https://github.com/songtianyi/golory
  # set 1pb-club/golory as upstream for your fork
  git remote add upstream https://github.com/1pb-club/golory.git
  # sync upstream code
  cd $FORKED_REPO_PATH
  git pull upstream master
  git push origin master
  # set your fork as a remote
  cd $GOPATH/src/github.com/1pb-club/golory
  git remote -v
  git remote add fork https://github.com/songtianyi/golory.git
  git remote -v
  # push local commits to your forked repo
  git push fork master
  # at last issue a pr
  ```


