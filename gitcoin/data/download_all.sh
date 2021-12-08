rm -fR apps/*
rm -fR bals/*
rm -fR combined/*
rm -fR logs/*
rm -fR neighbors/*
rm -fR txs/*
rm -fR zips/*

cd apps
scp -pr "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/apps/*" .
cd -

cd bals
scp -pr "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/bals/*" .
cd -

cd combined
scp -pr "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/combined/*" .
cd -

cd logs
scp -pr "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/logs/*" .
cd -

cd neighbors
scp -pr "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/neighbors/*" .
cd -

cd txs
scp -pr "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/txs/*" .
cd -

cd zips
scp -pr "unchainedindex.io:/home/jrush/Development/tokenomics.io/gitcoin/data/zips/*" .
cd -
