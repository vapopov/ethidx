// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethidx

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

// TokenMetaData contains all meta data concerning the Token contract.
var TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGroupIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_groupId\",\"type\":\"uint256\"}],\"name\":\"getGroup\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_indexId\",\"type\":\"uint256\"}],\"name\":\"getIndex\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"ethPriceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdPriceInCents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdCapitalization\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"percentageChange\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMetaData.ABI instead.
var TokenABI = TokenMetaData.ABI

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// GetGroup is a free data retrieval call binding the contract method 0xceb60654.
//
// Solidity: function getGroup(uint256 _groupId) view returns(string name, uint256[] indexes)
func (_Token *TokenCaller) GetGroup(opts *bind.CallOpts, _groupId *big.Int) (struct {
	Name    string
	Indexes []*big.Int
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getGroup", _groupId)

	outstruct := new(struct {
		Name    string
		Indexes []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Indexes = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetGroup is a free data retrieval call binding the contract method 0xceb60654.
//
// Solidity: function getGroup(uint256 _groupId) view returns(string name, uint256[] indexes)
func (_Token *TokenSession) GetGroup(_groupId *big.Int) (struct {
	Name    string
	Indexes []*big.Int
}, error) {
	return _Token.Contract.GetGroup(&_Token.CallOpts, _groupId)
}

// GetGroup is a free data retrieval call binding the contract method 0xceb60654.
//
// Solidity: function getGroup(uint256 _groupId) view returns(string name, uint256[] indexes)
func (_Token *TokenCallerSession) GetGroup(_groupId *big.Int) (struct {
	Name    string
	Indexes []*big.Int
}, error) {
	return _Token.Contract.GetGroup(&_Token.CallOpts, _groupId)
}

// GetGroupIds is a free data retrieval call binding the contract method 0xc02b8b2c.
//
// Solidity: function getGroupIds() view returns(uint256[])
func (_Token *TokenCaller) GetGroupIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getGroupIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetGroupIds is a free data retrieval call binding the contract method 0xc02b8b2c.
//
// Solidity: function getGroupIds() view returns(uint256[])
func (_Token *TokenSession) GetGroupIds() ([]*big.Int, error) {
	return _Token.Contract.GetGroupIds(&_Token.CallOpts)
}

// GetGroupIds is a free data retrieval call binding the contract method 0xc02b8b2c.
//
// Solidity: function getGroupIds() view returns(uint256[])
func (_Token *TokenCallerSession) GetGroupIds() ([]*big.Int, error) {
	return _Token.Contract.GetGroupIds(&_Token.CallOpts)
}

// GetIndex is a free data retrieval call binding the contract method 0x8e7cb6e1.
//
// Solidity: function getIndex(uint256 _indexId) view returns(string name, uint256 ethPriceInWei, uint256 usdPriceInCents, uint256 usdCapitalization, int256 percentageChange)
func (_Token *TokenCaller) GetIndex(opts *bind.CallOpts, _indexId *big.Int) (struct {
	Name              string
	EthPriceInWei     *big.Int
	UsdPriceInCents   *big.Int
	UsdCapitalization *big.Int
	PercentageChange  *big.Int
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getIndex", _indexId)

	outstruct := new(struct {
		Name              string
		EthPriceInWei     *big.Int
		UsdPriceInCents   *big.Int
		UsdCapitalization *big.Int
		PercentageChange  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.EthPriceInWei = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UsdPriceInCents = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UsdCapitalization = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PercentageChange = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetIndex is a free data retrieval call binding the contract method 0x8e7cb6e1.
//
// Solidity: function getIndex(uint256 _indexId) view returns(string name, uint256 ethPriceInWei, uint256 usdPriceInCents, uint256 usdCapitalization, int256 percentageChange)
func (_Token *TokenSession) GetIndex(_indexId *big.Int) (struct {
	Name              string
	EthPriceInWei     *big.Int
	UsdPriceInCents   *big.Int
	UsdCapitalization *big.Int
	PercentageChange  *big.Int
}, error) {
	return _Token.Contract.GetIndex(&_Token.CallOpts, _indexId)
}

// GetIndex is a free data retrieval call binding the contract method 0x8e7cb6e1.
//
// Solidity: function getIndex(uint256 _indexId) view returns(string name, uint256 ethPriceInWei, uint256 usdPriceInCents, uint256 usdCapitalization, int256 percentageChange)
func (_Token *TokenCallerSession) GetIndex(_indexId *big.Int) (struct {
	Name              string
	EthPriceInWei     *big.Int
	UsdPriceInCents   *big.Int
	UsdCapitalization *big.Int
	PercentageChange  *big.Int
}, error) {
	return _Token.Contract.GetIndex(&_Token.CallOpts, _indexId)
}
