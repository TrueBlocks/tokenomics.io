import { DownloadIcon } from './Utils';
import { DateHeader, DateCell, TagHeader, TagCell, NameHeader, NameCell, MatchedHeader, MatchedCell, BalanceHeader, BalanceCell, AppearanceHeader, TransactionHeader, EventLogsHeader, NeighborsHeader } from "./ColumnCells"

export const columns = [
  {
    title: <DateHeader />,
    dataIndex: 'date',
    key: 'date',
    width: '12%',
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => a.latestAppearance.bn - b.latestAppearance.bn,
    },
    render: (u, record) => {
      return <DateCell record={record} />
    },
  },
  // {
  //   title: <TagHeader />,
  //   dataIndex: 'types',
  //   key: 'types',
  //   width: '8%',
  //   render: (u, record) => {
  //     return (<TagCell record={record} />);
  //   },
  // },
  {
    title: <NameHeader />,
    dataIndex: 'name',
    key: 'name',
    width: '16%',
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => a.address - b.address,
    },
    render: function (u, record) {
      return <NameCell record={record} />
    },
  },
  {
    title: <MatchedHeader />,
    dataIndex: 'matched',
    key: 'matched',
    width: '6%',
    align: 'right',
    showSorterTooltip: false,
    sorter: {
      compare: function (a, b) {
        const unclaimedA = a.matched - a.claimed;
        const unclaimedB = b.matched - b.claimed;
        const diff = unclaimedB - unclaimedA;
        if (diff === 0)
          return (a.matched - b.matched)
        return diff;
      },
    },
    render: (u, record) => {
      return <MatchedCell record={record} />
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
        return b.balances[0].balance - a.balances[0].balance;
      },
    },
    render: (text, record) => {
      return <BalanceCell record={record} />;
    },
  },
  {
    title: <AppearanceHeader />,
    dataIndex: 'appearanceCount',
    key: 'appearanceCount',
    width: '6%',
    align: 'right',
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => {
        return b.appearanceCount - a.appearanceCount;
      },
    },
    render: function (text, record) {
      return <DownloadIcon address={record.address} count={record.appearanceCount} path='apps/' type='csv' />
    },
  },
  {
    title: <TransactionHeader />,
    dataIndex: 'appearanceCount',
    key: 'appearanceCount',
    width: '6%',
    align: 'right',
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => {
        return b.appearanceCount - a.appearanceCount;
      },
    },
    render: function (text, record) {
      return <DownloadIcon address={record.address} count={record.appearanceCount} path='txs/' type='csv' />
    },
  },
  {
    title: <EventLogsHeader />,
    dataIndex: 'logCount',
    key: 'logCount',
    width: '6%',
    align: 'right',
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => {
        return b.logCount - a.logCount;
      },
    },
    render: function (text, record) {
      return <DownloadIcon address={record.address} count={record.logCount} path='logs/' type='csv' />;
    },
  },
  {
    title: <NeighborsHeader />,
    dataIndex: 'logCount',
    key: 'logCount',
    width: '6%',
    align: 'right',
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => {
        return b.logCount - a.logCount;
      },
    },
    render: function (text, record) {
      return <DownloadIcon address={record.address} count={record.neighborCount} path='neighbors/' type='csv' />;
    },
  },
];
