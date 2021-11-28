import React from 'react';
import {Layout, Card, Popover} from 'antd';

import './App.css';
import 'antd/dist/antd.css';

import {faq_title, faq_text} from './FAQ.jsx';

const {Sider} = Layout;

const Strikeout = ({text}) => {
  return <div style={{display: 'inline', textDecoration: 'line-through'}}>{text}</div>;
};

const MyCard = ({children}) => {
  return (
    <Card
      style={{
        marginLeft: '4px',
        marginRight: '4px',
        marginBottom: '4px',
        border: '1px dotted gray',
        backgroundColor: 'antiquewhite',
      }}>
      {children}
    </Card>
  );
};

export const RightSider = () => {
  const hover1_text = (
    <ul style={{marginLeft: '-20px'}}>
      <li>
        <Strikeout text='Enable search by name' />
      </li>
      <li>Update the data after every block</li>
      <li>
        Export transactions, traces, reconciliations, etc.
      </li>
      <li>
        Allow for better filtering (e.g. remove highly-funded grants)
      </li>
      <li>
        Make data pouches for other ecosytems
      </li>
    </ul>
  );
  const hover1_title = 'Future Work';
  return (
    <Sider style={{backgroundColor: 'lightblue'}}>
      <MyCard>
        <div style={{fontSize: '12pt', fontWeight: '800'}}>
          <a rel='noreferrer' target='_blank' href='https://gitcoin.co/grants/184/trueblocks'>
            Donate to the TrueBlocks Grant
          </a>
        </div>
      </MyCard>
      <MyCard>
        Wondering how this website works? Read about the <b>Unchained Index</b>{' '}
        <a rel='noreferrer' target='_blank' href='https://unchainedindex.io'>
          here
        </a>
      </MyCard>
      <MyCard>
        Interested in learning more?{' '}
        <a rel='noreferrer' target='_blank' href='https://discord.gg/RAz6DJ6xkf'>
          Join our discord discussions
        </a>
      </MyCard>
      <MyCard>
        <Popover
          style={{border: '1px dashed black'}}
          color='lightblue'
          placement='left'
          title={hover1_title}
          content={hover1_text}
          trigger='hover'>
          <div style={{width: '100%', height: '100%'}}>Future Work</div>
        </Popover>
      </MyCard>
      <MyCard>
        <Popover
          style={{border: '1px dashed black'}}
          color='lightblue'
          placement='left'
          title={faq_title}
          content={faq_text}
          trigger='hover'>
          <div style={{width: '100%', height: '100%'}}>FAQ</div>
        </Popover>
      </MyCard>
    </Sider>
  );
};
