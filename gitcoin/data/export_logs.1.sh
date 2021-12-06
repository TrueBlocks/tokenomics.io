#!/usr/bin/env bash

addr=`echo $1 | tr [:upper:] [:lower:]`
chifra export \
    --logs --articulate --relevant --cache --cache_traces \
    --emitter 0xdf869fad6db91f437b59f1edefab319493d4c4ce \
    --emitter 0xf2354570be2fb420832fb7ff6ff0ae0df80cf2c6 \
    --emitter 0x7d655c57f71464b6f83811c55d84009cd9f5221c \
    --fmt csv $addr >logs/$addr.csv
./fixHeaders $1
