#!/usr/bin/env bash

# ./scripts/download_raw.sh

cat ./exports/mainnet/raw/*.json | \
    sed 's/|/ /g' | \
    jq '.[] | "\(.admin_address)|\(.id)|\(.title)|31-GitCoin:Grants|\(.active)"' | \
    sed 's/"//g' | tr '|' '\t' | \
    sort --ignore-case | \
    grep "^0x"
