import React, { useCallback } from 'react';
import { Switch } from 'antd';
import { useGlobalState } from './GlobalState';

import './ViewOptions.css';

export function ViewOptions() {
  const {
    sidebarEnabled,
    setSidebarEnabled,
    localExplorer,
    setLocalExplorer,
    showZero,
    setShowZero,
    chain,
    setChain,
  } = useGlobalState();

  const onExplorerClick = useCallback((checked) => setLocalExplorer(checked), [setLocalExplorer]);
  const onSidebarClick = useCallback((checked) => setSidebarEnabled(checked), [setSidebarEnabled]);
  const onShowZero = useCallback((checked) => setShowZero(checked), [setShowZero]);
  const onChain = useCallback((checked) => setChain(checked ? "gnosis" : "mainnet"), [setChain]);

  return (
    <div className='view-options'>
      <Switch
        checked={showZero}
        onChange={onShowZero}
      />
      <span onClick={() => onShowZero(!showZero)}>
        Show zeros
      </span>
      <Switch
        checked={chain === "gnosis"}
        onChange={onChain}
      />
      <span onClick={() => onChain(chain === "gnosis")}>
        Show gnosis
      </span>
      <Switch
        checked={localExplorer}
        onChange={onExplorerClick}
      />
      <span onClick={() => onExplorerClick(!localExplorer)}>
        Local explorer
      </span>
      <Switch
        defaultChecked={true}
        checked={sidebarEnabled}
        onChange={onSidebarClick}
      />
      <span onClick={() => onSidebarClick(!sidebarEnabled)}>
        Show sidebar
      </span>
    </div>
  );
}