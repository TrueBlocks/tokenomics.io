#!/usr/bin/env bash

# chifra when --timestamps
./test_apps.sh
./test_txs.sh
./test_logs.sh
./test_neighbors.sh

./fixHeaders

# ./zipem.sh
