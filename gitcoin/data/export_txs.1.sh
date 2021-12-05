#!/usr/bin/env bash

addr=`echo $1 | tr [:upper:] [:lower:]`
chifra export --articulate --cache --cache_traces --fmt csv $addr >txs/$addr.csv
./fixHeaders $1
