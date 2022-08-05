#!/bin/bash

export NOMICS_DIR=/root/tokenomics.io

echo "Updating websites"

# bash /etc/periodic/15min/update_webistes.sh

# if [ $? -gt 0 ]
# then
#     echo "Error while updating websites, exiting"
#     exit 1
# fi

# echo "Done. Cron will be updating the site every 15 minutes"
# # This runs forever
# tail -f /dev/null

if [ -z "$WEBSITES" ]
then
    echo "This script requires WEBSITES env variable to be set as a comma-separated list"
    exit 1
fi

if [ -z "$HTML_DIR" ]
then
    echo "This script requires HTML_DIR env variable to be set"
    exit 1
fi

if [ -z "$NOMICS_DIR" ]
then
    echo "This script requires NOMICS_DIR env variable to be set"
    exit 1
fi

start() {
    bash $NOMICS_DIR/scripts/update.sh
    STATUS=$?

    if [ $STATUS -gt 0 ]
    then
        echo "Error while updating the data"
        exit 1
    fi
}

build_website() {
    # If the website has already been built, we don't need to do anything
    if [ -f /html/index.html ]
    then
        return
    fi

    echo "Building website"

    cd $NOMICS_DIR/ui
    yarn
    yarn build

    # Do not copy development data mocks
    echo "Copying over the artifacts..."
    cp -v `find build/ -type f` /html
    cp -rv build/static /html
    cd -

    echo "Website built"
}

build_website

while true
do
    start
    sleep $NOMICS_SLEEP
done
