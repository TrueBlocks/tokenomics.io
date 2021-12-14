import React from 'react';

import {
  CheckCircleTwoTone,
  CloseCircleTwoTone,
  CopyTwoTone,
} from '@ant-design/icons'
import { Button, Tag, Tooltip } from 'antd';

import './GrantDataRenderer.css';

export function GrantDataRenderer({ grantData }) {
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

  return (
    <section className='grant-data-renderer'>
      {renderCopyToClipboard(JSON.stringify(grantData), 'Copy as JSON', '')}
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
              {renderBoolean(grantData.active)}
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
              {grantData.appearanceCount}
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
              {grantData.slug}
              {renderCopyToClipboard(grantData.slug)}
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
              {grantData.firstAppearance.bn}
            </td>
            <th title='Transaction ID'>
              Transaction ID
            </th>
            <td>
              {grantData.firstAppearance.txId}
            </td>
            <th title='Timestamp'>
              Timestamp
            </th>
            <td>
              {grantData.firstAppearance.timestamp}
            </td>
            <th title='Date'>
              Date
            </th>
            <td>
              {grantData.firstAppearance.date}
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
              {grantData.latestAppearance.bn}
            </td>
            <th title='Transaction ID'>
              Transaction ID
            </th>
            <td>
              {grantData.latestAppearance.txId}
            </td>
            <th title='Timestamp'>
              Timestamp
            </th>
            <td>
              {grantData.latestAppearance.timestamp}
            </td>
            <th title='Date'>
              Date
            </th>
            <td>
              {grantData.latestAppearance.date}
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
              {grantData.blockRange}
            </td>
            <th title='File size'>
              File size
            </th>
            <td>
              {grantData.fileSize}
            </td>
            <th title='Log count'>
              Log count
            </th>
            <td>
              {grantData.logCount}
            </td>
          </tr>
          <tr>
            <th title='Neighbor count'>
              Neighbor count
            </th>
            <td>
              {grantData.neighborCount}
            </td>
            <th title='Types'>
              Types
            </th>
            <td colSpan={5}>
              {grantData.types.split(',').map((type) => <Tag key={type}>{type}</Tag>)}
            </td>
          </tr>
          <tr>
            <th title='Balances'>
              Balances
            </th>
            <td colSpan={3}>
              <ol>
                {grantData.balances.map((balance) => <li key={balance.asset}>{balance.asset} {balance.balance}</li>)}
              </ol>
            </td>
            <th title='Matched'>
              Matched
            </th>
            <td>
              {grantData.matched}
            </td>
            <th title='Claimed'>
              Claimed
            </th>
            <td>
              {grantData.claimed}
            </td>
          </tr>
        </tbody>
      </table>
    </section>
  );
}