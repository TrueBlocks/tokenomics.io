import React, { useCallback } from 'react';
import { Switch } from 'antd';
import { useGlobalState } from './GlobalState';

import './ViewOptions.css';

export function ViewOptions() {
  const {
    sidebarEnabled,
    setSidebarEnabled,
    localExplorer,
    setLocalExplorer
  } = useGlobalState();

  const onExplorerClick = useCallback((checked) => setLocalExplorer(checked), [setLocalExplorer]);
  const onSidebarClick = useCallback((checked) => setSidebarEnabled(checked), [setSidebarEnabled]);

  return (
    <div className='view-options'>
      <Switch
        checked={localExplorer}
        onChange={onExplorerClick}
      />
      <span onClick={() => onExplorerClick(!localExplorer)}>
        Use local TrueBlocks Explorer
      </span>
      <Switch
        defaultChecked={true}
        checked={sidebarEnabled}
        onChange={onSidebarClick}
      />
      <span onClick={() => onSidebarClick(!sidebarEnabled)}>
        Enable Sidebar
      </span>
    </div>
  );
}