#!/usr/bin/env bash

# chifra when --timestamps
./combine_apps.sh
#./combine_bals.sh
./combine_txs.sh
./combine_logs.sh
./combine_neighbors.sh

cd combined
gzip *.csv
