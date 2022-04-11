## Gitcoin Data Pouch - Data/Appearances

The files in this folder contain, for each address, the 'appearances' of that address on the chain.

An appearance may be any location anywhere in the data - this includes the obvious places such as the `miner` of a block, the `from` and `to` addresses of a transaction, or the address of newly minted `smart contracts`.

But, that is not everywhere addresses appear. Appearances also include less obvious places such as inside the bytes of a transaction's `input` or `output` data or even twenty levels deep down in a smart contract `trace`.

The number of appearances for any address will be the same as the number of transactions reported for that address in its assoicated `txs` data set.

The format of an appearance is purposefully very simple:

| Name             | Description                          |
| ---------------- | ------------------------------------ |
| blockNumber      | the transaction's block number       |
| transactionIndex | the transcation's index in the block |

### Notes

- Much of the subseqent process done by `chifra export` is based on having first produced this `appearances` data set. For example, the `neighbors` data set is every neighbor appearing in any transaction in an address's appearance list, and so on.

- Appearances are a slightly fuzzy concept. When an address appears in a known place in the data (`from` or `to`) it is certainly a valid apperances. However, we also search the `input` data of a transaction, the `data` field of events, the `output` field of traces as well as the `log topics` for each event. This sometimes produces *false positives* in our index, which means there are entries in our index that are not addresses, but the number of these false positives is very small.

- Notwithstanding the fact that there may be false positives in our index, a false positive is never generated in the extracted data for a particular address. If an appearance is reported for an address, the string of bytes represented by that address defniitely appears in the given transcaction. Sometimes, it may be hard to find (and that almost always means it's ten levels deep in a `trace`).

### How this Data is Created

Use the following chifra command to generate this data for each address:

```
chifra export --appearances --fmt csv <address> | cut -f2,3 -d','
```

### Where's the Data?

The actual data produced for this project is too big to put on GitHub. You may download the entire dataset (or any portion) from the [Downloads Tab](https://tokenomics.io/gitcoin) of the website.