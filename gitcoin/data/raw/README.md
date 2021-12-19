## Gitcoin Data Pouch - Data/Raw

This is the collection of data retreived from GitCoin's API and cached here for performance reasons and to avoid rate limiting. This data is not currently being freshened, although it is the source of addresses for all of the rest of the processing.

### Use of this data

We use this data in the `backend` folder to produce `./ui/src/grants-data.js` which is used to populate the database table in the UI. It is also used to generate the file `./data/addresses` with this command:

```
find raw -name "*.json" -exec grep -His admin_address {} ';' | \
	tr '/' '\t' | tr '.' '\t' | tr '"' '\t' | \
	tr [:upper:] [:lower:] | \
	cut -f2,6 | \
	sort -u -k 2 | \
	tr '\t' ',' | \
	tee addresses.csv
```

### Notes

- We remove duplicates above with the `sort -u` step.

- The GitCoin API data contains the a few invalid addresses which we remove by hand.

| grant_id | value                                      |
| -------- | ------------------------------------------ |
| 1911     | 0                                          |
| 2425     | 0x0                                        |
| 1921     | 0x0000000000000000000000000000000000000000 |
| 2019     | 0x01972                                    |
| 2237     | 37q2dixqjur5duhyhcirfntbyv4pvsvkhx         |
| 3070     | bc1qupk2u36zdm0fd8mmnu0ha33g0d2lgynwxw6j70 |

### To Do

Freshen this data and use it to automatically generate the processing scripts.

### Where's the Data?

The actual data produced for this project is too big to put on GitHub. You may download the entire dataset (or any portion) from the [Downloads Tab](https://tokenomics.io/gitcoin) of the website.