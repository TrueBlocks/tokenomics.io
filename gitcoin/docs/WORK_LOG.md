# 2021-11-28

#### Resurrected work from GitCoin Round 8 - Data Pouch

#### Automated deployment of the website

#### Scraped GitCoin API for a list of all grants by ID using:

```
curl "https://gitcoin.co/api/v0.1/grants/?pk=$1" | jq | tee data/raw/$1.json
sleep 2
```
