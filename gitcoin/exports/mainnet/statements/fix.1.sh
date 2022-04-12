#!/usr/bin/env bash

echo "-------- fixing $1 --------------"
cat $1.csv | uniq >tmp
cat tmp > $1.csv
rm -f tmp
