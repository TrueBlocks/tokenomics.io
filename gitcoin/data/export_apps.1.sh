#!/usr/bin/env bash

echo "exporting $1"
chifra list --fmt csv --count $1 >apps/$1.csv
