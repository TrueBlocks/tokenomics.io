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

rm -fR recons/*
scp -pr "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/recons/*" recons/
