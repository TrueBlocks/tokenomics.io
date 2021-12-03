#!/usr/bin/env bash

cat logs/$1.csv | grep -v -i "^\"blocknumber\",\"transaction" | sed 's/^/\"'$1'\",/' | tee -a combined/logs.csv
