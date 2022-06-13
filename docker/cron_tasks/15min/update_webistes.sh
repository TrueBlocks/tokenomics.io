#!/bin/bash

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

bash $NOMICS_DIR/scripts/update.sh
bash $NOMICS_DIR/scripts/build.sh $WEBSITES $HTML_DIR