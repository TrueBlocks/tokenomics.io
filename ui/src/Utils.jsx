import { DownloadOutlined } from '@ant-design/icons';
import { config } from './Config';
import { useGlobalState } from './GlobalState';

export const DownloadIcon = ({ grantData, count, path }) => {
  const { project } = useGlobalState();
  const cc = count ? <>({count})<br /></> : <></>;
  return (
    <div style={{ display: 'grid', gridTemplateColumns: '1fr' }}>
      <div style={{ margin: '0px' }}>
        {cc}
        <a target={'blank'} href={config.buildPath(project, 'data') + grantData.curChain + "/" + path + grantData.address + '.csv'}>
          <DownloadOutlined /> {"csv"}
        </a>
      </div>
    </div>
  );
}
