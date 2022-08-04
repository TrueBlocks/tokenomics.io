import React from 'react';

const Strikeout = ({ text }) => {
  return <div style={{ display: 'inline', textDecoration: 'line-through' }}>{text}</div>;
};

export const ToDo = () => {
  return (
    <div style={{ textAlign: 'left' }}>
      <center>
        <h2>Future Work</h2>
      </center>
      <div style={{ display: 'grid', gridTemplateColumns: '1fr 12fr' }}>
        <div></div>
        <div>
          <b>List of things to do:</b> (<a target="top" href="https://discord.com/invite/RAz6DJ6xkf">help us build!</a>)
          <ul>
            <ToDoItem one="Add balance history charts" two="Charts showing history of balances per grant would be interesting" />
            <ToDoItem one="Make the anaysis extensible" two="The goal of the project is that data scientists can use -- and extend -- the data sets" />
            <ToDoItem one="Explore other ecosystems" two="We wrote this for much more than just this system -- any Ecosystem should have the same needs." />
            <ToDoItem one="Cover expenses" two="While the cost of running this system is very small, it is not zero. Find funding." />
            <ToDoItem one="Provide better filtering" two="Allow users to hide grants with less than X txs, for example." />
            <ToDoItem one="Enable search by name" two="Allow users to search for grants." strike={true} />
            <ToDoItem one="Continuously freshen data" two="As best as possible, freshen data as quickly as possible given size of data." strike={true} />
          </ul>
        </div>
      </div>
      <br />
      <div style={{ display: 'grid', gridTemplateColumns: '1fr 12fr' }}>
        <div></div>
        <div>
          <b>Ideas for other data sets:</b> (<a target="top" href="https://discord.com/invite/RAz6DJ6xkf">help us build!</a>)
          <ul>
            <ToDoItem one="Percent of total unique donors for each grant" two="Assuming X unique individual donors for all grants, this metric would show Y / X, where Y is the number of unique donors for a given grant." />
            <ToDoItem one="Percent of total unique grants donated to by each donor" two="Assuming X unique grants in a round, this metric would show Y / X, where Y is the number of grants a donor donated in the round." />
            <ToDoItem one="Uniq recipients by date" two="A data set showing the historical number of unique donors by date." />
            <ToDoItem one="Donor counts by day" two="A data set showing the total number of active donors per day per round." />
            <ToDoItem one="Recipient counts by day" two="A data set showing the total number of active grants (those being donated to) per day per round." />
            <ToDoItem one="Reciprocal pairs by day" two="Data sets showing the reciprical of much of the above data, where appropriate." />
            <ToDoItem one="Donation amount by bucket" two="Donation amounts bucketed by amount of donation." />
            <ToDoItem one="Comparison of Rounds" two="Per round data for many of the above metrics." />
          </ul>
        </div>
      </div>
    </div>);
};

const ToDoItem = ({ one, two, strike = false }) => {
  const o = strike ? <Strikeout text={one} /> : one
  const t = strike ? <Strikeout text={two} /> : two
  return (
    <>
      <li>{o}
        <div>- {t}</div>
      </li>
    </>
  );
}
