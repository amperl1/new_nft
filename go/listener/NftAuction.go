// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package listener

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// NftAuctionMetaData contains all meta data concerning the NftAuction contract.
var NftAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"bidEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"endAuctionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_startPrice\",\"type\":\"uint256\"}],\"name\":\"startAuctionEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ended\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getChainlinkDataFeedLatestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"priceFeeds\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startPrice\",\"type\":\"uint256\"}],\"name\":\"startAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// NftAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use NftAuctionMetaData.ABI instead.
var NftAuctionABI = NftAuctionMetaData.ABI

// NftAuction is an auto generated Go binding around an Ethereum contract.
type NftAuction struct {
	NftAuctionCaller     // Read-only binding to the contract
	NftAuctionTransactor // Write-only binding to the contract
	NftAuctionFilterer   // Log filterer for contract events
}

// NftAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type NftAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NftAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NftAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NftAuctionSession struct {
	Contract     *NftAuction       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NftAuctionCallerSession struct {
	Contract *NftAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NftAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NftAuctionTransactorSession struct {
	Contract     *NftAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NftAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type NftAuctionRaw struct {
	Contract *NftAuction // Generic contract binding to access the raw methods on
}

// NftAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NftAuctionCallerRaw struct {
	Contract *NftAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// NftAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NftAuctionTransactorRaw struct {
	Contract *NftAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNftAuction creates a new instance of NftAuction, bound to a specific deployed contract.
func NewNftAuction(address common.Address, backend bind.ContractBackend) (*NftAuction, error) {
	contract, err := bindNftAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NftAuction{NftAuctionCaller: NftAuctionCaller{contract: contract}, NftAuctionTransactor: NftAuctionTransactor{contract: contract}, NftAuctionFilterer: NftAuctionFilterer{contract: contract}}, nil
}

// NewNftAuctionCaller creates a new read-only instance of NftAuction, bound to a specific deployed contract.
func NewNftAuctionCaller(address common.Address, caller bind.ContractCaller) (*NftAuctionCaller, error) {
	contract, err := bindNftAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NftAuctionCaller{contract: contract}, nil
}

// NewNftAuctionTransactor creates a new write-only instance of NftAuction, bound to a specific deployed contract.
func NewNftAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*NftAuctionTransactor, error) {
	contract, err := bindNftAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NftAuctionTransactor{contract: contract}, nil
}

// NewNftAuctionFilterer creates a new log filterer instance of NftAuction, bound to a specific deployed contract.
func NewNftAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*NftAuctionFilterer, error) {
	contract, err := bindNftAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NftAuctionFilterer{contract: contract}, nil
}

