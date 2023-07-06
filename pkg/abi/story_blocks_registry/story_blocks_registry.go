// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package story_blocks_registry

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

// IStoryBlockStorageStoryBlockData is an auto generated low-level Go binding around an user-defined struct.
type IStoryBlockStorageStoryBlockData struct {
	Name        string
	Description string
	MediaUrl    string
	BlockType   uint8
}

// StoryBlocksRegistryMetaData contains all meta data concerning the StoryBlocksRegistry contract.
var StoryBlocksRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GroupedTypeNotGroupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IdOverBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumStoryBlock\",\"name\":\"sb\",\"type\":\"uint8\"}],\"name\":\"InvalidStoryBlock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyLinkedItems\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumStoryBlock\",\"name\":\"linkedType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"linkedItems\",\"type\":\"uint256[]\"}],\"name\":\"GroupedItems\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumStoryBlock\",\"name\":\"sb\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"StoryBlockMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"storyBlockId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumStoryBlock\",\"name\":\"blockType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"mediaUrl\",\"type\":\"string\"}],\"name\":\"StoryBlockWritten\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_LINKED_AT_ONCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mediaUrl\",\"type\":\"string\"},{\"internalType\":\"enumStoryBlock\",\"name\":\"linkedType\",\"type\":\"uint8\"},{\"internalType\":\"uint256[]\",\"name\":\"linkedItems\",\"type\":\"uint256[]\"}],\"name\":\"createGroup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumStoryBlock\",\"name\":\"sb\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mediaUrl\",\"type\":\"string\"}],\"name\":\"createStoryBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumStoryBlock\",\"name\":\"sb\",\"type\":\"uint8\"}],\"name\":\"currentIdFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"franchiseId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"linkedItems\",\"type\":\"uint256[]\"}],\"name\":\"groupItems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_franchiseId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"readGroup\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mediaUrl\",\"type\":\"string\"},{\"internalType\":\"enumStoryBlock\",\"name\":\"blockType\",\"type\":\"uint8\"}],\"internalType\":\"structIStoryBlockStorage.StoryBlockData\",\"name\":\"blockData\",\"type\":\"tuple\"},{\"internalType\":\"enumStoryBlock\",\"name\":\"linkedType\",\"type\":\"uint8\"},{\"internalType\":\"uint256[]\",\"name\":\"linkedItems\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"storyBlockId\",\"type\":\"uint256\"}],\"name\":\"readStoryBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"mediaUrl\",\"type\":\"string\"},{\"internalType\":\"enumStoryBlock\",\"name\":\"blockType\",\"type\":\"uint8\"}],\"internalType\":\"structIStoryBlockStorage.StoryBlockData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// StoryBlocksRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use StoryBlocksRegistryMetaData.ABI instead.
var StoryBlocksRegistryABI = StoryBlocksRegistryMetaData.ABI

// StoryBlocksRegistry is an auto generated Go binding around an Ethereum contract.
type StoryBlocksRegistry struct {
	StoryBlocksRegistryCaller     // Read-only binding to the contract
	StoryBlocksRegistryTransactor // Write-only binding to the contract
	StoryBlocksRegistryFilterer   // Log filterer for contract events
}

// StoryBlocksRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoryBlocksRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoryBlocksRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoryBlocksRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoryBlocksRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoryBlocksRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoryBlocksRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoryBlocksRegistrySession struct {
	Contract     *StoryBlocksRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StoryBlocksRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoryBlocksRegistryCallerSession struct {
	Contract *StoryBlocksRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// StoryBlocksRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoryBlocksRegistryTransactorSession struct {
	Contract     *StoryBlocksRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// StoryBlocksRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoryBlocksRegistryRaw struct {
	Contract *StoryBlocksRegistry // Generic contract binding to access the raw methods on
}

// StoryBlocksRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoryBlocksRegistryCallerRaw struct {
	Contract *StoryBlocksRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// StoryBlocksRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoryBlocksRegistryTransactorRaw struct {
	Contract *StoryBlocksRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStoryBlocksRegistry creates a new instance of StoryBlocksRegistry, bound to a specific deployed contract.
func NewStoryBlocksRegistry(address common.Address, backend bind.ContractBackend) (*StoryBlocksRegistry, error) {
	contract, err := bindStoryBlocksRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StoryBlocksRegistry{StoryBlocksRegistryCaller: StoryBlocksRegistryCaller{contract: contract}, StoryBlocksRegistryTransactor: StoryBlocksRegistryTransactor{contract: contract}, StoryBlocksRegistryFilterer: StoryBlocksRegistryFilterer{contract: contract}}, nil
}

// NewStoryBlocksRegistryCaller creates a new read-only instance of StoryBlocksRegistry, bound to a specific deployed contract.
func NewStoryBlocksRegistryCaller(address common.Address, caller bind.ContractCaller) (*StoryBlocksRegistryCaller, error) {
	contract, err := bindStoryBlocksRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoryBlocksRegistryCaller{contract: contract}, nil
}

