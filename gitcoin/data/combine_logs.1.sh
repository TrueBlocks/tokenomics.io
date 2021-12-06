#!/usr/bin/env bash

# remove the header, add in a column for the address, append to combined file
cat logs/$1.csv | \
    grep -v -i "^\"blocknumber\",\"transaction" | \
    sed 's/^/\"'$1'\",/' | \
    tee -a combined/logs.csv
