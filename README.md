# Tokenomics.io

A website and research platform associated with TrueBlocks. We're trying to show people what's possible when the world finally has globally accessible permissionless data. We no longer have to "ask permission" to the see the books of our charitable organizations.

"Isn't this a privacy invasion?"

No. It's an experiment trying to figure out what types of organizations can exist in a world where anyone can "look into the books." This is a natural thing that happens with permissionless blockchains.

## Running in Docker

```bash
cd docker
cp env.example .env
# Adjust the settings
$EDITOR .env
docker compose up
```

## Current Experiments

- [Ecosystem-Wide Accounting]
  - [GitCoin](./gitcoin)
  - [Giveth](./giveth)
  - [CLRFund](./clrfund)

- [R Scripts Investigating Other Things](https://github.com/TrueBlocks/tokenomics)

## Adding experiment

1. Add experiment name to `WEBSITES` in your `docker/.env` file.
2. To the same file, add: `NOMICS_EXPERIMENT-NAME-HERE_CHAINS=chain1,chain2`
3. Make sure both chains from step 2 are configured (their chain id and RPC provider URL have to defined as variables in `docker/.env` file as well)
4. In `docker/docker-compose.yml`, add new volume per each experiment's chain:
```yaml
volumes:
  experiment-name_exports_chain1:
  experiment-name_exports_chain2:
```
5. Add these volumes to `tokenomics` service:
```yaml
volumes:
  - experiment-name_exports_chain1:/root/tokenomics.io/experiment-name/exports/chain1
  - experiment-name_exports_chain2:/root/tokenomics.io/experiment-name/exports/chain2
```
6. Copy and paste existing `monitor_` service and adjust it for the new experiment. Mount volume from step 4 as `/exports` (for multiple volumes, you need multiple monitor services).
7. Create a new directory with `addresses.tsv` file in it, add addresses you want to watch to the file and mount the directory as `/addresses`:
```yaml
volumes:
  - ../experiment-name:/addresses:ro
```
8. Call `docker compose up` to run the containers.

## Folder Structure

This is the structure of the folders along with a short explanation of each.

.
├── ./ui
│
│ A react application for showing the various data
│
├── ./R
│
│ R scripts for playing around with the data
│
└── ./data
    │
    │ collections of scraped data and exports for various ecosystems
    │
    ├── ./gitcoin
    ├── ./giveth
    └── ./clrfund
