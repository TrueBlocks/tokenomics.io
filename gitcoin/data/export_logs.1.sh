#!/usr/bin/env bash

# convert to lower case
addr=`echo $1 | tr [:upper:] [:lower:]`

export DISPLAY_LOG="\"[{blockNumber}]\",\"[{transactionIndex}]\",\"[{logIndex}]\",\"[{address}]\",\"[{topic0}]\",\"[{topic1}]\",\"[{topic2}]\",\"[{topic3}]\",\"[{data}]\",\"[{type}]\",\"[{timestamp}]\""

# generate the data
chifra export \
    --logs --relevant --articulate --cache --cache_traces \
    --emitter 0xdf869fad6db91f437b59f1edefab319493d4c4ce \
    --emitter 0xf2354570be2fb420832fb7ff6ff0ae0df80cf2c6 \
    --emitter 0x7d655c57f71464b6f83811c55d84009cd9f5221c \
    --fmt csv $addr >logs/$addr.csv

export DISPLAY_FORMAT=

# echo "Exporting logs for $addr"
chifra export --logs --articulate $addr --fmt json | jq .data[].compressedLog | grep -v null >logs/articulated/$addr.json
