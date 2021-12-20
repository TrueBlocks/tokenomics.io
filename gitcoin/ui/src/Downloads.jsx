import React from 'react';
import { Typography } from 'antd';

const { Link } = Typography;

export const Downloads = () => {
  return (
    <div style={{ textAlign: 'left' }}>
      <center>
        <h2>Easy Access to the Data Sets</h2>
        <a rel='noreferrer' target='_blank' href='https://github.com/TrueBlocks/tokenomics.io/issues'>
          <Link>Have a suggestion?</Link>
        </a>
      </center>
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
      <br />
      <b>Some ideas for future data sets:</b> (<a target="top" href="https://discord.com/invite/RAz6DJ6xkf">help us build!</a>)
      <ul>
        <li>Percent of total unique donors for each grant</li>
        <block>- Assuming X unique individual donors for all grants, this metric would show Y / X, where Y is the number of unique donors for a given grant.</block>
        <li>Percent of total unique grants donated to by each donor</li>
        <block>- Assuming X unique grants in a round, this metric would show Y / X, where Y is the number of grants a donor donated in the round.</block>
        <li>Uniq recipients by date</li>
        <block>- A data set showing the historical number of unique donors by date.</block>
        <li>Donor counts by day</li>
        <block>- A data set showing the total number of active donors per day per round.</block>
        <li>Recipient counts by day</li>
        <block>- A data set showing the total number of active grants (those being donated to) per day per round.</block>
        <li>Reciprocal pairs by day</li>
        <block>- Data sets showing the reciprical of much of the above data, where appropriate.</block>
        <li>Donation amount by bucket</li>
        <block>- Donation amounts bucketed by amount of donation.</block>
        <li>Comparison of Rounds</li>
        <block>- Per round data for many of the above metrics.</block>
      </ul>
    </div>
  );
};
