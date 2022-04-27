import React from 'react';

import {
  CheckCircleTwoTone,
  CloseCircleTwoTone,
  CopyTwoTone,
} from '@ant-design/icons'
import { Button, Tag, Tooltip } from 'antd';

import './GrantDataRenderer.css';

export function GrantDataRenderer({ chain, grantData }) {
  const renderBoolean = (boolean) => boolean
    ? <Tooltip title="Yes"><CheckCircleTwoTone twoToneColor="#52c41a" /></Tooltip>
    : <Tooltip title="No"><CloseCircleTwoTone twoToneColor="#eb2f96" /></Tooltip>
  const renderCopyToClipboard = (textToCopy, label, type = "text") => (
    <Button
      type={type}
      onClick={() => {
        navigator.clipboard.writeText(textToCopy);
      }}
    >
      {label ?? null}
      <CopyTwoTone />
    </Button>
  );

  var chainData = grantData.chainData[0]
  if (!chainData) {
    return <div>Cannot display empty record</div>
  }

  var slug = "https://gitcoin.co/grants/" +
    grantData.grantId +
    "/" +
    grantData.name.replace(/[^a-z0-9 -]/g, '').replace(/\s+/g, '-').replace(/-+/g, '-');

  return (
    <section className='grant-data-renderer'>
      <table>
        <tbody>
          <tr>
            <th title='ID'>
              ID
            </th>
            <td>
              {grantData.grantId}
            </td>
            <th title='Active?'>
              Active?
            </th>
            <td>
              {renderBoolean(grantData.isActive)}
            </td>
            <th title='Core'>
              Core
            </th>
            <td>
              {renderBoolean(grantData.core)}
            </td>
            <th title='Appearance count'>
              Appearance count
            </th>
            <td>
              {chainData.counts.appearanceCount}
            </td>
          </tr>
          <tr>
            <th title='Name'>
              Name
            </th>
            <td colSpan={7}>
              {grantData.name}
            </td>
          </tr>
          <tr>
            <th title='Address'>
              Address
            </th>
            <td colSpan={3}>
              {grantData.address}
              {renderCopyToClipboard(grantData.address)}
            </td>
            <th title='Slug'>
              Slug
            </th>
            <td colSpan={3}>
              {slug
                ? (<>
                  {slug}
                  {renderCopyToClipboard(slug)}
                </>)
                : null
              }
            </td>
          </tr>
          <tr className='appearance-header'>
            <th title='First Appearance' colSpan={8}>
              First Appearance
            </th>
          </tr>
          <tr className='appearance-details'>
            <th title='Block'>
              Block
            </th>
            <td>
              {chainData.firstAppearance.bn}
            </td>
            <th title='Transaction ID'>
              Transaction ID
            </th>
            <td>
              {chainData.firstAppearance.txId}
            </td>
            <th title='Timestamp'>
              Timestamp
            </th>
            <td>
              {chainData.firstAppearance.timestamp}
            </td>
            <th title='Date'>
              Date
            </th>
            <td>
              {chainData.firstAppearance.date}
            </td>
          </tr>
          <tr className='appearance-header'>
            <th title='Latest Appearance' colSpan={8}>
              Latest Appearance
            </th>
          </tr>
          <tr className='appearance-details'>
            <th title='Block'>
              Block
            </th>
            <td>
              {chainData.latestAppearance.bn}
            </td>
            <th title='Transaction ID'>
              Transaction ID
            </th>
            <td>
              {chainData.latestAppearance.txId}
            </td>
            <th title='Timestamp'>
              Timestamp
            </th>
            <td>
              {chainData.latestAppearance.timestamp}
            </td>
            <th title='Date'>
              Date
            </th>
            <td>
              {chainData.latestAppearance.date}
            </td>
          </tr>
          <tr>
            <th title='Last updated'>
              Last updated
            </th>
            <td>
              {grantData.lastUpdated}
            </td>
            <th title='Block range'>
              Block range
            </th>
            <td>
              {chainData.blockRange}
            </td>
            <th title='File size'>
              File size
            </th>
            <td>
              {chainData.fileSize}
            </td>
            <th title='Log count'>
              Log count
            </th>
            <td>
              {chainData.counts.logCount}
            </td>
          </tr>
          <tr>
            <th title='Neighbor count'>
              Neighbor count
            </th>
            <td>
              {chainData.counts.neighborCount}
            </td>
            <th title='Types'>
              Types
            </th>
            <td colSpan={2}>
              {chainData.types.split(',').map((type) => <Tag key={type}>{type}</Tag>)}
            </td>
            <th title='Balances'>
              Balances
            </th>
            <td colSpan={2}>
              <ol>
                {chainData.balances && chainData.balances.map((balance) => <li key={balance.asset}>{balance.asset} {balance.balance}</li>)}
              </ol>
            </td>
          </tr>
        </tbody>
      </table>
      {renderCopyToClipboard(JSON.stringify(grantData), 'Copy as JSON', '')}
    </section>
  );
}