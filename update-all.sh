#!/usr/bin/env bash

# ./update-one.sh wallets "gnosis --chain mainnet" txt
./update-one.sh giveth  "gnosis --chain mainnet" csv
./update-one.sh gitcoin "mainnet"                csv
scripts/build.sh giveth,gitcoin ~/Websites/tokenomics.io/
