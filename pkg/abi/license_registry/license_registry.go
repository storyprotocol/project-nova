// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package license_registry

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

// ILicenseRegistryNFTLicenseInfo is an auto generated low-level Go binding around an user-defined struct.
type ILicenseRegistryNFTLicenseInfo struct {
	Policy     common.Address
	PolicyData []byte
}

// LicenseRegistryMetaData contains all meta data concerning the LicenseRegistry contract.
var LicenseRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"contractLicenseRepository\",\"name\":\"_licenseRepo\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"}],\"name\":\"LicenseAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"grantedId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetIds\",\"type\":\"uint256\"}],\"name\":\"LicenseGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assetsUnderLicense\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractILicenseGrantingPolicy\",\"name\":\"policy\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"policyData\",\"type\":\"bytes\"}],\"internalType\":\"structILicenseRegistryNFT.LicenseInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"assignLicense\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"assignedLicenseFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"contractILicenseGrantingPolicy\",\"name\":\"policy\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"policyData\",\"type\":\"bytes\"}],\"internalType\":\"structILicenseRegistryNFT.LicenseInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"assetIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"grantLicense\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"grantedLicenseForAssetAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"grantedLicenses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"grantedLicensesForAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"grantedId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"isAssetIncludedInLicense\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"licenseParams\",\"outputs\":[{\"internalType\":\"contractILicenseGrantingPolicy\",\"name\":\"policy\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"policyData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LicenseRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use LicenseRegistryMetaData.ABI instead.
var LicenseRegistryABI = LicenseRegistryMetaData.ABI

// LicenseRegistry is an auto generated Go binding around an Ethereum contract.
type LicenseRegistry struct {
	LicenseRegistryCaller     // Read-only binding to the contract
	LicenseRegistryTransactor // Write-only binding to the contract
	LicenseRegistryFilterer   // Log filterer for contract events
}

// LicenseRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LicenseRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LicenseRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LicenseRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LicenseRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LicenseRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LicenseRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LicenseRegistrySession struct {
	Contract     *LicenseRegistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LicenseRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LicenseRegistryCallerSession struct {
	Contract *LicenseRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// LicenseRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LicenseRegistryTransactorSession struct {
	Contract     *LicenseRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// LicenseRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LicenseRegistryRaw struct {
	Contract *LicenseRegistry // Generic contract binding to access the raw methods on
}

// LicenseRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LicenseRegistryCallerRaw struct {
	Contract *LicenseRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// LicenseRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LicenseRegistryTransactorRaw struct {
	Contract *LicenseRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLicenseRegistry creates a new instance of LicenseRegistry, bound to a specific deployed contract.
func NewLicenseRegistry(address common.Address, backend bind.ContractBackend) (*LicenseRegistry, error) {
	contract, err := bindLicenseRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LicenseRegistry{LicenseRegistryCaller: LicenseRegistryCaller{contract: contract}, LicenseRegistryTransactor: LicenseRegistryTransactor{contract: contract}, LicenseRegistryFilterer: LicenseRegistryFilterer{contract: contract}}, nil
}

// NewLicenseRegistryCaller creates a new read-only instance of LicenseRegistry, bound to a specific deployed contract.
func NewLicenseRegistryCaller(address common.Address, caller bind.ContractCaller) (*LicenseRegistryCaller, error) {
	contract, err := bindLicenseRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LicenseRegistryCaller{contract: contract}, nil
}

// NewLicenseRegistryTransactor creates a new write-only instance of LicenseRegistry, bound to a specific deployed contract.
func NewLicenseRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*LicenseRegistryTransactor, error) {
	contract, err := bindLicenseRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LicenseRegistryTransactor{contract: contract}, nil
}

// NewLicenseRegistryFilterer creates a new log filterer instance of LicenseRegistry, bound to a specific deployed contract.
func NewLicenseRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*LicenseRegistryFilterer, error) {
	contract, err := bindLicenseRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LicenseRegistryFilterer{contract: contract}, nil
}

