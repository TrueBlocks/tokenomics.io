#!/usr/bin/env bash

echo "exporting $1"
chifra export --cache --cache_traces --articulate --fmt csv $1 >txs/$1.csv
