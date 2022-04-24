#!/usr/bin/env bash

# chifra when --timestamps
./combine_apps.sh
./combine_txs.sh
./combine_logs.sh
./combine_neighbors.sh
./combine_statements.sh

cd exports/mainnet/combined
rm -f *.gz
yes | gzip *.csv

