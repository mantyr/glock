#!/bin/bash
#
# Executes `go test` on all go-kit packages, ignoring the vendor/ directory.
#
# Usage:
#    ./test.sh

cd $GOPATH/src/github.com/KyleBanks/glock

go test -cover $@ $(go list ./... | grep -v vendor)