// bindLicenseRegistry binds a generic wrapper to an already deployed contract.
func bindLicenseRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LicenseRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LicenseRegistry *LicenseRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LicenseRegistry.Contract.LicenseRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LicenseRegistry *LicenseRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LicenseRegistry.Contract.LicenseRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LicenseRegistry *LicenseRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LicenseRegistry.Contract.LicenseRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LicenseRegistry *LicenseRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LicenseRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LicenseRegistry *LicenseRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LicenseRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LicenseRegistry *LicenseRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LicenseRegistry.Contract.contract.Transact(opts, method, params...)
}

// AssetsUnderLicense is a free data retrieval call binding the contract method 0xf6b5ee42.
//
// Solidity: function assetsUnderLicense(address , uint256 ) view returns(uint256)
func (_LicenseRegistry *LicenseRegistryCaller) AssetsUnderLicense(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LicenseRegistry.contract.Call(opts, &out, "assetsUnderLicense", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssetsUnderLicense is a free data retrieval call binding the contract method 0xf6b5ee42.
//
// Solidity: function assetsUnderLicense(address , uint256 ) view returns(uint256)
func (_LicenseRegistry *LicenseRegistrySession) AssetsUnderLicense(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _LicenseRegistry.Contract.AssetsUnderLicense(&_LicenseRegistry.CallOpts, arg0, arg1)
}

// AssetsUnderLicense is a free data retrieval call binding the contract method 0xf6b5ee42.
//
// Solidity: function assetsUnderLicense(address , uint256 ) view returns(uint256)
func (_LicenseRegistry *LicenseRegistryCallerSession) AssetsUnderLicense(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _LicenseRegistry.Contract.AssetsUnderLicense(&_LicenseRegistry.CallOpts, arg0, arg1)
}

// AssignedLicenseFor is a free data retrieval call binding the contract method 0x0f83af70.
//
// Solidity: function assignedLicenseFor(address collection, uint256 assetId) view returns(uint256 licenseId, (address,bytes))
func (_LicenseRegistry *LicenseRegistryCaller) AssignedLicenseFor(opts *bind.CallOpts, collection common.Address, assetId *big.Int) (*big.Int, ILicenseRegistryNFTLicenseInfo, error) {
	var out []interface{}
	err := _LicenseRegistry.contract.Call(opts, &out, "assignedLicenseFor", collection, assetId)

	if err != nil {
		return *new(*big.Int), *new(ILicenseRegistryNFTLicenseInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(ILicenseRegistryNFTLicenseInfo)).(*ILicenseRegistryNFTLicenseInfo)

	return out0, out1, err

}

// AssignedLicenseFor is a free data retrieval call binding the contract method 0x0f83af70.
//
// Solidity: function assignedLicenseFor(address collection, uint256 assetId) view returns(uint256 licenseId, (address,bytes))
func (_LicenseRegistry *LicenseRegistrySession) AssignedLicenseFor(collection common.Address, assetId *big.Int) (*big.Int, ILicenseRegistryNFTLicenseInfo, error) {
	return _LicenseRegistry.Contract.AssignedLicenseFor(&_LicenseRegistry.CallOpts, collection, assetId)
}

// AssignedLicenseFor is a free data retrieval call binding the contract method 0x0f83af70.
//
// Solidity: function assignedLicenseFor(address collection, uint256 assetId) view returns(uint256 licenseId, (address,bytes))
func (_LicenseRegistry *LicenseRegistryCallerSession) AssignedLicenseFor(collection common.Address, assetId *big.Int) (*big.Int, ILicenseRegistryNFTLicenseInfo, error) {
	return _LicenseRegistry.Contract.AssignedLicenseFor(&_LicenseRegistry.CallOpts, collection, assetId)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LicenseRegistry *LicenseRegistryCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LicenseRegistry.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LicenseRegistry *LicenseRegistrySession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LicenseRegistry.Contract.BalanceOf(&_LicenseRegistry.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LicenseRegistry *LicenseRegistryCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LicenseRegistry.Contract.BalanceOf(&_LicenseRegistry.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LicenseRegistry *LicenseRegistryCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LicenseRegistry.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LicenseRegistry *LicenseRegistrySession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LicenseRegistry.Contract.GetApproved(&_LicenseRegistry.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LicenseRegistry *LicenseRegistryCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LicenseRegistry.Contract.GetApproved(&_LicenseRegistry.CallOpts, tokenId)
}

// GrantedLicenseForAssetAt is a free data retrieval call binding the contract method 0xbf76cf26.
//
// Solidity: function grantedLicenseForAssetAt(address collection, uint256 assetId, uint256 index) view returns(uint256)
func (_LicenseRegistry *LicenseRegistryCaller) GrantedLicenseForAssetAt(opts *bind.CallOpts, collection common.Address, assetId *big.Int, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LicenseRegistry.contract.Call(opts, &out, "grantedLicenseForAssetAt", collection, assetId, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GrantedLicenseForAssetAt is a free data retrieval call binding the contract method 0xbf76cf26.
//
// Solidity: function grantedLicenseForAssetAt(address collection, uint256 assetId, uint256 index) view returns(uint256)
func (_LicenseRegistry *LicenseRegistrySession) GrantedLicenseForAssetAt(collection common.Address, assetId *big.Int, index *big.Int) (*big.Int, error) {
	return _LicenseRegistry.Contract.GrantedLicenseForAssetAt(&_LicenseRegistry.CallOpts, collection, assetId, index)
}

// GrantedLicenseForAssetAt is a free data retrieval call binding the contract method 0xbf76cf26.
//
// Solidity: function grantedLicenseForAssetAt(address collection, uint256 assetId, uint256 index) view returns(uint256)
func (_LicenseRegistry *LicenseRegistryCallerSession) GrantedLicenseForAssetAt(collection common.Address, assetId *big.Int, index *big.Int) (*big.Int, error) {
	return _LicenseRegistry.Contract.GrantedLicenseForAssetAt(&_LicenseRegistry.CallOpts, collection, assetId, index)
}

// GrantedLicenses is a free data retrieval call binding the contract method 0xcdbafc0e.
//
// Solidity: function grantedLicenses(uint256 ) view returns(uint256)
func (_LicenseRegistry *LicenseRegistryCaller) GrantedLicenses(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LicenseRegistry.contract.Call(opts, &out, "grantedLicenses", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GrantedLicenses is a free data retrieval call binding the contract method 0xcdbafc0e.
//
// Solidity: function grantedLicenses(uint256 ) view returns(uint256)
func (_LicenseRegistry *LicenseRegistrySession) GrantedLicenses(arg0 *big.Int) (*big.Int, error) {
	return _LicenseRegistry.Contract.GrantedLicenses(&_LicenseRegistry.CallOpts, arg0)
}

// GrantedLicenses is a free data retrieval call binding the contract method 0xcdbafc0e.
//
// Solidity: function grantedLicenses(uint256 ) view returns(uint256)
func (_LicenseRegistry *LicenseRegistryCallerSession) GrantedLicenses(arg0 *big.Int) (*big.Int, error) {
	return _LicenseRegistry.Contract.GrantedLicenses(&_LicenseRegistry.CallOpts, arg0)
}

// GrantedLicensesForAsset is a free data retrieval call binding the contract method 0x581356c7.
//
// Solidity: function grantedLicensesForAsset(address collection, uint256 assetId) view returns(uint256)
func (_LicenseRegistry *LicenseRegistryCaller) GrantedLicensesForAsset(opts *bind.CallOpts, collection common.Address, assetId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LicenseRegistry.contract.Call(opts, &out, "grantedLicensesForAsset", collection, assetId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GrantedLicensesForAsset is a free data retrieval call binding the contract method 0x581356c7.
//
// Solidity: function grantedLicensesForAsset(address collection, uint256 assetId) view returns(uint256)
func (_LicenseRegistry *LicenseRegistrySession) GrantedLicensesForAsset(collection common.Address, assetId *big.Int) (*big.Int, error) {
	return _LicenseRegistry.Contract.GrantedLicensesForAsset(&_LicenseRegistry.CallOpts, collection, assetId)
}

// GrantedLicensesForAsset is a free data retrieval call binding the contract method 0x581356c7.
//
// Solidity: function grantedLicensesForAsset(address collection, uint256 assetId) view returns(uint256)
func (_LicenseRegistry *LicenseRegistryCallerSession) GrantedLicensesForAsset(collection common.Address, assetId *big.Int) (*big.Int, error) {
	return _LicenseRegistry.Contract.GrantedLicensesForAsset(&_LicenseRegistry.CallOpts, collection, assetId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LicenseRegistry *LicenseRegistryCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _LicenseRegistry.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LicenseRegistry *LicenseRegistrySession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LicenseRegistry.Contract.IsApprovedForAll(&_LicenseRegistry.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LicenseRegistry *LicenseRegistryCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LicenseRegistry.Contract.IsApprovedForAll(&_LicenseRegistry.CallOpts, owner, operator)
}

// IsAssetIncludedInLicense is a free data retrieval call binding the contract method 0xfe8aef3a.
//
// Solidity: function isAssetIncludedInLicense(uint256 grantedId, address collection, uint256 assetId) view returns(bool)
func (_LicenseRegistry *LicenseRegistryCaller) IsAssetIncludedInLicense(opts *bind.CallOpts, grantedId *big.Int, collection common.Address, assetId *big.Int) (bool, error) {
	var out []interface{}
	err := _LicenseRegistry.contract.Call(opts, &out, "isAssetIncludedInLicense", grantedId, collection, assetId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssetIncludedInLicense is a free data retrieval call binding the contract method 0xfe8aef3a.
//
// Solidity: function isAssetIncludedInLicense(uint256 grantedId, address collection, uint256 assetId) view returns(bool)
func (_LicenseRegistry *LicenseRegistrySession) IsAssetIncludedInLicense(grantedId *big.Int, collection common.Address, assetId *big.Int) (bool, error) {
	return _LicenseRegistry.Contract.IsAssetIncludedInLicense(&_LicenseRegistry.CallOpts, grantedId, collection, assetId)
}

// IsAssetIncludedInLicense is a free data retrieval call binding the contract method 0xfe8aef3a.
//
// Solidity: function isAssetIncludedInLicense(uint256 grantedId, address collection, uint256 assetId) view returns(bool)
func (_LicenseRegistry *LicenseRegistryCallerSession) IsAssetIncludedInLicense(grantedId *big.Int, collection common.Address, assetId *big.Int) (bool, error) {
	return _LicenseRegistry.Contract.IsAssetIncludedInLicense(&_LicenseRegistry.CallOpts, grantedId, collection, assetId)
}

// LicenseParams is a free data retrieval call binding the contract method 0xce8d19d0.
//
// Solidity: function licenseParams(uint256 ) view returns(address policy, bytes policyData)
func (_LicenseRegistry *LicenseRegistryCaller) LicenseParams(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Policy     common.Address
	PolicyData []byte
}, error) {
	var out []interface{}
	err := _LicenseRegistry.contract.Call(opts, &out, "licenseParams", arg0)

	outstruct := new(struct {
		Policy     common.Address
		PolicyData []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Policy = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PolicyData = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// LicenseParams is a free data retrieval call binding the contract method 0xce8d19d0.
//
// Solidity: function licenseParams(uint256 ) view returns(address policy, bytes policyData)
func (_LicenseRegistry *LicenseRegistrySession) LicenseParams(arg0 *big.Int) (struct {
	Policy     common.Address
	PolicyData []byte
}, error) {
	return _LicenseRegistry.Contract.LicenseParams(&_LicenseRegistry.CallOpts, arg0)
}

// LicenseParams is a free data retrieval call binding the contract method 0xce8d19d0.
//
// Solidity: function licenseParams(uint256 ) view returns(address policy, bytes policyData)
func (_LicenseRegistry *LicenseRegistryCallerSession) LicenseParams(arg0 *big.Int) (struct {
	Policy     common.Address
	PolicyData []byte
}, error) {
	return _LicenseRegistry.Contract.LicenseParams(&_LicenseRegistry.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LicenseRegistry *LicenseRegistryCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LicenseRegistry.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LicenseRegistry *LicenseRegistrySession) Name() (string, error) {
	return _LicenseRegistry.Contract.Name(&_LicenseRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LicenseRegistry *LicenseRegistryCallerSession) Name() (string, error) {
	return _LicenseRegistry.Contract.Name(&_LicenseRegistry.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LicenseRegistry *LicenseRegistryCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LicenseRegistry.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LicenseRegistry *LicenseRegistrySession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LicenseRegistry.Contract.OwnerOf(&_LicenseRegistry.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LicenseRegistry *LicenseRegistryCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LicenseRegistry.Contract.OwnerOf(&_LicenseRegistry.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LicenseRegistry *LicenseRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LicenseRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LicenseRegistry *LicenseRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LicenseRegistry.Contract.SupportsInterface(&_LicenseRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LicenseRegistry *LicenseRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LicenseRegistry.Contract.SupportsInterface(&_LicenseRegistry.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LicenseRegistry *LicenseRegistryCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LicenseRegistry.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LicenseRegistry *LicenseRegistrySession) Symbol() (string, error) {
	return _LicenseRegistry.Contract.Symbol(&_LicenseRegistry.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LicenseRegistry *LicenseRegistryCallerSession) Symbol() (string, error) {
	return _LicenseRegistry.Contract.Symbol(&_LicenseRegistry.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LicenseRegistry *LicenseRegistryCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _LicenseRegistry.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LicenseRegistry *LicenseRegistrySession) TokenURI(tokenId *big.Int) (string, error) {
	return _LicenseRegistry.Contract.TokenURI(&_LicenseRegistry.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LicenseRegistry *LicenseRegistryCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _LicenseRegistry.Contract.TokenURI(&_LicenseRegistry.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LicenseRegistry *LicenseRegistryTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LicenseRegistry.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LicenseRegistry *LicenseRegistrySession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LicenseRegistry.Contract.Approve(&_LicenseRegistry.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LicenseRegistry *LicenseRegistryTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LicenseRegistry.Contract.Approve(&_LicenseRegistry.TransactOpts, to, tokenId)
}

// AssignLicense is a paid mutator transaction binding the contract method 0x7145f8ce.
//
// Solidity: function assignLicense(uint256 licenseId, address collection, uint256 assetId, (address,bytes) info) returns()
func (_LicenseRegistry *LicenseRegistryTransactor) AssignLicense(opts *bind.TransactOpts, licenseId *big.Int, collection common.Address, assetId *big.Int, info ILicenseRegistryNFTLicenseInfo) (*types.Transaction, error) {
	return _LicenseRegistry.contract.Transact(opts, "assignLicense", licenseId, collection, assetId, info)
}

// AssignLicense is a paid mutator transaction binding the contract method 0x7145f8ce.
//
// Solidity: function assignLicense(uint256 licenseId, address collection, uint256 assetId, (address,bytes) info) returns()
func (_LicenseRegistry *LicenseRegistrySession) AssignLicense(licenseId *big.Int, collection common.Address, assetId *big.Int, info ILicenseRegistryNFTLicenseInfo) (*types.Transaction, error) {
	return _LicenseRegistry.Contract.AssignLicense(&_LicenseRegistry.TransactOpts, licenseId, collection, assetId, info)
}

// AssignLicense is a paid mutator transaction binding the contract method 0x7145f8ce.
//
// Solidity: function assignLicense(uint256 licenseId, address collection, uint256 assetId, (address,bytes) info) returns()
func (_LicenseRegistry *LicenseRegistryTransactorSession) AssignLicense(licenseId *big.Int, collection common.Address, assetId *big.Int, info ILicenseRegistryNFTLicenseInfo) (*types.Transaction, error) {
	return _LicenseRegistry.Contract.AssignLicense(&_LicenseRegistry.TransactOpts, licenseId, collection, assetId, info)
}

// GrantLicense is a paid mutator transaction binding the contract method 0xfb84f537.
//
// Solidity: function grantLicense(uint256 licenseId, address collection, uint256[] assetIds, address to) returns(uint256)
func (_LicenseRegistry *LicenseRegistryTransactor) GrantLicense(opts *bind.TransactOpts, licenseId *big.Int, collection common.Address, assetIds []*big.Int, to common.Address) (*types.Transaction, error) {
	return _LicenseRegistry.contract.Transact(opts, "grantLicense", licenseId, collection, assetIds, to)
}

// GrantLicense is a paid mutator transaction binding the contract method 0xfb84f537.
//
// Solidity: function grantLicense(uint256 licenseId, address collection, uint256[] assetIds, address to) returns(uint256)
func (_LicenseRegistry *LicenseRegistrySession) GrantLicense(licenseId *big.Int, collection common.Address, assetIds []*big.Int, to common.Address) (*types.Transaction, error) {
	return _LicenseRegistry.Contract.GrantLicense(&_LicenseRegistry.TransactOpts, licenseId, collection, assetIds, to)
}

// GrantLicense is a paid mutator transaction binding the contract method 0xfb84f537.
//
// Solidity: function grantLicense(uint256 licenseId, address collection, uint256[] assetIds, address to) returns(uint256)
func (_LicenseRegistry *LicenseRegistryTransactorSession) GrantLicense(licenseId *big.Int, collection common.Address, assetIds []*big.Int, to common.Address) (*types.Transaction, error) {
	return _LicenseRegistry.Contract.GrantLicense(&_LicenseRegistry.TransactOpts, licenseId, collection, assetIds, to)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LicenseRegistry *LicenseRegistryTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LicenseRegistry.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LicenseRegistry *LicenseRegistrySession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LicenseRegistry.Contract.SafeTransferFrom(&_LicenseRegistry.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LicenseRegistry *LicenseRegistryTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LicenseRegistry.Contract.SafeTransferFrom(&_LicenseRegistry.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LicenseRegistry *LicenseRegistryTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LicenseRegistry.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LicenseRegistry *LicenseRegistrySession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LicenseRegistry.Contract.SafeTransferFrom0(&_LicenseRegistry.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LicenseRegistry *LicenseRegistryTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LicenseRegistry.Contract.SafeTransferFrom0(&_LicenseRegistry.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LicenseRegistry *LicenseRegistryTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _LicenseRegistry.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LicenseRegistry *LicenseRegistrySession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LicenseRegistry.Contract.SetApprovalForAll(&_LicenseRegistry.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LicenseRegistry *LicenseRegistryTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LicenseRegistry.Contract.SetApprovalForAll(&_LicenseRegistry.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LicenseRegistry *LicenseRegistryTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LicenseRegistry.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LicenseRegistry *LicenseRegistrySession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LicenseRegistry.Contract.TransferFrom(&_LicenseRegistry.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LicenseRegistry *LicenseRegistryTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LicenseRegistry.Contract.TransferFrom(&_LicenseRegistry.TransactOpts, from, to, tokenId)
}

// LicenseRegistryApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LicenseRegistry contract.
type LicenseRegistryApprovalIterator struct {
	Event *LicenseRegistryApproval // Event containing the contract specifics and raw log

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
func (it *LicenseRegistryApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenseRegistryApproval)
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
		it.Event = new(LicenseRegistryApproval)
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
func (it *LicenseRegistryApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenseRegistryApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenseRegistryApproval represents a Approval event raised by the LicenseRegistry contract.
type LicenseRegistryApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LicenseRegistry *LicenseRegistryFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*LicenseRegistryApprovalIterator, error) {

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

	logs, sub, err := _LicenseRegistry.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LicenseRegistryApprovalIterator{contract: _LicenseRegistry.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LicenseRegistry *LicenseRegistryFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LicenseRegistryApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _LicenseRegistry.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenseRegistryApproval)
				if err := _LicenseRegistry.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_LicenseRegistry *LicenseRegistryFilterer) ParseApproval(log types.Log) (*LicenseRegistryApproval, error) {
	event := new(LicenseRegistryApproval)
	if err := _LicenseRegistry.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LicenseRegistryApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the LicenseRegistry contract.
type LicenseRegistryApprovalForAllIterator struct {
	Event *LicenseRegistryApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LicenseRegistryApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenseRegistryApprovalForAll)
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
		it.Event = new(LicenseRegistryApprovalForAll)
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
func (it *LicenseRegistryApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenseRegistryApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenseRegistryApprovalForAll represents a ApprovalForAll event raised by the LicenseRegistry contract.
type LicenseRegistryApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LicenseRegistry *LicenseRegistryFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*LicenseRegistryApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LicenseRegistry.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LicenseRegistryApprovalForAllIterator{contract: _LicenseRegistry.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LicenseRegistry *LicenseRegistryFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LicenseRegistryApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LicenseRegistry.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenseRegistryApprovalForAll)
				if err := _LicenseRegistry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_LicenseRegistry *LicenseRegistryFilterer) ParseApprovalForAll(log types.Log) (*LicenseRegistryApprovalForAll, error) {
	event := new(LicenseRegistryApprovalForAll)
	if err := _LicenseRegistry.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LicenseRegistryLicenseAssignedIterator is returned from FilterLicenseAssigned and is used to iterate over the raw logs and unpacked data for LicenseAssigned events raised by the LicenseRegistry contract.
type LicenseRegistryLicenseAssignedIterator struct {
	Event *LicenseRegistryLicenseAssigned // Event containing the contract specifics and raw log

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
func (it *LicenseRegistryLicenseAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenseRegistryLicenseAssigned)
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
		it.Event = new(LicenseRegistryLicenseAssigned)
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
func (it *LicenseRegistryLicenseAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenseRegistryLicenseAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenseRegistryLicenseAssigned represents a LicenseAssigned event raised by the LicenseRegistry contract.
type LicenseRegistryLicenseAssigned struct {
	Collection common.Address
	AssetId    *big.Int
	LicenseId  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLicenseAssigned is a free log retrieval operation binding the contract event 0x155144a683d52e73f7b0b1df5446c90b63121b52ccaedc4bcb3e0a13503c827a.
//
// Solidity: event LicenseAssigned(address indexed collection, uint256 indexed assetId, uint256 indexed licenseId)
func (_LicenseRegistry *LicenseRegistryFilterer) FilterLicenseAssigned(opts *bind.FilterOpts, collection []common.Address, assetId []*big.Int, licenseId []*big.Int) (*LicenseRegistryLicenseAssignedIterator, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var licenseIdRule []interface{}
	for _, licenseIdItem := range licenseId {
		licenseIdRule = append(licenseIdRule, licenseIdItem)
	}

	logs, sub, err := _LicenseRegistry.contract.FilterLogs(opts, "LicenseAssigned", collectionRule, assetIdRule, licenseIdRule)
	if err != nil {
		return nil, err
	}
	return &LicenseRegistryLicenseAssignedIterator{contract: _LicenseRegistry.contract, event: "LicenseAssigned", logs: logs, sub: sub}, nil
}

// WatchLicenseAssigned is a free log subscription operation binding the contract event 0x155144a683d52e73f7b0b1df5446c90b63121b52ccaedc4bcb3e0a13503c827a.
//
// Solidity: event LicenseAssigned(address indexed collection, uint256 indexed assetId, uint256 indexed licenseId)
func (_LicenseRegistry *LicenseRegistryFilterer) WatchLicenseAssigned(opts *bind.WatchOpts, sink chan<- *LicenseRegistryLicenseAssigned, collection []common.Address, assetId []*big.Int, licenseId []*big.Int) (event.Subscription, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var licenseIdRule []interface{}
	for _, licenseIdItem := range licenseId {
		licenseIdRule = append(licenseIdRule, licenseIdItem)
	}

	logs, sub, err := _LicenseRegistry.contract.WatchLogs(opts, "LicenseAssigned", collectionRule, assetIdRule, licenseIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenseRegistryLicenseAssigned)
				if err := _LicenseRegistry.contract.UnpackLog(event, "LicenseAssigned", log); err != nil {
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

// ParseLicenseAssigned is a log parse operation binding the contract event 0x155144a683d52e73f7b0b1df5446c90b63121b52ccaedc4bcb3e0a13503c827a.
//
// Solidity: event LicenseAssigned(address indexed collection, uint256 indexed assetId, uint256 indexed licenseId)
func (_LicenseRegistry *LicenseRegistryFilterer) ParseLicenseAssigned(log types.Log) (*LicenseRegistryLicenseAssigned, error) {
	event := new(LicenseRegistryLicenseAssigned)
	if err := _LicenseRegistry.contract.UnpackLog(event, "LicenseAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LicenseRegistryLicenseGrantedIterator is returned from FilterLicenseGranted and is used to iterate over the raw logs and unpacked data for LicenseGranted events raised by the LicenseRegistry contract.
type LicenseRegistryLicenseGrantedIterator struct {
	Event *LicenseRegistryLicenseGranted // Event containing the contract specifics and raw log

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
func (it *LicenseRegistryLicenseGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenseRegistryLicenseGranted)
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
		it.Event = new(LicenseRegistryLicenseGranted)
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
func (it *LicenseRegistryLicenseGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenseRegistryLicenseGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenseRegistryLicenseGranted represents a LicenseGranted event raised by the LicenseRegistry contract.
type LicenseRegistryLicenseGranted struct {
	GrantedId  *big.Int
	LicenseId  *big.Int
	Collection common.Address
	AssetIds   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLicenseGranted is a free log retrieval operation binding the contract event 0x5449a3ff3b5cdd7df42db67dad2ae1a31b8e887feea5a67ac8054ccabdf84a42.
//
// Solidity: event LicenseGranted(uint256 indexed grantedId, uint256 indexed licenseId, address indexed collection, uint256 assetIds)
func (_LicenseRegistry *LicenseRegistryFilterer) FilterLicenseGranted(opts *bind.FilterOpts, grantedId []*big.Int, licenseId []*big.Int, collection []common.Address) (*LicenseRegistryLicenseGrantedIterator, error) {

	var grantedIdRule []interface{}
	for _, grantedIdItem := range grantedId {
		grantedIdRule = append(grantedIdRule, grantedIdItem)
	}
	var licenseIdRule []interface{}
	for _, licenseIdItem := range licenseId {
		licenseIdRule = append(licenseIdRule, licenseIdItem)
	}
	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}

	logs, sub, err := _LicenseRegistry.contract.FilterLogs(opts, "LicenseGranted", grantedIdRule, licenseIdRule, collectionRule)
	if err != nil {
		return nil, err
	}
	return &LicenseRegistryLicenseGrantedIterator{contract: _LicenseRegistry.contract, event: "LicenseGranted", logs: logs, sub: sub}, nil
}

// WatchLicenseGranted is a free log subscription operation binding the contract event 0x5449a3ff3b5cdd7df42db67dad2ae1a31b8e887feea5a67ac8054ccabdf84a42.
//
// Solidity: event LicenseGranted(uint256 indexed grantedId, uint256 indexed licenseId, address indexed collection, uint256 assetIds)
func (_LicenseRegistry *LicenseRegistryFilterer) WatchLicenseGranted(opts *bind.WatchOpts, sink chan<- *LicenseRegistryLicenseGranted, grantedId []*big.Int, licenseId []*big.Int, collection []common.Address) (event.Subscription, error) {

	var grantedIdRule []interface{}
	for _, grantedIdItem := range grantedId {
		grantedIdRule = append(grantedIdRule, grantedIdItem)
	}
	var licenseIdRule []interface{}
	for _, licenseIdItem := range licenseId {
		licenseIdRule = append(licenseIdRule, licenseIdItem)
	}
	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}

	logs, sub, err := _LicenseRegistry.contract.WatchLogs(opts, "LicenseGranted", grantedIdRule, licenseIdRule, collectionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenseRegistryLicenseGranted)
				if err := _LicenseRegistry.contract.UnpackLog(event, "LicenseGranted", log); err != nil {
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

// ParseLicenseGranted is a log parse operation binding the contract event 0x5449a3ff3b5cdd7df42db67dad2ae1a31b8e887feea5a67ac8054ccabdf84a42.
//
// Solidity: event LicenseGranted(uint256 indexed grantedId, uint256 indexed licenseId, address indexed collection, uint256 assetIds)
func (_LicenseRegistry *LicenseRegistryFilterer) ParseLicenseGranted(log types.Log) (*LicenseRegistryLicenseGranted, error) {
	event := new(LicenseRegistryLicenseGranted)
	if err := _LicenseRegistry.contract.UnpackLog(event, "LicenseGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LicenseRegistryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LicenseRegistry contract.
type LicenseRegistryTransferIterator struct {
	Event *LicenseRegistryTransfer // Event containing the contract specifics and raw log

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
func (it *LicenseRegistryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenseRegistryTransfer)
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
		it.Event = new(LicenseRegistryTransfer)
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
func (it *LicenseRegistryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenseRegistryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenseRegistryTransfer represents a Transfer event raised by the LicenseRegistry contract.
type LicenseRegistryTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LicenseRegistry *LicenseRegistryFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LicenseRegistryTransferIterator, error) {

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

	logs, sub, err := _LicenseRegistry.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LicenseRegistryTransferIterator{contract: _LicenseRegistry.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LicenseRegistry *LicenseRegistryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LicenseRegistryTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _LicenseRegistry.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenseRegistryTransfer)
				if err := _LicenseRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_LicenseRegistry *LicenseRegistryFilterer) ParseTransfer(log types.Log) (*LicenseRegistryTransfer, error) {
	event := new(LicenseRegistryTransfer)
	if err := _LicenseRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
