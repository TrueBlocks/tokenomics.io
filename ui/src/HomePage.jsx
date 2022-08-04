import React, { useCallback, useEffect, useState, useMemo } from 'react';
import { useStatePersist } from 'use-state-persist';

import { Input, Layout, Tabs, Row, Col, notification } from 'antd';

import './App.css';
import 'antd/dist/antd.css';

import { columns as columnDefinitions } from './ColumnDefs';
import { BaseTable } from './BaseTable';

import { Downloads } from './Downloads';
import { ToDo } from './ToDo';

import { Sidebar } from './Sidebar';
import { useGlobalState, getChainData } from './GlobalState';
import { loadData } from './DataSource'
import { ViewOptions } from './ViewOptions';
import { config } from './Config';
import { LastUpdate } from './LastUpdate';

const { Content } = Layout;
const { Search } = Input;
const { TabPane } = Tabs;

export const HomePage = () => {
  const [lastTab, setLastTab] = useStatePersist('@lastTab', 1);
  const [searchText, setSearchText] = useState('');
  const {
    project,
    selectGrant,
    sidebarEnabled,
    sidebarVisible,
    showZero,
    chain,
    setSidebarVisible,
  } = useGlobalState();

  const chartsPath = config.buildPath(project, 'charts');

  const onSearch = (value) => {
    setSearchText(value.toLowerCase());
  };

  const [grantsData, setGrantsData] = useState([]);
  const [loading, setLoading] = useState(false);
  useEffect(() => {
    (async () => {

      if (!project) return;

      setLoading(true);
      try {
        const data = await loadData(project, chain);
        setGrantsData(data);
      } catch (error) {
        notification.error({
          message: 'Error while fetching data',
          description: error.toString(),
          duration: 0,
        });
      }
      setLoading(false);
    })();
  }, [chain, project]);

  const coreData = useMemo(() => grantsData.filter((grantData) => {
    grantData['curChain'] = chain
    const n = grantData.name.toLowerCase();
    const a = grantData.address.toLowerCase();
    const show = showZero || getChainData(grantData).counts.appearanceCount > 0;
    return show && grantData.core && (searchText === '' || n.includes(searchText) || a.includes(searchText));
  }), [searchText, showZero, chain, grantsData]);

  const grantData = useMemo(() => grantsData.filter((grantData) => {
    grantData['curChain'] = chain
    const n = grantData.name.toLowerCase();
    const a = grantData.address.toLowerCase();
    const show = showZero || getChainData(grantData).counts.appearanceCount > 0;
    return show && !grantData.core && (searchText === '' || n.includes(searchText) || a.includes(searchText));
  }), [searchText, showZero, chain, grantsData]);

  const columns = useMemo(() => columnDefinitions, []);
  const sidebarShown = useMemo(() => sidebarEnabled && sidebarVisible, [sidebarEnabled, sidebarVisible]);
  const dataRowSpan = useMemo(() => sidebarShown ? 17 : 24, [sidebarShown]);
  const sideRowSpan = 24 - dataRowSpan

  var val = lastTab;
  const tabSwitch = (key, event) => {
    selectGrant(null);
    setSidebarVisible(false);
    val = key;
    setLastTab(val);
  };

  useEffect(() => {
    setLastTab(val);
  }, [val, setLastTab]);

  const onSelectionChange = useCallback((grant) => {
    selectGrant(grant);
    if (sidebarEnabled) {
      setSidebarVisible(true);
    }
  }, [selectGrant, setSidebarVisible, sidebarEnabled]);

  const [sortField, setSortField] = useState('');
  const [sortOrder, setSortOrder] = useState(null);
  const [sortedData, setSortedData] = useState(grantData)

  useEffect(() => {
    const sorted = grantData.sort(function (a, b) {
      return (a[sortField] - b[sortField]) * (sortOrder === 'ascend' ? -1 : 1);
    });
    setSortedData(sorted)
  }, [grantData, sortField, sortOrder])

  const changeSort = useCallback((field, order) => {
    setSortField(field)
    setSortOrder(order)
  }, [setSortField, setSortOrder])

  const tab1Title = 'Core Contracts (' + coreData.length + ')';
  const tab2Title = 'Individual Grants (' + grantData.length + ')';

  return (
    <Content>
      <ViewOptions />
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
      <Tabs defaultActiveKey={lastTab} onChange={tabSwitch} style={{ border: '1px dotted gray', padding: '1px', paddingLeft: '8px', paddingRight: '8px' }}>
        <TabPane tab={tab2Title} key='1' style={{ paddingLeft: '8px', margin: '-25px 0px 0px 0px' }}>
          <Row className='with-sidebar'>
            <Col span={dataRowSpan} className='col-table'>
              <BaseTable
                chain={chain}
                changeSort={changeSort}
                dataSource={sortedData}
                columns={columns}
                rowKey={(record) => record.grantId}
                onSelectionChange={onSelectionChange}
              />
            </Col>
            {
              sidebarShown
                ? (
                  <Col span={sideRowSpan} className='col-sidebar'>
                    <Sidebar />
                  </Col>
                )
                : null
            }
          </Row>
        </TabPane>
        <TabPane tab={tab1Title} key='2' style={{ paddingLeft: '8px' }}>
          <Row className='with-sidebar'>
            <Col span={dataRowSpan} className='col-table'>
              <BaseTable
                loading={loading}
                chain={chain}
                dataSource={coreData}
                columns={columns}
                rowKey={(record) => record.grantId}
                onSelectionChange={onSelectionChange}
              />
            </Col>
            {
              sidebarShown
                ? (
                  <Col span={sideRowSpan} className='col-sidebar'>
                    <Sidebar />
                  </Col>
                )
                : null
            }
          </Row>
        </TabPane>
        <TabPane tab={'Data Sets'} key='3' style={{ paddingLeft: '8px' }}>
          <Downloads chain={chain} />
        </TabPane>
        <TabPane tab='Charts' key='4' style={{ paddingLeft: '8px' }}>
          <img
            width='800px'
            alt='Unclaimed'
            src={chartsPath + 'Unclaimed%20Match%20Round%208.png'}
          />
          <br />
          <br />
          <img width='800px' alt='Count By Date' src={chartsPath + 'Counts.png'} />
        </TabPane>
        <TabPane tab={'Future Work'} key='5' style={{ paddingLeft: '8px' }}>
          <ToDo />
        </TabPane>
      </Tabs>
      <LastUpdate />
    </Content>
  );
};
