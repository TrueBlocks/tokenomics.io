## Gitcoin Data Pouch - Data/Logs

The `./logs` folder contains exports of all the logs for a given address. Note that this may be many multiples of the number of transactions / appearances for that address.

The data format for a log is:

| Name             | Description                          |
| ---------------- | ------------------------------------ |
| blockNumber      | the transaction's block number       |
| transactionIndex | the transcation's index in the block |
| logIndex         |                                      |
| address          |                                      |
| topic0           |                                      |
| topic1           |                                      |
| topic2           |                                      |
| topic3           |                                      |
| data             |                                      |
| type             |                                      |
| compressedLog    |                                      |
| timestamp        |                                      |
