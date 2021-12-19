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

### Where's the Data?

The actual data produced for this project is too big to put on GitHub. You may download the entire dataset (or any portion) from the [Downloads Tab](https://tokenomics.io/gitcoin) of the website.