#!/usr/bin/env bash

#echo "Turned off"
#exit

git push
ssh jrush@unchainedindex.io "cd Development/tokenomics.io/gitcoin ; git pull ; ./scripts/deploy_ui.sh" &
#ssh jrush@unchainedindex.io "cd Development/tokenomics.io/gitcoin ; git pull ; ls -l"
