#!/bin/bash
#
# Generates all compiled binaries for distribution, and publishes a new Docker image.
#
# Usage:
#   ./create-release.sh VERSION
#   ./create-release.sh 1.0.0

VERSION="$1"
ROOT_DIR="$GOPATH/src/github.com/KyleBanks/glock"

if [[ -z  $VERSION  ]]; then
    echo "You must specify a VERSION"
    exit 1
elif [[ $VERSION == v* ]]; then
    echo "VERSION should be in the format of 1.0.0 (no preceding 'v')."
    exit 1
fi

# Build the images
cd $ROOT_DIR
docker build -t kylebanks/glock:$VERSION .
docker build -t kylebanks/glock:latest .

# Publish them
docker push kylebanks/glock:$VERSION
docker push kylebanks/glock:latest

# Generate the git tag
git tag -a v$VERSION -m "Version v$VERSION"
git push origin --tags

# Generate the binaries
cd $ROOT_DIR/scripts
bash -x cross-compile.sh
