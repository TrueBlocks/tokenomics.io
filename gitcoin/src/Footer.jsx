import React from 'react';
import { Layout } from 'antd';
import {TwitterOutlined, GithubOutlined, MailOutlined} from '@ant-design/icons';

import './App.css';
import 'antd/dist/antd.css';
const {Footer} = Layout;

export const Foot = () => {
  var share =
    'https://twitter.com/intent/tweet?text=A%20permissionlessly-generated%20gift%20to%20the%20@gitcoin%20community%20powered%20by%20@trueblocks.%20Gitcoin%20Grant%20Data%20Pouch%20(https://tokenomics.io/gitcoin).';
  return (
    <Footer>
      <div style={{fontSize: '9pt', display: 'grid', gridTemplateColumns: '1fr 10fr 1fr'}}>
        <div style={{textAlign: 'left'}}></div>
        <div>
          <i>Every log from every grant directly from mainnet</i>
          <br />
          A project of TrueBlocks, LLC &diams; Powered by <a href='https://twitter.com/trueblocks'>@trueblocks</a>{' '}
          and <a href='https://tokenomics.io'>tokenomics.io</a>
        </div>
        <div style={{textAlign: 'right'}}>
          <a rel='noreferrer' target='_blank' href={share}>
            <TwitterOutlined style={{fontSize: '14pt', color: '#282c34'}} />
          </a>
          {'  '}
          <a rel='noreferrer' target='_blank' href='https://github.com/TrueBlocks/tokenomics.io'>
            <GithubOutlined style={{fontSize: '14pt', color: '#282c34'}} />
          </a>
          {'  '}
          <a rel='noreferrer' target='_blank' href='mailto:info@trueblocks.io'>
            <MailOutlined style={{fontSize: '14pt', color: '#282c34'}} />
          </a>
        </div>
      </div>
    </Footer>
  );
};
