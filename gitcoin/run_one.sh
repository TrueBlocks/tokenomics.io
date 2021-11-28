#!/usr/bin/env bash

export DICT_MODE=true

chifra export --appearances $1 | cut -f2,3 >pouch/data/apps/$1.txt

chifra export \
    --logs --articulate --relevant --cache --cache_traces \
    --emitted_by 0xdf869fad6db91f437b59f1edefab319493d4c4ce \
    --emitted_by 0xf2354570be2fb420832fb7ff6ff0ae0df80cf2c6 \
    --emitted_by 0x7d655c57f71464b6f83811c55d84009cd9f5221c \
    --fmt csv $1 >pouch/data/$1.csv

#chifra export \
#    --logs --articulate --relevant --cache --cache_traces \
#    --emitted_by 0xdf869fad6db91f437b59f1edefab319493d4c4ce \
#    --emitted_by 0xf2354570be2fb420832fb7ff6ff0ae0df80cf2c6 \
#    --emitted_by 0x7d655c57f71464b6f83811c55d84009cd9f5221c \
#    --fmt json $1 | jq . >pouch/data/$1.json

sleep .3

#echo "-------------------- $1 -------------------"
#touch pouch/data/$1.csv
