#!/usr/bin/env bash

# convert to lower case
addr=`echo $1 | tr [:upper:] [:lower:]`

# generate the data
chifra list --count $addr
time chifra export --neighbors --cache --cache_traces --fmt csv $addr >neighbors/$addr.csv
