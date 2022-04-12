#!/usr/bin/env bash

for file in *.csv
do
    addr=$(echo $file | sed "s/.csv//")
    ./fix.1.sh $addr
done
