import React, { useEffect, useState } from 'react';
import { config } from './Config';
import { useGlobalState } from "./GlobalState"

export function LastUpdate() {
  const { project } = useGlobalState();
  const [lastUpdate, setLastUpdate] = useState('loading...');

  useEffect(() => {
    (async () => {
      const dataDirectory = config.buildPath(project, 'data');
      const response = await fetch(
        new URL(`${dataDirectory}lastUpdate.json`, window.location)
      );
      const data = await response.json();
      setLastUpdate(data);
    })();
  }, [project]);

  return (
    <i>
      <small>Last updated at block: {lastUpdate[0]} {lastUpdate[1]} UTC</small>
    </i>
    )
}