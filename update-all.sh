#!/usr/bin/env bash

PATH=$PATH:/home/jrush/Development/trueblocks-core/bin/
cd /home/jrush/Development/tokenomics.io/

./update-one.sh giveth  "gnosis --chain mainnet" csv
./update-one.sh gitcoin "mainnet"                csv
chifra when latest --no_header | cut -f1,3 | tr '\t' ' ' | sed 's/^/export const lastUpdate = \"Last updated at block: /' | sed 's/$/\";/' >giveth/ui/src/last-update.js
chifra when latest --no_header | cut -f1,3 | tr '\t' ' ' | sed 's/^/export const lastUpdate = \"Last updated at block: /' | sed 's/$/\";/' >gitcoin/ui/src/last-update.js
scripts/build.sh giveth,gitcoin ~/Websites/tokenomics.io/
