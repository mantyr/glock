#!/bin/bash
#
# Executes `go test` on all go-kit packages, ignoring the vendor/ directory.
#
# Usage:
#    ./test.sh

go test -cover $@ $(go list ./... | grep -v vendor)
