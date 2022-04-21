#!/usr/bin/env bash

# ./scripts/download_raw.sh

cat gitcoin/exports/mainnet/raw/*.json | \
    sed 's/|/ /g' | \
    jq '.[] | "\(.admin_address)|\(.id)|\(.title)|\(.slug)|\(.active)"' | \
    sed 's/"//g' | tr '|' '\t'
