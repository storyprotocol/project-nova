// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package story_registry

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

// NftToken is an auto generated low-level Go binding around an user-defined struct.
type NftToken struct {
	Collection common.Address
	TokenId    *big.Int
}

// StoryInfo is an auto generated low-level Go binding around an user-defined struct.
type StoryInfo struct {
	Author          []common.Address
	Title           string
	Characters      []NftToken
	AutographModule common.Address
}

// StoryRegistryMetaData contains all meta data concerning the StoryRegistry contract.
var StoryRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_franchise\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_canonCollection\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_fanFictionCollection\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"StoryAdded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FRANCHISE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"author\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structNftToken[]\",\"name\":\"characters\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"autographModule\",\"type\":\"address\"}],\"internalType\":\"structStoryInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"addStory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canonCollection\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"collect\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fanFictionCollection\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"franchise\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCannonCollection\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getCollectors\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFanFictionCollection\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"story\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"author\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structNftToken[]\",\"name\":\"characters\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"autographModule\",\"type\":\"address\"}],\"internalType\":\"structStoryInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"storyCollections\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"autographModule\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StoryRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use StoryRegistryMetaData.ABI instead.
var StoryRegistryABI = StoryRegistryMetaData.ABI

// StoryRegistry is an auto generated Go binding around an Ethereum contract.
type StoryRegistry struct {
	StoryRegistryCaller     // Read-only binding to the contract
	StoryRegistryTransactor // Write-only binding to the contract
	StoryRegistryFilterer   // Log filterer for contract events
}

// StoryRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoryRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoryRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoryRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoryRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoryRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoryRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoryRegistrySession struct {
	Contract     *StoryRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoryRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoryRegistryCallerSession struct {
	Contract *StoryRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StoryRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoryRegistryTransactorSession struct {
	Contract     *StoryRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StoryRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoryRegistryRaw struct {
	Contract *StoryRegistry // Generic contract binding to access the raw methods on
}

// StoryRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoryRegistryCallerRaw struct {
	Contract *StoryRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// StoryRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoryRegistryTransactorRaw struct {
	Contract *StoryRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStoryRegistry creates a new instance of StoryRegistry, bound to a specific deployed contract.
func NewStoryRegistry(address common.Address, backend bind.ContractBackend) (*StoryRegistry, error) {
	contract, err := bindStoryRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StoryRegistry{StoryRegistryCaller: StoryRegistryCaller{contract: contract}, StoryRegistryTransactor: StoryRegistryTransactor{contract: contract}, StoryRegistryFilterer: StoryRegistryFilterer{contract: contract}}, nil
}

// NewStoryRegistryCaller creates a new read-only instance of StoryRegistry, bound to a specific deployed contract.
func NewStoryRegistryCaller(address common.Address, caller bind.ContractCaller) (*StoryRegistryCaller, error) {
	contract, err := bindStoryRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoryRegistryCaller{contract: contract}, nil
}

// NewStoryRegistryTransactor creates a new write-only instance of StoryRegistry, bound to a specific deployed contract.
func NewStoryRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*StoryRegistryTransactor, error) {
	contract, err := bindStoryRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoryRegistryTransactor{contract: contract}, nil
}

// NewStoryRegistryFilterer creates a new log filterer instance of StoryRegistry, bound to a specific deployed contract.
func NewStoryRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*StoryRegistryFilterer, error) {
	contract, err := bindStoryRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoryRegistryFilterer{contract: contract}, nil
}

// bindStoryRegistry binds a generic wrapper to an already deployed contract.
func bindStoryRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoryRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StoryRegistry *StoryRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StoryRegistry.Contract.StoryRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StoryRegistry *StoryRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StoryRegistry.Contract.StoryRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StoryRegistry *StoryRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StoryRegistry.Contract.StoryRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StoryRegistry *StoryRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StoryRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StoryRegistry *StoryRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StoryRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StoryRegistry *StoryRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StoryRegistry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StoryRegistry *StoryRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StoryRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StoryRegistry *StoryRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StoryRegistry.Contract.DEFAULTADMINROLE(&_StoryRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StoryRegistry *StoryRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StoryRegistry.Contract.DEFAULTADMINROLE(&_StoryRegistry.CallOpts)
}

