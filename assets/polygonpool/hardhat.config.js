require("dotenv").config();
require("@nomicfoundation/hardhat-toolbox");

module.exports = {
  solidity: "0.8.18",
  networks: {
    zkEVM_testnet: {
      url: process.env.YOUR_CHAINSTACK_ENDPOINT,
      accounts: [process.env.YOUR_PRIVATE_KEY],
    },
  },
};
