#!/usr/bin/env bash

cat apps/$1.csv | grep -v -i "^\"blocknumber\",\"transaction" | sed 's/^/\"'$1'\",/' | tee -a combined/apps.csv
