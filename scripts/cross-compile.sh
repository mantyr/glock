#!/bin/bash
#
# Uses Docker to cross-compile glock for Mac/Windows/Linux with 32 & 64-bit variations of each.
#
# After compilation, the binaries can be found in the dist/ directory.
#
# Usage:
#    ./cross-compile.sh

cd $GOPATH/src/github.com/KyleBanks/glock

docker run --rm -it -v "$GOPATH":/go -w /go/src/github.com/KyleBanks/glock golang:1.6 bash
for GOOS in darwin linux windows; do
    for GOARCH in 386 amd64; do
        go build -v -o dist/glock-$GOOS-$GOARCH
    done
done