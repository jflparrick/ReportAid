# ReportAid

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](/docs/prs.md) [![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](/docs/COPYING.txt)

This is the repository of **ReportAid** - a tool for reporting on humanitarian aid. **ReportAid** is a blockchain implementation of the [IATI Standard](https://iatistandard.org/en/).

The idea for **ReportAid** originally came via a [University of Sussex](https://www.sussex.ac.uk/) Masters Student studying within the university's Science Policy Research Unit (SPRU), who wanted to explore humanitarian blockchains as a means of entering SPRU's [2018 Science, Technology and Innovation Policy Challenge](http://www.sussex.ac.uk/spru/newsandevents/2018/awards/sti-challenge). The original intention of that entry was to use blockchains as a means of delivering trustworthy aid reports, and the [first draft of the policy proposal](docs/SPRUChallenge/SPRUSTIPolicyProposal.md) reflected that intention. However, after reading that first draft, the Masters Student decided to change tack, as she wanted to present blockchains as a novel means of actually providing finance during a humanitarian crisis; in other words, she wanted to leverage the cryptocurrency capabilities of blockchains directly. However, while that is a laudable aim, it was not something the original [maintainer](#maintainer) of *ReportAid*, thought viable then. Instead, he could see the immediate benefit of a blockchain-based tool that could add trust to aid reports. The result is **ReportAid**, a blockchain-based proof of concept for reporting on humanitarian aid.

The intention is to explain the ideas behind **ReportAid** in an academic paper with a proposed title _Humanitarian Aid - a Blockchain Application That Adds Trust to Aid Provisioning_. The current version of the [abstract](docs/abstract.md) gives more detail.

## Table of Contents

- [Usage](#usage)
- [Demo](#demo)
  - [Demo Dependencies](#demo-dependencies)  
- [Built Using](#built-using)  
- [Install](#install)
  - [Getting Started](#getting-started)
  - [Dependencies](#dependencies)    
  - [Detailed Instructions](#detailed-instructions)
- [Maintainer](#maintainer)
- [Contributing](#contributing)
- [License](#license)

## Usage

Below describes a demo' of **ReportAid**.

## Demo

Before you can see the demo , you must install the [demo dependencies](#demo-dependencies).

A live demo' of [ReportAid](http://4b1bdf7b0f6beeadab5dadaf019cddbc94f618792ea30b8a2f5d957267d5bd92/) is running on the [Rinkeby Ethereum Testnet](https://www.rinkeby.io/), via the [dat://](https://dat.foundation/) peer-to-peer network. It is available at [http://4b1bdf7b0f6beeadab5dadaf019cddbc94f618792ea30b8a2f5d957267d5bd92/](http://4b1bdf7b0f6beeadab5dadaf019cddbc94f618792ea30b8a2f5d957267d5bd92/).

[ReportAid](http://4b1bdf7b0f6beeadab5dadaf019cddbc94f618792ea30b8a2f5d957267d5bd92/) is an early proof of concept, so apologies in advance - help for using the app' is currently 'minimal', at best. You should expect some bugs, too (if you find any, feel free to raise an [issue](https://github.com/glowkeeper/ReportAid/issues)). Furthermore, the app' is missing some functionality as it does not, currently, implement all of the [IATI Standard](https://iatistandard.org/en/). That said, as a proof of concept, the app' serves its purpose well. Hopefully, you agree and want to get involved - if so, please get in touch with s dot huckle at sussex dot ac dot uk.

### Demo Dependencies

To see the demo' (and use it), you will need to be running [Firefox](https://www.mozilla.org/) with the [DAT P2P Protocol](https://addons.mozilla.org/en-GB/firefox/addon/dat-p2p-protocol/) and [MetaMask](https://metamask.io/) extensions installed. [MetaMask](https://metamask.io/) should be pointing at the Rinkeby Test Network, and you will need a few test Ether in your [MetaMask](https://metamask.io/) wallet, which you can get from the [Rinkeby Faucet](https://faucet.rinkeby.io/). Those test Ether will allow you to sign the transactions necessary to create [IATI Standard](https://iatistandard.org/en/) records that update the blockchain.

## Built Using

- [node](https://nodejs.org/en/)
- [npm](https://www.npmjs.com/)
- [Ganache](https://github.com/trufflesuite/ganache)
- [Truffle](https://github.com/trufflesuite/truffle)
- [React](https://reactjs.org/) + [Redux](https://redux.js.org/) + [TypeScript](https://www.typescriptlang.org/)
- [MetaMask](https://metamask.io/)

## Install

The instruction below enable you to run **ReportAid** on a local, private, Ethereum test network (via [Ganache](https://github.com/trufflesuite/ganache)). Before following the instructions below, please install the [dependencies](#dependencies).

Follow the instructions in the [Ganache](https://github.com/trufflesuite/ganache) repository for downloading and installing Ganache; tl;dr - you need to clone the [Ganache](https://github.com/trufflesuite/ganache) repository, then run `npm install && npm start`.

In the **ReportAid** repository's [/app](/app) directory, type `npm install`. That should install everything listed in [package.json](/app/package.json), which form the components of the REACT-based web frontend to this application.

Now, publish the contracts to your local blockchain (via [Ganache](https://github.com/trufflesuite/ganache)):

1. Change directory to the [Ganache](https://github.com/trufflesuite/ganache) repository.
2. Start [Ganache](https://github.com/trufflesuite/ganache) by typing `npm start`.
3. Ensure [Ganache](https://github.com/trufflesuite/ganache) is running on [http://localhost:8545](http://localhost:8545) (you may need to change its settings to do so).
4. Change to the **ReportAid** repository's [/app](/app) directory.
5. Build the **ReportAid** smart contracts by running `npm run generate`.
6. Deploy the **ReportAid** smart contracts by running `npm run develMigrate`.
6. Edit the **ReportAid** config file [/app/src/app/utils/config.ts](/app/src/app/utils/config.ts) so that the contract static variables contain the addresses generated by `npm run develMigrate`, above.

Now create the web application:

1. Build the REACT frontend by typing `npm run compile` (you can also watch for any changes by running `npm run watch`).
2. Startup an instance of [http-server](https://www.npmjs.com/package/http-server) by typing `npm run start`.

Then fire up a [Firefox](https://www.mozilla.org/) browser and go to the URL [http://localhost:8081](http://localhost:8081). You must have the [Firefox](https://www.mozilla.org/) extension [MetaMask](https://metamask.io/) installed. [Ganache](https://github.com/trufflesuite/ganache) will have generated some test addresses with 100 test Ether that will allow you to sign the transactions necessary to create [IATI Standard](https://iatistandard.org/en/) records that update the blockchain. You need to import one of those addresses' private keys into [MetaMask](https://metamask.io/).

### Dependencies

After cloning this repository, download and install the dependencies (if you have not already done so):

- [node](https://nodejs.org/en/)
- [npm](https://www.npmjs.com/)
- [Ganache](https://github.com/trufflesuite/ganache)
- [Truffle](https://github.com/trufflesuite/truffle)
- [http-server](https://www.npmjs.com/package/http-server)
- [MetaMask](https://metamask.io/)

## Maintainer

[Steve Huckle](https://glowkeeper.github.io/).

## Contributing

Contributions welcome - please email [Steve Huckle](https://glowkeeper.github.io/).

## License

GNU General Public License v3.0

Please refer to the file: [COPYING](/docs/COPYING.txt) for the full text.
