// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// TraceabilityStep is an auto generated low-level Go binding around an user-defined struct.
type TraceabilityStep struct {
	Location    string
	Description string
	Timestamp   *big.Int
	Actor       common.Address
	Status      uint8
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"productId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"enumTraceability.StepStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"addStep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"authorize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"productId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"enumTraceability.ProductStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"createProduct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Authorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"productId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumTraceability.ProductStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ProductCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Revoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"productId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumTraceability.ProductStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"StatusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"productId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumTraceability.StepStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"StepAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"productId\",\"type\":\"string\"},{\"internalType\":\"enumTraceability.ProductStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"updateProductStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"productId\",\"type\":\"string\"}],\"name\":\"getProduct\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"enumTraceability.ProductStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"},{\"internalType\":\"enumTraceability.StepStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structTraceability.Step[]\",\"name\":\"steps\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"getProductsByCreator\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"productId\",\"type\":\"string\"}],\"name\":\"getProductStatus\",\"outputs\":[{\"internalType\":\"enumTraceability.ProductStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"productId\",\"type\":\"string\"}],\"name\":\"getSteps\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"},{\"internalType\":\"enumTraceability.StepStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structTraceability.Step[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"productId\",\"type\":\"string\"}],\"name\":\"isProductExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// GetProduct is a free data retrieval call binding the contract method 0x68111cce.
//
// Solidity: function getProduct(string productId) view returns(string name, string ipfsHash, address creator, uint8 status, (string,string,uint256,address,uint8)[] steps, string location)
func (_Contracts *ContractsCaller) GetProduct(opts *bind.CallOpts, productId string) (struct {
	Name     string
	IpfsHash string
	Creator  common.Address
	Status   uint8
	Steps    []TraceabilityStep
	Location string
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getProduct", productId)

	outstruct := new(struct {
		Name     string
		IpfsHash string
		Creator  common.Address
		Status   uint8
		Steps    []TraceabilityStep
		Location string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.IpfsHash = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Creator = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Steps = *abi.ConvertType(out[4], new([]TraceabilityStep)).(*[]TraceabilityStep)
	outstruct.Location = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// GetProduct is a free data retrieval call binding the contract method 0x68111cce.
//
// Solidity: function getProduct(string productId) view returns(string name, string ipfsHash, address creator, uint8 status, (string,string,uint256,address,uint8)[] steps, string location)
func (_Contracts *ContractsSession) GetProduct(productId string) (struct {
	Name     string
	IpfsHash string
	Creator  common.Address
	Status   uint8
	Steps    []TraceabilityStep
	Location string
}, error) {
	return _Contracts.Contract.GetProduct(&_Contracts.CallOpts, productId)
}

// GetProduct is a free data retrieval call binding the contract method 0x68111cce.
//
// Solidity: function getProduct(string productId) view returns(string name, string ipfsHash, address creator, uint8 status, (string,string,uint256,address,uint8)[] steps, string location)
func (_Contracts *ContractsCallerSession) GetProduct(productId string) (struct {
	Name     string
	IpfsHash string
	Creator  common.Address
	Status   uint8
	Steps    []TraceabilityStep
	Location string
}, error) {
	return _Contracts.Contract.GetProduct(&_Contracts.CallOpts, productId)
}

// GetProductStatus is a free data retrieval call binding the contract method 0x9f62af07.
//
// Solidity: function getProductStatus(string productId) view returns(uint8)
func (_Contracts *ContractsCaller) GetProductStatus(opts *bind.CallOpts, productId string) (uint8, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getProductStatus", productId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetProductStatus is a free data retrieval call binding the contract method 0x9f62af07.
//
// Solidity: function getProductStatus(string productId) view returns(uint8)
func (_Contracts *ContractsSession) GetProductStatus(productId string) (uint8, error) {
	return _Contracts.Contract.GetProductStatus(&_Contracts.CallOpts, productId)
}

// GetProductStatus is a free data retrieval call binding the contract method 0x9f62af07.
//
// Solidity: function getProductStatus(string productId) view returns(uint8)
func (_Contracts *ContractsCallerSession) GetProductStatus(productId string) (uint8, error) {
	return _Contracts.Contract.GetProductStatus(&_Contracts.CallOpts, productId)
}

// GetProductsByCreator is a free data retrieval call binding the contract method 0x473cbdfc.
//
// Solidity: function getProductsByCreator(address creator) view returns(string[])
func (_Contracts *ContractsCaller) GetProductsByCreator(opts *bind.CallOpts, creator common.Address) ([]string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getProductsByCreator", creator)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetProductsByCreator is a free data retrieval call binding the contract method 0x473cbdfc.
//
// Solidity: function getProductsByCreator(address creator) view returns(string[])
func (_Contracts *ContractsSession) GetProductsByCreator(creator common.Address) ([]string, error) {
	return _Contracts.Contract.GetProductsByCreator(&_Contracts.CallOpts, creator)
}

// GetProductsByCreator is a free data retrieval call binding the contract method 0x473cbdfc.
//
// Solidity: function getProductsByCreator(address creator) view returns(string[])
func (_Contracts *ContractsCallerSession) GetProductsByCreator(creator common.Address) ([]string, error) {
	return _Contracts.Contract.GetProductsByCreator(&_Contracts.CallOpts, creator)
}

// GetSteps is a free data retrieval call binding the contract method 0xc11b668a.
//
// Solidity: function getSteps(string productId) view returns((string,string,uint256,address,uint8)[])
func (_Contracts *ContractsCaller) GetSteps(opts *bind.CallOpts, productId string) ([]TraceabilityStep, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getSteps", productId)

	if err != nil {
		return *new([]TraceabilityStep), err
	}

	out0 := *abi.ConvertType(out[0], new([]TraceabilityStep)).(*[]TraceabilityStep)

	return out0, err

}

// GetSteps is a free data retrieval call binding the contract method 0xc11b668a.
//
// Solidity: function getSteps(string productId) view returns((string,string,uint256,address,uint8)[])
func (_Contracts *ContractsSession) GetSteps(productId string) ([]TraceabilityStep, error) {
	return _Contracts.Contract.GetSteps(&_Contracts.CallOpts, productId)
}

// GetSteps is a free data retrieval call binding the contract method 0xc11b668a.
//
// Solidity: function getSteps(string productId) view returns((string,string,uint256,address,uint8)[])
func (_Contracts *ContractsCallerSession) GetSteps(productId string) ([]TraceabilityStep, error) {
	return _Contracts.Contract.GetSteps(&_Contracts.CallOpts, productId)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address user) view returns(bool)
func (_Contracts *ContractsCaller) IsAuthorized(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isAuthorized", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address user) view returns(bool)
func (_Contracts *ContractsSession) IsAuthorized(user common.Address) (bool, error) {
	return _Contracts.Contract.IsAuthorized(&_Contracts.CallOpts, user)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address user) view returns(bool)
func (_Contracts *ContractsCallerSession) IsAuthorized(user common.Address) (bool, error) {
	return _Contracts.Contract.IsAuthorized(&_Contracts.CallOpts, user)
}

// IsProductExists is a free data retrieval call binding the contract method 0x7bdf8556.
//
// Solidity: function isProductExists(string productId) view returns(bool)
func (_Contracts *ContractsCaller) IsProductExists(opts *bind.CallOpts, productId string) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isProductExists", productId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProductExists is a free data retrieval call binding the contract method 0x7bdf8556.
//
// Solidity: function isProductExists(string productId) view returns(bool)
func (_Contracts *ContractsSession) IsProductExists(productId string) (bool, error) {
	return _Contracts.Contract.IsProductExists(&_Contracts.CallOpts, productId)
}

// IsProductExists is a free data retrieval call binding the contract method 0x7bdf8556.
//
// Solidity: function isProductExists(string productId) view returns(bool)
func (_Contracts *ContractsCallerSession) IsProductExists(productId string) (bool, error) {
	return _Contracts.Contract.IsProductExists(&_Contracts.CallOpts, productId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// AddStep is a paid mutator transaction binding the contract method 0xa69a26f6.
//
// Solidity: function addStep(string productId, string location, string description, uint8 status) returns()
func (_Contracts *ContractsTransactor) AddStep(opts *bind.TransactOpts, productId string, location string, description string, status uint8) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addStep", productId, location, description, status)
}

// AddStep is a paid mutator transaction binding the contract method 0xa69a26f6.
//
// Solidity: function addStep(string productId, string location, string description, uint8 status) returns()
func (_Contracts *ContractsSession) AddStep(productId string, location string, description string, status uint8) (*types.Transaction, error) {
	return _Contracts.Contract.AddStep(&_Contracts.TransactOpts, productId, location, description, status)
}

// AddStep is a paid mutator transaction binding the contract method 0xa69a26f6.
//
// Solidity: function addStep(string productId, string location, string description, uint8 status) returns()
func (_Contracts *ContractsTransactorSession) AddStep(productId string, location string, description string, status uint8) (*types.Transaction, error) {
	return _Contracts.Contract.AddStep(&_Contracts.TransactOpts, productId, location, description, status)
}

// Authorize is a paid mutator transaction binding the contract method 0xb6a5d7de.
//
// Solidity: function authorize(address user) returns()
func (_Contracts *ContractsTransactor) Authorize(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "authorize", user)
}

// Authorize is a paid mutator transaction binding the contract method 0xb6a5d7de.
//
// Solidity: function authorize(address user) returns()
func (_Contracts *ContractsSession) Authorize(user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Authorize(&_Contracts.TransactOpts, user)
}

// Authorize is a paid mutator transaction binding the contract method 0xb6a5d7de.
//
// Solidity: function authorize(address user) returns()
func (_Contracts *ContractsTransactorSession) Authorize(user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Authorize(&_Contracts.TransactOpts, user)
}

// CreateProduct is a paid mutator transaction binding the contract method 0x4255f38a.
//
// Solidity: function createProduct(string productId, string name, string ipfsHash, string location, uint8 status) returns()
func (_Contracts *ContractsTransactor) CreateProduct(opts *bind.TransactOpts, productId string, name string, ipfsHash string, location string, status uint8) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createProduct", productId, name, ipfsHash, location, status)
}

// CreateProduct is a paid mutator transaction binding the contract method 0x4255f38a.
//
// Solidity: function createProduct(string productId, string name, string ipfsHash, string location, uint8 status) returns()
func (_Contracts *ContractsSession) CreateProduct(productId string, name string, ipfsHash string, location string, status uint8) (*types.Transaction, error) {
	return _Contracts.Contract.CreateProduct(&_Contracts.TransactOpts, productId, name, ipfsHash, location, status)
}

// CreateProduct is a paid mutator transaction binding the contract method 0x4255f38a.
//
// Solidity: function createProduct(string productId, string name, string ipfsHash, string location, uint8 status) returns()
func (_Contracts *ContractsTransactorSession) CreateProduct(productId string, name string, ipfsHash string, location string, status uint8) (*types.Transaction, error) {
	return _Contracts.Contract.CreateProduct(&_Contracts.TransactOpts, productId, name, ipfsHash, location, status)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address user) returns()
func (_Contracts *ContractsTransactor) Revoke(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "revoke", user)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address user) returns()
func (_Contracts *ContractsSession) Revoke(user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Revoke(&_Contracts.TransactOpts, user)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address user) returns()
func (_Contracts *ContractsTransactorSession) Revoke(user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Revoke(&_Contracts.TransactOpts, user)
}

// UpdateProductStatus is a paid mutator transaction binding the contract method 0xadbe15bc.
//
// Solidity: function updateProductStatus(string productId, uint8 newStatus) returns()
func (_Contracts *ContractsTransactor) UpdateProductStatus(opts *bind.TransactOpts, productId string, newStatus uint8) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateProductStatus", productId, newStatus)
}

// UpdateProductStatus is a paid mutator transaction binding the contract method 0xadbe15bc.
//
// Solidity: function updateProductStatus(string productId, uint8 newStatus) returns()
func (_Contracts *ContractsSession) UpdateProductStatus(productId string, newStatus uint8) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateProductStatus(&_Contracts.TransactOpts, productId, newStatus)
}

// UpdateProductStatus is a paid mutator transaction binding the contract method 0xadbe15bc.
//
// Solidity: function updateProductStatus(string productId, uint8 newStatus) returns()
func (_Contracts *ContractsTransactorSession) UpdateProductStatus(productId string, newStatus uint8) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateProductStatus(&_Contracts.TransactOpts, productId, newStatus)
}

// ContractsAuthorizedIterator is returned from FilterAuthorized and is used to iterate over the raw logs and unpacked data for Authorized events raised by the Contracts contract.
type ContractsAuthorizedIterator struct {
	Event *ContractsAuthorized // Event containing the contract specifics and raw log

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
func (it *ContractsAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAuthorized)
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
		it.Event = new(ContractsAuthorized)
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
func (it *ContractsAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAuthorized represents a Authorized event raised by the Contracts contract.
type ContractsAuthorized struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAuthorized is a free log retrieval operation binding the contract event 0xdc84e3a4c83602050e3865df792a4e6800211a79ac60db94e703a820ce892924.
//
// Solidity: event Authorized(address indexed user)
func (_Contracts *ContractsFilterer) FilterAuthorized(opts *bind.FilterOpts, user []common.Address) (*ContractsAuthorizedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Authorized", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAuthorizedIterator{contract: _Contracts.contract, event: "Authorized", logs: logs, sub: sub}, nil
}

// WatchAuthorized is a free log subscription operation binding the contract event 0xdc84e3a4c83602050e3865df792a4e6800211a79ac60db94e703a820ce892924.
//
// Solidity: event Authorized(address indexed user)
func (_Contracts *ContractsFilterer) WatchAuthorized(opts *bind.WatchOpts, sink chan<- *ContractsAuthorized, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Authorized", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAuthorized)
				if err := _Contracts.contract.UnpackLog(event, "Authorized", log); err != nil {
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

// ParseAuthorized is a log parse operation binding the contract event 0xdc84e3a4c83602050e3865df792a4e6800211a79ac60db94e703a820ce892924.
//
// Solidity: event Authorized(address indexed user)
func (_Contracts *ContractsFilterer) ParseAuthorized(log types.Log) (*ContractsAuthorized, error) {
	event := new(ContractsAuthorized)
	if err := _Contracts.contract.UnpackLog(event, "Authorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsProductCreatedIterator is returned from FilterProductCreated and is used to iterate over the raw logs and unpacked data for ProductCreated events raised by the Contracts contract.
type ContractsProductCreatedIterator struct {
	Event *ContractsProductCreated // Event containing the contract specifics and raw log

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
func (it *ContractsProductCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsProductCreated)
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
		it.Event = new(ContractsProductCreated)
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
func (it *ContractsProductCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsProductCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsProductCreated represents a ProductCreated event raised by the Contracts contract.
type ContractsProductCreated struct {
	ProductId string
	Name      string
	Location  string
	Creator   common.Address
	Status    uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProductCreated is a free log retrieval operation binding the contract event 0x1e09055a595d5262ea3988c6785224ad94bdfb07835ee5ee0fc098c211f4e7c9.
//
// Solidity: event ProductCreated(string productId, string name, string location, address indexed creator, uint8 status)
func (_Contracts *ContractsFilterer) FilterProductCreated(opts *bind.FilterOpts, creator []common.Address) (*ContractsProductCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ProductCreated", creatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractsProductCreatedIterator{contract: _Contracts.contract, event: "ProductCreated", logs: logs, sub: sub}, nil
}

// WatchProductCreated is a free log subscription operation binding the contract event 0x1e09055a595d5262ea3988c6785224ad94bdfb07835ee5ee0fc098c211f4e7c9.
//
// Solidity: event ProductCreated(string productId, string name, string location, address indexed creator, uint8 status)
func (_Contracts *ContractsFilterer) WatchProductCreated(opts *bind.WatchOpts, sink chan<- *ContractsProductCreated, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ProductCreated", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsProductCreated)
				if err := _Contracts.contract.UnpackLog(event, "ProductCreated", log); err != nil {
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

// ParseProductCreated is a log parse operation binding the contract event 0x1e09055a595d5262ea3988c6785224ad94bdfb07835ee5ee0fc098c211f4e7c9.
//
// Solidity: event ProductCreated(string productId, string name, string location, address indexed creator, uint8 status)
func (_Contracts *ContractsFilterer) ParseProductCreated(log types.Log) (*ContractsProductCreated, error) {
	event := new(ContractsProductCreated)
	if err := _Contracts.contract.UnpackLog(event, "ProductCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRevokedIterator is returned from FilterRevoked and is used to iterate over the raw logs and unpacked data for Revoked events raised by the Contracts contract.
type ContractsRevokedIterator struct {
	Event *ContractsRevoked // Event containing the contract specifics and raw log

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
func (it *ContractsRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRevoked)
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
		it.Event = new(ContractsRevoked)
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
func (it *ContractsRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRevoked represents a Revoked event raised by the Contracts contract.
type ContractsRevoked struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRevoked is a free log retrieval operation binding the contract event 0xb6fa8b8bd5eab60f292eca876e3ef90722275b785309d84b1de113ce0b8c4e74.
//
// Solidity: event Revoked(address indexed user)
func (_Contracts *ContractsFilterer) FilterRevoked(opts *bind.FilterOpts, user []common.Address) (*ContractsRevokedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Revoked", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRevokedIterator{contract: _Contracts.contract, event: "Revoked", logs: logs, sub: sub}, nil
}

// WatchRevoked is a free log subscription operation binding the contract event 0xb6fa8b8bd5eab60f292eca876e3ef90722275b785309d84b1de113ce0b8c4e74.
//
// Solidity: event Revoked(address indexed user)
func (_Contracts *ContractsFilterer) WatchRevoked(opts *bind.WatchOpts, sink chan<- *ContractsRevoked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Revoked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRevoked)
				if err := _Contracts.contract.UnpackLog(event, "Revoked", log); err != nil {
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

// ParseRevoked is a log parse operation binding the contract event 0xb6fa8b8bd5eab60f292eca876e3ef90722275b785309d84b1de113ce0b8c4e74.
//
// Solidity: event Revoked(address indexed user)
func (_Contracts *ContractsFilterer) ParseRevoked(log types.Log) (*ContractsRevoked, error) {
	event := new(ContractsRevoked)
	if err := _Contracts.contract.UnpackLog(event, "Revoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsStatusUpdatedIterator is returned from FilterStatusUpdated and is used to iterate over the raw logs and unpacked data for StatusUpdated events raised by the Contracts contract.
type ContractsStatusUpdatedIterator struct {
	Event *ContractsStatusUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsStatusUpdated)
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
		it.Event = new(ContractsStatusUpdated)
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
func (it *ContractsStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsStatusUpdated represents a StatusUpdated event raised by the Contracts contract.
type ContractsStatusUpdated struct {
	ProductId string
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStatusUpdated is a free log retrieval operation binding the contract event 0x4e407d4f938199e743ff3f9c87db0c2b408db2c74b3b83b01299015302ef0532.
//
// Solidity: event StatusUpdated(string productId, uint8 newStatus)
func (_Contracts *ContractsFilterer) FilterStatusUpdated(opts *bind.FilterOpts) (*ContractsStatusUpdatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "StatusUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractsStatusUpdatedIterator{contract: _Contracts.contract, event: "StatusUpdated", logs: logs, sub: sub}, nil
}

// WatchStatusUpdated is a free log subscription operation binding the contract event 0x4e407d4f938199e743ff3f9c87db0c2b408db2c74b3b83b01299015302ef0532.
//
// Solidity: event StatusUpdated(string productId, uint8 newStatus)
func (_Contracts *ContractsFilterer) WatchStatusUpdated(opts *bind.WatchOpts, sink chan<- *ContractsStatusUpdated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "StatusUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsStatusUpdated)
				if err := _Contracts.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
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

// ParseStatusUpdated is a log parse operation binding the contract event 0x4e407d4f938199e743ff3f9c87db0c2b408db2c74b3b83b01299015302ef0532.
//
// Solidity: event StatusUpdated(string productId, uint8 newStatus)
func (_Contracts *ContractsFilterer) ParseStatusUpdated(log types.Log) (*ContractsStatusUpdated, error) {
	event := new(ContractsStatusUpdated)
	if err := _Contracts.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsStepAddedIterator is returned from FilterStepAdded and is used to iterate over the raw logs and unpacked data for StepAdded events raised by the Contracts contract.
type ContractsStepAddedIterator struct {
	Event *ContractsStepAdded // Event containing the contract specifics and raw log

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
func (it *ContractsStepAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsStepAdded)
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
		it.Event = new(ContractsStepAdded)
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
func (it *ContractsStepAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsStepAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsStepAdded represents a StepAdded event raised by the Contracts contract.
type ContractsStepAdded struct {
	ProductId   string
	Location    string
	Description string
	Actor       common.Address
	Status      uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStepAdded is a free log retrieval operation binding the contract event 0x46c1dcae6942c89865a9d3daeecc5b88cbc4cbf13b255d5ecc03091a7788ddfa.
//
// Solidity: event StepAdded(string productId, string location, string description, address indexed actor, uint8 status)
func (_Contracts *ContractsFilterer) FilterStepAdded(opts *bind.FilterOpts, actor []common.Address) (*ContractsStepAddedIterator, error) {

	var actorRule []interface{}
	for _, actorItem := range actor {
		actorRule = append(actorRule, actorItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "StepAdded", actorRule)
	if err != nil {
		return nil, err
	}
	return &ContractsStepAddedIterator{contract: _Contracts.contract, event: "StepAdded", logs: logs, sub: sub}, nil
}

// WatchStepAdded is a free log subscription operation binding the contract event 0x46c1dcae6942c89865a9d3daeecc5b88cbc4cbf13b255d5ecc03091a7788ddfa.
//
// Solidity: event StepAdded(string productId, string location, string description, address indexed actor, uint8 status)
func (_Contracts *ContractsFilterer) WatchStepAdded(opts *bind.WatchOpts, sink chan<- *ContractsStepAdded, actor []common.Address) (event.Subscription, error) {

	var actorRule []interface{}
	for _, actorItem := range actor {
		actorRule = append(actorRule, actorItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "StepAdded", actorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsStepAdded)
				if err := _Contracts.contract.UnpackLog(event, "StepAdded", log); err != nil {
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

// ParseStepAdded is a log parse operation binding the contract event 0x46c1dcae6942c89865a9d3daeecc5b88cbc4cbf13b255d5ecc03091a7788ddfa.
//
// Solidity: event StepAdded(string productId, string location, string description, address indexed actor, uint8 status)
func (_Contracts *ContractsFilterer) ParseStepAdded(log types.Log) (*ContractsStepAdded, error) {
	event := new(ContractsStepAdded)
	if err := _Contracts.contract.UnpackLog(event, "StepAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
