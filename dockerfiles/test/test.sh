#!/bin/bash

source /environment.sh

cd $pkgPath

# Run tests (ignoring vendor/)
go test -v $(go list ./... | grep -v '/vendor/')
