#!/bin/bash
set -e
set -o pipefail

if [ -z "$WEBSITES" ]
then
    echo "This script requires WEBSITES env variable to be set as a comma-separated list"
    exit 1
fi

WHEN_RESPONSE=`curl "$CHIFRA_URL/when?blocks=latest&fmt=txt&noHeader"`
CURL_EXIT=$?

if [ $CURL_EXIT -gt 0 ]
then
    echo "Error while calling chifra: $WHEN_RESPONSE"
    exit $CURL_EXIT
fi

WHEN=`echo $WHEN_RESPONSE | cut -f1,3 | tr '\t' ' ' | sed 's/^/export const lastUpdate = \"Last updated at block: /' | sed 's/$/\";/'`

update_project() {
    FOLDER=$1
    CHAINS=${2:-mainnet}
    FMT=${3:-json}

    for CHAIN in ${CHAINS//,/ }
    do
        echo "Folder: " $FOLDER
        echo "Chain:  " $CHAIN
        echo "Format: " $FMT

        nomics combine --folder $FOLDER --chain $CHAIN --fmt $FMT
        nomics compress --folder $FOLDER --chain $CHAIN --fmt $FMT
        nomics update --folder $FOLDER --chain $CHAIN --fmt $FMT | jq > $NOMICS_DIR/$FOLDER/ui/src/theData.json
    done
}

for WEBSITE in ${WEBSITES//,/ }
do
    echo "For website ${WEBSITE}"
    WEBSITE_UP=`echo $WEBSITE | tr '[:lower:]' '[:upper:]'`
    varname="NOMICS_${WEBSITE_UP}_CHAINS"
    CHAINS="${!varname}"
    # if [ -z "$CHAINS" ]
    # then
    #     echo "Env variable $varname missing"
    #     exit 1
    # fi

    update_project $WEBSITE $CHAINS
done
