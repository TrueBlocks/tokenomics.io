#!/usr/bin/env bash

# convert to lower case
addr=`echo $1 | tr [:upper:] [:lower:]`

# generate the data
chifra export --balances --fmt csv $addr >bals/$addr.csv
