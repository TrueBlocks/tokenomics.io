#!/usr/bin/env bash

cat neighbors/$1.csv | grep -v -i "^\"blockNumber\",\"transaction" | grep -v -i "^\"bn\",\"tx" | sed 's/^/\"'$1'\",/' | tee -a combined/neighbors.csv
