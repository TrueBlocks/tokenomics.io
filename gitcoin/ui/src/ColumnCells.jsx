import React, { useMemo, useState } from 'react';
import { ColumnTitle } from "./ColumnTitle";
import { CloudDownloadOutlined, CopyTwoTone } from '@ant-design/icons';
import { useGlobalState, getChainData } from './GlobalState';
import { config } from './Config';

//--------------------------------------------------
export const DateHeader = () => (
  <ColumnTitle
    title='Last Activity'
    tooltip='The most recent interaction this address had.'
  />
)
export const DateCell = ({ grantData }) => {
  var chainData = getChainData(grantData);
  if (chainData.counts.appearanceCount === 0) {
    return <div>N/A</div>
  }
  const dd = (<div>
    {chainData.latestAppearance.date.substr(0, 16)}
    <div>
      {dateDisplay(chainData.latestAppearance.bn)}
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
export const NameCell = ({ grantData }) => {
  const [copied, setCopied] = useState(false);
  const { localExplorer } = useGlobalState();

  let name = !!grantData.grantId ? grantData.name + ' (#' + grantData.grantId + ')' : grantData.name;
  name = name.replace('&#39;', "'");

  const explorerAddress = localExplorer
    ? "http://localhost:1234/address/"
    : 'http://etherscan.io/address/';

  const explorerLink = useMemo(() => (<>
    <a target={'top'} href={explorerAddress + grantData.address}>
      {grantData.address}
    </a>{' '}
    <CopyTwoTone onClick={() => {
      navigator.clipboard.writeText(grantData.address); setCopied(true); setTimeout(() => {
        setCopied(false)
      }, 500)
    }} />
    <div style={{ display: "inline", fontStyle: "italic", fontSize: "small", color: 'red' }}>{' '}
      {copied ? " *copied*" : ""}
    </div>
  </>), [copied, explorerAddress, grantData.address]);

  var slug = "https://gitcoin.co/grants/" +
    grantData.grantId +
    "/" +
    grantData.name.replace(/[^a-z0-9 -]/g, '').replace(/\s+/g, '-').replace(/-+/g, '-');

  if (!slug)
    return (
      <div>
        {name} <ZipLink grantData={grantData} />
        <br />
        {explorerLink}
      </div>
    );
  return (
    <div>
      <div>
        <a target={'top'} href={slug}>
          {name}
        </a> <ZipLink grantData={grantData} />
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

const ZipLink = ({ grantData }) => {
  var link = config.Urls.Data + grantData.curChain + "/zips/" + grantData.address + ".tar.gz";
  return (
    <>
      <a href={link} title="Download zip file">
        <CloudDownloadOutlined />
      </a>
    </>);
}