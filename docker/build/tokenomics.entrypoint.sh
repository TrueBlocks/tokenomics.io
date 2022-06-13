#!/bin/bash

export NOMICS_DIR=/root/tokenomics.io

echo "Container started, updating websites"

bash /etc/periodic/15min/update_websites.sh

if [ $? -gt 0 ]
then
    echo "Error while updating websites, exiting"
    exit 1
fi

echo "Done. Cron will be updating the site every 15 minutes"
while :
do
    ;
done