// bindNftAuction binds a generic wrapper to an already deployed contract.
func bindNftAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NftAuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NftAuction *NftAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NftAuction.Contract.NftAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NftAuction *NftAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftAuction.Contract.NftAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NftAuction *NftAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NftAuction.Contract.NftAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NftAuction *NftAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NftAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NftAuction *NftAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NftAuction *NftAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NftAuction.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NftAuction *NftAuctionCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NftAuction.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NftAuction *NftAuctionSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _NftAuction.Contract.UPGRADEINTERFACEVERSION(&_NftAuction.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NftAuction *NftAuctionCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _NftAuction.Contract.UPGRADEINTERFACEVERSION(&_NftAuction.CallOpts)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(uint256 tokenId, bool ended, uint256 startPrice, address _nftAddress, uint256 highestBid, address highestBidder, address _tokenAddress)
func (_NftAuction *NftAuctionCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenId       *big.Int
	Ended         bool
	StartPrice    *big.Int
	NftAddress    common.Address
	HighestBid    *big.Int
	HighestBidder common.Address
	TokenAddress  common.Address
}, error) {
	var out []interface{}
	err := _NftAuction.contract.Call(opts, &out, "auctions", arg0)

	outstruct := new(struct {
		TokenId       *big.Int
		Ended         bool
		StartPrice    *big.Int
		NftAddress    common.Address
		HighestBid    *big.Int
		HighestBidder common.Address
		TokenAddress  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Ended = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.StartPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.NftAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.HighestBid = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HighestBidder = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.TokenAddress = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(uint256 tokenId, bool ended, uint256 startPrice, address _nftAddress, uint256 highestBid, address highestBidder, address _tokenAddress)
func (_NftAuction *NftAuctionSession) Auctions(arg0 *big.Int) (struct {
	TokenId       *big.Int
	Ended         bool
	StartPrice    *big.Int
	NftAddress    common.Address
	HighestBid    *big.Int
	HighestBidder common.Address
	TokenAddress  common.Address
}, error) {
	return _NftAuction.Contract.Auctions(&_NftAuction.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(uint256 tokenId, bool ended, uint256 startPrice, address _nftAddress, uint256 highestBid, address highestBidder, address _tokenAddress)
func (_NftAuction *NftAuctionCallerSession) Auctions(arg0 *big.Int) (struct {
	TokenId       *big.Int
	Ended         bool
	StartPrice    *big.Int
	NftAddress    common.Address
	HighestBid    *big.Int
	HighestBidder common.Address
	TokenAddress  common.Address
}, error) {
	return _NftAuction.Contract.Auctions(&_NftAuction.CallOpts, arg0)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_NftAuction *NftAuctionCaller) GetAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NftAuction.contract.Call(opts, &out, "getAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_NftAuction *NftAuctionSession) GetAdmin() (common.Address, error) {
	return _NftAuction.Contract.GetAdmin(&_NftAuction.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_NftAuction *NftAuctionCallerSession) GetAdmin() (common.Address, error) {
	return _NftAuction.Contract.GetAdmin(&_NftAuction.CallOpts)
}

// GetChainlinkDataFeedLatestAnswer is a free data retrieval call binding the contract method 0xe7078f92.
//
// Solidity: function getChainlinkDataFeedLatestAnswer(address tokenAddress) view returns(int256)
func (_NftAuction *NftAuctionCaller) GetChainlinkDataFeedLatestAnswer(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NftAuction.contract.Call(opts, &out, "getChainlinkDataFeedLatestAnswer", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainlinkDataFeedLatestAnswer is a free data retrieval call binding the contract method 0xe7078f92.
//
// Solidity: function getChainlinkDataFeedLatestAnswer(address tokenAddress) view returns(int256)
func (_NftAuction *NftAuctionSession) GetChainlinkDataFeedLatestAnswer(tokenAddress common.Address) (*big.Int, error) {
	return _NftAuction.Contract.GetChainlinkDataFeedLatestAnswer(&_NftAuction.CallOpts, tokenAddress)
}

// GetChainlinkDataFeedLatestAnswer is a free data retrieval call binding the contract method 0xe7078f92.
//
// Solidity: function getChainlinkDataFeedLatestAnswer(address tokenAddress) view returns(int256)
func (_NftAuction *NftAuctionCallerSession) GetChainlinkDataFeedLatestAnswer(tokenAddress common.Address) (*big.Int, error) {
	return _NftAuction.Contract.GetChainlinkDataFeedLatestAnswer(&_NftAuction.CallOpts, tokenAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NftAuction *NftAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NftAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NftAuction *NftAuctionSession) Owner() (common.Address, error) {
	return _NftAuction.Contract.Owner(&_NftAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NftAuction *NftAuctionCallerSession) Owner() (common.Address, error) {
	return _NftAuction.Contract.Owner(&_NftAuction.CallOpts)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_NftAuction *NftAuctionCaller) PriceFeeds(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _NftAuction.contract.Call(opts, &out, "priceFeeds", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_NftAuction *NftAuctionSession) PriceFeeds(arg0 common.Address) (common.Address, error) {
	return _NftAuction.Contract.PriceFeeds(&_NftAuction.CallOpts, arg0)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_NftAuction *NftAuctionCallerSession) PriceFeeds(arg0 common.Address) (common.Address, error) {
	return _NftAuction.Contract.PriceFeeds(&_NftAuction.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NftAuction *NftAuctionCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NftAuction.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NftAuction *NftAuctionSession) ProxiableUUID() ([32]byte, error) {
	return _NftAuction.Contract.ProxiableUUID(&_NftAuction.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NftAuction *NftAuctionCallerSession) ProxiableUUID() ([32]byte, error) {
	return _NftAuction.Contract.ProxiableUUID(&_NftAuction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x742f0a90.
//
// Solidity: function bid(uint256 _auctionId, uint256 amount, address _tokenAddress) payable returns()
func (_NftAuction *NftAuctionTransactor) Bid(opts *bind.TransactOpts, _auctionId *big.Int, amount *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _NftAuction.contract.Transact(opts, "bid", _auctionId, amount, _tokenAddress)
}

// Bid is a paid mutator transaction binding the contract method 0x742f0a90.
//
// Solidity: function bid(uint256 _auctionId, uint256 amount, address _tokenAddress) payable returns()
func (_NftAuction *NftAuctionSession) Bid(_auctionId *big.Int, amount *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _NftAuction.Contract.Bid(&_NftAuction.TransactOpts, _auctionId, amount, _tokenAddress)
}

// Bid is a paid mutator transaction binding the contract method 0x742f0a90.
//
// Solidity: function bid(uint256 _auctionId, uint256 amount, address _tokenAddress) payable returns()
func (_NftAuction *NftAuctionTransactorSession) Bid(_auctionId *big.Int, amount *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _NftAuction.Contract.Bid(&_NftAuction.TransactOpts, _auctionId, amount, _tokenAddress)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 _auctionId) returns()
func (_NftAuction *NftAuctionTransactor) EndAuction(opts *bind.TransactOpts, _auctionId *big.Int) (*types.Transaction, error) {
	return _NftAuction.contract.Transact(opts, "endAuction", _auctionId)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 _auctionId) returns()
func (_NftAuction *NftAuctionSession) EndAuction(_auctionId *big.Int) (*types.Transaction, error) {
	return _NftAuction.Contract.EndAuction(&_NftAuction.TransactOpts, _auctionId)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 _auctionId) returns()
func (_NftAuction *NftAuctionTransactorSession) EndAuction(_auctionId *big.Int) (*types.Transaction, error) {
	return _NftAuction.Contract.EndAuction(&_NftAuction.TransactOpts, _auctionId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_NftAuction *NftAuctionTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftAuction.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_NftAuction *NftAuctionSession) Initialize() (*types.Transaction, error) {
	return _NftAuction.Contract.Initialize(&_NftAuction.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_NftAuction *NftAuctionTransactorSession) Initialize() (*types.Transaction, error) {
	return _NftAuction.Contract.Initialize(&_NftAuction.TransactOpts)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_NftAuction *NftAuctionTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _NftAuction.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_NftAuction *NftAuctionSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _NftAuction.Contract.OnERC721Received(&_NftAuction.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_NftAuction *NftAuctionTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _NftAuction.Contract.OnERC721Received(&_NftAuction.TransactOpts, operator, from, tokenId, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NftAuction *NftAuctionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftAuction.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NftAuction *NftAuctionSession) RenounceOwnership() (*types.Transaction, error) {
	return _NftAuction.Contract.RenounceOwnership(&_NftAuction.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NftAuction *NftAuctionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NftAuction.Contract.RenounceOwnership(&_NftAuction.TransactOpts)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address tokenAddress, address _priceFeed) returns()
func (_NftAuction *NftAuctionTransactor) SetPriceFeed(opts *bind.TransactOpts, tokenAddress common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _NftAuction.contract.Transact(opts, "setPriceFeed", tokenAddress, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address tokenAddress, address _priceFeed) returns()
func (_NftAuction *NftAuctionSession) SetPriceFeed(tokenAddress common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _NftAuction.Contract.SetPriceFeed(&_NftAuction.TransactOpts, tokenAddress, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address tokenAddress, address _priceFeed) returns()
func (_NftAuction *NftAuctionTransactorSession) SetPriceFeed(tokenAddress common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _NftAuction.Contract.SetPriceFeed(&_NftAuction.TransactOpts, tokenAddress, _priceFeed)
}

// StartAuction is a paid mutator transaction binding the contract method 0x83c26e72.
//
// Solidity: function startAuction(uint256 _tokenId, address _nftAddress, uint256 _startPrice) returns()
func (_NftAuction *NftAuctionTransactor) StartAuction(opts *bind.TransactOpts, _tokenId *big.Int, _nftAddress common.Address, _startPrice *big.Int) (*types.Transaction, error) {
	return _NftAuction.contract.Transact(opts, "startAuction", _tokenId, _nftAddress, _startPrice)
}

// StartAuction is a paid mutator transaction binding the contract method 0x83c26e72.
//
// Solidity: function startAuction(uint256 _tokenId, address _nftAddress, uint256 _startPrice) returns()
func (_NftAuction *NftAuctionSession) StartAuction(_tokenId *big.Int, _nftAddress common.Address, _startPrice *big.Int) (*types.Transaction, error) {
	return _NftAuction.Contract.StartAuction(&_NftAuction.TransactOpts, _tokenId, _nftAddress, _startPrice)
}

// StartAuction is a paid mutator transaction binding the contract method 0x83c26e72.
//
// Solidity: function startAuction(uint256 _tokenId, address _nftAddress, uint256 _startPrice) returns()
func (_NftAuction *NftAuctionTransactorSession) StartAuction(_tokenId *big.Int, _nftAddress common.Address, _startPrice *big.Int) (*types.Transaction, error) {
	return _NftAuction.Contract.StartAuction(&_NftAuction.TransactOpts, _tokenId, _nftAddress, _startPrice)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NftAuction *NftAuctionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NftAuction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NftAuction *NftAuctionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NftAuction.Contract.TransferOwnership(&_NftAuction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NftAuction *NftAuctionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NftAuction.Contract.TransferOwnership(&_NftAuction.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NftAuction *NftAuctionTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NftAuction.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NftAuction *NftAuctionSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NftAuction.Contract.UpgradeToAndCall(&_NftAuction.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NftAuction *NftAuctionTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NftAuction.Contract.UpgradeToAndCall(&_NftAuction.TransactOpts, newImplementation, data)
}

// NftAuctionInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NftAuction contract.
type NftAuctionInitializedIterator struct {
	Event *NftAuctionInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftAuctionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftAuctionInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftAuctionInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftAuctionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftAuctionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftAuctionInitialized represents a Initialized event raised by the NftAuction contract.
type NftAuctionInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NftAuction *NftAuctionFilterer) FilterInitialized(opts *bind.FilterOpts) (*NftAuctionInitializedIterator, error) {

	logs, sub, err := _NftAuction.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NftAuctionInitializedIterator{contract: _NftAuction.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NftAuction *NftAuctionFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NftAuctionInitialized) (event.Subscription, error) {

	logs, sub, err := _NftAuction.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftAuctionInitialized)
				if err := _NftAuction.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NftAuction *NftAuctionFilterer) ParseInitialized(log types.Log) (*NftAuctionInitialized, error) {
	event := new(NftAuctionInitialized)
	if err := _NftAuction.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftAuctionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NftAuction contract.
type NftAuctionOwnershipTransferredIterator struct {
	Event *NftAuctionOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftAuctionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftAuctionOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftAuctionOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftAuctionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftAuctionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftAuctionOwnershipTransferred represents a OwnershipTransferred event raised by the NftAuction contract.
type NftAuctionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NftAuction *NftAuctionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NftAuctionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NftAuction.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NftAuctionOwnershipTransferredIterator{contract: _NftAuction.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NftAuction *NftAuctionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NftAuctionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NftAuction.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftAuctionOwnershipTransferred)
				if err := _NftAuction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NftAuction *NftAuctionFilterer) ParseOwnershipTransferred(log types.Log) (*NftAuctionOwnershipTransferred, error) {
	event := new(NftAuctionOwnershipTransferred)
	if err := _NftAuction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftAuctionUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the NftAuction contract.
type NftAuctionUpgradedIterator struct {
	Event *NftAuctionUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftAuctionUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftAuctionUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftAuctionUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftAuctionUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftAuctionUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftAuctionUpgraded represents a Upgraded event raised by the NftAuction contract.
type NftAuctionUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NftAuction *NftAuctionFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NftAuctionUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NftAuction.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NftAuctionUpgradedIterator{contract: _NftAuction.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NftAuction *NftAuctionFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NftAuctionUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NftAuction.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftAuctionUpgraded)
				if err := _NftAuction.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NftAuction *NftAuctionFilterer) ParseUpgraded(log types.Log) (*NftAuctionUpgraded, error) {
	event := new(NftAuctionUpgraded)
	if err := _NftAuction.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftAuctionBidEventIterator is returned from FilterBidEvent and is used to iterate over the raw logs and unpacked data for BidEvent events raised by the NftAuction contract.
type NftAuctionBidEventIterator struct {
	Event *NftAuctionBidEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftAuctionBidEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftAuctionBidEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftAuctionBidEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftAuctionBidEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftAuctionBidEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftAuctionBidEvent represents a BidEvent event raised by the NftAuction contract.
type NftAuctionBidEvent struct {
	AuctionId    *big.Int
	Amount       *big.Int
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBidEvent is a free log retrieval operation binding the contract event 0x520ecd486fe67e5832852e39ae9ac29fc6d00a70ed60ffc44e9a3c954c1af459.
//
// Solidity: event bidEvent(uint256 _auctionId, uint256 amount, address _tokenAddress)
func (_NftAuction *NftAuctionFilterer) FilterBidEvent(opts *bind.FilterOpts) (*NftAuctionBidEventIterator, error) {

	logs, sub, err := _NftAuction.contract.FilterLogs(opts, "bidEvent")
	if err != nil {
		return nil, err
	}
	return &NftAuctionBidEventIterator{contract: _NftAuction.contract, event: "bidEvent", logs: logs, sub: sub}, nil
}

// WatchBidEvent is a free log subscription operation binding the contract event 0x520ecd486fe67e5832852e39ae9ac29fc6d00a70ed60ffc44e9a3c954c1af459.
//
// Solidity: event bidEvent(uint256 _auctionId, uint256 amount, address _tokenAddress)
func (_NftAuction *NftAuctionFilterer) WatchBidEvent(opts *bind.WatchOpts, sink chan<- *NftAuctionBidEvent) (event.Subscription, error) {

	logs, sub, err := _NftAuction.contract.WatchLogs(opts, "bidEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftAuctionBidEvent)
				if err := _NftAuction.contract.UnpackLog(event, "bidEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBidEvent is a log parse operation binding the contract event 0x520ecd486fe67e5832852e39ae9ac29fc6d00a70ed60ffc44e9a3c954c1af459.
//
// Solidity: event bidEvent(uint256 _auctionId, uint256 amount, address _tokenAddress)
func (_NftAuction *NftAuctionFilterer) ParseBidEvent(log types.Log) (*NftAuctionBidEvent, error) {
	event := new(NftAuctionBidEvent)
	if err := _NftAuction.contract.UnpackLog(event, "bidEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftAuctionEndAuctionEventIterator is returned from FilterEndAuctionEvent and is used to iterate over the raw logs and unpacked data for EndAuctionEvent events raised by the NftAuction contract.
type NftAuctionEndAuctionEventIterator struct {
	Event *NftAuctionEndAuctionEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftAuctionEndAuctionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftAuctionEndAuctionEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftAuctionEndAuctionEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftAuctionEndAuctionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftAuctionEndAuctionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftAuctionEndAuctionEvent represents a EndAuctionEvent event raised by the NftAuction contract.
type NftAuctionEndAuctionEvent struct {
	AuctionId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEndAuctionEvent is a free log retrieval operation binding the contract event 0xb0f1f99f7695152e2a0b4e1aff49be2628e041685f1a006309add06de8a19535.
//
// Solidity: event endAuctionEvent(uint256 _auctionId)
func (_NftAuction *NftAuctionFilterer) FilterEndAuctionEvent(opts *bind.FilterOpts) (*NftAuctionEndAuctionEventIterator, error) {

	logs, sub, err := _NftAuction.contract.FilterLogs(opts, "endAuctionEvent")
	if err != nil {
		return nil, err
	}
	return &NftAuctionEndAuctionEventIterator{contract: _NftAuction.contract, event: "endAuctionEvent", logs: logs, sub: sub}, nil
}

// WatchEndAuctionEvent is a free log subscription operation binding the contract event 0xb0f1f99f7695152e2a0b4e1aff49be2628e041685f1a006309add06de8a19535.
//
// Solidity: event endAuctionEvent(uint256 _auctionId)
func (_NftAuction *NftAuctionFilterer) WatchEndAuctionEvent(opts *bind.WatchOpts, sink chan<- *NftAuctionEndAuctionEvent) (event.Subscription, error) {

	logs, sub, err := _NftAuction.contract.WatchLogs(opts, "endAuctionEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftAuctionEndAuctionEvent)
				if err := _NftAuction.contract.UnpackLog(event, "endAuctionEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEndAuctionEvent is a log parse operation binding the contract event 0xb0f1f99f7695152e2a0b4e1aff49be2628e041685f1a006309add06de8a19535.
//
// Solidity: event endAuctionEvent(uint256 _auctionId)
func (_NftAuction *NftAuctionFilterer) ParseEndAuctionEvent(log types.Log) (*NftAuctionEndAuctionEvent, error) {
	event := new(NftAuctionEndAuctionEvent)
	if err := _NftAuction.contract.UnpackLog(event, "endAuctionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftAuctionStartAuctionEventIterator is returned from FilterStartAuctionEvent and is used to iterate over the raw logs and unpacked data for StartAuctionEvent events raised by the NftAuction contract.
type NftAuctionStartAuctionEventIterator struct {
	Event *NftAuctionStartAuctionEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NftAuctionStartAuctionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftAuctionStartAuctionEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftAuctionStartAuctionEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftAuctionStartAuctionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftAuctionStartAuctionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftAuctionStartAuctionEvent represents a StartAuctionEvent event raised by the NftAuction contract.
type NftAuctionStartAuctionEvent struct {
	AuctionId  *big.Int
	TokenId    *big.Int
	NftAddress common.Address
	StartPrice *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStartAuctionEvent is a free log retrieval operation binding the contract event 0xbbcba62c1368f40f8b3d0f6f39f8da38ef6b547143727b31d56cca93b6c80b67.
//
// Solidity: event startAuctionEvent(uint256 _auctionId, uint256 _tokenId, address _nftAddress, uint256 _startPrice)
func (_NftAuction *NftAuctionFilterer) FilterStartAuctionEvent(opts *bind.FilterOpts) (*NftAuctionStartAuctionEventIterator, error) {

	logs, sub, err := _NftAuction.contract.FilterLogs(opts, "startAuctionEvent")
	if err != nil {
		return nil, err
	}
	return &NftAuctionStartAuctionEventIterator{contract: _NftAuction.contract, event: "startAuctionEvent", logs: logs, sub: sub}, nil
}

// WatchStartAuctionEvent is a free log subscription operation binding the contract event 0xbbcba62c1368f40f8b3d0f6f39f8da38ef6b547143727b31d56cca93b6c80b67.
//
// Solidity: event startAuctionEvent(uint256 _auctionId, uint256 _tokenId, address _nftAddress, uint256 _startPrice)
func (_NftAuction *NftAuctionFilterer) WatchStartAuctionEvent(opts *bind.WatchOpts, sink chan<- *NftAuctionStartAuctionEvent) (event.Subscription, error) {

	logs, sub, err := _NftAuction.contract.WatchLogs(opts, "startAuctionEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftAuctionStartAuctionEvent)
				if err := _NftAuction.contract.UnpackLog(event, "startAuctionEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStartAuctionEvent is a log parse operation binding the contract event 0xbbcba62c1368f40f8b3d0f6f39f8da38ef6b547143727b31d56cca93b6c80b67.
//
// Solidity: event startAuctionEvent(uint256 _auctionId, uint256 _tokenId, address _nftAddress, uint256 _startPrice)
func (_NftAuction *NftAuctionFilterer) ParseStartAuctionEvent(log types.Log) (*NftAuctionStartAuctionEvent, error) {
	event := new(NftAuctionStartAuctionEvent)
	if err := _NftAuction.contract.UnpackLog(event, "startAuctionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
