#!/usr/bin/env bash

echo "-------------------- $1 -----------------------"
echo -n "Total: "
cat pouch/data/$1.csv | grep DonationSent | wc | cut -c1-10
#echo -n "Apps: "
#cat pouch/data/$1.csv | grep DonationSent | grep $1 | wc | cut -c1-10
echo
echo -n "Added: "
cat pouch/data/0xf2354570be2fb420832fb7ff6ff0ae0df80cf2c6.csv | grep PayoutAdded | grep $1 | cut -d, -f1,2,3,11-100
echo
echo -n "Claimed: "
cat pouch/data/0xf2354570be2fb420832fb7ff6ff0ae0df80cf2c6.csv | grep PayoutClaimed | grep $1 | cut -d, -f1,2,3,11-100
val=`echo $1 | sed 's/0x/000000000000000000000000/'`
#echo $val
echo
echo -n "Available: "
chifra state --call 0xf2354570be2fb420832fb7ff6ff0ae0df80cf2c6!0x65bcfbe7!$val | tr '\n' ' ' | cut -d'"' -f36-37 | tr '\"' ' ' | cut -f1 -d'}'
echo
sleep .3
