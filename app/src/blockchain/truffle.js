module.exports = {
  networks: {
    development: {
      host: "localhost",
      port: 8545,
      network_id: "*",
      gas: "6721975",
    },
    rinkeby: {
      host: "localhost", // Connect to geth on the specified
      port: 8545,
      from: "0x8f03ca885434522d695735a28d6a8a93b4390da9", // default address to use for any transaction Truffle makes during migrations
      network_id: 4,
      gas: 4612388 // Gas limit used for deploys
    },
    ropsten: {
      host: "localhost", // Connect to geth on the specified
      port: 8545,
      from: "0x79b0e7de13a17a74ab23fd2e6c69aa3cf93f4e1c",
      network_id: 3,
      gas: 4612388      //make sure this gas allocation isn't over 4M, which is the max
    }
  },
  compilers: {
      solc: {
         version: "0.6.6",    // Fetch exact version from solc-bin (default: truffle's version)
        // docker: true,        // Use "0.5.1" you've installed locally with docker (default: false)
        // settings: {          // See the solidity docs for advice about optimization and evmVersion
        //  optimizer: {
        //    enabled: false,
        //    runs: 200
        //  },
        //  evmVersion: "byzantium"
        // }
      }
  }
};
