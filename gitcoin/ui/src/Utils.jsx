import {DownloadOutlined} from '@ant-design/icons';

export const DownloadIcon = ({ record, extra, type }) => {
  if (extra !== '') {
    return (
      <div style={{display: 'grid', gridTemplateColumns: '1fr'}}>
        <pre style={{margin: '0px'}}>
          <small>({record.tx_cnt})</small>
          <br />
          <small>
            <a target={'blank'} href={'https://tokenomics.io/gitcoin/data/' + extra + record.address + '.' + type}>
              <DownloadOutlined /> {type}
            </a>
          </small>
        </pre>
      </div>
    );
  }

  return (
    <div style={{display: 'grid', gridTemplateColumns: '1fr'}}>
      <pre style={{margin: '0px'}}>
        <small>({record.log_cnt})</small>
        <br />
        <small>
          <a target={'blank'} href={'https://tokenomics.io/gitcoin/data/' + extra + record.address + '.' + type}>
            <DownloadOutlined /> {type}
          </a>
          <br />
          <a target={'blank'} href={'https://tokenomics.io/gitcoin/data/' + extra + record.address + '.json'}>
            <DownloadOutlined /> json
          </a>
        </small>
      </pre>
    </div>
  );
}
