import React, { useCallback, useMemo } from 'react';
import { Switch } from 'antd';
import { useGlobalState } from './GlobalState';
import { config } from './Config';

import './ViewOptions.css';

export function ViewOptions() {
  const {
    project,
    sidebarEnabled,
    setSidebarEnabled,
    localExplorer,
    setLocalExplorer,
    showZero,
    setShowZero,
    chain,
    setChain,
  } = useGlobalState();

  const hasGnosis = useMemo(() => config.chains.get(project).includes('gnosis'), [project]);

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
      {hasGnosis
        ? <>
          <Switch
            checked={chain === "gnosis"}
            onChange={onChain}
          />
          <span onClick={() => onChain(chain === "gnosis")}>
            Show gnosis
          </span>
        </>
        : null
      }
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