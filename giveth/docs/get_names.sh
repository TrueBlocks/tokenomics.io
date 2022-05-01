#!/usr/bin/env bash

echo "address|grantId|name|tag|active|core" | \
    tr '|' '\t' >../addresses.tsv

cat ../exports/raw/*.json | \
    jq '.[] | .[] | "\(.walletAddress)|\(.id)|\(.title)|31-Contracts:Giveth|true|false"' | \
    sed 's/"//g' | \
    tr '|' '\t' | \
    sort --ignore-case | \
    grep "^0x" >>../addresses.tsv
