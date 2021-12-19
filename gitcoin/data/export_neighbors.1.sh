#!/usr/bin/env bash

# convert to lower case
addr=`echo $1 | tr [:upper:] [:lower:]`

# generate the data
chifra list --count $addr
chifra export --neighbors --deep --cache --cache_traces --fmt csv $addr >neighbors/$addr.csv
