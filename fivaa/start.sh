#!/bin/bash

# Change Directory Bin
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
echo "[Fivaa App] Changing directory $DIR"

echo "[Fivaa App] Checking Bin"
echo ""
appBin="$DIR/bin"
if [ -f "$appBin/app" ]
then
    eval "$appBin/app"
else
    echo "[Fivaa App] Trying to Install Fivaa Application"
    eval "./install.sh"
fi