import { DownloadIcon } from './Utils';
import { DateHeader, DateCell, NameHeader, NameCell, BalanceHeader, AppearanceHeader, TransactionHeader, EventLogsHeader, NeighborsHeader, SuspiciousHeader } from "./ColumnCells"
import { getChainData } from './GlobalState';

export const columns = [
  {
    title: <DateHeader />,
    dataIndex: 'date',
    key: 'date',
    width: '8%',
    showSorterTooltip: false,
    sorter: {
      compare: function (a, b) {
        return getChainData(a).latestAppearance.bn - getChainData(b).latestAppearance.bn;
      },
    },
    render: (u, grantData) => {
      return <DateCell grantData={grantData} />
    },
  },
  {
    title: <NameHeader />,
    dataIndex: 'name',
    key: 'name',
    width: '20%',
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => a.address - b.address,
    },
    render: function (u, grantData) {
      return <NameCell grantData={grantData} />
    },
  },
  {
    title: <BalanceHeader />,
    dataIndex: 'balance',
    key: 'balance',
    align: 'right',
    width: '6%',
    showSorterTooltip: false,
    sorter: {
      compare: function (a, b) {
        var chainDataA = getChainData(a)
        var chainDataB = getChainData(b)
        if (chainDataA.balances.length === 0 || chainDataB.balances.length === 0) {
          return 0;
        }
        return chainDataB.balances[0].balance - chainDataA.balances[0].balance;
      },
    },
    render: function (text, grantData) {
      const chainData = getChainData(grantData)
      if (chainData.balances && chainData.balances.length > 0) {
        return (
          <>
            {chainData.balances[0].balance + ' ETH'}
            < DownloadIcon grantData={grantData} count={''} path='statements/balances/' type='csv' />
          </>
        )
      }
      return <></>
    },
  },
  {
    title: <AppearanceHeader />,
    dataIndex: 'chainData.counts.appearanceCount',
    key: 'chainData.counts.appearanceCount',
    width: '6%',
    align: 'right',
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => {
        return getChainData(b).counts.appearanceCount - getChainData(a).counts.appearanceCount;
      },
    },
    render: function (text, grantData) {
      const chainData = getChainData(grantData)
      return (
        <>
          {chainData.firstAppearance.bn}.{chainData.firstAppearance.txId}
          <DownloadIcon grantData={grantData} count={chainData.counts.appearanceCount} path='apps/' type='csv' />
        </>
      )
    },
  },
  {
    title: <TransactionHeader />,
    dataIndex: 'chainData.counts.transactionCount',
    key: 'chainData.counts.transactionCount',
    width: '6%',
    align: 'right',
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => {
        return getChainData(b).counts.transactionCount - getChainData(a).counts.transactionCount;
      },
    },
    render: function (text, grantData) {
      const chainData = getChainData(grantData)
      return (
        <DownloadIcon grantData={grantData} count={chainData.counts.transactionCount} path='txs/' type='csv' />
      );
    },
  },
  {
    title: <EventLogsHeader />,
    dataIndex: 'chainData.counts.logCount',
    key: 'chainData.counts.logCount',
    width: '6%',
    align: 'right',
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => {
        return getChainData(b).counts.logCount - getChainData(a).counts.logCount;
      },
    },
    render: function (text, grantData) {
      const chainData = getChainData(grantData)
      return (
        <DownloadIcon grantData={grantData} count={chainData.counts.logCount} path='logs/' type='csv' />
      );
    },
  },
  {
    title: <NeighborsHeader />,
    dataIndex: 'chainData.counts.neighborCount',
    key: 'chainData.counts.neighborCount',
    width: '6%',
    align: 'right',
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => {
        return getChainData(b).counts.neighborCount - getChainData(a).counts.neighborCount;
      },
    },
    render: function (text, grantData) {
      const chainData = getChainData(grantData)
      return (
        <DownloadIcon grantData={grantData} count={chainData.counts.neighborCount} path='neighbors/' type='csv' />
      );
    },
  },
  {
    title: <SuspiciousHeader />,
    render(text, record, index) {
      const icon = (index + 1) % 5 === 0
        ? <span>😡</span>
        : null;

      return <div className='column-suspicious'>{icon}</div>
    }
  },
];
