## Data/Zips

The `./zips` folder, and therefore, the `zips` endpoint, holds combinations of the other data sets. You may access it from the website at [https://tokenomics.io/gitcoin/data/zips/](https://tokenomics.io/gitcoin/data/zips/).

#### Per Address Zips

For each address, we include a .gz file of all the data sets for that address (`apps`, `txs`, `logs`, etc.). Tihs file in the `./zips` folder and is called `<address>.tar.gz`. These files may be downloaded from the `zips` endpoint.

#### Per Data Type Zips

We also collect together each different data type for all addresses in a single file, also in `./zips`, called `apps.tar.gz`, `txs.tar.gz`, `logs.tar.gz` etc. These files may be downloaded on the **Data for Nerds** page.

### How this Data is Created

The files in this folder are created at the end of each scan across all of the addresses. A simple bash shell script called [./update_zips.1.sh](../update_zips.1.sh) is called for the various addresses and data types by another shell script called [./update_zips.sh](./../update_zips.sh).

### Where's the Data?

The actual data produced for this project is too big to put on GitHub. You may download the entire dataset (or any portion) from the [Downloads Tab](https://tokenomics.io/gitcoin) of the website.
