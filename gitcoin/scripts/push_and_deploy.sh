#!/usr/bin/env bash

git push
ssh jrush@unchainedindex.io "cd Development/tokenomics.io/gitcoin ; git pull ; ./scripts/deploy_ui.sh" &
#ssh jrush@unchainedindex.io "cd Development/tokenomics.io/gitcoin ; git pull ; ls -l"
