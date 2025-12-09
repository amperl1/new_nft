const {ethers, deployments, upgrades} = require('hardhat');
const { expect } = require("chai")

describe('Test upgrade', async function () {
    it('should deploy', async function () {
        // mock一个代币一个nft
        const [signer, buyer] = await ethers.getSigners()
        const TestERC20 = await ethers.getContractFactory("TestERC20");
        const testERC20 = await TestERC20.deploy();
        await testERC20.waitForDeployment();
        const UsdcAddress = await testERC20.getAddress();
        let tx = await testERC20.connect(signer).transfer(buyer, ethers.parseEther("1000"))
        await tx.wait()
        const aggreagatorV3 = await ethers.getContractFactory("AggreagatorV3")
        const priceFeedEthDeploy = await aggreagatorV3.deploy(ethers.parseEther("10000"))
        const priceFeedEth = await priceFeedEthDeploy.waitForDeployment()
        const priceFeedEthAddress = await priceFeedEth.getAddress()
        console.log("ethFeed: ", priceFeedEthAddress)
        const priceFeedUSDCDeploy = await aggreagatorV3.deploy(ethers.parseEther("1"))
        const priceFeedUSDC = await priceFeedUSDCDeploy.waitForDeployment()
        const priceFeedUSDCAddress = await priceFeedUSDC.getAddress()

        const TestERC721 = await ethers.getContractFactory("TestERC721");
        const testERC721 = await TestERC721.deploy();
        await testERC721.waitForDeployment();
        const testERC721Address = await testERC721.getAddress();
        console.log("testERC721Address::", testERC721Address);

        // mint 10个 NFT
        for (let i = 0; i < 10; i++) {
            await testERC721.mint(signer.address, i + 1);
        }    

        console.log("usdcFeed: ", await priceFeedUSDCAddress)
        //1.部署合约
        const a = await deployments.fixture(["deployNFT"])
        // console.log("a::", a)
        const nftAuctionProxy = await deployments.get("NftAuctionProxy")
        const nftAuction = await ethers.getContractAt("NftAuction", nftAuctionProxy.address)
         // 给代理合约授权
        await testERC721.connect(signer).setApprovalForAll(nftAuctionProxy.address, true);

        console.log("nftAuction::", nftAuction)
        // console.log("nftAuction.address::", nftAuction.address)
        // 初始化
        // nftAuction.initialize()
        //2.调用方法创建拍卖
        const admin = await nftAuction.getAdmin()
        console.log("admin::", admin)
        console.log("signer::", signer.address)
        const tokenId = 1;
        await nftAuction.connect(signer).startAuction(
            tokenId,
            testERC721Address,
            ethers.parseEther("0.01")
        )
        console.log("塞入预言机价格start")
        // 塞入预言机价格
        const token2Usd = [{
            token: ethers.ZeroAddress,
            priceFeed: priceFeedEthAddress
        }, {
            token: UsdcAddress,
            priceFeed: priceFeedUSDCAddress
        }]
        for (let i = 0; i < token2Usd.length; i++) {
            const { token, priceFeed } = token2Usd[i];
            await nftAuction.setPriceFeed(token, priceFeed);
        }
        console.log("塞入预言机价格end")
        // 调用语言机接口塞入ETH=》USD的汇率
        // await nftAuction.setPriceFeed(ethers.ZeroAddress,"0x694AA1769357215DE4FAC081bf1f309aDC325306")
        // 拿塞入的汇率
        // const ethToUsd = await nftAuction.getChainlinkDataFeedLatestAnswer(ethers.ZeroAddress)
        // console.log("ethToUsd::", ethToUsd)
        // 竞价
        // nftAuction.bid(0,0,ethers.ZeroAddress, {value: ethers.parseEther("0.01")})
        console.log("参与竞价start")
        // await testERC721.connect(buyer).approve(nftAuctionProxy.address, tokenId);
        // ETH参与竞价
        // tx = await nftAuction.connect(buyer).bid(0, 0, ethers.ZeroAddress, { value: ethers.parseEther("0.01") });
        // await tx.wait()

        // USDC参与竞价
        tx = await testERC20.connect(buyer).approve(nftAuctionProxy.address, ethers.MaxUint256)
        await tx.wait()
        tx = await nftAuction.connect(buyer).bid(0, ethers.parseEther("101"), UsdcAddress);
        await tx.wait()
        // 结束拍卖
        await nftAuction.connect(signer).endAuction(0)
        const auctionResult = await nftAuction.auctions(0);
        console.log("createAuction success::", auctionResult)
        // 验证结果
        expect(auctionResult.highestBidder).to.equal(buyer.address);
        expect(auctionResult.highestBid).to.equal(ethers.parseEther("101"));

        // 验证 NFT 所有权
        const owner = await testERC721.ownerOf(tokenId);
        console.log("owner::", owner);
        expect(owner).to.equal(buyer.address);

        // const act1 = await upgrades.erc1967.getImplementationAddress(nftAuctionProxy.address)
        // console.log("act1::", act1)
        //3.升级合约
        const b = await deployments.fixture(["upgradeNftAuction"])
        // console.log("b::", b)
        const act2 = await upgrades.erc1967.getImplementationAddress(nftAuctionProxy.address)
        console.log("act2::", act2)
        //4.读取合约的action[0]
        const auction2 = await nftAuction.auctions(0)
        console.log("auction2::", auction2)
    })
})
