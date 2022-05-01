import { DownloadOutlined } from '@ant-design/icons';
import { config } from './Config';

export const DownloadIcon = ({ grantData, count, path }) => {
  const cc = count ? <>({count})<br /></> : <></>;
  return (
    <div style={{ display: 'grid', gridTemplateColumns: '1fr' }}>
      <div style={{ margin: '0px' }}>
        {cc}
        <a target={'blank'} href={config.Urls.Data + grantData.curChain + "/" + path + grantData.address + '.csv'}>
          <DownloadOutlined /> {"csv"}
        </a>
      </div>
    </div>
  );
}
