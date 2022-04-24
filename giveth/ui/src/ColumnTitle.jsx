import { Popover, Typography } from 'antd';
const { Text } = Typography;

export const ColumnTitle = ({ title, tooltip }) => {
  return (
    <div>
      {title}{' '}
      <Popover
        color='lightblue'
        content={
          <div style={{ width: '250px', wordWrap: 'break-word' }}>
            {tooltip}
          </div>
        }>
        <Text style={{ fontWeight: '800', color: '#006666' }}>?</Text>
      </Popover>
    </div>
  );
}
