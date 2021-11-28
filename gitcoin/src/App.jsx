import React from 'react';
import {Layout} from 'antd';
import {Affix} from 'antd';

import './App.css';
import 'antd/dist/antd.css';
import {Head} from './Header';

import {RightSider} from './RightSider';
import {LeftSider} from './LeftSider';
import {Foot} from './Footer';
import {HomePage} from './HomePage';

const {Content} = Layout;

function App() {
  return (
    <Layout className='App' style={{backgroundColor: '#e7dddc'}}>
      <Affix offsetTop={0}>
        <Head />
      </Affix>
      <Layout>
        <LeftSider />
        <Content style={{padding: '10px', paddingLeft: '10px', backgroundColor: 'lightblue'}}>
          <Layout>
            <HomePage />
            <RightSider />
          </Layout>
        </Content>
      </Layout>
      <Affix offsetBottom={0}>
        <Foot />
      </Affix>
    </Layout>
  );
}
export default App;
