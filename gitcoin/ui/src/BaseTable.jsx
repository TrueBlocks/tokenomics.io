import { Table } from 'antd';
import React, { useRef, useState } from 'react';
import { useCallback, useEffect, useMemo } from 'react/cjs/react.development';
import { useKeyNav } from './useKeyNav';

const pageSize = 10;

export function BaseTable(tableParams) {
  const tableWrapper = useRef();
  const {
    page,
    row,
    keyNavOn,
    selectRow,
    setPosition
  } = useKeyNav({
    pageSize,
    maxItems: tableParams.dataSource.length - 1
  });
  const [mouseHoverBlocked, setMouseHoverBlocked] = useState(false);
  // If the hover color is hidden, we want restore it as soon as
  // the user moves their mouse.
  const onMouseMove = useCallback(() => setMouseHoverBlocked(false), []);
  const classes = useMemo(() => [
    'main-rows',
    mouseHoverBlocked ? 'no-hover' : '',
  ].join(' '), [mouseHoverBlocked])
  // We don't want to have hover color visible when using keyboard nav
  useEffect(() => setMouseHoverBlocked(keyNavOn), [keyNavOn, row]);

  return (
    <div
      ref={tableWrapper}
      onMouseMove={onMouseMove}
    >
      <Table
        tableLayout='fixed'
        className={classes}
        size='small'
        bordered={true}
        pagination={{
          size: 'small',
          position: ['topRight', 'none'],
          hideOnSinglePage: true,
          showSizeChanger: false,
          showTotal: (total, range) => '(' + total + ' grants) ',
          current: page,
          onChange(page, pageSize) {
            setPosition((page - 1) * pageSize);
          },
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