import { DownloadIcon } from './Utils';
import { DateHeader, DateCell, NameHeader, NameCell, BalanceHeader, AppearanceHeader, TransactionHeader, EventLogsHeader, NeighborsHeader } from "./ColumnCells"

export const columns = [
  {
    title: <DateHeader />,
    dataIndex: 'date',
    key: 'date',
    width: '8%',
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => a.chainData[0].latestAppearance.bn - b.chainData[0].latestAppearance.bn,
    },
    render: (u, record) => {
      return <DateCell record={record} />
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
    render: function (u, record) {
      return <NameCell record={record} />
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
        return b.chainData[0].balances[0].balance - a.chainData[0].balances[0].balance;
      },
    },
    render: function (text, record) {
      if (record.chainData[0].balances && record.chainData[0].balances.length > 0) {
        return (
          <>
            {record.chainData[0].balances[0].balance + ' ETH'}
            < DownloadIcon address={record.address} count={''} path='statements/balances/' type='csv' />
          </>
        )
      }
      return <></>
    },
  },
  {
    title: <AppearanceHeader />,
    dataIndex: 'chainData[0].counts.appearanceCount',
    key: 'chainData[0].counts.appearanceCount',
    width: '6%',
    align: 'right',
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => {
        return b.chainData[0].counts.appearanceCount - a.chainData[0].counts.appearanceCount;
      },
    },
    render: function (text, record) {
      return (
        <>
          {record.chainData[0].firstAppearance.bn}.{record.chainData[0].firstAppearance.txId}
          <DownloadIcon address={record.address} count={record.chainData[0].counts.appearanceCount} path='apps/' type='csv' />
        </>
      )
    },
  },
  {
    title: <TransactionHeader />,
    dataIndex: 'chainData[0].counts.transactionCount',
    key: 'chainData[0].counts.transactionCount',
    width: '6%',
    align: 'right',
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => {
        return b.chainData[0].counts.transactionCount - a.chainData[0].counts.transactionCount;
      },
    },
    render: function (text, record) {
      return <DownloadIcon address={record.address} count={record.chainData[0].counts.transactionCount} path='txs/' type='csv' />
    },
  },
  {
    title: <EventLogsHeader />,
    dataIndex: 'chainData[0].counts.logCount',
    key: 'chainData[0].counts.logCount',
    width: '6%',
    align: 'right',
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => {
        return b.chainData[0].counts.logCount - a.chainData[0].counts.logCount;
      },
    },
    render: function (text, record) {
      return <DownloadIcon address={record.address} count={record.chainData[0].counts.logCount} path='logs/' type='csv' />;
    },
  },
  {
    title: <NeighborsHeader />,
    dataIndex: 'chainData[0].counts.neighborCount',
    key: 'chainData[0].counts.neighborCount',
    width: '6%',
    align: 'right',
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => {
        return b.chainData[0].counts.neighborCount - a.chainData[0].counts.neighborCount;
      },
    },
    render: function (text, record) {
      return <DownloadIcon address={record.address} count={record.chainData[0].counts.neighborCount} path='neighbors/' type='csv' />;
    },
  },
];
