export const config = {
  projects: [
    'gitcoin',
    'giveth',
  ],
  defaultProject() {
    return config.projects[0];
  },
  chains: new Map([
    ['gitcoin', ['mainnet']],
    ['giveth', ['mainnet', 'gnosis']],
  ]),
  urls: {
    'data': 'data/',
    'charts': 'charts/',
    'TrueBlocks': 'https://gitcoin.co/grants/184/trueblocks',
  },
  buildPath(project, urlKey) {
    const url = config.urls[urlKey];

    if (!url) throw new Error(`Wrong URL key: ${urlKey}`);

    if (url === 'TrueBlocks') return url;

    return `/${project}/${url}`;
  }
};
