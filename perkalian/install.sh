#!/bin/bash

# Change Directory Bin
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
echo "[Perkalian App] Changing directory $DIR"
appBin="$DIR/bin/app"
if [ -f "$appBin" ]
then
	rm $appBin
fi

# Downloading Dependency
echo "[Perkalian App] Resolving Dependency"
eval "go get -v"

echo "[Perkalian App] Build App"
# Build/Compile
buildCmd="go build -o $DIR/bin/app ."
eval $buildCmd > $DIR/build.log

eval "./start.sh"