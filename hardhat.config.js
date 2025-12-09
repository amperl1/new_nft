require("@nomicfoundation/hardhat-toolbox");
require("hardhat-deploy");
require("@openzeppelin/hardhat-upgrades");
require("solidity-coverage")

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
  // namedAccounts: {
  //   deployer: 0,
  //   user1: 1,
  //   user2: 2,
  // },
  networks: {
    sepolia: {
      url:"https://sepolia.infura.io/v3/773ce546239143aeac7f6f41aae96d57",
      accounts: ["a32d1ad9927d1dc5dd0f26af26a476d50bcb085eb0367b359332f7d9891ce626"],
      live: false,
      saveDeployments: true,
      tags: ["deployNFT"]
    }
    
  }
};
