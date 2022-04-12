## Data/Statements

This folder is experimental and currently holds pre-alpha data. Do not use.

### How this Data is Created

The following script (found in [../export_statements.1.sh](../../export_statements.1.sh)) produces the files in this folder and its subfolders.

```
chifra export --statements --fmt csv $addr >statements/$addr.csv
cat statements/$addr.csv | cut -d, -f1,2,3,4,5,6,9,25,26,30-33 | tee statements/balances/$addr.csv
echo "count,assetAddr,assetSymbol" | tee statements/tx_counts/$addr.csv
cat statements/balances/$addr.csv | grep -v assetAddr | cut -d, -f1,2 | sort | uniq -c | sort -n -r | sed 's/ //g' | sed 's/"/,/g' | cut -d, -f1,2,5 | tee -a statements/tx_counts/$addr.csv
```

### Where's the Data?

The actual data produced for this project is too big to put on GitHub. You may download the entire dataset (or any portion) from the [Downloads Tab](https://tokenomics.io/gitcoin) of the website.