## Gitcoin Data Pouch - Data/Logs

The `./logs` folder contains exports of all the logs for a given address. Note that this may be many multiples of the number of transactions / appearances for that address.

The data format for a log is:

| Name             | Description                                                        |
| ---------------- | ------------------------------------------------------------------ |
| blockNumber      | the transaction's block number                                     |
| transactionIndex | the transcation's index in the block                               |
| logIndex         | The index of the log in the block (not in the transaction)         |
| timestamp        | The timestamp of the block                                         |
| address          | The address of the smart contract that generated the event         |
| topic0           | The signature of the event represented by this event hashed        |
| topic1           | The first indexed parameter to the event (if any)                  |
| topic2           | The second indexed parameter to the event (if any)                 |
| topic3           | The third indexed parameter to the event (if any)                  |
| data             | The remainder of the parameters to the event (if any), RLP encoded |
| type             | Unknown                                                            |
| compressedLog    | The articulated event log presented in a single field              |

### How this Data is Created

Two different data sets are created from the log data. The first contains just the raw log data from the chain. The second, called `articulated` shows the 'meaning' of the log data by converting the raw data into human readable text. The `articulated` data is presented on the website and is stored in a folder called `./logs/articulated`. The articulated data is more useful.

The chifra command used to create this data is found in [../export_logs.1.sh](../export_logs.1.sh) and is called for each address.

```
donate_rnd_05=0xdf869fad6db91f437b59f1edefab319493d4c4ce
donate_rnd_06=0x7d655c57f71464b6f83811c55d84009cd9f5221c
pay_rnd_08=0xf2354570be2fb420832fb7ff6ff0ae0df80cf2c6
pay_rnd_09=0x3342e3737732d879743f2682a3953a730ae4f47c	
pay_rnd_10=0x3ebaffe01513164e638480404c651e885cca0aa4

chifra export \
    --logs --relevant --articulate --cache --cache_traces \
    --emitter $donate_rnd_05 \
    --emitter $donate_rnd_06 \
    --emitter $pay_rnd_08 \
    --emitter $pay_rnd_09 \
    --emitter $pay_rnd_10 \
    --fmt csv $addr >/tmp/$addr.csv
cat /tmp/$addr.csv | cut -f1-10 -d, >logs/$addr.csv
cat /tmp/$addr.csv | cut -f1-5,11-2000 -d, >logs/articulated/$addr.csv
```


### Where's the Data?

The actual data produced for this project is too big to put on GitHub. You may download the entire dataset (or any portion) from the [Downloads Tab](https://tokenomics.io/gitcoin) of the website.