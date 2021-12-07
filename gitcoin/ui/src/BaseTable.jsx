import { Table } from 'antd';
import React, { useRef } from 'react';
import { useKeyNav } from './useKeyNav';

const pageSize = 10;
const selectedRowStyle = {
  color: 'darkblue',
  backgroundColor: 'antiquewhite'
};

export function BaseTable(tableParams) {
  const tableWrapper = useRef();
  const { page, row, keyNavOn } = useKeyNav({
    pageSize,
    maxItems: tableParams.dataSource.length - 1
  });

  return (
    <div
      ref={tableWrapper}
    >
      <Table
        className='main-rows'
        size='small'
        bordered={true}
        pagination={{
          size: 'small',
          position: ['topRight', 'none'],
          hideOnSinglePage: true,
          showSizeChanger: false,
          showTotal: (total, range) => '(' + total + ' grants) ',
          current: page
        }}
        onRow={(record, rowIndex) => {
          return {
            style: keyNavOn && rowIndex === row ? selectedRowStyle : {},
          };
        }}
        {...tableParams}
      />
      </div>
  );
}