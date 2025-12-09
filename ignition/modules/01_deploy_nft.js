const {deployments, upgrades, ethers} = require("hardhat");
const path = require("path");
const fs = require("fs");

module.exports = async ({getNamedAccounts,deployments}) => {
    const {save} = deployments;
    const {deployer} = await getNamedAccounts();

    console.log("Deployer: ", deployer);
    const NftAuction = await ethers.getContractFactory("NftAuction");

    const nftAuctionProxy = await upgrades.deployProxy(NftAuction, [], {
        initializer: "initialize"
    })

    await nftAuctionProxy.waitForDeployment();
    const proxyAddress = await nftAuctionProxy.getAddress();
    console.log("NFT Auction Proxy: ", proxyAddress);
    const actAddress = await upgrades.erc1967.getImplementationAddress(proxyAddress);
    console.log("NFT Auction: ", actAddress);
    
    const storePath = path.resolve(__dirname, "./.cache/proxyNftAuction.json");
    console.log("Store Path: ", storePath);
    fs.writeFileSync(
        storePath,
        JSON.stringify({
            proxyAddress,
            actAddress,
            abi: NftAuction.interface.format("json"),      
        })
    );
    await save("NftAuctionProxy", {
        abi: NftAuction.interface.format("json"),
        address: proxyAddress,
    })
    // await deploy("NFT", {
    //     from: deployer,
    //     args: ["Hello"],
    //     log: true,
    // })
};

module.exports.tags = ["deployNFT"];