#!/bin/bash
#
# Cross-compile glock for Mac/Windows/Linux with 32 & 64-bit variations of each.
#
# After compilation, the binaries can be found in the dist/ directory.
#
# Usage:
#    ./cross-compile.sh

cd $GOPATH/src/github.com/KyleBanks/glock

for GOOS in darwin linux windows; do
    for GOARCH in 386 amd64; do
        go build -v -o dist/glock-$GOOS-$GOARCH
    done
done