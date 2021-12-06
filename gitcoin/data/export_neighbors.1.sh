#!/usr/bin/env bash

addr=`echo $1 | tr [:upper:] [:lower:]`
chifra list --count $addr
time chifra export --neighbors --cache --cache_traces --fmt csv $addr >neighbors/$addr.csv
./fixHeaders $addr
