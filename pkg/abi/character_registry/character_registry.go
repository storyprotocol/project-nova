// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package character_registry

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

// ICharacterRegistryAuthor is an auto generated low-level Go binding around an user-defined struct.
type ICharacterRegistryAuthor struct {
	FeeCollector common.Address
	Name         string
}

// ICharacterRegistryCharacterInfo is an auto generated low-level Go binding around an user-defined struct.
type ICharacterRegistryCharacterInfo struct {
	Name          string
	Description   string
	Author        ICharacterRegistryAuthor
	LicenseModule common.Address
	CharType      uint8
}

// CharacterRegistryMetaData contains all meta data concerning the CharacterRegistry contract.
var CharacterRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_canonCollection\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_externalCanon\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_fanficCollection\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_externalFanfic\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"charId\",\"type\":\"uint256\"}],\"name\":\"CharacterAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structICharacterRegistry.Author\",\"name\":\"author\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"licenseModule\",\"type\":\"address\"},{\"internalType\":\"enumICharacterRegistry.CharacterType\",\"name\":\"charType\",\"type\":\"uint8\"}],\"internalType\":\"structICharacterRegistry.CharacterInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"externalId\",\"type\":\"uint256\"}],\"name\":\"addCharacter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canonCollection\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"charInfoId\",\"type\":\"uint256\"}],\"name\":\"character\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structICharacterRegistry.Author\",\"name\":\"author\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"licenseModule\",\"type\":\"address\"},{\"internalType\":\"enumICharacterRegistry.CharacterType\",\"name\":\"charType\",\"type\":\"uint8\"}],\"internalType\":\"structICharacterRegistry.CharacterInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"collectionAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"externalCanon\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"externalFanfic\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fanficCollection\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"}],\"name\":\"isExternalCollection\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCanonCharacters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCollections\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFanficCharacters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CharacterRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use CharacterRegistryMetaData.ABI instead.
var CharacterRegistryABI = CharacterRegistryMetaData.ABI

// CharacterRegistry is an auto generated Go binding around an Ethereum contract.
type CharacterRegistry struct {
	CharacterRegistryCaller     // Read-only binding to the contract
	CharacterRegistryTransactor // Write-only binding to the contract
	CharacterRegistryFilterer   // Log filterer for contract events
}

// CharacterRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CharacterRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CharacterRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CharacterRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CharacterRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CharacterRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CharacterRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CharacterRegistrySession struct {
	Contract     *CharacterRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CharacterRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CharacterRegistryCallerSession struct {
	Contract *CharacterRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CharacterRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CharacterRegistryTransactorSession struct {
	Contract     *CharacterRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CharacterRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CharacterRegistryRaw struct {
	Contract *CharacterRegistry // Generic contract binding to access the raw methods on
}

// CharacterRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CharacterRegistryCallerRaw struct {
	Contract *CharacterRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// CharacterRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CharacterRegistryTransactorRaw struct {
	Contract *CharacterRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCharacterRegistry creates a new instance of CharacterRegistry, bound to a specific deployed contract.
func NewCharacterRegistry(address common.Address, backend bind.ContractBackend) (*CharacterRegistry, error) {
	contract, err := bindCharacterRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CharacterRegistry{CharacterRegistryCaller: CharacterRegistryCaller{contract: contract}, CharacterRegistryTransactor: CharacterRegistryTransactor{contract: contract}, CharacterRegistryFilterer: CharacterRegistryFilterer{contract: contract}}, nil
}

// NewCharacterRegistryCaller creates a new read-only instance of CharacterRegistry, bound to a specific deployed contract.
func NewCharacterRegistryCaller(address common.Address, caller bind.ContractCaller) (*CharacterRegistryCaller, error) {
	contract, err := bindCharacterRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CharacterRegistryCaller{contract: contract}, nil
}

// NewCharacterRegistryTransactor creates a new write-only instance of CharacterRegistry, bound to a specific deployed contract.
func NewCharacterRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*CharacterRegistryTransactor, error) {
	contract, err := bindCharacterRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CharacterRegistryTransactor{contract: contract}, nil
}

// NewCharacterRegistryFilterer creates a new log filterer instance of CharacterRegistry, bound to a specific deployed contract.
func NewCharacterRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*CharacterRegistryFilterer, error) {
	contract, err := bindCharacterRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CharacterRegistryFilterer{contract: contract}, nil
}

