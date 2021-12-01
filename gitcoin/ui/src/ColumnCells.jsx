import React from 'react';
import { Tag } from 'antd';
import { ColumnTitle } from "./ColumnTitle";

//--------------------------------------------------
export const DateHeader = () => (
  <ColumnTitle
    title='Last Activity'
    tooltip='The most recent interaction this address had with a GitCoin-related contract.'
  />
)
export const DateCell = ({ record }) => {
  const dd = (<div>
    {record.latestAppearance.date}
    < div >
      <small>
        <i>
          {dateDisplay(record.latestAppearance.bn)}
        </i>
      </small>
    </div >
  </div >
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
    <pre>
      <br />
      <Tag color='blue' key={record.address}>
        <small>{record.types}</small>
      </Tag>
    </pre>
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
  var name = !!record.grantId ? record.name + ' (#' + record.grantId + ')' : record.name;
  name = name.replace('&#39;', "'");
  if (!record.slug)
    return (
      <pre>
        <small>{name}</small>
        <br />
        <a target={'top'} href={'https://etherscan.io/address/' + record.address}>
          <small>{record.address}</small>
        </a>
      </pre>
    );
  return (
    <div>
      <pre>
        <a target={'top'} href={record.slug}>
          <small>{name}</small>
        </a>
        <br />
        <a target={'top'} href={'https://etherscan.io/address/' + record.address}>
          <small>{record.address}</small>
        </a>
      </pre>
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
export const EventLogsHeader = () => (
  <ColumnTitle
    title='Event Logs'
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
      <pre>
        <small>{text}</small>
      </pre>
    </div>
  );
};
