// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bscspendcoin

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
)

// BscspendcoinMetaData contains all meta data concerning the Bscspendcoin contract.
var BscspendcoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUSDTBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"spendUserCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BscspendcoinABI is the input ABI used to generate the binding from.
// Deprecated: Use BscspendcoinMetaData.ABI instead.
var BscspendcoinABI = BscspendcoinMetaData.ABI

// Bscspendcoin is an auto generated Go binding around an Ethereum contract.
type Bscspendcoin struct {
	BscspendcoinCaller     // Read-only binding to the contract
	BscspendcoinTransactor // Write-only binding to the contract
	BscspendcoinFilterer   // Log filterer for contract events
}

// BscspendcoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type BscspendcoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscspendcoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BscspendcoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscspendcoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BscspendcoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscspendcoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BscspendcoinSession struct {
	Contract     *Bscspendcoin     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BscspendcoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BscspendcoinCallerSession struct {
	Contract *BscspendcoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BscspendcoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BscspendcoinTransactorSession struct {
	Contract     *BscspendcoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BscspendcoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type BscspendcoinRaw struct {
	Contract *Bscspendcoin // Generic contract binding to access the raw methods on
}

// BscspendcoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BscspendcoinCallerRaw struct {
	Contract *BscspendcoinCaller // Generic read-only contract binding to access the raw methods on
}

// BscspendcoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BscspendcoinTransactorRaw struct {
	Contract *BscspendcoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBscspendcoin creates a new instance of Bscspendcoin, bound to a specific deployed contract.
func NewBscspendcoin(address common.Address, backend bind.ContractBackend) (*Bscspendcoin, error) {
	contract, err := bindBscspendcoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bscspendcoin{BscspendcoinCaller: BscspendcoinCaller{contract: contract}, BscspendcoinTransactor: BscspendcoinTransactor{contract: contract}, BscspendcoinFilterer: BscspendcoinFilterer{contract: contract}}, nil
}

// NewBscspendcoinCaller creates a new read-only instance of Bscspendcoin, bound to a specific deployed contract.
func NewBscspendcoinCaller(address common.Address, caller bind.ContractCaller) (*BscspendcoinCaller, error) {
	contract, err := bindBscspendcoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BscspendcoinCaller{contract: contract}, nil
}

// NewBscspendcoinTransactor creates a new write-only instance of Bscspendcoin, bound to a specific deployed contract.
func NewBscspendcoinTransactor(address common.Address, transactor bind.ContractTransactor) (*BscspendcoinTransactor, error) {
	contract, err := bindBscspendcoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BscspendcoinTransactor{contract: contract}, nil
}

// NewBscspendcoinFilterer creates a new log filterer instance of Bscspendcoin, bound to a specific deployed contract.
func NewBscspendcoinFilterer(address common.Address, filterer bind.ContractFilterer) (*BscspendcoinFilterer, error) {
	contract, err := bindBscspendcoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BscspendcoinFilterer{contract: contract}, nil
}

// bindBscspendcoin binds a generic wrapper to an already deployed contract.
func bindBscspendcoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BscspendcoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bscspendcoin *BscspendcoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bscspendcoin.Contract.BscspendcoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bscspendcoin *BscspendcoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bscspendcoin.Contract.BscspendcoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bscspendcoin *BscspendcoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bscspendcoin.Contract.BscspendcoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bscspendcoin *BscspendcoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bscspendcoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bscspendcoin *BscspendcoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bscspendcoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bscspendcoin *BscspendcoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bscspendcoin.Contract.contract.Transact(opts, method, params...)
}

// GetUSDTBalance is a free data retrieval call binding the contract method 0xedbb9f34.
//
// Solidity: function getUSDTBalance(address tokenAddress, address user) view returns(uint256)
func (_Bscspendcoin *BscspendcoinCaller) GetUSDTBalance(opts *bind.CallOpts, tokenAddress common.Address, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bscspendcoin.contract.Call(opts, &out, "getUSDTBalance", tokenAddress, user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUSDTBalance is a free data retrieval call binding the contract method 0xedbb9f34.
//
// Solidity: function getUSDTBalance(address tokenAddress, address user) view returns(uint256)
func (_Bscspendcoin *BscspendcoinSession) GetUSDTBalance(tokenAddress common.Address, user common.Address) (*big.Int, error) {
	return _Bscspendcoin.Contract.GetUSDTBalance(&_Bscspendcoin.CallOpts, tokenAddress, user)
}

// GetUSDTBalance is a free data retrieval call binding the contract method 0xedbb9f34.
//
// Solidity: function getUSDTBalance(address tokenAddress, address user) view returns(uint256)
func (_Bscspendcoin *BscspendcoinCallerSession) GetUSDTBalance(tokenAddress common.Address, user common.Address) (*big.Int, error) {
	return _Bscspendcoin.Contract.GetUSDTBalance(&_Bscspendcoin.CallOpts, tokenAddress, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bscspendcoin *BscspendcoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bscspendcoin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bscspendcoin *BscspendcoinSession) Owner() (common.Address, error) {
	return _Bscspendcoin.Contract.Owner(&_Bscspendcoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bscspendcoin *BscspendcoinCallerSession) Owner() (common.Address, error) {
	return _Bscspendcoin.Contract.Owner(&_Bscspendcoin.CallOpts)
}

// SpendUserCoin is a paid mutator transaction binding the contract method 0x49afc972.
//
// Solidity: function spendUserCoin(address tokenAddress, address from, address to, uint256 amount) returns()
func (_Bscspendcoin *BscspendcoinTransactor) SpendUserCoin(opts *bind.TransactOpts, tokenAddress common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bscspendcoin.contract.Transact(opts, "spendUserCoin", tokenAddress, from, to, amount)
}

// SpendUserCoin is a paid mutator transaction binding the contract method 0x49afc972.
//
// Solidity: function spendUserCoin(address tokenAddress, address from, address to, uint256 amount) returns()
func (_Bscspendcoin *BscspendcoinSession) SpendUserCoin(tokenAddress common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bscspendcoin.Contract.SpendUserCoin(&_Bscspendcoin.TransactOpts, tokenAddress, from, to, amount)
}

// SpendUserCoin is a paid mutator transaction binding the contract method 0x49afc972.
//
// Solidity: function spendUserCoin(address tokenAddress, address from, address to, uint256 amount) returns()
func (_Bscspendcoin *BscspendcoinTransactorSession) SpendUserCoin(tokenAddress common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bscspendcoin.Contract.SpendUserCoin(&_Bscspendcoin.TransactOpts, tokenAddress, from, to, amount)
}
