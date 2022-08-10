import React from 'react';

import './ProjectList.css';

export default function ProjectList() {
    return (
        <main className="project-list">
            <h1>
                Tokenomics.io Website
            </h1>

            <p>
                A website and research platform associated with the <a href="https://trueblocks.io">TrueBlocks</a> project.
            </p>

            <nav>
                <h2>
                    Current Experiments
                </h2>
                <ol>
                    <li>
                        <a href="/gitcoin">
                            GitCoin Grants Data Pouch
                        </a>
                    </li>
                    <li>
                        <a href="/giveth">
                            Giveth Data Pouch
                        </a>
                    </li>
                </ol>
            </nav>
        </main>
    );
}