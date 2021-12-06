import React from 'react';
import { Typography } from 'antd';

const { Link } = Typography;

export const DataForNerds = () => {
  return (
    <div style={{ textAlign: 'left' }}>
      <center>
        <h3>Coming soon...</h3>
        <a rel='noreferrer' target='_blank' href='https://github.com/TrueBlocks/tokenomics.io/issues'>
          <Link>Have a suggestion?</Link>
        </a>
      </center>
      <b>Zips:</b>
      <ul>
        <li>Every file in a single zip</li>
        <li>Appearances for [all address | all grants | all core ]</li>
        <li>Transactions for [all address | all grants | all core ]</li>
        <li>Logs for [all address | all grants | all core ]</li>
        <li>Neighbors for [all address | all grants | all core ]</li>
        <li><b>Note:</b> Data sets containing all of the data models for a single address as a zip files are linked on from the main display.</li>
      </ul>
      <b>Ideas:</b> (<a target="top" href="https://discord.com/invite/RAz6DJ6xkf">help us build!</a>)
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
