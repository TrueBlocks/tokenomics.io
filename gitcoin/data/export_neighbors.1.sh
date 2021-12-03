#!/usr/bin/env bash

addr=`echo $1 | tr [:upper:] [:lower:]`
chifra export --neighbors --cache --cache_traces --fmt csv $addr | \
    sed 's/'$addr'/-----------------self---------------------/' >neighbors/$addr.csv
