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
for GOOS in darwin linux windows; do
  GOOS=${GOOS} GOARCH=amd64 go build -o "swe_${GOOS}_amd64" github.com/sandhillgeo/go-swe/cmd/swe
done
if [[ "$?" != 0 ]] ; then
    echo "Error building program for go-swe"
    exit 1
fi
echo "Executable built at $(realpath $DIR/../bin/swe)"
