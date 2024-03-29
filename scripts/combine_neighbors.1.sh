#!/usr/bin/env bash

# remove the header, add in a column for the address, append to combined file
cat neighbors/$1.csv | \
    grep -v -i "^\"block" | \
    grep -v -i "^\"bn\",\"tx" | \
    sed 's/^/\"'$1'\",/' | \
    tee -a combined/neighbors.csv
