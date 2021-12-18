import React, { useMemo, useState } from 'react';
import { Tag } from 'antd';
import { ColumnTitle } from "./ColumnTitle";
import { CloudDownloadOutlined, CopyTwoTone } from '@ant-design/icons';
import { useGlobalState } from './GlobalState';

//--------------------------------------------------
export const DateHeader = () => (
  <ColumnTitle
    title='Last Activity'
    tooltip='The most recent interaction this address had with a GitCoin-related contract.'
  />
)
export const DateCell = ({ record }) => {
  const dd = (<div>
    {record.latestAppearance.date.substr(0, 16)}
    <div>
      {dateDisplay(record.latestAppearance.bn)}
    </div>
  </div>
  );
  return <Cell2 text={dd} />;
}

//--------------------------------------------------
export const TagHeader = () => (
  <ColumnTitle
    title='Types'
    tooltip='The types of data to download. Includes logs, transactions, and neighbors.'
  />
)
export const TagCell = ({ record }) => {
  return (<div style={{ marginTop: '-20px' }}>
    <div>
      <br />
      <Tag color='blue' key={record.address}>
        {record.types}
      </Tag>
    </div>
  </div>
  );
}

//--------------------------------------------------
export const NameHeader = () => (
  <ColumnTitle
    title='Name'
    tooltip='The name and address of the grant or core contract linked either to the GitCoin grant or Etherscan.'
  />
)
export const NameCell = ({ record }) => {
  const [copied, setCopied] = useState(false);
  const { localExplorer } = useGlobalState();

  let name = !!record.grantId ? record.name + ' (#' + record.grantId + ')' : record.name;
  name = name.replace('&#39;', "'");

  const explorerAddress = localExplorer
    ? "http://localhost:1234/dashboard/accounts?address="
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
export const MatchedHeader = () => (
  <ColumnTitle
    title='CLR'
    tooltip='The match, claimed, and unclaimed amounts for the grant from Round 8. Sorts by unclaimed then match. (Some of the unmatched payouts may have gone through other channels.)'
  />
)
export const MatchedCell = ({ record }) => {
  const diff = record.matched - record.claimed > 0;
  var unclaimed = <div style={{ border: '1px dashed orange' }}>claimed</div>;
  if (diff)
    unclaimed = <div style={{ border: '1px dashed orange' }}>unclaimed</div>;
  if (record.matched === 0)
    unclaimed = <div>-</div>;
  const tt = (
    <div>
      <div>{record.matched + ' DAI'}</div>
      <div>{record.claimed + ' DAI'}</div>
      <div>{unclaimed}</div>
    </div>
  );
  return <Cell2 text={tt} />;
}

//--------------------------------------------------
export const BalanceHeader = () => (
  <ColumnTitle
    title='Balances'
    tooltip='The balances for the account in ETH. We will add DAI and other tokens later'
  />
)
export const BalanceCell = ({ record }) => {
  if (record.balances && record.balances.length > 0) {
    return <Cell2 text={record.balances[0].balance + ' ETH'} />
  }
  return <></>
}

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
    tooltip='The number of GitCoin related logs in which this address appears.'
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
  var link = "https://tokenomics.io/gitcoin/data/zips/" + addr + ".tar.gz";
  return (
    <>
      <a href={link} title="Download zip file">
        <CloudDownloadOutlined />
      </a>
    </>);
}