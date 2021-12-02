#!/usr/bin/env bash

tar -cvf zips/$1.tar apps/$1.csv
tar -rvf zips/$1.tar txs/$1.csv
tar -rvf zips/$1.tar logs/$1.csv
tar -rvf zips/$1.tar neighbors/$1.csv
#tar -rvf zips/$1.tar bals/$1.csv
cd zips
yes | gzip $1.tar
