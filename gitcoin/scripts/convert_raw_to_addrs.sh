#!/usr/bin/env bash

# ./scripts/download_raw.sh

cat gitcoin/exports/mainnet/raw/*.json | \
    sed 's/|/ /g' | \
    jq '.[] | "\(.admin_address)|\(.id)|\(.title)|\(.slug)|\(.active)"' | \
    sed 's/"//g' | tr '|' '\t'

#find exports/raw -name "*.json" -exec grep -His admin_address {} ';' | \
#    grep 0x | \
#    tr '/' '\t' | \
#    tr '.' '\t' | \
#    tr '"' '\t' | \
#    grep -v "0x0\t" | \
#    grep -v 0x0000000000000000000000000000000000000000 | \
#    tr [:upper:] [:lower:] | \
#    cut -f3,7 | \
#    tr '\t' ',' | \
#    awk -F, '{print $2,$1}' OFS=, | \
#    sort | \
#    tee addresses.csv
