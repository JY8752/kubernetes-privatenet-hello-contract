// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hello

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

// HelloMetaData contains all meta data concerning the Hello contract.
var HelloMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"hello\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// HelloABI is the input ABI used to generate the binding from.
// Deprecated: Use HelloMetaData.ABI instead.
var HelloABI = HelloMetaData.ABI

// Hello is an auto generated Go binding around an Ethereum contract.
type Hello struct {
	HelloCaller     // Read-only binding to the contract
	HelloTransactor // Write-only binding to the contract
	HelloFilterer   // Log filterer for contract events
}

// HelloCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelloCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelloTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelloFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelloSession struct {
	Contract     *Hello            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelloCallerSession struct {
	Contract *HelloCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HelloTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelloTransactorSession struct {
	Contract     *HelloTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelloRaw struct {
	Contract *Hello // Generic contract binding to access the raw methods on
}

// HelloCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelloCallerRaw struct {
	Contract *HelloCaller // Generic read-only contract binding to access the raw methods on
}

// HelloTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelloTransactorRaw struct {
	Contract *HelloTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHello creates a new instance of Hello, bound to a specific deployed contract.
func NewHello(address common.Address, backend bind.ContractBackend) (*Hello, error) {
	contract, err := bindHello(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hello{HelloCaller: HelloCaller{contract: contract}, HelloTransactor: HelloTransactor{contract: contract}, HelloFilterer: HelloFilterer{contract: contract}}, nil
}

// NewHelloCaller creates a new read-only instance of Hello, bound to a specific deployed contract.
func NewHelloCaller(address common.Address, caller bind.ContractCaller) (*HelloCaller, error) {
	contract, err := bindHello(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloCaller{contract: contract}, nil
}

// NewHelloTransactor creates a new write-only instance of Hello, bound to a specific deployed contract.
func NewHelloTransactor(address common.Address, transactor bind.ContractTransactor) (*HelloTransactor, error) {
	contract, err := bindHello(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloTransactor{contract: contract}, nil
}

// NewHelloFilterer creates a new log filterer instance of Hello, bound to a specific deployed contract.
func NewHelloFilterer(address common.Address, filterer bind.ContractFilterer) (*HelloFilterer, error) {
	contract, err := bindHello(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloFilterer{contract: contract}, nil
}

// bindHello binds a generic wrapper to an already deployed contract.
func bindHello(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hello *HelloRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hello.Contract.HelloCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hello *HelloRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hello.Contract.HelloTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hello *HelloRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hello.Contract.HelloTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hello *HelloCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hello.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hello *HelloTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hello.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hello *HelloTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hello.Contract.contract.Transact(opts, method, params...)
}

// Hello is a free data retrieval call binding the contract method 0x19ff1d21.
//
// Solidity: function hello() pure returns(string)
func (_Hello *HelloCaller) Hello(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hello.contract.Call(opts, &out, "hello")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Hello is a free data retrieval call binding the contract method 0x19ff1d21.
//
// Solidity: function hello() pure returns(string)
func (_Hello *HelloSession) Hello() (string, error) {
	return _Hello.Contract.Hello(&_Hello.CallOpts)
}

// Hello is a free data retrieval call binding the contract method 0x19ff1d21.
//
// Solidity: function hello() pure returns(string)
func (_Hello *HelloCallerSession) Hello() (string, error) {
	return _Hello.Contract.Hello(&_Hello.CallOpts)
}
