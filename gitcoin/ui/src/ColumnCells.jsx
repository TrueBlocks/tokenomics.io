import React, { useMemo, useState } from 'react';
import { Tag } from 'antd';
import { ColumnTitle } from "./ColumnTitle";
import { CloudDownloadOutlined, CopyTwoTone } from '@ant-design/icons';
import { useGlobalState } from './GlobalState';
import { configUrls } from './Config';

//--------------------------------------------------
export const DateHeader = () => (
  <ColumnTitle
    title='Last Activity'
    tooltip='The most recent interaction this address had.'
  />
)
export const DateCell = ({ record }) => {
  if (record.chainData[0].counts.appearanceCount === 0) {
    return <div>N/A</div>
  }
  const dd = (<div>
    {record.chainData[0].latestAppearance.date.substr(0, 16)}
    <div>
      {dateDisplay(record.chainData[0].latestAppearance.bn)}
    </div>
  </div>
  );
  return <Cell2 text={dd} />;
}

//--------------------------------------------------
export const NameHeader = () => (
  <ColumnTitle
    title='Name'
    tooltip='The name and address of the grant or core contract.'
  />
)
export const NameCell = ({ record }) => {
  const [copied, setCopied] = useState(false);
  const { localExplorer } = useGlobalState();

  let name = !!record.grantId ? record.name + ' (#' + record.grantId + ')' : record.name;
  name = name.replace('&#39;', "'");

  const explorerAddress = localExplorer
    ? "http://localhost:1234/address/"
    : 'http://etherscan.io/address/';

  const explorerLink = useMemo(() => (<>
    <a target={'top'} href={explorerAddress + record.address}>
      {record.address}
    </a>{' '}
    <CopyTwoTone onClick={() => {
      navigator.clipboard.writeText(record.address); setCopied(true); setTimeout(() => {
        setCopied(false)
      }, 500)
    }} />
    <div style={{ display: "inline", fontStyle: "italic", fontSize: "small", color: 'red' }}>{' '}
      {copied ? " *copied*" : ""}
    </div>
  </>), [copied, explorerAddress, record.address]);

  if (!record.slug)
    return (
      <div>
        {name} <ZipLink addr={record.address} />
        <br />
        {explorerLink}
      </div>
    );
  return (
    <div>
      <div>
        <a target={'top'} href={record.slug}>
          {name}
        </a> <ZipLink addr={record.address} />
        <br />
        {explorerLink}
      </div>
    </div>
  );
}

//--------------------------------------------------
export const BalanceHeader = () => (
  <ColumnTitle
    title='Balances'
    tooltip='The balances for the account in ETH. We will add DAI and other tokens later'
  />
)

//--------------------------------------------------
export const AppearanceHeader = () => (
  <ColumnTitle
    title='Appearances'
    tooltip='The total number of appearances for this address. (An appearance is any transaction the account has ever appeared in including internal txs.)'
  />
)

//--------------------------------------------------
export const TransactionHeader = () => (
  <ColumnTitle
    title='Transactions'
    tooltip='Every transaction in which this address appears.'
  />
)

//--------------------------------------------------
export const EventLogsHeader = () => (
  <ColumnTitle
    title='Logs'
    tooltip='The number of related logs in which this address appears.'
  />
)

//--------------------------------------------------
export const NeighborsHeader = () => (
  <ColumnTitle
    title='Neighbors'
    tooltip='The list of addresses that appear in the same transactions as this grant.'
  />
)

//--------------------------------------------------
const dateDisplay = (block) => {
  if (block === '0')
    return '';
  return ('(block: ' + block + ')');
}

//--------------------------------------------------
const Cell2 = ({ text }) => {
  return (
    <div>
      <div>
        {text}
      </div>
    </div>
  );
};

const ZipLink = ({ addr }) => {
  var link = configUrls.Zips + addr + ".tar.gz";
  return (
    <>
      <a href={link} title="Download zip file">
        <CloudDownloadOutlined />
      </a>
    </>);
}