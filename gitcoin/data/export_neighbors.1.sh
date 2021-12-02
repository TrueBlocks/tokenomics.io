#!/usr/bin/env bash

echo "exporting $1"
chifra export --neighbors --fmt csv $1 >neighbors/$1.csv
