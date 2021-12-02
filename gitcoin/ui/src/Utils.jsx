import { DownloadOutlined } from '@ant-design/icons';

export const DownloadIcon = ({ address, count, path }) => {
  return (
    <div style={{ display: 'grid', gridTemplateColumns: '1fr' }}>
      <pre style={{ margin: '0px' }}>
        <small>({count})</small>
        <br />
        <small>
          <a target={'blank'} href={'https://tokenomics.io/gitcoin/data/' + path + address + '.csv'}>
            <DownloadOutlined /> {"csv"}
          </a>
        </small>
      </pre>
    </div>
  );
}
