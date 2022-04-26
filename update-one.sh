#!/usr/bin/env bash

FOLDER=`echo $1 | sed 's/\///g'`
if [ -z "$FOLDER" ]
then
    echo "FOLDER required"
fi

CHAIN=$2
if [ -z "$CHAIN" ]
then
    CHAIN="mainnet"
fi

FMT=$3
if [ -z "$FMT" ]
then
    FMT="csv"
fi

echo "Folder: " $FOLDER
echo "Chain:  " $CHAIN
echo "Format: " $FMT

echo "Running ./update-$folder.sh at " `date`

cd $FOLDER
pwd
RUN_ONCE=true chifra scrape monitors --chain $CHAIN --file commands.fil --fmt $FMT
cd - 2>&1 >/dev/null
./nomics combine --folder $FOLDER --chain $CHAIN --fmt $FMT
./nomics compress --folder $FOLDER --chain $CHAIN --fmt $FMT
./nomics update --folder $FOLDER --chain $CHAIN --fmt $FMT | jq >$FOLDER/ui/src/theData.json

#for dir in ${SUBDIRS//,/ }
#do
#    echo "Building $dir"
#    cd $dir/ui
#    yarn
#    yarn build
#    echo "Copying over the artifacts..."
#
#    if [ ! -d $TARGET/$dir ]
#    then
#        mkdir -p $TARGET/$dir/charts/
#    fi
#
#    cp -pr build/* $TARGET/$dir/
#    cd -
#    cp -pr charts/* $TARGET/$dir/charts/
#done
#
#echo "Copying main index.html"
#cp index.html $TARGET
#
#echo "Done..."
