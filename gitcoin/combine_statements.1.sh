#!/usr/bin/env bash

# remove the header, add in a column for the address, append to combined file
cat exports/mainnet/statements/$1.csv | \
    grep -v -i "^\"asset" | \
    sed 's/^/\"'$1'\",/' | \
    tee -a exports/mainnet/combined/statements.csv

cat exports/mainnet/statements/balances/$1.csv | \
    grep -v -i "^\"asset" | \
    sed 's/^/\"'$1'\",/' | \
    tee -a exports/mainnet/combined/statements_balances.csv

cat exports/mainnet/statements/tx_counts/$1.csv | \
    grep -v -i "^count" | \
    sed 's/^/\"'$1'\",/' | \
    tee -a exports/mainnet/combined/statements_tx_counts.csv

