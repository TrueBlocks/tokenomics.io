import React from 'react';

import {
  CheckCircleTwoTone,
  CloseCircleTwoTone,
  CopyTwoTone,
} from '@ant-design/icons'
import { Button, Tag } from 'antd';

import './GrantDataRenderer.css';

export function GrantDataRenderer({ grantData }) {
  const renderBoolean = (boolean) => boolean
  ? <CheckCircleTwoTone twoToneColor="#52c41a" />
  : <CloseCircleTwoTone twoToneColor="#eb2f96" />
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
            <th>
              ID
            </th>
            <td>
              {grantData.grantId}
            </td>
            <th>
              Active?
            </th>
            <td>
              {renderBoolean(grantData.active)}
            </td>
            <th>
              Core
            </th>
            <td>
            {renderBoolean(grantData.core)}
            </td>
            <th>
              Appearance count
            </th>
            <td>
              {grantData.appearanceCount}
            </td>
          </tr>
          <tr>
            <th>
              Name
            </th>
            <td colSpan={7}>
              {grantData.name}
            </td>
          </tr>
          <tr>
            <th>
              Address
            </th>
            <td colSpan={3}>
              {grantData.address}
              {renderCopyToClipboard(grantData.address)}
            </td>
            <th>
              Slug
            </th>
            <td colSpan={3}>
              {grantData.slug}
              {renderCopyToClipboard(grantData.slug)}
            </td>
          </tr>
          <tr className='appearance-header'>
            <th colSpan={8}>
              First Appearance
            </th>
          </tr>
          <tr className='appearance-details'>
            <th>
              Block
            </th>
            <td>
              {grantData.firstAppearance.bn}
            </td>
            <th>
              Transaction ID
            </th>
            <td>
              {grantData.firstAppearance.txId}
            </td>
            <th>
              Timestamp
            </th>
            <td>
              {grantData.firstAppearance.timestamp}
            </td>
            <th>
              Date
            </th>
            <td>
              {grantData.firstAppearance.date}
            </td>
          </tr>
          <tr className='appearance-header'>
            <th colSpan={8}>
              Latest Appearance
            </th>
          </tr>
          <tr className='appearance-details'>
            <th>
              Block
            </th>
            <td>
              {grantData.latestAppearance.bn}
            </td>
            <th>
              Transaction ID
            </th>
            <td>
              {grantData.latestAppearance.txId}
            </td>
            <th>
              Timestamp
            </th>
            <td>
              {grantData.latestAppearance.timestamp}
            </td>
            <th>
              Date
            </th>
            <td>
              {grantData.latestAppearance.date}
            </td>
          </tr>
          <tr>
            <th>
              Last updated
            </th>
            <td>
              {grantData.lastUpdated}
            </td>
            <th>
              Block range
            </th>
            <td>
              {grantData.blockRange}
            </td>
            <th>
              File size
            </th>
            <td>
              {grantData.fileSize}
            </td>
            <th>
              Log count
            </th>
            <td>
              {grantData.logCount}
            </td>
          </tr>
          <tr>
            <th>
              Neighbor count
            </th>
            <td>
              {grantData.neighborCount}
            </td>
            <th>
              Types
            </th>
            <td colSpan={5}>
              {grantData.types.split(',').map((type) => <Tag key={type}>{type}</Tag>)}
            </td>
          </tr>
          <tr>
            <th>
              Balances
            </th>
            <td colSpan={3}>
              <ol>
                {grantData.balances.map((balance) => <li key={balance.asset}>{balance.asset} {balance.balance}</li>)}
              </ol>
            </td>
            <th>
              Matched
            </th>
            <td>
              {grantData.matched}
            </td>
            <th>
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