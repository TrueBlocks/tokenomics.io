import { Table } from 'antd';
import React, { useRef, useState } from 'react';
import { useEffect, useMemo } from 'react/cjs/react.development';
import { useKeyNav } from './useKeyNav';

const pageSize = 10;

export function BaseTable(tableParams) {
  const tableWrapper = useRef();
  const { page, row, keyNavOn, selectRow } = useKeyNav({
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
            onClick() {
              selectRow(rowIndex)
            },
            className: keyNavOn && rowIndex === row ? 'selected' : '',
          };
        }}
        {...tableParams}
      />
    </div>
  );
}