// NewStoryBlocksRegistryTransactor creates a new write-only instance of StoryBlocksRegistry, bound to a specific deployed contract.
func NewStoryBlocksRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*StoryBlocksRegistryTransactor, error) {
	contract, err := bindStoryBlocksRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoryBlocksRegistryTransactor{contract: contract}, nil
}

// NewStoryBlocksRegistryFilterer creates a new log filterer instance of StoryBlocksRegistry, bound to a specific deployed contract.
func NewStoryBlocksRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*StoryBlocksRegistryFilterer, error) {
	contract, err := bindStoryBlocksRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoryBlocksRegistryFilterer{contract: contract}, nil
}

// bindStoryBlocksRegistry binds a generic wrapper to an already deployed contract.
func bindStoryBlocksRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoryBlocksRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StoryBlocksRegistry *StoryBlocksRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StoryBlocksRegistry.Contract.StoryBlocksRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StoryBlocksRegistry *StoryBlocksRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.StoryBlocksRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StoryBlocksRegistry *StoryBlocksRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.StoryBlocksRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StoryBlocksRegistry *StoryBlocksRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StoryBlocksRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.contract.Transact(opts, method, params...)
}

// MAXLINKEDATONCE is a free data retrieval call binding the contract method 0x0f7966c0.
//
// Solidity: function MAX_LINKED_AT_ONCE() view returns(uint256)
func (_StoryBlocksRegistry *StoryBlocksRegistryCaller) MAXLINKEDATONCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StoryBlocksRegistry.contract.Call(opts, &out, "MAX_LINKED_AT_ONCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLINKEDATONCE is a free data retrieval call binding the contract method 0x0f7966c0.
//
// Solidity: function MAX_LINKED_AT_ONCE() view returns(uint256)
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) MAXLINKEDATONCE() (*big.Int, error) {
	return _StoryBlocksRegistry.Contract.MAXLINKEDATONCE(&_StoryBlocksRegistry.CallOpts)
}

