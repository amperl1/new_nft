// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity >0.8.0 <0.9.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";

import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

contract NftAuction is Initializable,OwnableUpgradeable, UUPSUpgradeable, IERC721Receiver {
    mapping(uint256 => Auction) public auctions;
    uint256 auctionId;
    address admin;
    
    mapping(address => AggregatorV3Interface) public priceFeeds; // price feeds for each token

    struct Auction {
        uint256 tokenId;
        bool ended;
        uint256 startPrice;
        address _nftAddress;
        uint highestBid;
        address highestBidder;
        address _tokenAddress;
    }

    function initialize() initializer public {
        admin = msg.sender;
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();
    }
    // 3112.16057600
    function getChainlinkDataFeedLatestAnswer(
        address tokenAddress
    ) public view returns (int) {
        AggregatorV3Interface priceFeed = priceFeeds[tokenAddress];
        // prettier-ignore
        (
            /* uint80 roundId */,
            int256 answer,
            /*uint256 startedAt*/,
            /*uint256 updatedAt*/,
            /*uint80 answeredInRound*/
        ) = priceFeed.latestRoundData();
        return answer;
    }

    function setPriceFeed(
        address tokenAddress,
        address _priceFeed
    ) public {
        priceFeeds[tokenAddress] = AggregatorV3Interface(_priceFeed);
    }
    function getAdmin() public view returns(address) {
        return admin;
    }

    event startAuctionEvent(uint256 _auctionId, uint256 _tokenId, address _nftAddress, uint256 _startPrice);
    function startAuction(uint256 _tokenId, address _nftAddress, uint256 _startPrice) public {
        require(msg.sender == admin, "You are not the owner");
        
        // 转移NFT所有权到合约
        IERC721(_nftAddress).safeTransferFrom(admin, address(this), _tokenId);

        // 设置NFT合约地址
        auctions[auctionId] = Auction(_tokenId, false, _startPrice, _nftAddress,0, address(0), address(0));
        auctionId++;
        emit startAuctionEvent(auctionId, _tokenId, _nftAddress, _startPrice);
    }
    event bidEvent(uint256 _auctionId, uint256 amount, address _tokenAddress);
    function bid(uint256 _auctionId, uint256 amount, address _tokenAddress) public payable {
        Auction storage auction = auctions[_auctionId];
        require(auction.ended == false, "The auction has ended");
        // 获取美金价格，校验价格是否合法
        uint price;
        if(_tokenAddress != address(0)) {
            // ERC20
            price = amount * uint(getChainlinkDataFeedLatestAnswer(_tokenAddress));
        } else {
            // ETH
            amount = msg.value;
            price = amount * uint(getChainlinkDataFeedLatestAnswer(address(0)));
        }
        
        uint startPriceValue = auction.startPrice *
            uint(getChainlinkDataFeedLatestAnswer(auction._tokenAddress));

        uint highestBidValue = auction.highestBid *
            uint(getChainlinkDataFeedLatestAnswer(auction._tokenAddress));
        require(price > highestBidValue && price > startPriceValue, "There is already a higher bid");

        // 转移 ERC20 到合约
        if (_tokenAddress != address(0)) {
            IERC20(_tokenAddress).transferFrom(msg.sender, address(this), amount);
        } 

        // 退还前最高价
        if (auction.highestBid > 0) {
            if (auction._tokenAddress == address(0)) {
                // auction.tokenAddress = _tokenAddress;
                (bool success, ) = payable(auction.highestBidder).call{value: auction.highestBid}("");
                require(success, "Transfer failed");

            } else {
                // 退回之前的ERC20
                IERC20(auction._tokenAddress).transfer(
                    auction.highestBidder,
                    auction.highestBid
                );
            }
        }

        auction.highestBidder = msg.sender;
        auction.highestBid = amount;
        auction._tokenAddress = _tokenAddress;
        emit bidEvent(_auctionId, amount, _tokenAddress);
    }
    event endAuctionEvent(uint256 _auctionId);
    function endAuction(uint256 _auctionId) public {
        Auction storage auction = auctions[_auctionId];
        require(msg.sender == admin, "You are not the owner");
        require(!auction.ended, "The auction has already ended");
        auction.ended = true;
        // 转移NFT到最高出价者
        IERC721(auction._nftAddress).safeTransferFrom(
            address(this),
            auction.highestBidder,
            auction.tokenId
        );
        // 资金给买家（owner)
        if (auction._tokenAddress == address(0)) {
            (bool success, ) = payable(admin).call{value: auction.highestBid}("");
            require(success, "Transfer failed");

        } else {
            IERC20(auction._tokenAddress).transfer(admin, auction.highestBid);
        }
        emit endAuctionEvent(_auctionId);
    }

     function _authorizeUpgrade(address newImplementation)
        internal override onlyOwner {}
    
    function onERC721Received(
        address operator,
        address from,
        uint256 tokenId,
        bytes calldata data
    ) external override returns (bytes4) {
        // 自定义接收逻辑（如验证数据、更新状态等）
        // ...
        return IERC721Receiver.onERC721Received.selector; // 必须返回此值以确认接收
    }
}