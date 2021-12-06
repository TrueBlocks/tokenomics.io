#!/usr/bin/env bash

addr=`echo $1 | tr [:upper:] [:lower:]`
chifra export --appearances --fmt csv $addr | cut -f2,3 -d',' >apps/$addr.csv
./fixHeaders $1
