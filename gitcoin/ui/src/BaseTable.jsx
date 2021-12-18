import { Table } from 'antd';
import React, { useRef, useState, useCallback, useEffect, useMemo } from 'react';
import { GrantDataRenderer } from './GrantDataRenderer';
import { useKeyNav } from './useKeyNav';

const pageSize = 10;

export function BaseTable({ onSelectionChange, ...tableParams }) {
  const tableWrapper = useRef();
  const {
    page,
    row,
    position,
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
  // Propagate selection change to the parent
  const notifySelectionChange = useCallback(
    (selectedItem) => onSelectionChange?.(keyNavOn ? selectedItem : null),
    [keyNavOn, onSelectionChange]
  );
  useEffect(
    () => keyNavOn && notifySelectionChange(tableParams.dataSource[position]),
    [keyNavOn, notifySelectionChange, position, tableParams.dataSource]
  );



  return (
    <div
      ref={tableWrapper}
      onMouseMove={onMouseMove}
    >
      <Table
        className={classes}
        size='small'
        bordered={true}
        expandable={{
          columnWidth: "2%",
          expandedRowRender(record) {
            return <GrantDataRenderer grantData={record} />
          }
        }}
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
              selectRow(rowIndex);
              notifySelectionChange(record);
            },
            className: keyNavOn && rowIndex === row ? 'selected' : '',
          };
        }}
        {...tableParams}
      />
    </div >
  );
}