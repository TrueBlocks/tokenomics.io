#!/usr/bin/env bash

rm -f zips/*

tar -cvf zips/apps.tar apps
tar -cvf zips/logs.tar logs
tar -cvf zips/neighbors.tar neighbors
tar -cvf zips/raw.tar raw
tar -cvf zips/txs.tar txs

cd zips
gzip -v *
ls -l

cd ..
./update_zips.sh
