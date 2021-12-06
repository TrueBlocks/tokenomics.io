#!/usr/bin/env bash

# convert to lower case
addr=`echo $1 | tr [:upper:] [:lower:]`

# generate the data
chifra export --articulate --cache --cache_traces --fmt csv $addr >txs/$addr.csv

# clean up headers
./fixHeaders $1
