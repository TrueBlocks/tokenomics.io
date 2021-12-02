#!/usr/bin/env bash

chifra export --neighbors --cache --cache_traces --fmt csv $1 | \
    sed 's/'$1'/-----------------self---------------------/' >neighbors/$1.csv
