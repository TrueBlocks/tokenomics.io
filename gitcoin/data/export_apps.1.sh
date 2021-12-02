#!/usr/bin/env bash

echo "exporting $1"
chifra export --appearances $1 | cut -f2,3 >apps/$1.csv
