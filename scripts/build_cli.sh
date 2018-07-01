#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

mkdir -p $DIR/../bin

echo "******************"
echo "Formatting $DIR/../cmd/swe"
cd $DIR/../cmd/swe
go fmt
echo "Done formatting."
echo "******************"
echo "Building program for go-swe"
cd $DIR/../bin
####################################################
#echo "Building program for darwin"
#GOTAGS= CGO_ENABLED=1 GOOS=${GOOS} GOARCH=amd64 go build --tags "darwin" -o "swe_darwin_amd64" github.com/sandhillgeo/go-swe/cmd/swe
#if [[ "$?" != 0 ]] ; then
#    echo "Error building program for go-swe"
#    exit 1
#fi
#echo "Executable built at $(realpath $DIR/../bin/swe_darwin_amd64)"
####################################################
echo "Building program for linux"
GOTAGS= CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build --tags "linux" -o "swe_linux_amd64" github.com/sandhillgeo/go-swe/cmd/swe
if [[ "$?" != 0 ]] ; then
    echo "Error building program for go-swe"
    exit 1
fi
echo "Executable built at $(realpath $DIR/../bin/swe_linux_amd64)"
####################################################
echo "Building program for Windows"
GOTAGS= CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build -o "swe_windows_amd64.exe" github.com/sandhillgeo/go-swe/cmd/swe
if [[ "$?" != 0 ]] ; then
    echo "Error building program for go-swe"
    exit 1
fi
echo "Executable built at $(realpath $DIR/../bin/swe_windows_amd64.exe)"
