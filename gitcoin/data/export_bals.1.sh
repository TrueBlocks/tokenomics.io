#!/usr/bin/env bash

chifra export --balances --fmt csv $1 >bals/$1.csv
