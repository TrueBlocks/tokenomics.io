#!/usr/bin/env bash

# This script is called by download_raw to process each grant id pulling
# data from GitCoin's API endpoint. We have to sleep to avoid rage limiting.

curl -s "https://gitcoin.co/api/v0.1/grants/?pk=$1" | jq | tee raw/$1.json
echo "Fetched $1 --> Sleeping for four seconds"
sleep 4
