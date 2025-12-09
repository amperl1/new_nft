const {ethers,upgrades} = require("hardhat");
const fs = require("fs");
const path = require("path");

module.exports = async function ({getNamedAccounts,deployments}) {
    console.log("upgrade nft auction");
    const {save} = deployments;
    const {deployer} = await getNamedAccounts();
    
    console.log("deployer:",deployer);
    const storePath = path.resolve(__dirname, "./.cache/proxyNftAuction.json");
    const storeData = fs.readFileSync(storePath, "utf-8");
    console.log("storeData:",storeData);
    const {proxyAddress, actAddress, abi} = JSON.parse(storeData);

    const NftAuctionv2 = await ethers.getContractFactory("NFTAuctionv2");

    const nftAuctionProxyv2 = await upgrades.upgradeProxy(proxyAddress, NftAuctionv2);
    await nftAuctionProxyv2.waitForDeployment();
    const proxyAddressV2 = await nftAuctionProxyv2.getAddress();

    // fs.writeFileSync
    await save("NftAuctionV2", {
        abi,
        address: proxyAddressV2
    })
}

module.exports.tags = ["upgradeNftAuction"];