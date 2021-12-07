#!/usr/bin/env bash

# chifra when --timestamps
./export_apps.sh
# ./export_bals.sh
./export_txs.sh
./export_logs.sh
./export_neighbors.sh

./update_zips.sh
