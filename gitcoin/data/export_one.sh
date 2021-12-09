#!/usr/bin/env bash

chifra names $1
./export_apps.1.sh $1
# ./export_bals.1.sh $1
./export_txs.1.sh $1
./export_logs.1.sh $1
./export_neighbors.1.sh $1

#./combine_update.sh
#./update_zips.sh
