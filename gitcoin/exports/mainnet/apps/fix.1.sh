#!/usr/bin/env bash

echo "-------- fixing $1 --------------"
echo "\"blockNumber\",\"transactionIndex\"" >tmp
cat $1.csv | cut -d, -f2,3 | \
    grep -v blockNumber | \
    sed 's/\"//g' | \
    tr ',' '\t' | \
    sort -k 1 -n -k 2 -n | \
    uniq | \
    tr '\t' '|' | \
    sed 's/^/\"/' | \
    sed 's/$/\"/' | \
    sed 's/|/\",\"/' >>tmp
cat tmp > $1.csv
rm -f tmp
