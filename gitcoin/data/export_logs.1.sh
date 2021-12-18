#!/usr/bin/env bash

# convert to lower case
addr=`echo $1 | tr [:upper:] [:lower:]`

# donation contracts
donate_rnd_05=0xdf869fad6db91f437b59f1edefab319493d4c4ce
donate_rnd_06=0x7d655c57f71464b6f83811c55d84009cd9f5221c

# payout contracts
pay_rnd_08=0xf2354570be2fb420832fb7ff6ff0ae0df80cf2c6
pay_rnd_09=0x3342e3737732d879743f2682a3953a730ae4f47c	
pay_rnd_10=0x3ebaffe01513164e638480404c651e885cca0aa4

# Generate the logs for this address into a temp file
chifra export \
    --logs --relevant --articulate --cache --cache_traces \
    --emitter $donate_rnd_05 \
    --emitter $donate_rnd_06 \
    --emitter $pay_rnd_08 \
    --emitter $pay_rnd_09 \
    --emitter $pay_rnd_10 \
    --fmt csv $addr >/tmp/$addr.csv

# Separate out regular fields into the logs file
cat /tmp/$addr.csv | cut -f1-10 >logs/$addr.csv

# Sparate out the compressed log field into its own file
cat /tmp/$addr.csv | cut -f1-4,11 >logs/articulated/$addr.csv

# Cleanup
rm -f /tmp/$addr.csv
