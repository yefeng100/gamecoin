// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trcspendcoin

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

// TrcspendcoinMetaData contains all meta data concerning the Trcspendcoin contract.
var TrcspendcoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"collectTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TrcspendcoinABI is the input ABI used to generate the binding from.
// Deprecated: Use TrcspendcoinMetaData.ABI instead.
var TrcspendcoinABI = TrcspendcoinMetaData.ABI

// Trcspendcoin is an auto generated Go binding around an Ethereum contract.
type Trcspendcoin struct {
	TrcspendcoinCaller     // Read-only binding to the contract
	TrcspendcoinTransactor // Write-only binding to the contract
	TrcspendcoinFilterer   // Log filterer for contract events
}

// TrcspendcoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrcspendcoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrcspendcoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrcspendcoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrcspendcoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrcspendcoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrcspendcoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrcspendcoinSession struct {
	Contract     *Trcspendcoin     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TrcspendcoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrcspendcoinCallerSession struct {
	Contract *TrcspendcoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TrcspendcoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrcspendcoinTransactorSession struct {
	Contract     *TrcspendcoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TrcspendcoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrcspendcoinRaw struct {
	Contract *Trcspendcoin // Generic contract binding to access the raw methods on
}

// TrcspendcoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrcspendcoinCallerRaw struct {
	Contract *TrcspendcoinCaller // Generic read-only contract binding to access the raw methods on
}

// TrcspendcoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrcspendcoinTransactorRaw struct {
	Contract *TrcspendcoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrcspendcoin creates a new instance of Trcspendcoin, bound to a specific deployed contract.
func NewTrcspendcoin(address common.Address, backend bind.ContractBackend) (*Trcspendcoin, error) {
	contract, err := bindTrcspendcoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trcspendcoin{TrcspendcoinCaller: TrcspendcoinCaller{contract: contract}, TrcspendcoinTransactor: TrcspendcoinTransactor{contract: contract}, TrcspendcoinFilterer: TrcspendcoinFilterer{contract: contract}}, nil
}

// NewTrcspendcoinCaller creates a new read-only instance of Trcspendcoin, bound to a specific deployed contract.
func NewTrcspendcoinCaller(address common.Address, caller bind.ContractCaller) (*TrcspendcoinCaller, error) {
	contract, err := bindTrcspendcoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrcspendcoinCaller{contract: contract}, nil
}

// NewTrcspendcoinTransactor creates a new write-only instance of Trcspendcoin, bound to a specific deployed contract.
func NewTrcspendcoinTransactor(address common.Address, transactor bind.ContractTransactor) (*TrcspendcoinTransactor, error) {
	contract, err := bindTrcspendcoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrcspendcoinTransactor{contract: contract}, nil
}

// NewTrcspendcoinFilterer creates a new log filterer instance of Trcspendcoin, bound to a specific deployed contract.
func NewTrcspendcoinFilterer(address common.Address, filterer bind.ContractFilterer) (*TrcspendcoinFilterer, error) {
	contract, err := bindTrcspendcoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrcspendcoinFilterer{contract: contract}, nil
}

// bindTrcspendcoin binds a generic wrapper to an already deployed contract.
func bindTrcspendcoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TrcspendcoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trcspendcoin *TrcspendcoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trcspendcoin.Contract.TrcspendcoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trcspendcoin *TrcspendcoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trcspendcoin.Contract.TrcspendcoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trcspendcoin *TrcspendcoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trcspendcoin.Contract.TrcspendcoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trcspendcoin *TrcspendcoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trcspendcoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trcspendcoin *TrcspendcoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trcspendcoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trcspendcoin *TrcspendcoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trcspendcoin.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trcspendcoin *TrcspendcoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trcspendcoin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trcspendcoin *TrcspendcoinSession) Owner() (common.Address, error) {
	return _Trcspendcoin.Contract.Owner(&_Trcspendcoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trcspendcoin *TrcspendcoinCallerSession) Owner() (common.Address, error) {
	return _Trcspendcoin.Contract.Owner(&_Trcspendcoin.CallOpts)
}

// CollectTokens is a paid mutator transaction binding the contract method 0x237f4b66.
//
// Solidity: function collectTokens(address tokenAddress, address from, uint256 amount) returns()
func (_Trcspendcoin *TrcspendcoinTransactor) CollectTokens(opts *bind.TransactOpts, tokenAddress common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Trcspendcoin.contract.Transact(opts, "collectTokens", tokenAddress, from, amount)
}

// CollectTokens is a paid mutator transaction binding the contract method 0x237f4b66.
//
// Solidity: function collectTokens(address tokenAddress, address from, uint256 amount) returns()
func (_Trcspendcoin *TrcspendcoinSession) CollectTokens(tokenAddress common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Trcspendcoin.Contract.CollectTokens(&_Trcspendcoin.TransactOpts, tokenAddress, from, amount)
}

// CollectTokens is a paid mutator transaction binding the contract method 0x237f4b66.
//
// Solidity: function collectTokens(address tokenAddress, address from, uint256 amount) returns()
func (_Trcspendcoin *TrcspendcoinTransactorSession) CollectTokens(tokenAddress common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Trcspendcoin.Contract.CollectTokens(&_Trcspendcoin.TransactOpts, tokenAddress, from, amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address tokenAddress) returns()
func (_Trcspendcoin *TrcspendcoinTransactor) WithdrawTokens(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _Trcspendcoin.contract.Transact(opts, "withdrawTokens", tokenAddress)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address tokenAddress) returns()
func (_Trcspendcoin *TrcspendcoinSession) WithdrawTokens(tokenAddress common.Address) (*types.Transaction, error) {
	return _Trcspendcoin.Contract.WithdrawTokens(&_Trcspendcoin.TransactOpts, tokenAddress)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address tokenAddress) returns()
func (_Trcspendcoin *TrcspendcoinTransactorSession) WithdrawTokens(tokenAddress common.Address) (*types.Transaction, error) {
	return _Trcspendcoin.Contract.WithdrawTokens(&_Trcspendcoin.TransactOpts, tokenAddress)
}
