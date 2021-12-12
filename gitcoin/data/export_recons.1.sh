#!/usr/bin/env bash

# convert to lower case
addr=`echo $1 | tr [:upper:] [:lower:]`

# generate the data
chifra export --statements --fmt csv $addr >recons/$addr.csv
