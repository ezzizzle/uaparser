#!/bin/bash

mkdir -p build

# Build the linux binary and Docker image
echo "Building for Mac at build/uaparserserver.mac"
GOOS=darwin GOARCH=amd64 go build -o build/uaparserserver.mac

echo "Building for Linux at build/uaparserserver.linux"
GOOS=linux GOARCH=amd64 go build -o build/uaparserserver.linux

echo "Building for Linux32 at build/uaparserserver.linux32"
GOOS=linux GOARCH=386 go build -o build/uaparserserver.linux32

echo "Building for Windows at build/uaparserserver.win"
GOOS=windows GOARCH=amd64 go build -o build/uaparserserver.win

echo "Building for Windows32 at build/uaparserserver.win32"
GOOS=windows GOARCH=386 go build -o build/uaparserserver.win32

echo "Building docker container 'uaparserserver'"
docker build -t uaparserserver .