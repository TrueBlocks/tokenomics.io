#!/usr/bin/env bash

curl \
    --location \
    --request POST \
    --header 'Content-Type: application/json' \
    --data-raw \
        '{"query": "{projectByAddress(address: \"'$1'\" ) {title,listed,description,id,slug,slugHistory,walletAddress,slugHistory}}", "variables": {}}' \
        'https://mainnet.serve.giveth.io/graphql' | jq | tee ../exports/raw/$1.json
sleep 1
