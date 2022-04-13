#!/usr/bin/env bash

# echo "Turned off"
# exit

# build the site and deploy it to the website

echo "Changing into the ui folder..."
cd gitcoin/ui

echo "Building..."
yarn build

echo "Copying over the artifacts..."
yes | cp -pr build/* /home/jrush/Websites/tokenomics.io/gitcoin/
cd - 2>/dev/null
yes | cp -pR charts/* /home/jrush/Websites/tokenomics.io/gitcoin/charts/

echo "Done..."

echo "Changing into the ui folder..."
cd giveth/ui

echo "Building..."
yarn build

echo "Copying over the artifacts..."
yes | cp -pr build/* /home/jrush/Websites/tokenomics.io/giveth/
cd - 2>/dev/null
yes | cp -pR charts/* /home/jrush/Websites/tokenomics.io/giveth/charts/

echo "Done..."
