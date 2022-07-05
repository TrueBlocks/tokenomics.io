#!/bin/bash
set -e

if [ -z "$WEBSITES" ]
then
    echo "This script requires WEBSITES env variable to be set as a comma-separated list"
    exit 1
fi

update_per_file_data() {
    FOLDER=$1
    CHAIN=$2

    EXPORTS_DIR="/exports/$FOLDER/$CHAIN"

    # Turn filenames into addresses
    find $EXPORTS_DIR/statements -type f -name *.csv | xargs basename -a -s .csv > /tmp/addresses

    cd ${EXPORTS_DIR}/neighbors/networks
    cat /tmp/addresses | parallel --jobs 200% "python /app/neighbor_networks.py"
}

update_project() {
    FOLDER=$1
    CHAINS=${2:-mainnet}
    FMT=${3:-csv}

    for CHAIN in ${CHAINS//,/ }
    do
        echo "Folder: " $FOLDER
        echo "Chain:  " $CHAIN

        update_per_file_data $FOLDER $CHAIN
    done
}

start() {
    for WEBSITE in ${WEBSITES//,/ }
    do
        echo "Building neighbor network for ${WEBSITE}"
        WEBSITE_UP=`echo $WEBSITE | tr '[:lower:]' '[:upper:]'`
        varname="NOMICS_${WEBSITE_UP}_CHAINS"
        CHAINS="${!varname}"

        update_project $WEBSITE $CHAINS
    done
}

while true
do
    start
    echo "Images built, sleeping for 15 minutes"
    sleep 900
done