## Data/Txs

This folder contains the actual transactional data for each appearance. Note that in some cases, the address may not be appearant. For example, an address that appears in an internal trace of a smart contract may not appear in the exported fields of the rows of this file.

The data format for a transaction is:

| Name             | Description                                                                          |
| ---------------- | ------------------------------------------------------------------------------------ |
| blockNumber      | the transaction's block number                                                       |
| transactionIndex | the transcation's index in the block                                                 |
| date             | the date and time of the block                                                       |
| timestamp        | the timestamp of the block                                                           |
| from             | the sender of the transaction                                                        |
| to               | the recipient of the transaction                                                     |
| ether            | the amount of the transaction in ether                                               |
| ethGasPrice      | the gasPrice of the transcation in ether                                             |
| gasUsed          | the amount of gasUsed by the transaction (not part of raw tx, picked up from receipt |
| hash             | the transaction's hash                                                               |
| isError          | true if the transaction ended in error                                               |
| encoding         | the four-byte encoding of the transaction's function call                            |
| compressedTx     | the compressed version of the articulated transaction in a single field              |

### How this Data is Created

The following chifra command is used to create the data in this folder. It is called for each address.

```
chifra export --articulate --cache --cache_traces --fmt csv <address>
```

### Where's the Data?

The actual data produced for this project is too big to put on GitHub. You may download the entire dataset (or any portion) from the [Downloads Tab](https://tokenomics.io/gitcoin) of the website.