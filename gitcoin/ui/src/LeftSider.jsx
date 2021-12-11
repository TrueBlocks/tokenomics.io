import React, { useState, Fragment } from 'react';
import { Typography, Layout } from 'antd';

import './App.css';
import 'antd/dist/antd.css';

const { Sider } = Layout;
const { Paragraph, Text } = Typography;

const LeftSideItem = ({ question, answer }) => {
  return (
    <Fragment>
      <Paragraph style={{ textDecoration: 'underline', fontWeight: 'bold', color: 'lightblue' }}>
        {question}
      </Paragraph>
      <Text style={{ color: 'lightblue' }}>
        {answer}
      </Text>
      <br />
      <br />
    </Fragment>
  );
}

const LocalCheckbox = () => {
  const [local, setLocal] = useState(localStorage.getItem("local") === "on" ? true : false)
  var box = <input type="checkbox" id="option3" onChange={() => { localStorage.setItem("local", "on"); setLocal("on"); }} />
  if (local) {
    box = <input checked type="checkbox" id="option3" onChange={() => { localStorage.setItem("local", ""); setLocal(""); }} />
  }
  return (
    <div id="checkbox-container">
      <div>{box} <label for="option">http://localhost:1234</label></div >
    </div>
  );
}

export const LeftSider = () => {
  return (
    <Sider style={{ paddingLeft: '20px', paddingRight: '20px' }}>
      {/* <LeftSideItem
        question='Do you have an unclaimed match?'
        answer={
          <div>
            {'Search for your grant to the right, then look in the '}
            <Text type='warning'>CLR</Text> column for un-claimed matching funds.
          </div>
        }
      /> */}
      <LeftSideItem
        question='What is a data pouch?'
        answer={`
          A data pouch is a place to store and disseminate data. This pouch was created with TrueBlocks
          and a locally-running Ethereum node.
        `}
      />
      <LeftSideItem
        question='Why did we build this?'
        answer={`
          To demonstrate what's possible with local first software. Also, to provide others in the community an independant source of
          transparent data about the GitCoin Grant program.
        `}
      />
      <br />
      <br />
      <small>
        <Text style={{ color: 'lightblue' }}>
          This website is <i>alpha</i>, which means you should use the data with caution.
          <br />
          <br />
          <a rel='noreferrer' target='_blank' href='https://github.com/TrueBlocks/tokenomics.io/issues'>
            <Text style={{ color: 'lightblue', textDecoration: 'underline' }}>Have a suggestion?</Text>
          </a>
        </Text>
      </small>
      <div style={{ position: "absolute", bottom: 20, left: 20 }}>
        <Text style={{ color: 'lightblue' }}>
          <LocalCheckbox />
        </Text>
      </div>
    </Sider>
  );
};
