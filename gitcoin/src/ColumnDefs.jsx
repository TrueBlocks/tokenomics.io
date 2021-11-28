import {Popover, Tag, Typography} from 'antd';
import {DownloadIcon} from './Utils';

const { Text } = Typography;

const widths = {
  date: '15%',
  type: '8%',
  name: '21%',
  matched: '6%',
  claimed: '6%',
  balance: '6%',
  tx_cnt: '8%',
  log_cnt: '8%',
};

const blk = (block) => {
  if (block === '0')
    return '';
  return ('(block: ' + block + ')');
}
const ColumnTitle = ({title, tooltip}) => {
  return (
    <div>
      {title}{' '}
      <Popover
        color='lightblue'
        content={
          <div style={{width: '250px', wordWrap: 'break-word'}}>
            {tooltip}
          </div>
        }>
        <Text style={{fontWeight: '800', color: '#006666'}}>?</Text>
      </Popover>
    </div>
  );
}

export const columns = [
  {
    title: (
      <ColumnTitle
        title='Last Activity'
        tooltip='The most recent interaction this address had with a GitCoin-related contract.'
      />
    ),
    dataIndex: 'date',
    key: 'date',
    width: widths['date'],
    render: (text, record) => {
      const dd = (
        <div>
          {text}
          <div>
            <small>
              <i>
                {blk(record.last_block)}
              </i>
            </small>
          </div>
        </div>
      );
      return renderCell(dd);
    },
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => a.last_block - b.last_block,
    },
  },
  {
    title: (
      <ColumnTitle
        title='Type'
        tooltip='The type of data to download. Currently, only logs, but coming soon txs, traces, accounting, etc.'
      />
    ),
    dataIndex: 'type',
    key: 'type',
    width: widths['type'],
    filters: [
      {text: 'Transactions', value: 'txs'},
      {text: 'Logs', value: 'logs'},
      {text: 'Traces', value: 'traces'},
    ],
    onFilter: (value, record) => record.type.includes(value),
    render: (text, record) => (
      <div style={{marginTop: '-20px'}}>
        <pre>
          <br />
          <Tag color='blue' key={record.address}>
            <small>{text}</small>
          </Tag>
        </pre>{' '}
      </div>
    ),
  },
  {
    title: (
      <ColumnTitle
        title='Name'
        tooltip='The name and address of the grant or core contract linked either to the GitCoin grant or Etherscan.'
      />
    ),
    dataIndex: 'name',
    key: 'name',
    width: widths['name'],
    render: function (text, record) {
      var name = !!record.grant_id ? record.name + ' (#' + record.grant_id + ')' : record.name;
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
    },
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => a.address - b.address,
    },
  },
  {
    title: <ColumnTitle title='CLR' tooltip='The match, claimed, and unclaimed amounts for the grant from Round 8. Sorts by unclaimed then match. (Some of the unmatched payouts may have gone through other channels.)' />,
    dataIndex: 'matched',
    key: 'matched',
    align: 'right',
    width: widths['matched'],
    render: (text, record) => {
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
      return renderCell(tt);
    },
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
  },
  {
    title: (
      <ColumnTitle
        title='Balances'
        tooltip='The balances for the account in ETH. We will add DAI and other tokens later'
      />
    ),
    dataIndex: 'balance',
    key: 'balance',
    align: 'right',
    width: widths['balance'],
    render: (text, record) => {
      return renderCell(record.balances[0].balance + ' ETH');
    },
    showSorterTooltip: false,
    sorter: {
      compare: function (a, b) {
        return b.balances[0].balance - a.balances[0].balance;
      },
    },
  },
  {
    title: (
      <ColumnTitle
        title='Appearances'
        tooltip='The total number of appearances for this address. (An appearance is any transaction the account has ever appeared in including internal txs.)'
      />
    ),
    dataIndex: 'tx_cnt',
    key: 'tx_cnt',
    width: widths['tx_cnt'],
    align: 'right',
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => {
        return b.tx_cnt - a.tx_cnt;
      },
    },
    render: function (text, record) {
      return <DownloadIcon record={record} extra='apps/' type='txt' />
    },
  },
  {
    title: (
      <ColumnTitle title='Event Logs' tooltip='The number of GitCoin related logs in which this address appears.' />
    ),
    dataIndex: 'log_cnt',
    key: 'log_cnt',
    width: widths['log_cnt'],
    align: 'right',
    showSorterTooltip: false,
    sorter: {
      compare: (a, b) => {
        return b.log_cnt - a.log_cnt;
      },
    },
    render: function (text, record) {
      return <DownloadIcon record={record} extra='' type='csv' />;
    },
  },
];

const renderCell = (text) => {
  return (
    <div>
      <pre>
        <small>{text}</small>
      </pre>
    </div>
  );
};
