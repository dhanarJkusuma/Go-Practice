#!/bin/bash

# Change Directory Bin
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
echo "[Cart App] Changing directory $DIR"

echo "[Cart App] Checking Bin"
echo ""
appBin="$DIR/bin"
if [ -f "$appBin/app" ]
then
    eval "$appBin/app"
else
    echo "[Cart App] Trying to Install Cart Application"
    eval "./install.sh"
fi