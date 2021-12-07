## Gitcoin Data Pouch - Data/Neighbors

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

### Notes

- `log_topic` and `log_data` are numbered by log and in the case of topic by topic, so you will see `log_1_topic_0` or `log_2_data` as an example.

- It is possible to generate much more detailed information for each neighbor. For example, the third topic on the fifteenth log in the tenth-level deep trace of the ninth trace, but we don't do that here for preformance reasons. See the `chifra export --neighbors` option called `--deep` for more information.

- In this version of the data, we show only the first appearance of a given address. In many cases, an address may appear multiple times in a transcation. We eliminated these second and subsquent appearances to keep the data size manageable. Again, see the `--deep` option.
