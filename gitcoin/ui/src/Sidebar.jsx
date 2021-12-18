import { CloseOutlined } from '@ant-design/icons';
import { Descriptions, Button, Spin, Card } from 'antd';
import React, { useState, useMemo, useEffect, useCallback } from 'react';

import { useGlobalState } from './GlobalState';

// We will use this value as base for URLs of all PNGs
const svgSourceBaseUrl = 'https://www.tokenomics.io/gitcoin/data/neighbors/images/pngs/';

export function Sidebar() {
  const {
    selectedGrant,
    setSidebarVisible,
  } = useGlobalState();
  // Anything we need from the selected grant
  const {
    name,
    address,
    appearanceCount,
    logCount,
    neighborCount
  } = useMemo(() => selectedGrant || {}, [selectedGrant]);
  // Construct the selected grant image's URL
  const imageSource = useMemo(() => new URL(`${address}.png`, svgSourceBaseUrl), [address]);
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
            src={imageSource}
            alt={address}
            className='graph-image'
            onLoad={() => setLoading(false)}
            onError={() => setLoading(false)}
          />
        </Spin>
      </div>
    </Card>
  );
}