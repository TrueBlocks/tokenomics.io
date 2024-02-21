## Description of GitCoin API at https://gitcoin.co/api/v0.1/grants

We used with `pk` ranging from 0-4000:

```
curl "https://gitcoin.co/api/v0.1/grants/?pk=$1" | jq | tee data/raw/$1.json
sleep 2
```

This produced a database of raw JSON for the grants which included these fields:

| Field Name                          | Description                                                                |
| ----------------------------------- | -------------------------------------------------------------------------- |
| id                                  | the grant id                                                               |
| active                              | `true` if this grant is active                                             |
| title                               | the name of the grant and source for the slug                              |
| slug                                | grant title with non-alphanumeric characters removed and spaces as dashes  |
| description                         | a description of the grant provided by the grantee                         |
| reference_url                       | the grant's URL, if any                                                    |
| logo                                | the grant's log, if any                                                    |
| admin_address                       | Ethereum address of the admin user for the grant                           |
| amount_received                     | the amount received by this grant                                          |
| token_address                       | token address of particular token accepted or 0x0                          |
| token_symbol                        | the symbol associated with the `token_address` or `Any Token`              |
| contract_address                    | the grant's contract address, if any                                       |
| metadata                            | meta data about the grant (see below)                                      |
| network                             | the network where the project runs                                         |
| required\_gas\_price                | unclear                                                                    |
| admin_profile                       | the profile of the admin of the grant (see below)                          |
| team_members                        | an array of profiles of the team member profiles (see below) for the grant |
| clr\_prediction\_curve              | an array of three-dimensional points with unclear meaning                  |
| clr\_round\_num                     | unclear                                                                    |
| is\_clr\_active                     | unclear                                                                    |
| amount\_received\_in\_round         | the amount of grant money received in the current round?                   |
| positive\_round\_contributor\_count | unclear                                                                    |


### Other Data Types

These additional data types are part of the `v0.1/grants` endpoint:

**MetaData:**

| Field Name                                     | Description |
| ---------------------------------------------- | ----------- |
| gem                                            | unknown     |
| upcoming                                       | unknown     |
| wall\_of\_love                                 | unknown     |
| last\_calc\_time\_contributor\_counts          | unknown     |
| last\_calc\_time\_sybil\_and\_contrib\_amounts | unknown     |

**Profile:**

| Field Name    | Description                                    |
| ------------- | ---------------------------------------------- |
| id            | the profile's id                               |
| url           | URL to the profile on github.com               |
| name          | the profile's name                             |
| handle        | the profile's handle on github.com             |
| keywords      | array of keywords associated with this profile |
| position      | unknown                                        |
| avatar_url    | github.com avatar URL                          |
| github_url    | github.com profile URL                         |
| total_earned  | total earned from GitCoin?                     |
| organizations | array of github organizations                  |

### What we discovered:

- We wish to include deep links into each grant to allow visitors to our site to donate to grants they find interesting.

- The Gitcoin Grants website presents the grants at http://gitcoin.co/grants/{id}/{slug}.

- The API provides no easy way to aquire an association of `grant id` and `slug`. This makes producing deep links into the grant website more difficult than they should be.

- The only way we could figure out how to get this information programatically is using a shell script that runs through the individual grant id's and call into the API using `pk`. We found that a large percentage of the grant ids (about 55%) did not refer to a valid grant. This places an unnecessary burder on the server.

- **Suggestion:** Add an end point to return an assoication of `grant_ids` and `slugs` in order to make deep linking into the grant website easier.
