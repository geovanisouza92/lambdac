#!/bin/bash

set -eo pipefail

# echo 'Set $tagName'
tagName=$1

# echo 'Find src/'
if ( find /src -maxdepth 0 -empty | read v ); then
    echo 'Error: Must mount Go source code into /src directory'
    exit 990
fi

# Grab go package name
# echo 'Set $pkgName'
pkgName="$(go list -e -f '{{.ImportComment}}' 2>/dev/null || true)"
if [ -z "$pkgName" ]; then
    echo 'Error: Must add canonical import path to root package'
    exit 992
fi

# Grab just first path listed in GOPATH
# echo 'Set $goPath'
goPath="${GOPATH%%:*}"

# Grab just first path listed in GOPATH
# echo 'Set $pkgPath'
pkgPath="$goPath/src/$pkgName"

# Setup src directory in GOPATH
# echo 'Making package directory inside GOPATH'
mkdir -p "$(dirname "$pkgPath")"

# Link source dir into GOPATH
# echo 'Linking /src to GOPATH'
ln -sf /src "$pkgPath"

# echo 'Checking Go dependencies'
if [ $GO15VENDOREXPERIMENT ]; then
    echo 'Using Go 1.5 vendor experiment'
elif [ -e "$pkgPath/Godeps/_workspace" ]; then
    # Add local Godeps dir to GOPATH
    GOPATH=$pkgPath/Godeps/_workspace:$GOPATH
else
    # Get all package dependencies
    go get -t -d -v ./...
fi