// FRANCHISEROLE is a free data retrieval call binding the contract method 0x71be958d.
//
// Solidity: function FRANCHISE_ROLE() view returns(bytes32)
func (_StoryRegistry *StoryRegistryCaller) FRANCHISEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StoryRegistry.contract.Call(opts, &out, "FRANCHISE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FRANCHISEROLE is a free data retrieval call binding the contract method 0x71be958d.
//
// Solidity: function FRANCHISE_ROLE() view returns(bytes32)
func (_StoryRegistry *StoryRegistrySession) FRANCHISEROLE() ([32]byte, error) {
	return _StoryRegistry.Contract.FRANCHISEROLE(&_StoryRegistry.CallOpts)
}

// FRANCHISEROLE is a free data retrieval call binding the contract method 0x71be958d.
//
// Solidity: function FRANCHISE_ROLE() view returns(bytes32)
func (_StoryRegistry *StoryRegistryCallerSession) FRANCHISEROLE() ([32]byte, error) {
	return _StoryRegistry.Contract.FRANCHISEROLE(&_StoryRegistry.CallOpts)
}

// CanonCollection is a free data retrieval call binding the contract method 0xf928b1df.
//
// Solidity: function canonCollection() view returns(address)
func (_StoryRegistry *StoryRegistryCaller) CanonCollection(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StoryRegistry.contract.Call(opts, &out, "canonCollection")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CanonCollection is a free data retrieval call binding the contract method 0xf928b1df.
//
// Solidity: function canonCollection() view returns(address)
func (_StoryRegistry *StoryRegistrySession) CanonCollection() (common.Address, error) {
	return _StoryRegistry.Contract.CanonCollection(&_StoryRegistry.CallOpts)
}

// CanonCollection is a free data retrieval call binding the contract method 0xf928b1df.
//
// Solidity: function canonCollection() view returns(address)
func (_StoryRegistry *StoryRegistryCallerSession) CanonCollection() (common.Address, error) {
	return _StoryRegistry.Contract.CanonCollection(&_StoryRegistry.CallOpts)
}

// FanFictionCollection is a free data retrieval call binding the contract method 0xe210d2cf.
//
// Solidity: function fanFictionCollection() view returns(address)
func (_StoryRegistry *StoryRegistryCaller) FanFictionCollection(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StoryRegistry.contract.Call(opts, &out, "fanFictionCollection")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FanFictionCollection is a free data retrieval call binding the contract method 0xe210d2cf.
//
// Solidity: function fanFictionCollection() view returns(address)
func (_StoryRegistry *StoryRegistrySession) FanFictionCollection() (common.Address, error) {
	return _StoryRegistry.Contract.FanFictionCollection(&_StoryRegistry.CallOpts)
}

// FanFictionCollection is a free data retrieval call binding the contract method 0xe210d2cf.
//
// Solidity: function fanFictionCollection() view returns(address)
func (_StoryRegistry *StoryRegistryCallerSession) FanFictionCollection() (common.Address, error) {
	return _StoryRegistry.Contract.FanFictionCollection(&_StoryRegistry.CallOpts)
}

// Franchise is a free data retrieval call binding the contract method 0xc9761dd9.
//
// Solidity: function franchise() view returns(address)
func (_StoryRegistry *StoryRegistryCaller) Franchise(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StoryRegistry.contract.Call(opts, &out, "franchise")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Franchise is a free data retrieval call binding the contract method 0xc9761dd9.
//
// Solidity: function franchise() view returns(address)
func (_StoryRegistry *StoryRegistrySession) Franchise() (common.Address, error) {
	return _StoryRegistry.Contract.Franchise(&_StoryRegistry.CallOpts)
}

// Franchise is a free data retrieval call binding the contract method 0xc9761dd9.
//
// Solidity: function franchise() view returns(address)
func (_StoryRegistry *StoryRegistryCallerSession) Franchise() (common.Address, error) {
	return _StoryRegistry.Contract.Franchise(&_StoryRegistry.CallOpts)
}

// GetCannonCollection is a free data retrieval call binding the contract method 0x5adbcdb5.
//
// Solidity: function getCannonCollection() view returns(address)
func (_StoryRegistry *StoryRegistryCaller) GetCannonCollection(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StoryRegistry.contract.Call(opts, &out, "getCannonCollection")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCannonCollection is a free data retrieval call binding the contract method 0x5adbcdb5.
//
// Solidity: function getCannonCollection() view returns(address)
func (_StoryRegistry *StoryRegistrySession) GetCannonCollection() (common.Address, error) {
	return _StoryRegistry.Contract.GetCannonCollection(&_StoryRegistry.CallOpts)
}

// GetCannonCollection is a free data retrieval call binding the contract method 0x5adbcdb5.
//
// Solidity: function getCannonCollection() view returns(address)
func (_StoryRegistry *StoryRegistryCallerSession) GetCannonCollection() (common.Address, error) {
	return _StoryRegistry.Contract.GetCannonCollection(&_StoryRegistry.CallOpts)
}

// GetCollectors is a free data retrieval call binding the contract method 0xca45115a.
//
// Solidity: function getCollectors(address collection, uint256 id) view returns(address[])
func (_StoryRegistry *StoryRegistryCaller) GetCollectors(opts *bind.CallOpts, collection common.Address, id *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _StoryRegistry.contract.Call(opts, &out, "getCollectors", collection, id)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCollectors is a free data retrieval call binding the contract method 0xca45115a.
//
// Solidity: function getCollectors(address collection, uint256 id) view returns(address[])
func (_StoryRegistry *StoryRegistrySession) GetCollectors(collection common.Address, id *big.Int) ([]common.Address, error) {
	return _StoryRegistry.Contract.GetCollectors(&_StoryRegistry.CallOpts, collection, id)
}

// GetCollectors is a free data retrieval call binding the contract method 0xca45115a.
//
// Solidity: function getCollectors(address collection, uint256 id) view returns(address[])
func (_StoryRegistry *StoryRegistryCallerSession) GetCollectors(collection common.Address, id *big.Int) ([]common.Address, error) {
	return _StoryRegistry.Contract.GetCollectors(&_StoryRegistry.CallOpts, collection, id)
}

// GetFanFictionCollection is a free data retrieval call binding the contract method 0x6204d173.
//
// Solidity: function getFanFictionCollection() view returns(address)
func (_StoryRegistry *StoryRegistryCaller) GetFanFictionCollection(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StoryRegistry.contract.Call(opts, &out, "getFanFictionCollection")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFanFictionCollection is a free data retrieval call binding the contract method 0x6204d173.
//
// Solidity: function getFanFictionCollection() view returns(address)
func (_StoryRegistry *StoryRegistrySession) GetFanFictionCollection() (common.Address, error) {
	return _StoryRegistry.Contract.GetFanFictionCollection(&_StoryRegistry.CallOpts)
}

// GetFanFictionCollection is a free data retrieval call binding the contract method 0x6204d173.
//
// Solidity: function getFanFictionCollection() view returns(address)
func (_StoryRegistry *StoryRegistryCallerSession) GetFanFictionCollection() (common.Address, error) {
	return _StoryRegistry.Contract.GetFanFictionCollection(&_StoryRegistry.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StoryRegistry *StoryRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StoryRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StoryRegistry *StoryRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StoryRegistry.Contract.GetRoleAdmin(&_StoryRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StoryRegistry *StoryRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StoryRegistry.Contract.GetRoleAdmin(&_StoryRegistry.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_StoryRegistry *StoryRegistryCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StoryRegistry.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_StoryRegistry *StoryRegistrySession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _StoryRegistry.Contract.GetRoleMember(&_StoryRegistry.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_StoryRegistry *StoryRegistryCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _StoryRegistry.Contract.GetRoleMember(&_StoryRegistry.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_StoryRegistry *StoryRegistryCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _StoryRegistry.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_StoryRegistry *StoryRegistrySession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _StoryRegistry.Contract.GetRoleMemberCount(&_StoryRegistry.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_StoryRegistry *StoryRegistryCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _StoryRegistry.Contract.GetRoleMemberCount(&_StoryRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StoryRegistry *StoryRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StoryRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StoryRegistry *StoryRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StoryRegistry.Contract.HasRole(&_StoryRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StoryRegistry *StoryRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StoryRegistry.Contract.HasRole(&_StoryRegistry.CallOpts, role, account)
}

// Story is a free data retrieval call binding the contract method 0x0b3b2a67.
//
// Solidity: function story(address collection, uint256 tokenId) view returns((address[],string,(address,uint256)[],address))
func (_StoryRegistry *StoryRegistryCaller) Story(opts *bind.CallOpts, collection common.Address, tokenId *big.Int) (StoryInfo, error) {
	var out []interface{}
	err := _StoryRegistry.contract.Call(opts, &out, "story", collection, tokenId)

	if err != nil {
		return *new(StoryInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(StoryInfo)).(*StoryInfo)

	return out0, err

}

// Story is a free data retrieval call binding the contract method 0x0b3b2a67.
//
// Solidity: function story(address collection, uint256 tokenId) view returns((address[],string,(address,uint256)[],address))
func (_StoryRegistry *StoryRegistrySession) Story(collection common.Address, tokenId *big.Int) (StoryInfo, error) {
	return _StoryRegistry.Contract.Story(&_StoryRegistry.CallOpts, collection, tokenId)
}

// Story is a free data retrieval call binding the contract method 0x0b3b2a67.
//
// Solidity: function story(address collection, uint256 tokenId) view returns((address[],string,(address,uint256)[],address))
func (_StoryRegistry *StoryRegistryCallerSession) Story(collection common.Address, tokenId *big.Int) (StoryInfo, error) {
	return _StoryRegistry.Contract.Story(&_StoryRegistry.CallOpts, collection, tokenId)
}

// StoryCollections is a free data retrieval call binding the contract method 0x05fe7368.
//
// Solidity: function storyCollections(address ) view returns(string title, address autographModule)
func (_StoryRegistry *StoryRegistryCaller) StoryCollections(opts *bind.CallOpts, arg0 common.Address) (struct {
	Title           string
	AutographModule common.Address
}, error) {
	var out []interface{}
	err := _StoryRegistry.contract.Call(opts, &out, "storyCollections", arg0)

	outstruct := new(struct {
		Title           string
		AutographModule common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Title = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.AutographModule = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// StoryCollections is a free data retrieval call binding the contract method 0x05fe7368.
//
// Solidity: function storyCollections(address ) view returns(string title, address autographModule)
func (_StoryRegistry *StoryRegistrySession) StoryCollections(arg0 common.Address) (struct {
	Title           string
	AutographModule common.Address
}, error) {
	return _StoryRegistry.Contract.StoryCollections(&_StoryRegistry.CallOpts, arg0)
}

// StoryCollections is a free data retrieval call binding the contract method 0x05fe7368.
//
// Solidity: function storyCollections(address ) view returns(string title, address autographModule)
func (_StoryRegistry *StoryRegistryCallerSession) StoryCollections(arg0 common.Address) (struct {
	Title           string
	AutographModule common.Address
}, error) {
	return _StoryRegistry.Contract.StoryCollections(&_StoryRegistry.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StoryRegistry *StoryRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StoryRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StoryRegistry *StoryRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StoryRegistry.Contract.SupportsInterface(&_StoryRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StoryRegistry *StoryRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StoryRegistry.Contract.SupportsInterface(&_StoryRegistry.CallOpts, interfaceId)
}

// AddStory is a paid mutator transaction binding the contract method 0x65a199af.
//
// Solidity: function addStory(address collection, (address[],string,(address,uint256)[],address) info, uint256 tokenId) returns()
func (_StoryRegistry *StoryRegistryTransactor) AddStory(opts *bind.TransactOpts, collection common.Address, info StoryInfo, tokenId *big.Int) (*types.Transaction, error) {
	return _StoryRegistry.contract.Transact(opts, "addStory", collection, info, tokenId)
}

// AddStory is a paid mutator transaction binding the contract method 0x65a199af.
//
// Solidity: function addStory(address collection, (address[],string,(address,uint256)[],address) info, uint256 tokenId) returns()
func (_StoryRegistry *StoryRegistrySession) AddStory(collection common.Address, info StoryInfo, tokenId *big.Int) (*types.Transaction, error) {
	return _StoryRegistry.Contract.AddStory(&_StoryRegistry.TransactOpts, collection, info, tokenId)
}

// AddStory is a paid mutator transaction binding the contract method 0x65a199af.
//
// Solidity: function addStory(address collection, (address[],string,(address,uint256)[],address) info, uint256 tokenId) returns()
func (_StoryRegistry *StoryRegistryTransactorSession) AddStory(collection common.Address, info StoryInfo, tokenId *big.Int) (*types.Transaction, error) {
	return _StoryRegistry.Contract.AddStory(&_StoryRegistry.TransactOpts, collection, info, tokenId)
}

// Collect is a paid mutator transaction binding the contract method 0xc8fea2fb.
//
// Solidity: function collect(address sender, address collection, uint256 id) returns()
func (_StoryRegistry *StoryRegistryTransactor) Collect(opts *bind.TransactOpts, sender common.Address, collection common.Address, id *big.Int) (*types.Transaction, error) {
	return _StoryRegistry.contract.Transact(opts, "collect", sender, collection, id)
}

// Collect is a paid mutator transaction binding the contract method 0xc8fea2fb.
//
// Solidity: function collect(address sender, address collection, uint256 id) returns()
func (_StoryRegistry *StoryRegistrySession) Collect(sender common.Address, collection common.Address, id *big.Int) (*types.Transaction, error) {
	return _StoryRegistry.Contract.Collect(&_StoryRegistry.TransactOpts, sender, collection, id)
}

// Collect is a paid mutator transaction binding the contract method 0xc8fea2fb.
//
// Solidity: function collect(address sender, address collection, uint256 id) returns()
func (_StoryRegistry *StoryRegistryTransactorSession) Collect(sender common.Address, collection common.Address, id *big.Int) (*types.Transaction, error) {
	return _StoryRegistry.Contract.Collect(&_StoryRegistry.TransactOpts, sender, collection, id)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StoryRegistry *StoryRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StoryRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StoryRegistry *StoryRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StoryRegistry.Contract.GrantRole(&_StoryRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StoryRegistry *StoryRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StoryRegistry.Contract.GrantRole(&_StoryRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StoryRegistry *StoryRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StoryRegistry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StoryRegistry *StoryRegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StoryRegistry.Contract.RenounceRole(&_StoryRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StoryRegistry *StoryRegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StoryRegistry.Contract.RenounceRole(&_StoryRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StoryRegistry *StoryRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StoryRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StoryRegistry *StoryRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StoryRegistry.Contract.RevokeRole(&_StoryRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StoryRegistry *StoryRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StoryRegistry.Contract.RevokeRole(&_StoryRegistry.TransactOpts, role, account)
}

// StoryRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StoryRegistry contract.
type StoryRegistryRoleAdminChangedIterator struct {
	Event *StoryRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StoryRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoryRegistryRoleAdminChanged)
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
		it.Event = new(StoryRegistryRoleAdminChanged)
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
func (it *StoryRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoryRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoryRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the StoryRegistry contract.
type StoryRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StoryRegistry *StoryRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StoryRegistryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _StoryRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StoryRegistryRoleAdminChangedIterator{contract: _StoryRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StoryRegistry *StoryRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StoryRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _StoryRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoryRegistryRoleAdminChanged)
				if err := _StoryRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StoryRegistry *StoryRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*StoryRegistryRoleAdminChanged, error) {
	event := new(StoryRegistryRoleAdminChanged)
	if err := _StoryRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoryRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StoryRegistry contract.
type StoryRegistryRoleGrantedIterator struct {
	Event *StoryRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *StoryRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoryRegistryRoleGranted)
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
		it.Event = new(StoryRegistryRoleGranted)
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
func (it *StoryRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoryRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoryRegistryRoleGranted represents a RoleGranted event raised by the StoryRegistry contract.
type StoryRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StoryRegistry *StoryRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StoryRegistryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StoryRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StoryRegistryRoleGrantedIterator{contract: _StoryRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StoryRegistry *StoryRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StoryRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StoryRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoryRegistryRoleGranted)
				if err := _StoryRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StoryRegistry *StoryRegistryFilterer) ParseRoleGranted(log types.Log) (*StoryRegistryRoleGranted, error) {
	event := new(StoryRegistryRoleGranted)
	if err := _StoryRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoryRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StoryRegistry contract.
type StoryRegistryRoleRevokedIterator struct {
	Event *StoryRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StoryRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoryRegistryRoleRevoked)
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
		it.Event = new(StoryRegistryRoleRevoked)
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
func (it *StoryRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoryRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoryRegistryRoleRevoked represents a RoleRevoked event raised by the StoryRegistry contract.
type StoryRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StoryRegistry *StoryRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StoryRegistryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StoryRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StoryRegistryRoleRevokedIterator{contract: _StoryRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StoryRegistry *StoryRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StoryRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StoryRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoryRegistryRoleRevoked)
				if err := _StoryRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StoryRegistry *StoryRegistryFilterer) ParseRoleRevoked(log types.Log) (*StoryRegistryRoleRevoked, error) {
	event := new(StoryRegistryRoleRevoked)
	if err := _StoryRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoryRegistryStoryAddedIterator is returned from FilterStoryAdded and is used to iterate over the raw logs and unpacked data for StoryAdded events raised by the StoryRegistry contract.
type StoryRegistryStoryAddedIterator struct {
	Event *StoryRegistryStoryAdded // Event containing the contract specifics and raw log

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
func (it *StoryRegistryStoryAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoryRegistryStoryAdded)
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
		it.Event = new(StoryRegistryStoryAdded)
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
func (it *StoryRegistryStoryAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoryRegistryStoryAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoryRegistryStoryAdded represents a StoryAdded event raised by the StoryRegistry contract.
type StoryRegistryStoryAdded struct {
	Collection common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStoryAdded is a free log retrieval operation binding the contract event 0x93dd642f9cdb4c408b6cf816f340e12e1e86aa102b1d14df8c68d23b1a921427.
//
// Solidity: event StoryAdded(address collection, uint256 tokenId)
func (_StoryRegistry *StoryRegistryFilterer) FilterStoryAdded(opts *bind.FilterOpts) (*StoryRegistryStoryAddedIterator, error) {

	logs, sub, err := _StoryRegistry.contract.FilterLogs(opts, "StoryAdded")
	if err != nil {
		return nil, err
	}
	return &StoryRegistryStoryAddedIterator{contract: _StoryRegistry.contract, event: "StoryAdded", logs: logs, sub: sub}, nil
}

// WatchStoryAdded is a free log subscription operation binding the contract event 0x93dd642f9cdb4c408b6cf816f340e12e1e86aa102b1d14df8c68d23b1a921427.
//
// Solidity: event StoryAdded(address collection, uint256 tokenId)
func (_StoryRegistry *StoryRegistryFilterer) WatchStoryAdded(opts *bind.WatchOpts, sink chan<- *StoryRegistryStoryAdded) (event.Subscription, error) {

	logs, sub, err := _StoryRegistry.contract.WatchLogs(opts, "StoryAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoryRegistryStoryAdded)
				if err := _StoryRegistry.contract.UnpackLog(event, "StoryAdded", log); err != nil {
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

// ParseStoryAdded is a log parse operation binding the contract event 0x93dd642f9cdb4c408b6cf816f340e12e1e86aa102b1d14df8c68d23b1a921427.
//
// Solidity: event StoryAdded(address collection, uint256 tokenId)
func (_StoryRegistry *StoryRegistryFilterer) ParseStoryAdded(log types.Log) (*StoryRegistryStoryAdded, error) {
	event := new(StoryRegistryStoryAdded)
	if err := _StoryRegistry.contract.UnpackLog(event, "StoryAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
