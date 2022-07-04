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

WHEN=`echo $WHEN_RESPONSE | cut -d ' ' -f1,3 | tr '\t' ' ' | sed 's/^/export const lastUpdate = \"Last updated at block: /' | sed 's/$/\";/'`

update_statement_data() {
    FOLDER=$1
    CHAIN=$2

    EXPORTS_DIR="$NOMICS_DIR/$FOLDER/exports/$CHAIN"

    for FILE in `find $EXPORTS_DIR/statements -type f -name *.csv`
    do
        FILE_NAME=`echo $FILE | sed 's;.*/;;g'`
        ADDRESS=`echo $FILE_NAME | sed 's/.csv//'`

        mkdir -p ${EXPORTS_DIR}/zips/${ADDRESS}/statements
        mkdir -p ${EXPORTS_DIR}/combined/statements

        cat $FILE | cut -d, -f1,2,3,4,5,6,9,25,26,30-33 > ${EXPORTS_DIR}/statements/balances/${FILE_NAME}
        echo "count,assetAddr,assetSymbol" > ${EXPORTS_DIR}/statements/tx_counts/${FILE_NAME}

        # Some datasets don't have assetAddr, so grep can fail
        set +e
        cat ${EXPORTS_DIR}/statements/balances/${FILE_NAME} | grep -v assetAddr | cut -d, -f1,2 | sort | uniq -c | sort -n -r | sed 's/ //g' | sed 's/"/,/g' | cut -d, -f1,2,5 | tee -a ${EXPORTS_DIR}/statements/tx_counts/${FILE_NAME}
        set -e
    done
}

update_project() {
    FOLDER=$1
    CHAINS=${2:-mainnet}
    FMT=${3:-csv}

    for CHAIN in ${CHAINS//,/ }
    do
        echo "Folder: " $FOLDER
        echo "Chain:  " $CHAIN
        echo "Format: " $FMT

        TEMP_FILE=/tmp/data-${RANDOM}.json

        update_statement_data $FOLDER $CHAIN

        nomics combine --folder $FOLDER --chain $CHAIN --fmt $FMT
        nomics compress --folder $FOLDER --chain $CHAIN --fmt $FMT
        nomics update --folder $FOLDER --chain $CHAIN --fmt $FMT > $TEMP_FILE
        cat $TEMP_FILE | jq > $NOMICS_DIR/$FOLDER/ui/src/theData.json
        echo $WHEN > $NOMICS_DIR/$FOLDER/ui/src/last-update.js
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
