import React, { useState } from 'react';
import { Typography, Layout } from 'antd';
import { TwitterOutlined, GithubOutlined, MailOutlined } from '@ant-design/icons';

import './App.css';
import 'antd/dist/antd.css';
const { Footer } = Layout;
const { Text } = Typography

export const Foot = () => {
  var share =
    'https://twitter.com/intent/tweet?text=A%20permissionlessly-generated%20gift%20to%20the%20@gitcoin%20community%20powered%20by%20@trueblocks.%20Gitcoin%20Grant%20Data%20Pouch%20(https://tokenomics.io/gitcoin).';
  return (
    <Footer>
      <div style={{ fontSize: '9pt', display: 'grid', gridTemplateColumns: '1fr 10fr 1fr', border: '1px solid black' }}>
        <div style={{ textAlign: 'left' }}>
          <div style={{}}>
            <LocalCheckbox />
          </div>
        </div>
        <div>
          <i>Every appearance, log, neighbor, and transaction from every grant pulled directly from mainnet</i>
          <br />
          A project of TrueBlocks, LLC &diams; Powered by <a href='https://twitter.com/trueblocks'>@trueblocks</a>{' '}
          and <a href='https://tokenomics.io'>tokenomics.io</a>
        </div>
        <div style={{ textAlign: 'right' }}>
          <a rel='noreferrer' target='_blank' href={share}>
            <TwitterOutlined style={{ fontSize: '14pt', color: '#282c34' }} />
          </a>
          {'  '}
          <a rel='noreferrer' target='_blank' href='https://github.com/TrueBlocks/tokenomics.io'>
            <GithubOutlined style={{ fontSize: '14pt', color: '#282c34' }} />
          </a>
          {'  '}
          <a rel='noreferrer' target='_blank' href='mailto:info@trueblocks.io'>
            <MailOutlined style={{ fontSize: '14pt', color: '#282c34' }} />
          </a>
        </div>
      </div>
    </Footer>
  );
};

export const LocalCheckbox = () => {
  const [local, setLocal] = useState(localStorage.getItem("local") === "on" ? true : false)
  var box = <input type="checkbox" id="option" onChange={() => { localStorage.setItem("local", "on"); setLocal("on"); }} />
  if (local) {
    box = <input checked type="checkbox" id="option" onChange={() => { localStorage.setItem("local", ""); setLocal(""); }} />
  }
  return (
    <Text style={{ color: "black", position: "absolute", bottom: 20, left: 10 }}>
      <div>{box} <label htmlFor="option">http://localhost:1234</label></div >
    </Text>
  );
}
