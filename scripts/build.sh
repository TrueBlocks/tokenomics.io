#!/usr/bin/env bash

# exit immediately if a simple command exits with a non-zero status
set -e

# syntax: build.sh input,directories,comma,separated target parent
SUBDIRS=$1
TARGET=$2

if [ -z "$SUBDIRS" ]
then
    echo "SUBDIRS required"
    exit 1
fi

if [ -z "$TARGET" ]
then
    echo "TARGET required"
    exit 1
fi

echo "Building tokenomics.io into $TARGET..."

for dir in ${SUBDIRS//,/ }
do
    echo "--------------------------------------------"
    echo "Building $dir"
    echo "--------------------------------------------"
    cd $dir/ui
    yarn
    yarn build
    echo "Copying over the artifacts..."

    if [ ! -d $TARGET/$dir ]
    then
        mkdir -p $TARGET/$dir/charts/
    fi

    cp -pr build/* $TARGET/$dir/
    cd -
    cp -pr charts/* $TARGET/$dir/charts/
done

echo "Copying main index.html"
cp index.html $TARGET

echo "Done..."
