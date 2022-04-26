#!/usr/bin/env bash

# convert to lower case
addr=`echo $1 | tr [:upper:] [:lower:]`

tar -cvf zips/$addr.tar apps/$addr.csv
tar -rvf zips/$addr.tar txs/$addr.csv
tar -rvf zips/$addr.tar logs/$addr.csv
tar -rvf zips/$addr.tar neighbors/$addr.csv
tar -rvf zips/$addr.tar neighbors/adjacencies/$addr.txt
tar -rvf zips/$addr.tar statements/$addr.csv
tar -rvf zips/$addr.tar statements/balances/$addr.csv
tar -rvf zips/$addr.tar statements/tx_counts/$addr.csv

cd zips
yes | gzip $addr.tar
