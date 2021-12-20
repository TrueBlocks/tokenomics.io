import React from 'react';
import { Typography } from 'antd';

const { Link } = Typography;

export const Downloads = () => {
  return (
    <div style={{ textAlign: 'left' }}>
      <center>
        <h2>Download Data Sets</h2>
        <a rel='noreferrer' target='_blank' href='https://github.com/TrueBlocks/tokenomics.io/issues'>
          <Link>Have a suggestion?</Link>
        </a>
      </center>
      <div style={{ display: 'grid', gridTemplateColumns: '1fr 12fr' }}>
        <div></div>
        <div>
          <b>Zips:</b>
          <ul>
            <li>All data for [ <a target="top" href="https://tokenomics.io/gitcoin/data/zips/combined/">all address</a> ]</li>
            <li>Appearances for [ <a target="top" href="https://tokenomics.io/gitcoin/data/zips/combined/apps.csv.gz">all address</a> ]</li>
            <li>Transactions for [ <a target="top" href="https://tokenomics.io/gitcoin/data/zips/combined/txs.csv.gz">all address</a> ]</li>
            <li>Logs for [ <a target="top" href="https://tokenomics.io/gitcoin/data/zips/combined/logs.csv.gz">raw data</a> | <a target="top" href="https://tokenomics.io/gitcoin/data/zips/combined/logs_articulated.csv.gz">articulated</a> ]</li>
            <li>Neighbors for [ <a target="top" href="https://tokenomics.io/gitcoin/data/zips/combined/neighbors.csv.gz">all address</a> ]</li>
            <li>Statements for [ <a target="top" href="https://tokenomics.io/gitcoin/data/zips/combined/statements.csv.gz">raw data</a> | <a target="top" href="https://tokenomics.io/gitcoin/data/zips/combined/statements_balances.csv.gz">summaries</a> | <a target="top" href="https://tokenomics.io/gitcoin/data/zips/combined/statements_tx_counts.csv.gz">asset tx counts</a> ]</li>
          </ul>
          <ul><b>Note:</b> Data sets containing each individual address are linked on from the main display.</ul>
          <p />
        </div>
      </div>
    </div>
  );
};
