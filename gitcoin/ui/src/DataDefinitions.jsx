import React from 'react';
import { Typography } from 'antd';

const { Link } = Typography;

export const DataDefinitions = () => {
  return (
    <div style={{ textAlign: 'left' }}>
      <center>
        <h2>Data Definitions</h2>
      </center>
      <div style={{ display: 'grid', gridTemplateColumns: "1fr 4fr 1fr" }}>
        <div></div>
        <div>
          Currently, we produce four data sets, thus. They are described on the linked README files:<p />
          <li><a target="top" href="https://github.com/TrueBlocks/tokenomics.io/blob/master/gitcoin/data/apps/README.md">Appearances</a> - a list of all appearances for an address</li>
          <li><a target="top" href="https://github.com/TrueBlocks/tokenomics.io/blob/master/gitcoin/data/txs/README.md">Transactions</a> - transactional details for an address</li>
          <li><a target="top" href="https://github.com/TrueBlocks/tokenomics.io/blob/master/gitcoin/data/logs/README.md">Logs</a> - all GitCoin-related logs for an address</li>
          <li><a target="top" href="https://github.com/TrueBlocks/tokenomics.io/blob/master/gitcoin/data/neighbors/README.md">Neighbors</a> - a list of address that appear in the same transactions as this address</li>
        </div>
        <div></div>
      </div>
    </div>);
};
