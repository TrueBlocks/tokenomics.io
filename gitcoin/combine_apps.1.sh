#!/usr/bin/env bash

# remove the header, add in a column for the address, append to combined file
cat exports/mainnet/apps/$1.csv | \
    grep -v -i "^\"block" | \
    sed 's/^/\"'$1'\",/' | \
    tee -a exports/mainnet/combined/apps.csv
