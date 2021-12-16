#!/usr/bin/env bash

# convert to lower case
addr=`echo $1 | tr [:upper:] [:lower:]`

# generate the data
#chifra export --statements --fmt csv $addr >statements/$addr.csv
echo -n $addr ","
cat statements/$addr.csv | grep -v assetSymbol | cut -f2 -d, | sort | uniq -c | sort -n -r | tr '\n' '|' | sed 's/ //g' | sed 's/\"|/|/g' | tr '"' ';'
echo
