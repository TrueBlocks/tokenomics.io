## Gitcoin Data Pouch - Data/Txs

This folder contains the actual transactional data for each appearance. Note that in some cases, the address may not be appearant. For example, an address that appears in an internal trace of a smart contract may not appear in the exported fields of the rows of this file.

The data format for a transaction is:

| Name             | Description                          |
| ---------------- | ------------------------------------ |
| blockNumber      | the transaction's block number       |
| transactionIndex | the transcation's index in the block |
| date             |                                      |
| timestamp        |                                      |
| from             |                                      |
| to               |                                      |
| ether            |                                      |
| ethGasPrice      |                                      |
| gasUsed          |                                      |
| hash             |                                      |
| isError          |                                      |
| encoding         |                                      |
| compressedTx     |                                      |
