#!/usr/bin/env bash

# switch the makefile to mac

cd pouch

curl https://gitcoin.co/grants/grants.json | jq >grants_from_gitcoin.json

# get the payout data from the smart contract
cat data/0xf2354570be2fb420832fb7ff6ff0ae0df80cf2c6.csv | cut -f1,2,3,11-20 -d, | grep Payout >app-data/payouts.csv

# build counts
#cat data/0x7d655c57f71464b6f83811c55d84009cd9f5221c.csv \
#    data/0xdf869fad6db91f437b59f1edefab319493d4c4ce.csv | \
#    cut -f1 -d, | sed 's/\"//g' | sed 's/^9/09/' | \
#    sed 's/^8/08/' | \
#    sed 's/^7/07/' | cut -c1-5 | sed 's/$/000/' | \
#    sort -n | uniq -c | grep -v blocks >../charts/counts.txt

make clean
makeClass -aorv
make -j 12

# update the json files
bin/pouch --csv2json
if [ $? -ne 0 ]; then
    cd -
    exit 1
fi

## update the website data
rm -f data/records.bin
bin/pouch --freshen

# return to normal
cd - >/dev/null 2>&1
