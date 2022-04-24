## Data/Neighbors

Neighbors are created by visiting each appearance for a given address and extracting a list of every address that also
appears in that same transaction. These lists of neighbors can get quite large.

In this data set, there are 10.6 times as many neighbors as there are appearances. That is, on average there are 10.6 other
addresses appearing in a transaciton as the address of interest itself.

The data format for a neighbor is:

| Name               | Description                                                                              |
| ------------------ | ---------------------------------------------------------------------------------------- |
| blockNumber        | the transaction's block number                                                           |
| transactionIndex   | the transcation's index in the block                                                     |
| neighbor           | the address of the account also in this transaction                                      |
| reason (see below) | one of `[ from, to, input, log_generator, log_topic, log_data, creation, trace, miner ]` |

### Reasons

An address may appear in a transaction for many reasons. Our best explanation is this image:

![](https://trueblocks.io/data-model/data-model-600.png)

"Obvious" addresses may appear in any of the pink locations in the Ethereum data. These items, after all, are specifically labeled as being of type address, however, TrueBlocks digs deeper.

In addition to searching all the obvious places in the data, TrueBlocks also digs into the bytes data notated above in green. For this reason, think of `neighbors` not necessarily as sources or destinations of funds into or out of an account, but as addresses that appear co-incident in a given transaction to the address in question.

Here's a short explanation of each `reason` (or location) where the address was found in a given transaction.

First we detail the "obvious" locations:

| Reason        | Description                                                                                                                                                                                                                                                                 |
| ------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| from          | The top-level `from` address of the transactions. Always an EOA. This address pays the gas.                                                                                                                                                                                 |
| to            | The top-level `to` address of the transaction.                                                                                                                                                                                                                              |
| log_generator | The address of any smart contract that may have, during the transaction, generated an event. Note that this frequently includes smart contracts other than `to`.                                                                                                            |
| log_topic     | Frequently, addresses appear in the `log_topics` of a transaction. While this is not technically labeled as an address, it is very obvious when a `log_topic` contains an address. Every valid token transfer includes both the sender and the recipient in the log_topics. |
| creation      | If a transaction results in the creation of a smart contract, this value will hold the address of the new contract.                                                                                                                                                         |
| miner         | This is the address of the winning `miner` for the block and/or the miner who mined any uncles in the block.                                                                                                                                                                |

And finally, the less obvious "bytes-parsed" locations:

| Reason   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| input    | Very, very frequently, addresses appear in the `input` data of a transtation but not elsewhere. In this case, that address will appear with this reason.                                                                                                                                                                                                                                                                                                     |
| log_data | Like the `input` data, addresses frequently appear in a log's `data` field. Again, in that case, the address will appear with this reason (if it has not already appeared previously).                                                                                                                                                                                                                                                                       |
| trace    | Addresses with this reason appear no-where else in the transaction other than in the traces. The trace data is very complex (see the above image) and takes a very long time to extract. For this reason, we include a reason of `trace` for any address that appears in a transaction without having been previously labeled. See the `--deep` option of the `chifra export` command, which digs as deeply as possible into `traces`, for more information. |

### Notes

- `log_topic` and `log_data` are numbered by log and in the case of topic by topic, so you will see `log_1_topic_0` or `log_2_data` as an example.

- It is possible to generate much more detailed information for each neighbor. For example, the third topic on the fifteenth log in the tenth-level deep trace of the ninth trace, but we don't do that here for preformance reasons. See the `chifra export --neighbors` option called `--deep` for more information.

- In this version of the data, we show only the first appearance of a given address. In many cases, an address may appear multiple times in a transcation. We eliminated these second and subsquent appearances to keep the data size manageable. Again, see the `--deep` option.


### How this Data is Created

1. Run the following chifra command for each address

```
chifra export --neighbors --deep --fmt csv <address>
```

### Where's the Data?

The actual data produced for this project is too big to put on GitHub. You may download the entire dataset (or any portion) from the [Downloads Tab](https://tokenomics.io/gitcoin) of the website.