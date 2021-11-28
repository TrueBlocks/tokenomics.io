#!/usr/bin/env bash

cd pouch

make

bin/pouch --summarize --bucketSize    500 | tee ../charts/counts_000500.txt
bin/pouch --summarize --bucketSize   1000 | tee ../charts/counts_001000.txt
bin/pouch --summarize --bucketSize   2500 | tee ../charts/counts_002500.txt
bin/pouch --summarize --bucketSize   5000 | tee ../charts/counts_005000.txt
bin/pouch --summarize --bucketSize  10000 | tee ../charts/counts_010000.txt
bin/pouch --summarize --bucketSize  25000 | tee ../charts/counts_025000.txt
bin/pouch --summarize --bucketSize  50000 | tee ../charts/counts_050000.txt
bin/pouch --summarize --bucketSize 100000 | tee ../charts/counts_100000.txt

cd -
