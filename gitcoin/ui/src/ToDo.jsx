import React from 'react';

export const ToDo = () => {
  return (
    <div style={{ textAlign: 'left' }}>
      <center>
        <h2>To Do List</h2>
      </center>
      <div style={{ display: 'grid', gridTemplateColumns: "1fr 4fr 1fr" }}>
        <div></div>
        <div>
          <li>Add Balance Histories</li>
          <li>Make Extensible</li>
          <li>Explore Other Ecosystems</li>
          <li>Cover Expenses</li>
        </div>
        <div></div>
      </div>
    </div>);
};
