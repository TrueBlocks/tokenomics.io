#!/bin/bash
set -e
set -o pipefail

if [ -z "$WEBSITES" ]
then
    echo "This script requires WEBSITES env variable to be set as a comma-separated list"
    exit 1
fi

CORE_HOST="${CORE_URL:-http://core}:${CORE_PORT:-8080}"
WHEN_RESPONSE=`curl "${CORE_HOST}/when?blocks=latest&fmt=txt&noHeader"`
CURL_EXIT=$?

if [ $CURL_EXIT -gt 0 ]
then
    echo "Error while calling chifra: $WHEN_RESPONSE"
    exit $CURL_EXIT
fi

WHEN=`echo $WHEN_RESPONSE | cut -d ' ' -f1,3 | tr '\t' ' ' | sed 's/^/[/' | sed 's/ /, "/' | sed 's/$/\"]/'`

update_statement_data() {
    EXPORTS_DIR=$1
    ADDRESS=$2
    FILE_NAME=$3
    FILE=$4

    mkdir -p ${EXPORTS_DIR}/zips/${ADDRESS}/statements

    cat $FILE | cut -d, -f1,2,3,4,5,6,9,25,26,30-33 > ${EXPORTS_DIR}/statements/balances/${FILE_NAME}
    echo "count,assetAddr,assetSymbol" > ${EXPORTS_DIR}/statements/tx_counts/${FILE_NAME}

    # Some datasets don't have assetAddr, so grep can fail
    set +e
    cat ${EXPORTS_DIR}/statements/balances/${FILE_NAME} | grep -v assetAddr | cut -d, -f1,2 | sort | uniq -c | sort -n -r | sed 's/ //g' | sed 's/"/,/g' | cut -d, -f1,2,5 | tee -a ${EXPORTS_DIR}/statements/tx_counts/${FILE_NAME}
    set -e
}

update_per_file_data() {
    FOLDER=$1
    CHAIN=$2

    EXPORTS_DIR="$NOMICS_DIR/$FOLDER/exports/$CHAIN"

    for FILE in `find $EXPORTS_DIR/statements -type f -name *.csv`
    do
        FILE_NAME=`echo $FILE | sed 's;.*/;;g'`
        ADDRESS=`echo $FILE_NAME | sed 's/.csv//'`

        update_statement_data $EXPORTS_DIR $ADDRESS $FILE_NAME $FILE
    done
}

update_project() {
    FOLDER=$1
    CHAINS=${2:-mainnet}
    FMT=${3:-csv}

    chainsArg=""

    for CHAIN in ${CHAINS//,/ }
    do
        mkdir -p $FOLDER/exports/$CHAIN/zips/combined
        mkdir -p $FOLDER/exports/$CHAIN/combined/statements/{balances,tx_counts}
        mkdir -p /html/$FOLDER/data/$CHAIN
        chainsArg="$chainsArg $CHAIN"
        update_per_file_data $FOLDER $CHAIN
    done

    echo "Folder: " $FOLDER
    echo "Chain:  " $chainArg
    echo "Format: " $FMT

    TEMP_FILE=/tmp/data-${RANDOM}.json

    nomics combine --folder $FOLDER --chain $chainArg --fmt $FMT
    nomics compress --folder $FOLDER --chain $chainArg --fmt $FMT
    nomics update --folder $FOLDER --chain $chainArg --fmt $FMT > $TEMP_FILE
    cat $TEMP_FILE | jq > $NOMICS_DIR/$FOLDER/theData.json

    echo "Copying data file"
    cp $NOMICS_DIR/$FOLDER/theData.json /html/$FOLDER/data/theData.json

    echo $WHEN > /html/$FOLDER/data/lastUpdate.json

    for CHAIN in ${CHAINS//,/ }
    do
        echo "Copying static data"
        cp -rv $NOMICS_DIR/$FOLDER/exports/$CHAIN /html/$FOLDER/data/
    done

        rm $TEMP_FILE
    done
}

for WEBSITE in ${WEBSITES//,/ }
do
    echo "For website ${WEBSITE}"
    WEBSITE_UP=`echo $WEBSITE | tr '[:lower:]' '[:upper:]'`
    varname="NOMICS_${WEBSITE_UP}_CHAINS"
    CHAINS="${!varname}"

    update_project $WEBSITE $CHAINS
done
