#/usr/bin/env bash

curl -s "https://gitcoin.co/api/v0.1/grants/?pk=$1" | jq | tee raw/$1.json
echo "Fetched $1 --> Sleeping for four seconds"
sleep 4
