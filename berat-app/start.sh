#!/bin/bash

# Change Directory Bin
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
echo "[Weight App] Changing directory $DIR"

echo "[Weight App] Checking Bin"
appBin="$DIR/bin"
if [ -f "$appBin/app" ]
then
    eval "$appBin/app"
else
    echo "[Weight App] Trying to Install Weight Application"
    eval "./install.sh"
fi