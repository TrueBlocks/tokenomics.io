#!/usr/bin/env bash

./download_raw.sh
find raw -name "*.json" -exec grep -His admin_address {} ';' | \
    tr '/' '\t' | \
    tr '.' '\t' | \
    tr '"' '\t' | \
    tr [:upper:] [:lower:] | \
    cut -f2,6 | \
    sort -u -k 2 | \
    tr '\t' ',' | \
    tee addresses.csv
