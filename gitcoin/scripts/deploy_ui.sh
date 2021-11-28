#!/usr/bin/env bash

# build the site and deploy it to the website

cd ui
yarn build
yes | cp -pr build/* /home/jrush/Websites/tokenomics.io/gitcoin/
cd - 2>/dev/null
yes | cp -pR charts/* /home/jrush/Websites/tokenomics.io/gitcoin/charts/
