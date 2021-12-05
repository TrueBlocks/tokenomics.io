#!/usr/bin/env bash

addr=`echo $1 | tr [:upper:] [:lower:]`
chifra export --balances --fmt csv $addr >bals/$addr.csv
./fixHeaders $1
