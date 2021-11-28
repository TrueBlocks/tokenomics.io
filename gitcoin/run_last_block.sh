#!/usr/bin/env bash

export DICT_MODE=true

echo -n $1 "|" | tr '|' '\t'
tail -1 pouch/data/$1.csv | sed 's/\"//g' | cut -f1 -d, | sed 's/blocknumber/none/'
