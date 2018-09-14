#!/bin/bash

# Change Directory Bin
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
echo "[Ongkir App] Changing directory $DIR"

echo "[Ongkir App] Checking Bin"
appBin="$DIR/bin"
if [ -f "$appBin/app" ]
then
    eval "$appBin/app"
else
    echo "[Ongkir App] Trying to Install Ongkir Application"
    eval "./install.sh"
fi