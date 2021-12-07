## Gitcoin Data Pouch - Data/Appearances

The files in this folder display, for each address, the 'appearances' of the address on the chain. An appearance can be any location anywhere in the data - this includes the obvious places such as being a `miner`, the `from` or `to` address, or a newly minted `smart contract`, but also includes less obvious places such as inside of a transaction's `input` or `output` data fields or twenty levels deep down in a `trace`.

The number of appearances for any address is the same as the number of transactions reported in the `txs` data set.

The data format for an appearance is:

| Name             | Description                          |
| ---------------- | ------------------------------------ |
| blockNumber      | the transaction's block number       |
| transactionIndex | the transcation's index in the block |
