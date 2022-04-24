import React from 'react';
import { Typography, Layout } from 'antd';

import './App.css';
import 'antd/dist/antd.css';

const { Header } = Layout;
const { Title, Text } = Typography;

export const Head = () => {
  return (
    <Header style={{ height: '140px' }}>
      <div style={{ display: 'grid', gridTemplateColumns: '2fr 10fr 2fr' }}>
        <div>
          <i style={{ padding: '3px', color: 'black', backgroundColor: 'tomato' }}>** Not affiliated **</i>
        </div>
        <Typography>
          <Title style={{
            color: 'lightblue', paddingBottom: "0px", paddingTop: "0px", margin: "0px"
          }}>
            A <Text type='warning'>Permissionless</Text> Data Dump for the Community
            <br />
            <a className='install-tb' href="https://trueblocks.io/" target="top">You can run this code on your own computer!</a>
          </Title>
        </Typography>
        <div style={{ textAlign: 'right' }}>
          <a rel='noreferrer' target='_blank' href='https://discord.gg/RAz6DJ6xkf'>
            <Text style={{ color: 'lightblue', textDecoration: 'underline' }}>Like to help?</Text>
          </a>
        </div>
      </div>
    </Header>
  );
};
