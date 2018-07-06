#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

mkdir -p $DIR/../bin

echo "******************"
echo "Formatting $(realpath $DIR/../swe)"
cd $DIR/../swe
go fmt
echo "Done formatting."
echo "******************"
echo "Building AAR for go-swe and go-gpkg"
cd $DIR/../bin
gomobile bind -target android -javapkg=com.sandhillgeo github.com/sandhillgeo/go-swe/swe github.com/sandhillgeo/go-gpkg/gpkg
if [[ "$?" != 0 ]] ; then
    echo "Error building Android Archive for go-swe and go-gpkg"
    exit 1
fi
echo "Executable built at $(realpath $DIR/../bin)"
