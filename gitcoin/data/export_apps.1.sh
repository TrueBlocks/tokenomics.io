#!/usr/bin/env bash

# convert to lower case
addr=`echo $1 | tr [:upper:] [:lower:]`

# generate the data
chifra export --appearances --fmt csv $addr | cut -f2,3 -d',' >apps/$addr.csv
