#!/usr/bin/env bash

# first, get the latest repo
git pull

# make the tool
cd pouch

#curl https://gitcoin.co/grants/grants.json | jq >grants_from_gitcoin.json

# get the payout data from the smart contract
#cat data/0xf2354570be2fb420832fb7ff6ff0ae0df80cf2c6.csv | cut -f1,2,3,11-20 -d, | grep Payout >app-data/payouts.csv

#make clean
makeClass -aorv
make -j 3

# build the json data from the csv
bin/pouch --csv2json
if [ $? -ne 0 ]; then
    cd -
    exit 1
fi

# freshen the data
rm -f pouch/data/records.bin
bin/pouch --freshen

# update some summary data
# cat data/0x7d655c57f71464b6f83811c55d84009cd9f5221c.csv \
#    data/0xdf869fad6db91f437b59f1edefab319493d4c4ce.csv | \
#    cut -f1 -d, | sed 's/\"//g' | sed 's/^9/09/' | \
#    sed 's/^8/08/' | \
#    sed 's/^7/07/' | cut -c1-5 | sed 's/$/000/' | \
#    sort -n | uniq -c | grep -v blocks >../charts/counts.txt

# return to top folder
cd - >/dev/null 2>&1

# publish the site
yarn build
yes | cp -pr build/* /home/jrush/Websites/tokenomics.io/gitcoin/
yes | cp -pR charts/* /home/jrush/Websites/tokenomics.io/gitcoin/charts/
