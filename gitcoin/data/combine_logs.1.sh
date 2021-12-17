#!/usr/bin/env bash

# remove the header, add in a column for the address, append to combined file
cat logs/$1.csv | \
    grep -v -i "^\"block" | \
    sed 's/^/\"'$1'\",/' | \
    tee -a combined/logs.csv

cat logs/articulated/$.txt | \
    sed 's/^/\"'$1'\",/' | \
    tee -a combined/logs_articulated.csv

