#!/usr/bin/env bash

rm -fR apps/*
scp -pr "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/apps/*" apps/

rm -fR bals/*
scp -pr "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/bals/*" bals/

rm -fR combined/*
scp -pr "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/combined/*" combined/

rm -fR logs/*
scp -pr "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/logs/*" logs/

rm -fR neighbors/*
scp -pr "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/neighbors/*" neighbors/

rm -fR txs/*
scp -pr "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/txs/*" txs/

rm -fR zips/*
scp -pr "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/zips/*" zips/

rm -fR statements/*
scp -pr "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/statements/*" statements/

rm -fR updateCmd/*
scp -pr "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/updateCmd/*" updateCmd/

rm -fR raw/*
scp -pr "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/raw/*" raw/

scp -p "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/*.md" .
scp -p "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/addresses.csv" .
scp -p "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/*.sh" .

#assets
#disputes
