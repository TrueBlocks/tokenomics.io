#!/usr/bin/env bash

echo "exporting $1"
chifra list --fmt csv $1 >apps/$1.csv
