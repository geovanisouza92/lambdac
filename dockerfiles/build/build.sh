#!/bin/bash

source /environment.sh

# Compile statically linked version of package
echo "Building $pkgName"
`CGO_ENABLED=${CGO_ENABLED:-0} go build -a --installsuffix cgo --ldflags="${LDFLAGS:--s}" $pkgName`

# Grab the last segment from the package name
# echo 'Getting binary name'
name=${pkgName##*/}

if [[ $COMPRESS_BINARY == "true" ]]; then
    # echo 'Compressing binary'
    goupx $name
fi

if [ -e "/var/run/docker.sock" ] && [ -e "./Dockerfile" ]; then
    # echo 'Building Docker image'
    # Default TAG_NAME to package name if not set explictly
    tagName=${tagName:-"$name":latest}

    # Build the image from the Dockerfile in the package directory
    docker build -t $tagName .
fi
