#!/usr/bin/env bash

# remove the header, add in a column for the address, append to combined file
cat statements/$1.csv | \
    grep -v -i "^\"asset" | \
    sed 's/^/\"'$1'\",/' | \
    tee -a combined/statements.csv

cat statements/balances/$1.csv | \
    grep -v -i "^\"asset" | \
    sed 's/^/\"'$1'\",/' | \
    tee -a combined/statements_balances.csv

cat statements/tx_counts/$1.csv | \
    grep -v -i "^count" | \
    sed 's/^/\"'$1'\",/' | \
    tee -a combined/statements_tx_counts.csv