// MAXLINKEDATONCE is a free data retrieval call binding the contract method 0x0f7966c0.
//
// Solidity: function MAX_LINKED_AT_ONCE() view returns(uint256)
func (_StoryBlocksRegistry *StoryBlocksRegistryCallerSession) MAXLINKEDATONCE() (*big.Int, error) {
	return _StoryBlocksRegistry.Contract.MAXLINKEDATONCE(&_StoryBlocksRegistry.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_StoryBlocksRegistry *StoryBlocksRegistryCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StoryBlocksRegistry.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _StoryBlocksRegistry.Contract.BalanceOf(&_StoryBlocksRegistry.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_StoryBlocksRegistry *StoryBlocksRegistryCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _StoryBlocksRegistry.Contract.BalanceOf(&_StoryBlocksRegistry.CallOpts, owner)
}

// CurrentIdFor is a free data retrieval call binding the contract method 0x69cf2c37.
//
// Solidity: function currentIdFor(uint8 sb) view returns(uint256)
func (_StoryBlocksRegistry *StoryBlocksRegistryCaller) CurrentIdFor(opts *bind.CallOpts, sb uint8) (*big.Int, error) {
	var out []interface{}
	err := _StoryBlocksRegistry.contract.Call(opts, &out, "currentIdFor", sb)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentIdFor is a free data retrieval call binding the contract method 0x69cf2c37.
//
// Solidity: function currentIdFor(uint8 sb) view returns(uint256)
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) CurrentIdFor(sb uint8) (*big.Int, error) {
	return _StoryBlocksRegistry.Contract.CurrentIdFor(&_StoryBlocksRegistry.CallOpts, sb)
}

// CurrentIdFor is a free data retrieval call binding the contract method 0x69cf2c37.
//
// Solidity: function currentIdFor(uint8 sb) view returns(uint256)
func (_StoryBlocksRegistry *StoryBlocksRegistryCallerSession) CurrentIdFor(sb uint8) (*big.Int, error) {
	return _StoryBlocksRegistry.Contract.CurrentIdFor(&_StoryBlocksRegistry.CallOpts, sb)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_StoryBlocksRegistry *StoryBlocksRegistryCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StoryBlocksRegistry.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) Description() (string, error) {
	return _StoryBlocksRegistry.Contract.Description(&_StoryBlocksRegistry.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_StoryBlocksRegistry *StoryBlocksRegistryCallerSession) Description() (string, error) {
	return _StoryBlocksRegistry.Contract.Description(&_StoryBlocksRegistry.CallOpts)
}

// FranchiseId is a free data retrieval call binding the contract method 0xb9d0787c.
//
// Solidity: function franchiseId() view returns(uint256)
func (_StoryBlocksRegistry *StoryBlocksRegistryCaller) FranchiseId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StoryBlocksRegistry.contract.Call(opts, &out, "franchiseId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FranchiseId is a free data retrieval call binding the contract method 0xb9d0787c.
//
// Solidity: function franchiseId() view returns(uint256)
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) FranchiseId() (*big.Int, error) {
	return _StoryBlocksRegistry.Contract.FranchiseId(&_StoryBlocksRegistry.CallOpts)
}

// FranchiseId is a free data retrieval call binding the contract method 0xb9d0787c.
//
// Solidity: function franchiseId() view returns(uint256)
func (_StoryBlocksRegistry *StoryBlocksRegistryCallerSession) FranchiseId() (*big.Int, error) {
	return _StoryBlocksRegistry.Contract.FranchiseId(&_StoryBlocksRegistry.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_StoryBlocksRegistry *StoryBlocksRegistryCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StoryBlocksRegistry.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _StoryBlocksRegistry.Contract.GetApproved(&_StoryBlocksRegistry.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_StoryBlocksRegistry *StoryBlocksRegistryCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _StoryBlocksRegistry.Contract.GetApproved(&_StoryBlocksRegistry.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_StoryBlocksRegistry *StoryBlocksRegistryCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _StoryBlocksRegistry.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _StoryBlocksRegistry.Contract.IsApprovedForAll(&_StoryBlocksRegistry.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_StoryBlocksRegistry *StoryBlocksRegistryCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _StoryBlocksRegistry.Contract.IsApprovedForAll(&_StoryBlocksRegistry.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StoryBlocksRegistry *StoryBlocksRegistryCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StoryBlocksRegistry.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) Name() (string, error) {
	return _StoryBlocksRegistry.Contract.Name(&_StoryBlocksRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StoryBlocksRegistry *StoryBlocksRegistryCallerSession) Name() (string, error) {
	return _StoryBlocksRegistry.Contract.Name(&_StoryBlocksRegistry.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_StoryBlocksRegistry *StoryBlocksRegistryCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StoryBlocksRegistry.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _StoryBlocksRegistry.Contract.OwnerOf(&_StoryBlocksRegistry.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_StoryBlocksRegistry *StoryBlocksRegistryCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _StoryBlocksRegistry.Contract.OwnerOf(&_StoryBlocksRegistry.CallOpts, tokenId)
}

// ReadGroup is a free data retrieval call binding the contract method 0x3c7fc59d.
//
// Solidity: function readGroup(uint256 id) view returns((string,string,string,uint8) blockData, uint8 linkedType, uint256[] linkedItems)
func (_StoryBlocksRegistry *StoryBlocksRegistryCaller) ReadGroup(opts *bind.CallOpts, id *big.Int) (struct {
	BlockData   IStoryBlockStorageStoryBlockData
	LinkedType  uint8
	LinkedItems []*big.Int
}, error) {
	var out []interface{}
	err := _StoryBlocksRegistry.contract.Call(opts, &out, "readGroup", id)

	outstruct := new(struct {
		BlockData   IStoryBlockStorageStoryBlockData
		LinkedType  uint8
		LinkedItems []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockData = *abi.ConvertType(out[0], new(IStoryBlockStorageStoryBlockData)).(*IStoryBlockStorageStoryBlockData)
	outstruct.LinkedType = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.LinkedItems = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// ReadGroup is a free data retrieval call binding the contract method 0x3c7fc59d.
//
// Solidity: function readGroup(uint256 id) view returns((string,string,string,uint8) blockData, uint8 linkedType, uint256[] linkedItems)
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) ReadGroup(id *big.Int) (struct {
	BlockData   IStoryBlockStorageStoryBlockData
	LinkedType  uint8
	LinkedItems []*big.Int
}, error) {
	return _StoryBlocksRegistry.Contract.ReadGroup(&_StoryBlocksRegistry.CallOpts, id)
}

// ReadGroup is a free data retrieval call binding the contract method 0x3c7fc59d.
//
// Solidity: function readGroup(uint256 id) view returns((string,string,string,uint8) blockData, uint8 linkedType, uint256[] linkedItems)
func (_StoryBlocksRegistry *StoryBlocksRegistryCallerSession) ReadGroup(id *big.Int) (struct {
	BlockData   IStoryBlockStorageStoryBlockData
	LinkedType  uint8
	LinkedItems []*big.Int
}, error) {
	return _StoryBlocksRegistry.Contract.ReadGroup(&_StoryBlocksRegistry.CallOpts, id)
}

// ReadStoryBlock is a free data retrieval call binding the contract method 0x9484c103.
//
// Solidity: function readStoryBlock(uint256 storyBlockId) view returns((string,string,string,uint8))
func (_StoryBlocksRegistry *StoryBlocksRegistryCaller) ReadStoryBlock(opts *bind.CallOpts, storyBlockId *big.Int) (IStoryBlockStorageStoryBlockData, error) {
	var out []interface{}
	err := _StoryBlocksRegistry.contract.Call(opts, &out, "readStoryBlock", storyBlockId)

	if err != nil {
		return *new(IStoryBlockStorageStoryBlockData), err
	}

	out0 := *abi.ConvertType(out[0], new(IStoryBlockStorageStoryBlockData)).(*IStoryBlockStorageStoryBlockData)

	return out0, err

}

// ReadStoryBlock is a free data retrieval call binding the contract method 0x9484c103.
//
// Solidity: function readStoryBlock(uint256 storyBlockId) view returns((string,string,string,uint8))
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) ReadStoryBlock(storyBlockId *big.Int) (IStoryBlockStorageStoryBlockData, error) {
	return _StoryBlocksRegistry.Contract.ReadStoryBlock(&_StoryBlocksRegistry.CallOpts, storyBlockId)
}

// ReadStoryBlock is a free data retrieval call binding the contract method 0x9484c103.
//
// Solidity: function readStoryBlock(uint256 storyBlockId) view returns((string,string,string,uint8))
func (_StoryBlocksRegistry *StoryBlocksRegistryCallerSession) ReadStoryBlock(storyBlockId *big.Int) (IStoryBlockStorageStoryBlockData, error) {
	return _StoryBlocksRegistry.Contract.ReadStoryBlock(&_StoryBlocksRegistry.CallOpts, storyBlockId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StoryBlocksRegistry *StoryBlocksRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StoryBlocksRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StoryBlocksRegistry.Contract.SupportsInterface(&_StoryBlocksRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StoryBlocksRegistry *StoryBlocksRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StoryBlocksRegistry.Contract.SupportsInterface(&_StoryBlocksRegistry.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StoryBlocksRegistry *StoryBlocksRegistryCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StoryBlocksRegistry.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) Symbol() (string, error) {
	return _StoryBlocksRegistry.Contract.Symbol(&_StoryBlocksRegistry.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StoryBlocksRegistry *StoryBlocksRegistryCallerSession) Symbol() (string, error) {
	return _StoryBlocksRegistry.Contract.Symbol(&_StoryBlocksRegistry.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_StoryBlocksRegistry *StoryBlocksRegistryCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _StoryBlocksRegistry.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) TokenURI(tokenId *big.Int) (string, error) {
	return _StoryBlocksRegistry.Contract.TokenURI(&_StoryBlocksRegistry.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_StoryBlocksRegistry *StoryBlocksRegistryCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _StoryBlocksRegistry.Contract.TokenURI(&_StoryBlocksRegistry.CallOpts, tokenId)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_StoryBlocksRegistry *StoryBlocksRegistryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StoryBlocksRegistry.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) Version() (string, error) {
	return _StoryBlocksRegistry.Contract.Version(&_StoryBlocksRegistry.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_StoryBlocksRegistry *StoryBlocksRegistryCallerSession) Version() (string, error) {
	return _StoryBlocksRegistry.Contract.Version(&_StoryBlocksRegistry.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StoryBlocksRegistry.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.Approve(&_StoryBlocksRegistry.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.Approve(&_StoryBlocksRegistry.TransactOpts, to, tokenId)
}

// CreateGroup is a paid mutator transaction binding the contract method 0xa8185a2e.
//
// Solidity: function createGroup(string name, string _description, string mediaUrl, uint8 linkedType, uint256[] linkedItems) returns(uint256)
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactor) CreateGroup(opts *bind.TransactOpts, name string, _description string, mediaUrl string, linkedType uint8, linkedItems []*big.Int) (*types.Transaction, error) {
	return _StoryBlocksRegistry.contract.Transact(opts, "createGroup", name, _description, mediaUrl, linkedType, linkedItems)
}

// CreateGroup is a paid mutator transaction binding the contract method 0xa8185a2e.
//
// Solidity: function createGroup(string name, string _description, string mediaUrl, uint8 linkedType, uint256[] linkedItems) returns(uint256)
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) CreateGroup(name string, _description string, mediaUrl string, linkedType uint8, linkedItems []*big.Int) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.CreateGroup(&_StoryBlocksRegistry.TransactOpts, name, _description, mediaUrl, linkedType, linkedItems)
}

// CreateGroup is a paid mutator transaction binding the contract method 0xa8185a2e.
//
// Solidity: function createGroup(string name, string _description, string mediaUrl, uint8 linkedType, uint256[] linkedItems) returns(uint256)
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactorSession) CreateGroup(name string, _description string, mediaUrl string, linkedType uint8, linkedItems []*big.Int) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.CreateGroup(&_StoryBlocksRegistry.TransactOpts, name, _description, mediaUrl, linkedType, linkedItems)
}

// CreateStoryBlock is a paid mutator transaction binding the contract method 0x2336a84a.
//
// Solidity: function createStoryBlock(uint8 sb, string name, string _description, string mediaUrl) returns(uint256)
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactor) CreateStoryBlock(opts *bind.TransactOpts, sb uint8, name string, _description string, mediaUrl string) (*types.Transaction, error) {
	return _StoryBlocksRegistry.contract.Transact(opts, "createStoryBlock", sb, name, _description, mediaUrl)
}

// CreateStoryBlock is a paid mutator transaction binding the contract method 0x2336a84a.
//
// Solidity: function createStoryBlock(uint8 sb, string name, string _description, string mediaUrl) returns(uint256)
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) CreateStoryBlock(sb uint8, name string, _description string, mediaUrl string) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.CreateStoryBlock(&_StoryBlocksRegistry.TransactOpts, sb, name, _description, mediaUrl)
}

// CreateStoryBlock is a paid mutator transaction binding the contract method 0x2336a84a.
//
// Solidity: function createStoryBlock(uint8 sb, string name, string _description, string mediaUrl) returns(uint256)
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactorSession) CreateStoryBlock(sb uint8, name string, _description string, mediaUrl string) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.CreateStoryBlock(&_StoryBlocksRegistry.TransactOpts, sb, name, _description, mediaUrl)
}

// GroupItems is a paid mutator transaction binding the contract method 0x9dc6693c.
//
// Solidity: function groupItems(uint256 id, uint256[] linkedItems) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactor) GroupItems(opts *bind.TransactOpts, id *big.Int, linkedItems []*big.Int) (*types.Transaction, error) {
	return _StoryBlocksRegistry.contract.Transact(opts, "groupItems", id, linkedItems)
}

// GroupItems is a paid mutator transaction binding the contract method 0x9dc6693c.
//
// Solidity: function groupItems(uint256 id, uint256[] linkedItems) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) GroupItems(id *big.Int, linkedItems []*big.Int) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.GroupItems(&_StoryBlocksRegistry.TransactOpts, id, linkedItems)
}

// GroupItems is a paid mutator transaction binding the contract method 0x9dc6693c.
//
// Solidity: function groupItems(uint256 id, uint256[] linkedItems) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactorSession) GroupItems(id *big.Int, linkedItems []*big.Int) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.GroupItems(&_StoryBlocksRegistry.TransactOpts, id, linkedItems)
}

// Initialize is a paid mutator transaction binding the contract method 0x46193ccd.
//
// Solidity: function initialize(uint256 _franchiseId, string _name, string _symbol, string _description) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactor) Initialize(opts *bind.TransactOpts, _franchiseId *big.Int, _name string, _symbol string, _description string) (*types.Transaction, error) {
	return _StoryBlocksRegistry.contract.Transact(opts, "initialize", _franchiseId, _name, _symbol, _description)
}

// Initialize is a paid mutator transaction binding the contract method 0x46193ccd.
//
// Solidity: function initialize(uint256 _franchiseId, string _name, string _symbol, string _description) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) Initialize(_franchiseId *big.Int, _name string, _symbol string, _description string) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.Initialize(&_StoryBlocksRegistry.TransactOpts, _franchiseId, _name, _symbol, _description)
}

// Initialize is a paid mutator transaction binding the contract method 0x46193ccd.
//
// Solidity: function initialize(uint256 _franchiseId, string _name, string _symbol, string _description) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactorSession) Initialize(_franchiseId *big.Int, _name string, _symbol string, _description string) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.Initialize(&_StoryBlocksRegistry.TransactOpts, _franchiseId, _name, _symbol, _description)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _StoryBlocksRegistry.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.Multicall(&_StoryBlocksRegistry.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.Multicall(&_StoryBlocksRegistry.TransactOpts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StoryBlocksRegistry.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.SafeTransferFrom(&_StoryBlocksRegistry.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.SafeTransferFrom(&_StoryBlocksRegistry.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _StoryBlocksRegistry.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.SafeTransferFrom0(&_StoryBlocksRegistry.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.SafeTransferFrom0(&_StoryBlocksRegistry.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _StoryBlocksRegistry.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.SetApprovalForAll(&_StoryBlocksRegistry.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.SetApprovalForAll(&_StoryBlocksRegistry.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StoryBlocksRegistry.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistrySession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.TransferFrom(&_StoryBlocksRegistry.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_StoryBlocksRegistry *StoryBlocksRegistryTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StoryBlocksRegistry.Contract.TransferFrom(&_StoryBlocksRegistry.TransactOpts, from, to, tokenId)
}

// StoryBlocksRegistryApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StoryBlocksRegistry contract.
type StoryBlocksRegistryApprovalIterator struct {
	Event *StoryBlocksRegistryApproval // Event containing the contract specifics and raw log

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
func (it *StoryBlocksRegistryApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoryBlocksRegistryApproval)
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
		it.Event = new(StoryBlocksRegistryApproval)
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
func (it *StoryBlocksRegistryApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoryBlocksRegistryApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoryBlocksRegistryApproval represents a Approval event raised by the StoryBlocksRegistry contract.
type StoryBlocksRegistryApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*StoryBlocksRegistryApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _StoryBlocksRegistry.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StoryBlocksRegistryApprovalIterator{contract: _StoryBlocksRegistry.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StoryBlocksRegistryApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _StoryBlocksRegistry.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoryBlocksRegistryApproval)
				if err := _StoryBlocksRegistry.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) ParseApproval(log types.Log) (*StoryBlocksRegistryApproval, error) {
	event := new(StoryBlocksRegistryApproval)
	if err := _StoryBlocksRegistry.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoryBlocksRegistryApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the StoryBlocksRegistry contract.
type StoryBlocksRegistryApprovalForAllIterator struct {
	Event *StoryBlocksRegistryApprovalForAll // Event containing the contract specifics and raw log

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
func (it *StoryBlocksRegistryApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoryBlocksRegistryApprovalForAll)
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
		it.Event = new(StoryBlocksRegistryApprovalForAll)
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
func (it *StoryBlocksRegistryApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoryBlocksRegistryApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoryBlocksRegistryApprovalForAll represents a ApprovalForAll event raised by the StoryBlocksRegistry contract.
type StoryBlocksRegistryApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*StoryBlocksRegistryApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _StoryBlocksRegistry.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &StoryBlocksRegistryApprovalForAllIterator{contract: _StoryBlocksRegistry.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *StoryBlocksRegistryApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _StoryBlocksRegistry.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoryBlocksRegistryApprovalForAll)
				if err := _StoryBlocksRegistry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) ParseApprovalForAll(log types.Log) (*StoryBlocksRegistryApprovalForAll, error) {
	event := new(StoryBlocksRegistryApprovalForAll)
	if err := _StoryBlocksRegistry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoryBlocksRegistryGroupedItemsIterator is returned from FilterGroupedItems and is used to iterate over the raw logs and unpacked data for GroupedItems events raised by the StoryBlocksRegistry contract.
type StoryBlocksRegistryGroupedItemsIterator struct {
	Event *StoryBlocksRegistryGroupedItems // Event containing the contract specifics and raw log

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
func (it *StoryBlocksRegistryGroupedItemsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoryBlocksRegistryGroupedItems)
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
		it.Event = new(StoryBlocksRegistryGroupedItems)
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
func (it *StoryBlocksRegistryGroupedItemsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoryBlocksRegistryGroupedItemsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoryBlocksRegistryGroupedItems represents a GroupedItems event raised by the StoryBlocksRegistry contract.
type StoryBlocksRegistryGroupedItems struct {
	Id          *big.Int
	LinkedType  uint8
	LinkedItems []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGroupedItems is a free log retrieval operation binding the contract event 0x823d677ed42b321275c6b3aa7040f32e41d12aea91120455a19a60ae310b0af3.
//
// Solidity: event GroupedItems(uint256 indexed id, uint8 linkedType, uint256[] linkedItems)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) FilterGroupedItems(opts *bind.FilterOpts, id []*big.Int) (*StoryBlocksRegistryGroupedItemsIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _StoryBlocksRegistry.contract.FilterLogs(opts, "GroupedItems", idRule)
	if err != nil {
		return nil, err
	}
	return &StoryBlocksRegistryGroupedItemsIterator{contract: _StoryBlocksRegistry.contract, event: "GroupedItems", logs: logs, sub: sub}, nil
}

// WatchGroupedItems is a free log subscription operation binding the contract event 0x823d677ed42b321275c6b3aa7040f32e41d12aea91120455a19a60ae310b0af3.
//
// Solidity: event GroupedItems(uint256 indexed id, uint8 linkedType, uint256[] linkedItems)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) WatchGroupedItems(opts *bind.WatchOpts, sink chan<- *StoryBlocksRegistryGroupedItems, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _StoryBlocksRegistry.contract.WatchLogs(opts, "GroupedItems", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoryBlocksRegistryGroupedItems)
				if err := _StoryBlocksRegistry.contract.UnpackLog(event, "GroupedItems", log); err != nil {
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

// ParseGroupedItems is a log parse operation binding the contract event 0x823d677ed42b321275c6b3aa7040f32e41d12aea91120455a19a60ae310b0af3.
//
// Solidity: event GroupedItems(uint256 indexed id, uint8 linkedType, uint256[] linkedItems)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) ParseGroupedItems(log types.Log) (*StoryBlocksRegistryGroupedItems, error) {
	event := new(StoryBlocksRegistryGroupedItems)
	if err := _StoryBlocksRegistry.contract.UnpackLog(event, "GroupedItems", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoryBlocksRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StoryBlocksRegistry contract.
type StoryBlocksRegistryInitializedIterator struct {
	Event *StoryBlocksRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *StoryBlocksRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoryBlocksRegistryInitialized)
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
		it.Event = new(StoryBlocksRegistryInitialized)
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
func (it *StoryBlocksRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoryBlocksRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoryBlocksRegistryInitialized represents a Initialized event raised by the StoryBlocksRegistry contract.
type StoryBlocksRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*StoryBlocksRegistryInitializedIterator, error) {

	logs, sub, err := _StoryBlocksRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StoryBlocksRegistryInitializedIterator{contract: _StoryBlocksRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StoryBlocksRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _StoryBlocksRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoryBlocksRegistryInitialized)
				if err := _StoryBlocksRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) ParseInitialized(log types.Log) (*StoryBlocksRegistryInitialized, error) {
	event := new(StoryBlocksRegistryInitialized)
	if err := _StoryBlocksRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoryBlocksRegistryStoryBlockMintedIterator is returned from FilterStoryBlockMinted and is used to iterate over the raw logs and unpacked data for StoryBlockMinted events raised by the StoryBlocksRegistry contract.
type StoryBlocksRegistryStoryBlockMintedIterator struct {
	Event *StoryBlocksRegistryStoryBlockMinted // Event containing the contract specifics and raw log

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
func (it *StoryBlocksRegistryStoryBlockMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoryBlocksRegistryStoryBlockMinted)
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
		it.Event = new(StoryBlocksRegistryStoryBlockMinted)
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
func (it *StoryBlocksRegistryStoryBlockMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoryBlocksRegistryStoryBlockMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoryBlocksRegistryStoryBlockMinted represents a StoryBlockMinted event raised by the StoryBlocksRegistry contract.
type StoryBlocksRegistryStoryBlockMinted struct {
	To      common.Address
	Sb      uint8
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStoryBlockMinted is a free log retrieval operation binding the contract event 0x46c5798f976ec5e5dec617e7f27479bc578f63e6bcf4f693ade547151a3219cd.
//
// Solidity: event StoryBlockMinted(address indexed to, uint8 indexed sb, uint256 indexed tokenId)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) FilterStoryBlockMinted(opts *bind.FilterOpts, to []common.Address, sb []uint8, tokenId []*big.Int) (*StoryBlocksRegistryStoryBlockMintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var sbRule []interface{}
	for _, sbItem := range sb {
		sbRule = append(sbRule, sbItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _StoryBlocksRegistry.contract.FilterLogs(opts, "StoryBlockMinted", toRule, sbRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StoryBlocksRegistryStoryBlockMintedIterator{contract: _StoryBlocksRegistry.contract, event: "StoryBlockMinted", logs: logs, sub: sub}, nil
}

// WatchStoryBlockMinted is a free log subscription operation binding the contract event 0x46c5798f976ec5e5dec617e7f27479bc578f63e6bcf4f693ade547151a3219cd.
//
// Solidity: event StoryBlockMinted(address indexed to, uint8 indexed sb, uint256 indexed tokenId)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) WatchStoryBlockMinted(opts *bind.WatchOpts, sink chan<- *StoryBlocksRegistryStoryBlockMinted, to []common.Address, sb []uint8, tokenId []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var sbRule []interface{}
	for _, sbItem := range sb {
		sbRule = append(sbRule, sbItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _StoryBlocksRegistry.contract.WatchLogs(opts, "StoryBlockMinted", toRule, sbRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoryBlocksRegistryStoryBlockMinted)
				if err := _StoryBlocksRegistry.contract.UnpackLog(event, "StoryBlockMinted", log); err != nil {
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

// ParseStoryBlockMinted is a log parse operation binding the contract event 0x46c5798f976ec5e5dec617e7f27479bc578f63e6bcf4f693ade547151a3219cd.
//
// Solidity: event StoryBlockMinted(address indexed to, uint8 indexed sb, uint256 indexed tokenId)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) ParseStoryBlockMinted(log types.Log) (*StoryBlocksRegistryStoryBlockMinted, error) {
	event := new(StoryBlocksRegistryStoryBlockMinted)
	if err := _StoryBlocksRegistry.contract.UnpackLog(event, "StoryBlockMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoryBlocksRegistryStoryBlockWrittenIterator is returned from FilterStoryBlockWritten and is used to iterate over the raw logs and unpacked data for StoryBlockWritten events raised by the StoryBlocksRegistry contract.
type StoryBlocksRegistryStoryBlockWrittenIterator struct {
	Event *StoryBlocksRegistryStoryBlockWritten // Event containing the contract specifics and raw log

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
func (it *StoryBlocksRegistryStoryBlockWrittenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoryBlocksRegistryStoryBlockWritten)
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
		it.Event = new(StoryBlocksRegistryStoryBlockWritten)
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
func (it *StoryBlocksRegistryStoryBlockWrittenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoryBlocksRegistryStoryBlockWrittenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoryBlocksRegistryStoryBlockWritten represents a StoryBlockWritten event raised by the StoryBlocksRegistry contract.
type StoryBlocksRegistryStoryBlockWritten struct {
	StoryBlockId *big.Int
	BlockType    uint8
	Name         string
	Description  string
	MediaUrl     string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStoryBlockWritten is a free log retrieval operation binding the contract event 0x621ca783668a03c147f29e5ae01f40cc55df212956bde68afb53b1d6119a2ea7.
//
// Solidity: event StoryBlockWritten(uint256 indexed storyBlockId, uint8 indexed blockType, string name, string description, string mediaUrl)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) FilterStoryBlockWritten(opts *bind.FilterOpts, storyBlockId []*big.Int, blockType []uint8) (*StoryBlocksRegistryStoryBlockWrittenIterator, error) {

	var storyBlockIdRule []interface{}
	for _, storyBlockIdItem := range storyBlockId {
		storyBlockIdRule = append(storyBlockIdRule, storyBlockIdItem)
	}
	var blockTypeRule []interface{}
	for _, blockTypeItem := range blockType {
		blockTypeRule = append(blockTypeRule, blockTypeItem)
	}

	logs, sub, err := _StoryBlocksRegistry.contract.FilterLogs(opts, "StoryBlockWritten", storyBlockIdRule, blockTypeRule)
	if err != nil {
		return nil, err
	}
	return &StoryBlocksRegistryStoryBlockWrittenIterator{contract: _StoryBlocksRegistry.contract, event: "StoryBlockWritten", logs: logs, sub: sub}, nil
}

// WatchStoryBlockWritten is a free log subscription operation binding the contract event 0x621ca783668a03c147f29e5ae01f40cc55df212956bde68afb53b1d6119a2ea7.
//
// Solidity: event StoryBlockWritten(uint256 indexed storyBlockId, uint8 indexed blockType, string name, string description, string mediaUrl)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) WatchStoryBlockWritten(opts *bind.WatchOpts, sink chan<- *StoryBlocksRegistryStoryBlockWritten, storyBlockId []*big.Int, blockType []uint8) (event.Subscription, error) {

	var storyBlockIdRule []interface{}
	for _, storyBlockIdItem := range storyBlockId {
		storyBlockIdRule = append(storyBlockIdRule, storyBlockIdItem)
	}
	var blockTypeRule []interface{}
	for _, blockTypeItem := range blockType {
		blockTypeRule = append(blockTypeRule, blockTypeItem)
	}

	logs, sub, err := _StoryBlocksRegistry.contract.WatchLogs(opts, "StoryBlockWritten", storyBlockIdRule, blockTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoryBlocksRegistryStoryBlockWritten)
				if err := _StoryBlocksRegistry.contract.UnpackLog(event, "StoryBlockWritten", log); err != nil {
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

// ParseStoryBlockWritten is a log parse operation binding the contract event 0x621ca783668a03c147f29e5ae01f40cc55df212956bde68afb53b1d6119a2ea7.
//
// Solidity: event StoryBlockWritten(uint256 indexed storyBlockId, uint8 indexed blockType, string name, string description, string mediaUrl)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) ParseStoryBlockWritten(log types.Log) (*StoryBlocksRegistryStoryBlockWritten, error) {
	event := new(StoryBlocksRegistryStoryBlockWritten)
	if err := _StoryBlocksRegistry.contract.UnpackLog(event, "StoryBlockWritten", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoryBlocksRegistryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StoryBlocksRegistry contract.
type StoryBlocksRegistryTransferIterator struct {
	Event *StoryBlocksRegistryTransfer // Event containing the contract specifics and raw log

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
func (it *StoryBlocksRegistryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoryBlocksRegistryTransfer)
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
		it.Event = new(StoryBlocksRegistryTransfer)
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
func (it *StoryBlocksRegistryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoryBlocksRegistryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoryBlocksRegistryTransfer represents a Transfer event raised by the StoryBlocksRegistry contract.
type StoryBlocksRegistryTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*StoryBlocksRegistryTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _StoryBlocksRegistry.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StoryBlocksRegistryTransferIterator{contract: _StoryBlocksRegistry.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StoryBlocksRegistryTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _StoryBlocksRegistry.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoryBlocksRegistryTransfer)
				if err := _StoryBlocksRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_StoryBlocksRegistry *StoryBlocksRegistryFilterer) ParseTransfer(log types.Log) (*StoryBlocksRegistryTransfer, error) {
	event := new(StoryBlocksRegistryTransfer)
	if err := _StoryBlocksRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