// bindCharacterRegistry binds a generic wrapper to an already deployed contract.
func bindCharacterRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CharacterRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CharacterRegistry *CharacterRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CharacterRegistry.Contract.CharacterRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CharacterRegistry *CharacterRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CharacterRegistry.Contract.CharacterRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CharacterRegistry *CharacterRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CharacterRegistry.Contract.CharacterRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CharacterRegistry *CharacterRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CharacterRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CharacterRegistry *CharacterRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CharacterRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CharacterRegistry *CharacterRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CharacterRegistry.Contract.contract.Transact(opts, method, params...)
}

// CanonCollection is a free data retrieval call binding the contract method 0xf928b1df.
//
// Solidity: function canonCollection() view returns(address)
func (_CharacterRegistry *CharacterRegistryCaller) CanonCollection(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CharacterRegistry.contract.Call(opts, &out, "canonCollection")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CanonCollection is a free data retrieval call binding the contract method 0xf928b1df.
//
// Solidity: function canonCollection() view returns(address)
func (_CharacterRegistry *CharacterRegistrySession) CanonCollection() (common.Address, error) {
	return _CharacterRegistry.Contract.CanonCollection(&_CharacterRegistry.CallOpts)
}

// CanonCollection is a free data retrieval call binding the contract method 0xf928b1df.
//
// Solidity: function canonCollection() view returns(address)
func (_CharacterRegistry *CharacterRegistryCallerSession) CanonCollection() (common.Address, error) {
	return _CharacterRegistry.Contract.CanonCollection(&_CharacterRegistry.CallOpts)
}

// Character is a free data retrieval call binding the contract method 0x813f9d62.
//
// Solidity: function character(address collection, uint256 charInfoId) view returns((string,string,(address,string),address,uint8))
func (_CharacterRegistry *CharacterRegistryCaller) Character(opts *bind.CallOpts, collection common.Address, charInfoId *big.Int) (ICharacterRegistryCharacterInfo, error) {
	var out []interface{}
	err := _CharacterRegistry.contract.Call(opts, &out, "character", collection, charInfoId)

	if err != nil {
		return *new(ICharacterRegistryCharacterInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ICharacterRegistryCharacterInfo)).(*ICharacterRegistryCharacterInfo)

	return out0, err

}

// Character is a free data retrieval call binding the contract method 0x813f9d62.
//
// Solidity: function character(address collection, uint256 charInfoId) view returns((string,string,(address,string),address,uint8))
func (_CharacterRegistry *CharacterRegistrySession) Character(collection common.Address, charInfoId *big.Int) (ICharacterRegistryCharacterInfo, error) {
	return _CharacterRegistry.Contract.Character(&_CharacterRegistry.CallOpts, collection, charInfoId)
}

// Character is a free data retrieval call binding the contract method 0x813f9d62.
//
// Solidity: function character(address collection, uint256 charInfoId) view returns((string,string,(address,string),address,uint8))
func (_CharacterRegistry *CharacterRegistryCallerSession) Character(collection common.Address, charInfoId *big.Int) (ICharacterRegistryCharacterInfo, error) {
	return _CharacterRegistry.Contract.Character(&_CharacterRegistry.CallOpts, collection, charInfoId)
}

// CollectionAt is a free data retrieval call binding the contract method 0xb04d662a.
//
// Solidity: function collectionAt(uint256 index) view returns(address)
func (_CharacterRegistry *CharacterRegistryCaller) CollectionAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CharacterRegistry.contract.Call(opts, &out, "collectionAt", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollectionAt is a free data retrieval call binding the contract method 0xb04d662a.
//
// Solidity: function collectionAt(uint256 index) view returns(address)
func (_CharacterRegistry *CharacterRegistrySession) CollectionAt(index *big.Int) (common.Address, error) {
	return _CharacterRegistry.Contract.CollectionAt(&_CharacterRegistry.CallOpts, index)
}

// CollectionAt is a free data retrieval call binding the contract method 0xb04d662a.
//
// Solidity: function collectionAt(uint256 index) view returns(address)
func (_CharacterRegistry *CharacterRegistryCallerSession) CollectionAt(index *big.Int) (common.Address, error) {
	return _CharacterRegistry.Contract.CollectionAt(&_CharacterRegistry.CallOpts, index)
}

// ExternalCanon is a free data retrieval call binding the contract method 0x59c5ca69.
//
// Solidity: function externalCanon() view returns(bool)
func (_CharacterRegistry *CharacterRegistryCaller) ExternalCanon(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CharacterRegistry.contract.Call(opts, &out, "externalCanon")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExternalCanon is a free data retrieval call binding the contract method 0x59c5ca69.
//
// Solidity: function externalCanon() view returns(bool)
func (_CharacterRegistry *CharacterRegistrySession) ExternalCanon() (bool, error) {
	return _CharacterRegistry.Contract.ExternalCanon(&_CharacterRegistry.CallOpts)
}

// ExternalCanon is a free data retrieval call binding the contract method 0x59c5ca69.
//
// Solidity: function externalCanon() view returns(bool)
func (_CharacterRegistry *CharacterRegistryCallerSession) ExternalCanon() (bool, error) {
	return _CharacterRegistry.Contract.ExternalCanon(&_CharacterRegistry.CallOpts)
}

// ExternalFanfic is a free data retrieval call binding the contract method 0xbf8330e6.
//
// Solidity: function externalFanfic() view returns(bool)
func (_CharacterRegistry *CharacterRegistryCaller) ExternalFanfic(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CharacterRegistry.contract.Call(opts, &out, "externalFanfic")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExternalFanfic is a free data retrieval call binding the contract method 0xbf8330e6.
//
// Solidity: function externalFanfic() view returns(bool)
func (_CharacterRegistry *CharacterRegistrySession) ExternalFanfic() (bool, error) {
	return _CharacterRegistry.Contract.ExternalFanfic(&_CharacterRegistry.CallOpts)
}

// ExternalFanfic is a free data retrieval call binding the contract method 0xbf8330e6.
//
// Solidity: function externalFanfic() view returns(bool)
func (_CharacterRegistry *CharacterRegistryCallerSession) ExternalFanfic() (bool, error) {
	return _CharacterRegistry.Contract.ExternalFanfic(&_CharacterRegistry.CallOpts)
}

// FanficCollection is a free data retrieval call binding the contract method 0x754627a0.
//
// Solidity: function fanficCollection() view returns(address)
func (_CharacterRegistry *CharacterRegistryCaller) FanficCollection(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CharacterRegistry.contract.Call(opts, &out, "fanficCollection")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FanficCollection is a free data retrieval call binding the contract method 0x754627a0.
//
// Solidity: function fanficCollection() view returns(address)
func (_CharacterRegistry *CharacterRegistrySession) FanficCollection() (common.Address, error) {
	return _CharacterRegistry.Contract.FanficCollection(&_CharacterRegistry.CallOpts)
}

// FanficCollection is a free data retrieval call binding the contract method 0x754627a0.
//
// Solidity: function fanficCollection() view returns(address)
func (_CharacterRegistry *CharacterRegistryCallerSession) FanficCollection() (common.Address, error) {
	return _CharacterRegistry.Contract.FanficCollection(&_CharacterRegistry.CallOpts)
}

// IsExternalCollection is a free data retrieval call binding the contract method 0x3dcc4ac4.
//
// Solidity: function isExternalCollection(address collection) view returns(bool)
func (_CharacterRegistry *CharacterRegistryCaller) IsExternalCollection(opts *bind.CallOpts, collection common.Address) (bool, error) {
	var out []interface{}
	err := _CharacterRegistry.contract.Call(opts, &out, "isExternalCollection", collection)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExternalCollection is a free data retrieval call binding the contract method 0x3dcc4ac4.
//
// Solidity: function isExternalCollection(address collection) view returns(bool)
func (_CharacterRegistry *CharacterRegistrySession) IsExternalCollection(collection common.Address) (bool, error) {
	return _CharacterRegistry.Contract.IsExternalCollection(&_CharacterRegistry.CallOpts, collection)
}

// IsExternalCollection is a free data retrieval call binding the contract method 0x3dcc4ac4.
//
// Solidity: function isExternalCollection(address collection) view returns(bool)
func (_CharacterRegistry *CharacterRegistryCallerSession) IsExternalCollection(collection common.Address) (bool, error) {
	return _CharacterRegistry.Contract.IsExternalCollection(&_CharacterRegistry.CallOpts, collection)
}

// TotalCanonCharacters is a free data retrieval call binding the contract method 0x25f9fe5e.
//
// Solidity: function totalCanonCharacters() view returns(uint256)
func (_CharacterRegistry *CharacterRegistryCaller) TotalCanonCharacters(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CharacterRegistry.contract.Call(opts, &out, "totalCanonCharacters")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCanonCharacters is a free data retrieval call binding the contract method 0x25f9fe5e.
//
// Solidity: function totalCanonCharacters() view returns(uint256)
func (_CharacterRegistry *CharacterRegistrySession) TotalCanonCharacters() (*big.Int, error) {
	return _CharacterRegistry.Contract.TotalCanonCharacters(&_CharacterRegistry.CallOpts)
}

// TotalCanonCharacters is a free data retrieval call binding the contract method 0x25f9fe5e.
//
// Solidity: function totalCanonCharacters() view returns(uint256)
func (_CharacterRegistry *CharacterRegistryCallerSession) TotalCanonCharacters() (*big.Int, error) {
	return _CharacterRegistry.Contract.TotalCanonCharacters(&_CharacterRegistry.CallOpts)
}

// TotalCollections is a free data retrieval call binding the contract method 0x61d9db2d.
//
// Solidity: function totalCollections() view returns(uint256)
func (_CharacterRegistry *CharacterRegistryCaller) TotalCollections(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CharacterRegistry.contract.Call(opts, &out, "totalCollections")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCollections is a free data retrieval call binding the contract method 0x61d9db2d.
//
// Solidity: function totalCollections() view returns(uint256)
func (_CharacterRegistry *CharacterRegistrySession) TotalCollections() (*big.Int, error) {
	return _CharacterRegistry.Contract.TotalCollections(&_CharacterRegistry.CallOpts)
}

// TotalCollections is a free data retrieval call binding the contract method 0x61d9db2d.
//
// Solidity: function totalCollections() view returns(uint256)
func (_CharacterRegistry *CharacterRegistryCallerSession) TotalCollections() (*big.Int, error) {
	return _CharacterRegistry.Contract.TotalCollections(&_CharacterRegistry.CallOpts)
}

// TotalFanficCharacters is a free data retrieval call binding the contract method 0x5bce9a84.
//
// Solidity: function totalFanficCharacters() view returns(uint256)
func (_CharacterRegistry *CharacterRegistryCaller) TotalFanficCharacters(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CharacterRegistry.contract.Call(opts, &out, "totalFanficCharacters")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFanficCharacters is a free data retrieval call binding the contract method 0x5bce9a84.
//
// Solidity: function totalFanficCharacters() view returns(uint256)
func (_CharacterRegistry *CharacterRegistrySession) TotalFanficCharacters() (*big.Int, error) {
	return _CharacterRegistry.Contract.TotalFanficCharacters(&_CharacterRegistry.CallOpts)
}

// TotalFanficCharacters is a free data retrieval call binding the contract method 0x5bce9a84.
//
// Solidity: function totalFanficCharacters() view returns(uint256)
func (_CharacterRegistry *CharacterRegistryCallerSession) TotalFanficCharacters() (*big.Int, error) {
	return _CharacterRegistry.Contract.TotalFanficCharacters(&_CharacterRegistry.CallOpts)
}

// AddCharacter is a paid mutator transaction binding the contract method 0x218ecce9.
//
// Solidity: function addCharacter(address collection, (string,string,(address,string),address,uint8) info, uint256 externalId) returns(uint256)
func (_CharacterRegistry *CharacterRegistryTransactor) AddCharacter(opts *bind.TransactOpts, collection common.Address, info ICharacterRegistryCharacterInfo, externalId *big.Int) (*types.Transaction, error) {
	return _CharacterRegistry.contract.Transact(opts, "addCharacter", collection, info, externalId)
}

// AddCharacter is a paid mutator transaction binding the contract method 0x218ecce9.
//
// Solidity: function addCharacter(address collection, (string,string,(address,string),address,uint8) info, uint256 externalId) returns(uint256)
func (_CharacterRegistry *CharacterRegistrySession) AddCharacter(collection common.Address, info ICharacterRegistryCharacterInfo, externalId *big.Int) (*types.Transaction, error) {
	return _CharacterRegistry.Contract.AddCharacter(&_CharacterRegistry.TransactOpts, collection, info, externalId)
}

// AddCharacter is a paid mutator transaction binding the contract method 0x218ecce9.
//
// Solidity: function addCharacter(address collection, (string,string,(address,string),address,uint8) info, uint256 externalId) returns(uint256)
func (_CharacterRegistry *CharacterRegistryTransactorSession) AddCharacter(collection common.Address, info ICharacterRegistryCharacterInfo, externalId *big.Int) (*types.Transaction, error) {
	return _CharacterRegistry.Contract.AddCharacter(&_CharacterRegistry.TransactOpts, collection, info, externalId)
}

// CharacterRegistryCharacterAddedIterator is returned from FilterCharacterAdded and is used to iterate over the raw logs and unpacked data for CharacterAdded events raised by the CharacterRegistry contract.
type CharacterRegistryCharacterAddedIterator struct {
	Event *CharacterRegistryCharacterAdded // Event containing the contract specifics and raw log

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
func (it *CharacterRegistryCharacterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterRegistryCharacterAdded)
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
		it.Event = new(CharacterRegistryCharacterAdded)
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
func (it *CharacterRegistryCharacterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterRegistryCharacterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterRegistryCharacterAdded represents a CharacterAdded event raised by the CharacterRegistry contract.
type CharacterRegistryCharacterAdded struct {
	Collection common.Address
	CharId     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCharacterAdded is a free log retrieval operation binding the contract event 0x78c09245786488a58459f874caf967caf723628ed5857fc3a6c3855db36a6a22.
//
// Solidity: event CharacterAdded(address collection, uint256 charId)
func (_CharacterRegistry *CharacterRegistryFilterer) FilterCharacterAdded(opts *bind.FilterOpts) (*CharacterRegistryCharacterAddedIterator, error) {

	logs, sub, err := _CharacterRegistry.contract.FilterLogs(opts, "CharacterAdded")
	if err != nil {
		return nil, err
	}
	return &CharacterRegistryCharacterAddedIterator{contract: _CharacterRegistry.contract, event: "CharacterAdded", logs: logs, sub: sub}, nil
}

// WatchCharacterAdded is a free log subscription operation binding the contract event 0x78c09245786488a58459f874caf967caf723628ed5857fc3a6c3855db36a6a22.
//
// Solidity: event CharacterAdded(address collection, uint256 charId)
func (_CharacterRegistry *CharacterRegistryFilterer) WatchCharacterAdded(opts *bind.WatchOpts, sink chan<- *CharacterRegistryCharacterAdded) (event.Subscription, error) {

	logs, sub, err := _CharacterRegistry.contract.WatchLogs(opts, "CharacterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterRegistryCharacterAdded)
				if err := _CharacterRegistry.contract.UnpackLog(event, "CharacterAdded", log); err != nil {
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

// ParseCharacterAdded is a log parse operation binding the contract event 0x78c09245786488a58459f874caf967caf723628ed5857fc3a6c3855db36a6a22.
//
// Solidity: event CharacterAdded(address collection, uint256 charId)
func (_CharacterRegistry *CharacterRegistryFilterer) ParseCharacterAdded(log types.Log) (*CharacterRegistryCharacterAdded, error) {
	event := new(CharacterRegistryCharacterAdded)
	if err := _CharacterRegistry.contract.UnpackLog(event, "CharacterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
