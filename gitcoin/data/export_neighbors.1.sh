#!/usr/bin/env bash

echo "exporting $1"
chifra export --neighbors --fmt csv $1 | sed 's/'$1'/-----------------self---------------------/' >neighbors/$1.csv
