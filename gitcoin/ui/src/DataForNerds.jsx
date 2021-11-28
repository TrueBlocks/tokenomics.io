import React from 'react';
import {Typography} from 'antd';

const {Link} = Typography;

export const DataForNerds = () => {
  var record = { "tx_cnt": "1" };
  return (
    <div style={{textAlign: 'left'}}>
      <center>
        <h3>Coming soon...</h3>
        <a rel='noreferrer' target='_blank' href='https://github.com/TrueBlocks/tokenomics.io/issues'>
          <Link>Have a suggestion?</Link>
        </a>
      </center>
      <ul>
        <li>Percent of total unique donors</li>
        <li>Uniq recipients by date</li>
        <li>Uniq donors by count</li>
        <li>Uniq recipients by count</li>
        <li>Donor counts by day</li>
        <li>Recipient counts by day</li>
        <li>Reciprocal pairs by day</li>
        <li>Donation amount by bucket</li>
        <li>Comparison of Round 8 to Round 9</li>
      </ul>
    </div>
  );
};
