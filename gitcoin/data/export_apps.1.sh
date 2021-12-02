#!/usr/bin/env bash

chifra export --appearances --fmt csv $1 | cut -f2,3 -d',' >apps/$1.csv
