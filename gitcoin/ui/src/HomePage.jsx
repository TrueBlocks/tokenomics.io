import React, { useEffect, useState } from 'react';
import { useStatePersist } from 'use-state-persist';

import { Input, Layout, Tabs } from 'antd';

import './App.css';
import 'antd/dist/antd.css';

import { grantsData } from './grants-data';
import { columns } from './ColumnDefs';
import { BaseTable } from './BaseTable';

import { Downloads } from './Downloads';
import { DataDefinitions } from './DataDefinitions';

import { lastUpdate } from './last-update.js';

const { Content } = Layout;
const { Search } = Input;
const { TabPane } = Tabs;

export const HomePage = () => {
  const [lastTab, setLastTab] = useStatePersist('@lastTab', 1);
  const [searchText, setSearchText] = useState('');

  const onSearch = (value) => {
    setSearchText(value.toLowerCase());
    console.log(value);
  };

  const contractData = grantsData.filter((item) => {
    const n = item.name.toLowerCase();
    const a = item.address.toLowerCase();
    return item.core && (searchText === '' || n.includes(searchText) || a.includes(searchText));
  });
  const grantData = grantsData.filter((item) => {
    const n = item.name.toLowerCase();
    const a = item.address.toLowerCase();
    return !item.core && (searchText === '' || n.includes(searchText) || a.includes(searchText));
  });

  var val = lastTab;
  const tabSwitch = (key, event) => {
    val = key;
    setLastTab(val);
  };

  useEffect(() => {
    setLastTab(val);
  }, [val, setLastTab]);

  const tab1Title = 'Donation Contracts (' + contractData.length + ')';
  const tab2Title = 'Individual Grants (' + grantData.length + ')';
  return (
    <Content>
      <div style={{ display: 'grid', gridTemplateColumns: '3fr 1fr' }}>
        <div></div>
        <Search
          allowClear
          style={{ paddingRight: '2px' }}
          width='10px'
          placeholder='search grants by address or name...'
          onSearch={onSearch}
          enterButton></Search>
      </div>
      <Tabs defaultActiveKey={lastTab} onChange={tabSwitch} style={{ border: '1px dotted gray', padding: '1px' }}>
        <TabPane tab={tab2Title} key='1' style={{ paddingLeft: '8px', margin: '-25px 0px 0px 0px' }}>
          <BaseTable dataSource={grantData} columns={columns} rowKey={(record) => record.grantId} />
        </TabPane>
        <TabPane tab={tab1Title} key='2' style={{ paddingLeft: '8px' }}>
          <BaseTable dataSource={contractData} columns={columns} rowKey={(record) => record.grantId} />
        </TabPane>
        <TabPane tab={'Downloads'} key='3' style={{ paddingLeft: '8px' }}>
          <Downloads />
        </TabPane>
        <TabPane tab='Data Definitions' key='5' style={{ paddingLeft: '8px' }}>
          <DataDefinitions />
        </TabPane>
        <TabPane tab='Charts' key='4' style={{ paddingLeft: '8px' }}>
          <img
            width='800px'
            alt='Unclaimed'
            src='https://tokenomics.io/gitcoin/charts/Unclaimed%20Match%20Round%208.png'
          />
          <br />
          <br />
          <img width='800px' alt='Count By Date' src='https://tokenomics.io/gitcoin/charts/Counts.png' />
        </TabPane>
      </Tabs>
      <i>
        <small>{lastUpdate}</small>
      </i>
    </Content>
  );
};
