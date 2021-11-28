import React from "react"

const fq1 = 'What is this website?';
const fa1 = (
  <div>
    {`This websites gives access to two datasets. First, every log on any of the GitCoin Grant smart
    contracts. And second, corresponding logs for the individual grant. The website is an experiment
    in live monitoring and presenting chain data for a large collection of related addresses at a
    minimal cost. The permissionless nature of the data and the fact that we run locally on desktop
    computers allows us to provide this data for free.`}
  </div>
);

const fq2 = 'May I download the entire data set in a single file?';
const fa2 = (
  <div>
    {`All donations and payouts are available by downloading the data from the GitCoin Grant contracts on the `}<i>{`Dontation Contracts`}</i>{` tab.`}
  </div>
);

const fq2a = 'I\'ve gotten more donations than you\'re showing. Why is that?';
const fa2a = (
  <div>
    {`We're showing donations that appear on the Ethereum mainnet. This does not include donations
    sent through zkSync. Once the funds being donated on zkSync are brought back to the grant,
    they will appear here.`}
  </div>
);

const fq3 = 'What does the TrueBlocks infrastructure look like?';
const fa3 = (
  <div>
    {`Currently, we run two OpenEthereum archive nodes in-house. About five years ago, we
    spent $4,500 US to buy two "hefty" Linux computers. Subsequently, we've upgraded twice
    to increase the hard-drive space. (We're now running Raid 0, 12TB SSD's on both boxes.)
    We also spend about $30.00 US per month on this web server.`}
  </div>
);

const fq4 = 'Why do you only export event logs?';
const fa4 = (
  <div>
    {`Users are most familiar with that data. The full suite of TrueBlocks tools produces significantly
    more than that, however. TrueBlocks produces full tranactional details, full tracing data, and even
    full ETH accounting, but we do not include that on this website. If you're interested, contact join
    our discord to discuss.`}
  </div>
);

const fq5 = 'What\'s so difficult about what you did?';
const fa5 = (
  <ul style={{marginLeft: '-20px'}}>
    <li>The data comes directly from an Ethereum node.</li>
    <li>
      Without TrueBlocks, getting this same data{' '}
      <i>
        <u>directly from a node</u>
      </i>{' '}
      takes WEEKS!
    </li>
    <li>TrueBlocks runs on our local Mac desktop using absolutly no third-party APIs.</li>
    <li>TrueBlocks has no databases to install or maintain. It runs local-first and is very cheap to operate.</li>
    <li>TrueBlocks is perfectly private. There are no API keys to sign up for. You do not have to log in. No-one knows what your're doing.</li>
  </ul>
);

const fq6 = 'Why did you build the data pouch?';
const fa6 = (
  <div>
    {`Five years ago we fell in love with the idea of per-block, 18-decimal-place-accurate, permissionless,
    radically-transparent data. We've been trying to build on that dream ever since. Recently, we've
    expanded the idea into `}
    <i>ecosystem accounting</i>
    {` which means accounting for community-wide
    constellations of inter-related addresses such as the GitCoin grant community. Get in touch with us
    for more information.`}
  </div>
);

const FaqEntry = ({ question, answer }) => {
  return (
    <li>
      <b>{question}</b>
      <br />
      <div style={{width: '650px', wordWrap: 'break-word', paddingBottom: "10px"}}>{answer}</div>
    </li>
  );
};

export const faq_text = (
  <ul style={{marginLeft: '-20px'}}>
    <FaqEntry question={fq1} answer={fa1} />
    <FaqEntry question={fq2} answer={fa2} />
    <FaqEntry question={fq2a} answer={fa2a} />
    <FaqEntry question={fq4} answer={fa4} />
    <FaqEntry question={fq5} answer={fa5} />
    <FaqEntry question={fq3} answer={fa3} />
    <FaqEntry question={fq6} answer={fa6} />
  </ul>
);

export const faq_title = 'FAQ';
