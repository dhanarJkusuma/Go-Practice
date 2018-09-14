#!/bin/bash

# Change Directory Bin
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
echo "[Perkalian App] Changing directory $DIR"

echo "[Perkalian App] Checking Bin"
echo ""
appBin="$DIR/bin"
if [ -f "$appBin/app" ]
then
    eval "$appBin/app"
else
    echo "[Perkalian App] Trying to Install Perkalian Application"
    eval "./install.sh"
fi