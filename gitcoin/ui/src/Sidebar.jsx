import { CloseOutlined } from '@ant-design/icons';
import { Descriptions, Button, Spin, Card } from 'antd';
import React, { useState, useMemo, useEffect, useCallback } from 'react';

import { useGlobalState } from './GlobalState';
import { config } from './Config';

export function Sidebar() {
  const {
    selectedGrant,
    setSidebarVisible,
  } = useGlobalState();
  // Anything we need from the selected grant
  const {
    name,
    address,
    curChain,
    appearanceCount,
    transactionCount,
    logCount,
    neighborCount
  } = useMemo(() => selectedGrant || {}, [selectedGrant]);

  var neighborBase = config.Urls.Data + curChain + "/neighbors/"

  // Construct the selected grant image's URL
  const adjSource = useMemo(() => new URL(`${address}.txt`, neighborBase + "adjacencies/", window.location), [address, neighborBase]);
  const svgSource = useMemo(() => new URL(`${address}.svg`, neighborBase + "images/", window.location), [address, neighborBase]);
  const pngSource = useMemo(() => new URL(`${address}.png`, neighborBase + "images/pngs/", window.location), [address, neighborBase]);
  const [loading, setLoading] = useState(false);
  const onCloseClick = useCallback(() => {
    setSidebarVisible(false);
  }, [setSidebarVisible]);

  // Set loading when a different grant is selected
  useEffect(() => {
    if (!address) return;

    // setLoading(true);
  }, [address]);

  const closeButton = (
    <Button
      onClick={onCloseClick}
    >
      <CloseOutlined />
    </Button>
  );

  return (
    <Card title={name} className='sidebar' extra={closeButton}>
      <Descriptions bordered column={2}>
        <Descriptions.Item label="nApps">
          {appearanceCount}
        </Descriptions.Item>
        <Descriptions.Item label="nTxs">
          {transactionCount}
        </Descriptions.Item>
        <Descriptions.Item label="nLogs">
          {logCount}
        </Descriptions.Item>
        <Descriptions.Item label="nNeighbors">
          {neighborCount}
        </Descriptions.Item>
        <Descriptions.Item label="Other">
          {neighborCount}
        </Descriptions.Item>
      </Descriptions>

      <div>
        <Spin spinning={loading}>
          <img
            src={pngSource}
            alt={address}
            className='graph-image'
            onLoad={() => setLoading(false)}
            onError={() => setLoading(false)}
          />
        </Spin>
        <div>[<a href={svgSource} target="blank">zoom</a>] [<a href={adjSource} target="blank">adj</a>]</div>
      </div>
    </Card>
  );
}