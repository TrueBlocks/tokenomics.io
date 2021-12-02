#!/usr/bin/env bash

chifra export --articulate --cache --cache_traces --fmt csv $1 >txs/$1.csv
