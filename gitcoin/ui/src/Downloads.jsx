import React from 'react';
import { Typography } from 'antd';
import { configUrls } from './Config';

const { Link } = Typography;

export const Downloads = () => {
  return (
    <div style={{ textAlign: 'left' }}>
      <center>
        <h2>Data Sets</h2>
        <a rel='noreferrer' target='_blank' href='https://github.com/TrueBlocks/tokenomics.io/issues'>
          <Link>Have a suggestion?</Link>
        </a>
      </center>
      <div style={{ display: 'grid', gridTemplateColumns: '1fr 12fr' }}>
        <div></div>
        <div>
          <b>Download zip files of:</b>
          <ul>
            <li>All data for [ <a target="top" href={configUrls.DataZips + "combined/"}>all address</a> ]</li>
            <li>Appearances for [ <a target="top" href={configUrls.DataZips + "combined/apps.csv.gz"}>all address</a> ]</li>
            <li>Transactions for [ <a target="top" href={configUrls.DataZips + "combined/txs.csv.gz"}>all address</a> ]</li>
            <li>Logs for [ <a target="top" href={configUrls.DataZips + "combined/logs.csv.gz"}>raw data</a> | <a target="top" href={configUrls.DataZips + "combined/logs_articulated.csv.gz"}>articulated</a> ]</li >
            <li>Neighbors for [ <a target="top" href={configUrls.DataZips + "combined/neighbors.csv.gz"}>all address</a> ]</li >
            <li>Statements for [ <a target="top" href={configUrls.DataZips + "combined/statements.csv.gz"}>raw data</a> | <a target="top" href={configUrls.DataZips + "combined/statements_balances.csv.gz"} > summaries</a > | <a target="top" href={configUrls.DataZips + "combined/statements_tx_counts.csv.gz"}>asset tx counts</a> ]</li >
          </ul >
        </div >

        <div></div>
        <div>
          <b>Data Definitions: (click for more information):</b>
          <ul>
            <li><a target="top" href={configUrls.DataDefs + "apps/README.md"}>Appearances</a> - a list of all appearances for an address</li>
            <li><a target="top" href={configUrls.DataDefs + "txs/README.md"}>Transactions</a> - transactional details for an address</li>
            <li><a target="top" href={configUrls.DataDefs + "logs/README.md"}>Logs</a> - all GitCoin-related logs for an address</li>
            <li><a target="top" href={configUrls.DataDefs + "neighbors/README.md"}>Neighbors</a> - a list of address that appear in the same transactions as this address</li>
            <li><a target="top" href={configUrls.DataDefs + "statements/README.md"}>Statements</a> - reconciled banking statements for all assets owned by an address</li>
          </ul>
        </div>

        <div></div>
        <div>
          <p />
          <p />
          <b>Note:</b> Downloads containing complete data sets for individual grants are linked from the main display.
        </div>

      </div >
    </div >
  );
};
