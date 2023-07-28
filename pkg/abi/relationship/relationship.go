// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package relationship

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

// IRelationshipModuleRelationshipConfig is an auto generated low-level Go binding around an user-defined struct.
type IRelationshipModuleRelationshipConfig struct {
	SourceIPAssetTypeMask *big.Int
	DestIPAssetTypeMask   *big.Int
	OnlySameFranchise     bool
	Processor             common.Address
	Disputer              common.Address
	TimeConfig            IRelationshipModuleTimeConfig
}

// IRelationshipModuleRelationshipParams is an auto generated low-level Go binding around an user-defined struct.
type IRelationshipModuleRelationshipParams struct {
	SourceContract common.Address
	SourceId       *big.Int
	DestContract   common.Address
	DestId         *big.Int
	RelationshipId [32]byte
	Ttl            *big.Int
}

// IRelationshipModuleSetRelationshipConfigParams is an auto generated low-level Go binding around an user-defined struct.
type IRelationshipModuleSetRelationshipConfigParams struct {
	SourceIPAssets        []uint8
	AllowedExternalSource bool
	DestIPAssets          []uint8
	AllowedExternalDest   bool
	OnlySameFranchise     bool
	Processor             common.Address
	Disputer              common.Address
	TimeConfig            IRelationshipModuleTimeConfig
}

// IRelationshipModuleTimeConfig is an auto generated low-level Go binding around an user-defined struct.
type IRelationshipModuleTimeConfig struct {
	MaxTTL    *big.Int
	MinTTL    *big.Int
	Renewable bool
}

// RelationshipMetaData contains all meta data concerning the Relationship contract.
var RelationshipMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_franchiseRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CannotRelateToOtherFranchise\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IntentAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEndTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidIPAssetArray\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTTL\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MissingRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonExistingRelationship\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"UnsupportedInterface\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedRelationshipDst\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedRelationshipSrc\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accessControl\",\"type\":\"address\"}],\"name\":\"AccessControlUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"relationshipId\",\"type\":\"bytes32\"}],\"name\":\"RelationPendingProcessor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"relationshipId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"RelationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"relationshipId\",\"type\":\"bytes32\"}],\"name\":\"RelationUnset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"relationshipId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceIPAssetTypeMask\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destIPAssetTypeMask\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"onlySameFranchise\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"processor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxTTL\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minTTL\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"renewable\",\"type\":\"bool\"}],\"name\":\"RelationshipConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"relationshipId\",\"type\":\"bytes32\"}],\"name\":\"RelationshipConfigUnset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FRANCHISE_REGISTRY\",\"outputs\":[{\"internalType\":\"contractFranchiseRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accessControl\",\"type\":\"address\"}],\"name\":\"__RelationshipModuleBase_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"relationshipId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structIRelationshipModule.RelationshipParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"areTheyRelated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mask\",\"type\":\"uint256\"}],\"name\":\"convertFromMask\",\"outputs\":[{\"internalType\":\"enumIPAsset[]\",\"name\":\"ipAssets\",\"type\":\"uint8[]\"},{\"internalType\":\"bool\",\"name\":\"allowsExternal\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPAsset[]\",\"name\":\"ipAssets\",\"type\":\"uint8[]\"},{\"internalType\":\"bool\",\"name\":\"allowsExternal\",\"type\":\"bool\"}],\"name\":\"convertToMask\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccessControl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"relationshipId\",\"type\":\"bytes32\"}],\"name\":\"getRelationshipConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceIPAssetTypeMask\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destIPAssetTypeMask\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"onlySameFranchise\",\"type\":\"bool\"},{\"internalType\":\"contractIRelationshipProcessor\",\"name\":\"processor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"disputer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint112\",\"name\":\"maxTTL\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"minTTL\",\"type\":\"uint112\"},{\"internalType\":\"bool\",\"name\":\"renewable\",\"type\":\"bool\"}],\"internalType\":\"structIRelationshipModule.TimeConfig\",\"name\":\"timeConfig\",\"type\":\"tuple\"}],\"internalType\":\"structIRelationshipModule.RelationshipConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"relationshipId\",\"type\":\"bytes32\"}],\"name\":\"getRelationshipConfigDecoded\",\"outputs\":[{\"components\":[{\"internalType\":\"enumIPAsset[]\",\"name\":\"sourceIPAssets\",\"type\":\"uint8[]\"},{\"internalType\":\"bool\",\"name\":\"allowedExternalSource\",\"type\":\"bool\"},{\"internalType\":\"enumIPAsset[]\",\"name\":\"destIPAssets\",\"type\":\"uint8[]\"},{\"internalType\":\"bool\",\"name\":\"allowedExternalDest\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlySameFranchise\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"processor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"disputer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint112\",\"name\":\"maxTTL\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"minTTL\",\"type\":\"uint112\"},{\"internalType\":\"bool\",\"name\":\"renewable\",\"type\":\"bool\"}],\"internalType\":\"structIRelationshipModule.TimeConfig\",\"name\":\"timeConfig\",\"type\":\"tuple\"}],\"internalType\":\"structIRelationshipModule.SetRelationshipConfigParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getRelationshipId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accessControl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"relationshipId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structIRelationshipModule.RelationshipParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"isRelationshipExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"relationshipId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structIRelationshipModule.RelationshipParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"relate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accessControl\",\"type\":\"address\"}],\"name\":\"setAccessControl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"enumIPAsset[]\",\"name\":\"sourceIPAssets\",\"type\":\"uint8[]\"},{\"internalType\":\"bool\",\"name\":\"allowedExternalSource\",\"type\":\"bool\"},{\"internalType\":\"enumIPAsset[]\",\"name\":\"destIPAssets\",\"type\":\"uint8[]\"},{\"internalType\":\"bool\",\"name\":\"allowedExternalDest\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlySameFranchise\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"processor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"disputer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint112\",\"name\":\"maxTTL\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"minTTL\",\"type\":\"uint112\"},{\"internalType\":\"bool\",\"name\":\"renewable\",\"type\":\"bool\"}],\"internalType\":\"structIRelationshipModule.TimeConfig\",\"name\":\"timeConfig\",\"type\":\"tuple\"}],\"internalType\":\"structIRelationshipModule.SetRelationshipConfigParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"setRelationshipConfig\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"relationshipId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mask\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"assetType\",\"type\":\"uint8\"}],\"name\":\"supportsIPAssetType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"relationshipId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"internalType\":\"structIRelationshipModule.RelationshipParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"unrelate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"relationshipId\",\"type\":\"bytes32\"}],\"name\":\"unsetRelationshipConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// RelationshipABI is the input ABI used to generate the binding from.
// Deprecated: Use RelationshipMetaData.ABI instead.
var RelationshipABI = RelationshipMetaData.ABI

// Relationship is an auto generated Go binding around an Ethereum contract.
type Relationship struct {
	RelationshipCaller     // Read-only binding to the contract
	RelationshipTransactor // Write-only binding to the contract
	RelationshipFilterer   // Log filterer for contract events
}

// RelationshipCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelationshipCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelationshipTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelationshipTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelationshipFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelationshipFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelationshipSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelationshipSession struct {
	Contract     *Relationship     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelationshipCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelationshipCallerSession struct {
	Contract *RelationshipCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RelationshipTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelationshipTransactorSession struct {
	Contract     *RelationshipTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RelationshipRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelationshipRaw struct {
	Contract *Relationship // Generic contract binding to access the raw methods on
}

// RelationshipCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelationshipCallerRaw struct {
	Contract *RelationshipCaller // Generic read-only contract binding to access the raw methods on
}

// RelationshipTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelationshipTransactorRaw struct {
	Contract *RelationshipTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelationship creates a new instance of Relationship, bound to a specific deployed contract.
func NewRelationship(address common.Address, backend bind.ContractBackend) (*Relationship, error) {
	contract, err := bindRelationship(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Relationship{RelationshipCaller: RelationshipCaller{contract: contract}, RelationshipTransactor: RelationshipTransactor{contract: contract}, RelationshipFilterer: RelationshipFilterer{contract: contract}}, nil
}

// NewRelationshipCaller creates a new read-only instance of Relationship, bound to a specific deployed contract.
func NewRelationshipCaller(address common.Address, caller bind.ContractCaller) (*RelationshipCaller, error) {
	contract, err := bindRelationship(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelationshipCaller{contract: contract}, nil
}

// NewRelationshipTransactor creates a new write-only instance of Relationship, bound to a specific deployed contract.
func NewRelationshipTransactor(address common.Address, transactor bind.ContractTransactor) (*RelationshipTransactor, error) {
	contract, err := bindRelationship(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelationshipTransactor{contract: contract}, nil
}

// NewRelationshipFilterer creates a new log filterer instance of Relationship, bound to a specific deployed contract.
func NewRelationshipFilterer(address common.Address, filterer bind.ContractFilterer) (*RelationshipFilterer, error) {
	contract, err := bindRelationship(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelationshipFilterer{contract: contract}, nil
}

// bindRelationship binds a generic wrapper to an already deployed contract.
func bindRelationship(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RelationshipMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relationship *RelationshipRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Relationship.Contract.RelationshipCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relationship *RelationshipRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relationship.Contract.RelationshipTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relationship *RelationshipRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relationship.Contract.RelationshipTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relationship *RelationshipCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Relationship.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relationship *RelationshipTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relationship.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relationship *RelationshipTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relationship.Contract.contract.Transact(opts, method, params...)
}

// FRANCHISEREGISTRY is a free data retrieval call binding the contract method 0x94f99b40.
//
// Solidity: function FRANCHISE_REGISTRY() view returns(address)
func (_Relationship *RelationshipCaller) FRANCHISEREGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Relationship.contract.Call(opts, &out, "FRANCHISE_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FRANCHISEREGISTRY is a free data retrieval call binding the contract method 0x94f99b40.
//
// Solidity: function FRANCHISE_REGISTRY() view returns(address)
func (_Relationship *RelationshipSession) FRANCHISEREGISTRY() (common.Address, error) {
	return _Relationship.Contract.FRANCHISEREGISTRY(&_Relationship.CallOpts)
}

// FRANCHISEREGISTRY is a free data retrieval call binding the contract method 0x94f99b40.
//
// Solidity: function FRANCHISE_REGISTRY() view returns(address)
func (_Relationship *RelationshipCallerSession) FRANCHISEREGISTRY() (common.Address, error) {
	return _Relationship.Contract.FRANCHISEREGISTRY(&_Relationship.CallOpts)
}

// AreTheyRelated is a free data retrieval call binding the contract method 0x15023622.
//
// Solidity: function areTheyRelated((address,uint256,address,uint256,bytes32,uint256) params) view returns(bool)
func (_Relationship *RelationshipCaller) AreTheyRelated(opts *bind.CallOpts, params IRelationshipModuleRelationshipParams) (bool, error) {
	var out []interface{}
	err := _Relationship.contract.Call(opts, &out, "areTheyRelated", params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AreTheyRelated is a free data retrieval call binding the contract method 0x15023622.
//
// Solidity: function areTheyRelated((address,uint256,address,uint256,bytes32,uint256) params) view returns(bool)
func (_Relationship *RelationshipSession) AreTheyRelated(params IRelationshipModuleRelationshipParams) (bool, error) {
	return _Relationship.Contract.AreTheyRelated(&_Relationship.CallOpts, params)
}

// AreTheyRelated is a free data retrieval call binding the contract method 0x15023622.
//
// Solidity: function areTheyRelated((address,uint256,address,uint256,bytes32,uint256) params) view returns(bool)
func (_Relationship *RelationshipCallerSession) AreTheyRelated(params IRelationshipModuleRelationshipParams) (bool, error) {
	return _Relationship.Contract.AreTheyRelated(&_Relationship.CallOpts, params)
}

// ConvertFromMask is a free data retrieval call binding the contract method 0x8fa24e0b.
//
// Solidity: function convertFromMask(uint256 mask) pure returns(uint8[] ipAssets, bool allowsExternal)
func (_Relationship *RelationshipCaller) ConvertFromMask(opts *bind.CallOpts, mask *big.Int) (struct {
	IpAssets       []uint8
	AllowsExternal bool
}, error) {
	var out []interface{}
	err := _Relationship.contract.Call(opts, &out, "convertFromMask", mask)

	outstruct := new(struct {
		IpAssets       []uint8
		AllowsExternal bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IpAssets = *abi.ConvertType(out[0], new([]uint8)).(*[]uint8)
	outstruct.AllowsExternal = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// ConvertFromMask is a free data retrieval call binding the contract method 0x8fa24e0b.
//
// Solidity: function convertFromMask(uint256 mask) pure returns(uint8[] ipAssets, bool allowsExternal)
func (_Relationship *RelationshipSession) ConvertFromMask(mask *big.Int) (struct {
	IpAssets       []uint8
	AllowsExternal bool
}, error) {
	return _Relationship.Contract.ConvertFromMask(&_Relationship.CallOpts, mask)
}

// ConvertFromMask is a free data retrieval call binding the contract method 0x8fa24e0b.
//
// Solidity: function convertFromMask(uint256 mask) pure returns(uint8[] ipAssets, bool allowsExternal)
func (_Relationship *RelationshipCallerSession) ConvertFromMask(mask *big.Int) (struct {
	IpAssets       []uint8
	AllowsExternal bool
}, error) {
	return _Relationship.Contract.ConvertFromMask(&_Relationship.CallOpts, mask)
}

// ConvertToMask is a free data retrieval call binding the contract method 0xdd35c3ce.
//
// Solidity: function convertToMask(uint8[] ipAssets, bool allowsExternal) pure returns(uint256)
func (_Relationship *RelationshipCaller) ConvertToMask(opts *bind.CallOpts, ipAssets []uint8, allowsExternal bool) (*big.Int, error) {
	var out []interface{}
	err := _Relationship.contract.Call(opts, &out, "convertToMask", ipAssets, allowsExternal)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToMask is a free data retrieval call binding the contract method 0xdd35c3ce.
//
// Solidity: function convertToMask(uint8[] ipAssets, bool allowsExternal) pure returns(uint256)
func (_Relationship *RelationshipSession) ConvertToMask(ipAssets []uint8, allowsExternal bool) (*big.Int, error) {
	return _Relationship.Contract.ConvertToMask(&_Relationship.CallOpts, ipAssets, allowsExternal)
}

// ConvertToMask is a free data retrieval call binding the contract method 0xdd35c3ce.
//
// Solidity: function convertToMask(uint8[] ipAssets, bool allowsExternal) pure returns(uint256)
func (_Relationship *RelationshipCallerSession) ConvertToMask(ipAssets []uint8, allowsExternal bool) (*big.Int, error) {
	return _Relationship.Contract.ConvertToMask(&_Relationship.CallOpts, ipAssets, allowsExternal)
}

// GetAccessControl is a free data retrieval call binding the contract method 0xfc1dad81.
//
// Solidity: function getAccessControl() view returns(address)
func (_Relationship *RelationshipCaller) GetAccessControl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Relationship.contract.Call(opts, &out, "getAccessControl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccessControl is a free data retrieval call binding the contract method 0xfc1dad81.
//
// Solidity: function getAccessControl() view returns(address)
func (_Relationship *RelationshipSession) GetAccessControl() (common.Address, error) {
	return _Relationship.Contract.GetAccessControl(&_Relationship.CallOpts)
}

// GetAccessControl is a free data retrieval call binding the contract method 0xfc1dad81.
//
// Solidity: function getAccessControl() view returns(address)
func (_Relationship *RelationshipCallerSession) GetAccessControl() (common.Address, error) {
	return _Relationship.Contract.GetAccessControl(&_Relationship.CallOpts)
}

// GetRelationshipConfig is a free data retrieval call binding the contract method 0xfd0722d8.
//
// Solidity: function getRelationshipConfig(bytes32 relationshipId) view returns((uint256,uint256,bool,address,address,(uint112,uint112,bool)))
func (_Relationship *RelationshipCaller) GetRelationshipConfig(opts *bind.CallOpts, relationshipId [32]byte) (IRelationshipModuleRelationshipConfig, error) {
	var out []interface{}
	err := _Relationship.contract.Call(opts, &out, "getRelationshipConfig", relationshipId)

	if err != nil {
		return *new(IRelationshipModuleRelationshipConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IRelationshipModuleRelationshipConfig)).(*IRelationshipModuleRelationshipConfig)

	return out0, err

}

// GetRelationshipConfig is a free data retrieval call binding the contract method 0xfd0722d8.
//
// Solidity: function getRelationshipConfig(bytes32 relationshipId) view returns((uint256,uint256,bool,address,address,(uint112,uint112,bool)))
func (_Relationship *RelationshipSession) GetRelationshipConfig(relationshipId [32]byte) (IRelationshipModuleRelationshipConfig, error) {
	return _Relationship.Contract.GetRelationshipConfig(&_Relationship.CallOpts, relationshipId)
}

// GetRelationshipConfig is a free data retrieval call binding the contract method 0xfd0722d8.
//
// Solidity: function getRelationshipConfig(bytes32 relationshipId) view returns((uint256,uint256,bool,address,address,(uint112,uint112,bool)))
func (_Relationship *RelationshipCallerSession) GetRelationshipConfig(relationshipId [32]byte) (IRelationshipModuleRelationshipConfig, error) {
	return _Relationship.Contract.GetRelationshipConfig(&_Relationship.CallOpts, relationshipId)
}

// GetRelationshipConfigDecoded is a free data retrieval call binding the contract method 0xc833e736.
//
// Solidity: function getRelationshipConfigDecoded(bytes32 relationshipId) view returns((uint8[],bool,uint8[],bool,bool,address,address,(uint112,uint112,bool)))
func (_Relationship *RelationshipCaller) GetRelationshipConfigDecoded(opts *bind.CallOpts, relationshipId [32]byte) (IRelationshipModuleSetRelationshipConfigParams, error) {
	var out []interface{}
	err := _Relationship.contract.Call(opts, &out, "getRelationshipConfigDecoded", relationshipId)

	if err != nil {
		return *new(IRelationshipModuleSetRelationshipConfigParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IRelationshipModuleSetRelationshipConfigParams)).(*IRelationshipModuleSetRelationshipConfigParams)

	return out0, err

}

// GetRelationshipConfigDecoded is a free data retrieval call binding the contract method 0xc833e736.
//
// Solidity: function getRelationshipConfigDecoded(bytes32 relationshipId) view returns((uint8[],bool,uint8[],bool,bool,address,address,(uint112,uint112,bool)))
func (_Relationship *RelationshipSession) GetRelationshipConfigDecoded(relationshipId [32]byte) (IRelationshipModuleSetRelationshipConfigParams, error) {
	return _Relationship.Contract.GetRelationshipConfigDecoded(&_Relationship.CallOpts, relationshipId)
}

// GetRelationshipConfigDecoded is a free data retrieval call binding the contract method 0xc833e736.
//
// Solidity: function getRelationshipConfigDecoded(bytes32 relationshipId) view returns((uint8[],bool,uint8[],bool,bool,address,address,(uint112,uint112,bool)))
func (_Relationship *RelationshipCallerSession) GetRelationshipConfigDecoded(relationshipId [32]byte) (IRelationshipModuleSetRelationshipConfigParams, error) {
	return _Relationship.Contract.GetRelationshipConfigDecoded(&_Relationship.CallOpts, relationshipId)
}

// GetRelationshipId is a free data retrieval call binding the contract method 0x28654436.
//
// Solidity: function getRelationshipId(string name) pure returns(bytes32)
func (_Relationship *RelationshipCaller) GetRelationshipId(opts *bind.CallOpts, name string) ([32]byte, error) {
	var out []interface{}
	err := _Relationship.contract.Call(opts, &out, "getRelationshipId", name)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRelationshipId is a free data retrieval call binding the contract method 0x28654436.
//
// Solidity: function getRelationshipId(string name) pure returns(bytes32)
func (_Relationship *RelationshipSession) GetRelationshipId(name string) ([32]byte, error) {
	return _Relationship.Contract.GetRelationshipId(&_Relationship.CallOpts, name)
}

// GetRelationshipId is a free data retrieval call binding the contract method 0x28654436.
//
// Solidity: function getRelationshipId(string name) pure returns(bytes32)
func (_Relationship *RelationshipCallerSession) GetRelationshipId(name string) ([32]byte, error) {
	return _Relationship.Contract.GetRelationshipId(&_Relationship.CallOpts, name)
}

// IsRelationshipExpired is a free data retrieval call binding the contract method 0x77032784.
//
// Solidity: function isRelationshipExpired((address,uint256,address,uint256,bytes32,uint256) params) view returns(bool)
func (_Relationship *RelationshipCaller) IsRelationshipExpired(opts *bind.CallOpts, params IRelationshipModuleRelationshipParams) (bool, error) {
	var out []interface{}
	err := _Relationship.contract.Call(opts, &out, "isRelationshipExpired", params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRelationshipExpired is a free data retrieval call binding the contract method 0x77032784.
//
// Solidity: function isRelationshipExpired((address,uint256,address,uint256,bytes32,uint256) params) view returns(bool)
func (_Relationship *RelationshipSession) IsRelationshipExpired(params IRelationshipModuleRelationshipParams) (bool, error) {
	return _Relationship.Contract.IsRelationshipExpired(&_Relationship.CallOpts, params)
}

// IsRelationshipExpired is a free data retrieval call binding the contract method 0x77032784.
//
// Solidity: function isRelationshipExpired((address,uint256,address,uint256,bytes32,uint256) params) view returns(bool)
func (_Relationship *RelationshipCallerSession) IsRelationshipExpired(params IRelationshipModuleRelationshipParams) (bool, error) {
	return _Relationship.Contract.IsRelationshipExpired(&_Relationship.CallOpts, params)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Relationship *RelationshipCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Relationship.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Relationship *RelationshipSession) ProxiableUUID() ([32]byte, error) {
	return _Relationship.Contract.ProxiableUUID(&_Relationship.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Relationship *RelationshipCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Relationship.Contract.ProxiableUUID(&_Relationship.CallOpts)
}

// SupportsIPAssetType is a free data retrieval call binding the contract method 0xe1e1bacb.
//
// Solidity: function supportsIPAssetType(uint256 mask, uint8 assetType) pure returns(bool)
func (_Relationship *RelationshipCaller) SupportsIPAssetType(opts *bind.CallOpts, mask *big.Int, assetType uint8) (bool, error) {
	var out []interface{}
	err := _Relationship.contract.Call(opts, &out, "supportsIPAssetType", mask, assetType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsIPAssetType is a free data retrieval call binding the contract method 0xe1e1bacb.
//
// Solidity: function supportsIPAssetType(uint256 mask, uint8 assetType) pure returns(bool)
func (_Relationship *RelationshipSession) SupportsIPAssetType(mask *big.Int, assetType uint8) (bool, error) {
	return _Relationship.Contract.SupportsIPAssetType(&_Relationship.CallOpts, mask, assetType)
}

// SupportsIPAssetType is a free data retrieval call binding the contract method 0xe1e1bacb.
//
// Solidity: function supportsIPAssetType(uint256 mask, uint8 assetType) pure returns(bool)
func (_Relationship *RelationshipCallerSession) SupportsIPAssetType(mask *big.Int, assetType uint8) (bool, error) {
	return _Relationship.Contract.SupportsIPAssetType(&_Relationship.CallOpts, mask, assetType)
}

// RelationshipModuleBaseInit is a paid mutator transaction binding the contract method 0xb3ef1bbb.
//
// Solidity: function __RelationshipModuleBase_init(address accessControl) returns()
func (_Relationship *RelationshipTransactor) RelationshipModuleBaseInit(opts *bind.TransactOpts, accessControl common.Address) (*types.Transaction, error) {
	return _Relationship.contract.Transact(opts, "__RelationshipModuleBase_init", accessControl)
}

// RelationshipModuleBaseInit is a paid mutator transaction binding the contract method 0xb3ef1bbb.
//
// Solidity: function __RelationshipModuleBase_init(address accessControl) returns()
func (_Relationship *RelationshipSession) RelationshipModuleBaseInit(accessControl common.Address) (*types.Transaction, error) {
	return _Relationship.Contract.RelationshipModuleBaseInit(&_Relationship.TransactOpts, accessControl)
}

// RelationshipModuleBaseInit is a paid mutator transaction binding the contract method 0xb3ef1bbb.
//
// Solidity: function __RelationshipModuleBase_init(address accessControl) returns()
func (_Relationship *RelationshipTransactorSession) RelationshipModuleBaseInit(accessControl common.Address) (*types.Transaction, error) {
	return _Relationship.Contract.RelationshipModuleBaseInit(&_Relationship.TransactOpts, accessControl)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address accessControl) returns()
func (_Relationship *RelationshipTransactor) Initialize(opts *bind.TransactOpts, accessControl common.Address) (*types.Transaction, error) {
	return _Relationship.contract.Transact(opts, "initialize", accessControl)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address accessControl) returns()
func (_Relationship *RelationshipSession) Initialize(accessControl common.Address) (*types.Transaction, error) {
	return _Relationship.Contract.Initialize(&_Relationship.TransactOpts, accessControl)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address accessControl) returns()
func (_Relationship *RelationshipTransactorSession) Initialize(accessControl common.Address) (*types.Transaction, error) {
	return _Relationship.Contract.Initialize(&_Relationship.TransactOpts, accessControl)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Relationship *RelationshipTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Relationship.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Relationship *RelationshipSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Relationship.Contract.Multicall(&_Relationship.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Relationship *RelationshipTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Relationship.Contract.Multicall(&_Relationship.TransactOpts, data)
}

// Relate is a paid mutator transaction binding the contract method 0x82c5ae6e.
//
// Solidity: function relate((address,uint256,address,uint256,bytes32,uint256) params, bytes data) returns()
func (_Relationship *RelationshipTransactor) Relate(opts *bind.TransactOpts, params IRelationshipModuleRelationshipParams, data []byte) (*types.Transaction, error) {
	return _Relationship.contract.Transact(opts, "relate", params, data)
}

// Relate is a paid mutator transaction binding the contract method 0x82c5ae6e.
//
// Solidity: function relate((address,uint256,address,uint256,bytes32,uint256) params, bytes data) returns()
func (_Relationship *RelationshipSession) Relate(params IRelationshipModuleRelationshipParams, data []byte) (*types.Transaction, error) {
	return _Relationship.Contract.Relate(&_Relationship.TransactOpts, params, data)
}

// Relate is a paid mutator transaction binding the contract method 0x82c5ae6e.
//
// Solidity: function relate((address,uint256,address,uint256,bytes32,uint256) params, bytes data) returns()
func (_Relationship *RelationshipTransactorSession) Relate(params IRelationshipModuleRelationshipParams, data []byte) (*types.Transaction, error) {
	return _Relationship.Contract.Relate(&_Relationship.TransactOpts, params, data)
}

// SetAccessControl is a paid mutator transaction binding the contract method 0x19129e5a.
//
// Solidity: function setAccessControl(address accessControl) returns()
func (_Relationship *RelationshipTransactor) SetAccessControl(opts *bind.TransactOpts, accessControl common.Address) (*types.Transaction, error) {
	return _Relationship.contract.Transact(opts, "setAccessControl", accessControl)
}

// SetAccessControl is a paid mutator transaction binding the contract method 0x19129e5a.
//
// Solidity: function setAccessControl(address accessControl) returns()
func (_Relationship *RelationshipSession) SetAccessControl(accessControl common.Address) (*types.Transaction, error) {
	return _Relationship.Contract.SetAccessControl(&_Relationship.TransactOpts, accessControl)
}

// SetAccessControl is a paid mutator transaction binding the contract method 0x19129e5a.
//
// Solidity: function setAccessControl(address accessControl) returns()
func (_Relationship *RelationshipTransactorSession) SetAccessControl(accessControl common.Address) (*types.Transaction, error) {
	return _Relationship.Contract.SetAccessControl(&_Relationship.TransactOpts, accessControl)
}

// SetRelationshipConfig is a paid mutator transaction binding the contract method 0x44dd5e2a.
//
// Solidity: function setRelationshipConfig(string name, (uint8[],bool,uint8[],bool,bool,address,address,(uint112,uint112,bool)) params) returns(bytes32 relationshipId)
func (_Relationship *RelationshipTransactor) SetRelationshipConfig(opts *bind.TransactOpts, name string, params IRelationshipModuleSetRelationshipConfigParams) (*types.Transaction, error) {
	return _Relationship.contract.Transact(opts, "setRelationshipConfig", name, params)
}

// SetRelationshipConfig is a paid mutator transaction binding the contract method 0x44dd5e2a.
//
// Solidity: function setRelationshipConfig(string name, (uint8[],bool,uint8[],bool,bool,address,address,(uint112,uint112,bool)) params) returns(bytes32 relationshipId)
func (_Relationship *RelationshipSession) SetRelationshipConfig(name string, params IRelationshipModuleSetRelationshipConfigParams) (*types.Transaction, error) {
	return _Relationship.Contract.SetRelationshipConfig(&_Relationship.TransactOpts, name, params)
}

// SetRelationshipConfig is a paid mutator transaction binding the contract method 0x44dd5e2a.
//
// Solidity: function setRelationshipConfig(string name, (uint8[],bool,uint8[],bool,bool,address,address,(uint112,uint112,bool)) params) returns(bytes32 relationshipId)
func (_Relationship *RelationshipTransactorSession) SetRelationshipConfig(name string, params IRelationshipModuleSetRelationshipConfigParams) (*types.Transaction, error) {
	return _Relationship.Contract.SetRelationshipConfig(&_Relationship.TransactOpts, name, params)
}

// Unrelate is a paid mutator transaction binding the contract method 0x75f4252d.
//
// Solidity: function unrelate((address,uint256,address,uint256,bytes32,uint256) params) returns()
func (_Relationship *RelationshipTransactor) Unrelate(opts *bind.TransactOpts, params IRelationshipModuleRelationshipParams) (*types.Transaction, error) {
	return _Relationship.contract.Transact(opts, "unrelate", params)
}

// Unrelate is a paid mutator transaction binding the contract method 0x75f4252d.
//
// Solidity: function unrelate((address,uint256,address,uint256,bytes32,uint256) params) returns()
func (_Relationship *RelationshipSession) Unrelate(params IRelationshipModuleRelationshipParams) (*types.Transaction, error) {
	return _Relationship.Contract.Unrelate(&_Relationship.TransactOpts, params)
}

// Unrelate is a paid mutator transaction binding the contract method 0x75f4252d.
//
// Solidity: function unrelate((address,uint256,address,uint256,bytes32,uint256) params) returns()
func (_Relationship *RelationshipTransactorSession) Unrelate(params IRelationshipModuleRelationshipParams) (*types.Transaction, error) {
	return _Relationship.Contract.Unrelate(&_Relationship.TransactOpts, params)
}

// UnsetRelationshipConfig is a paid mutator transaction binding the contract method 0x35f52e92.
//
// Solidity: function unsetRelationshipConfig(bytes32 relationshipId) returns()
func (_Relationship *RelationshipTransactor) UnsetRelationshipConfig(opts *bind.TransactOpts, relationshipId [32]byte) (*types.Transaction, error) {
	return _Relationship.contract.Transact(opts, "unsetRelationshipConfig", relationshipId)
}

// UnsetRelationshipConfig is a paid mutator transaction binding the contract method 0x35f52e92.
//
// Solidity: function unsetRelationshipConfig(bytes32 relationshipId) returns()
func (_Relationship *RelationshipSession) UnsetRelationshipConfig(relationshipId [32]byte) (*types.Transaction, error) {
	return _Relationship.Contract.UnsetRelationshipConfig(&_Relationship.TransactOpts, relationshipId)
}

// UnsetRelationshipConfig is a paid mutator transaction binding the contract method 0x35f52e92.
//
// Solidity: function unsetRelationshipConfig(bytes32 relationshipId) returns()
func (_Relationship *RelationshipTransactorSession) UnsetRelationshipConfig(relationshipId [32]byte) (*types.Transaction, error) {
	return _Relationship.Contract.UnsetRelationshipConfig(&_Relationship.TransactOpts, relationshipId)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Relationship *RelationshipTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Relationship.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Relationship *RelationshipSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Relationship.Contract.UpgradeTo(&_Relationship.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Relationship *RelationshipTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Relationship.Contract.UpgradeTo(&_Relationship.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Relationship *RelationshipTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Relationship.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Relationship *RelationshipSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Relationship.Contract.UpgradeToAndCall(&_Relationship.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Relationship *RelationshipTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Relationship.Contract.UpgradeToAndCall(&_Relationship.TransactOpts, newImplementation, data)
}

// RelationshipAccessControlUpdatedIterator is returned from FilterAccessControlUpdated and is used to iterate over the raw logs and unpacked data for AccessControlUpdated events raised by the Relationship contract.
type RelationshipAccessControlUpdatedIterator struct {
	Event *RelationshipAccessControlUpdated // Event containing the contract specifics and raw log

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
func (it *RelationshipAccessControlUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelationshipAccessControlUpdated)
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
		it.Event = new(RelationshipAccessControlUpdated)
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
func (it *RelationshipAccessControlUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelationshipAccessControlUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelationshipAccessControlUpdated represents a AccessControlUpdated event raised by the Relationship contract.
type RelationshipAccessControlUpdated struct {
	AccessControl common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAccessControlUpdated is a free log retrieval operation binding the contract event 0xc8ebe3bc6cc2f03e640cefc0f0c872637b7d9492bd5e6318eaba1ab468da9709.
//
// Solidity: event AccessControlUpdated(address indexed accessControl)
func (_Relationship *RelationshipFilterer) FilterAccessControlUpdated(opts *bind.FilterOpts, accessControl []common.Address) (*RelationshipAccessControlUpdatedIterator, error) {

	var accessControlRule []interface{}
	for _, accessControlItem := range accessControl {
		accessControlRule = append(accessControlRule, accessControlItem)
	}

	logs, sub, err := _Relationship.contract.FilterLogs(opts, "AccessControlUpdated", accessControlRule)
	if err != nil {
		return nil, err
	}
	return &RelationshipAccessControlUpdatedIterator{contract: _Relationship.contract, event: "AccessControlUpdated", logs: logs, sub: sub}, nil
}

// WatchAccessControlUpdated is a free log subscription operation binding the contract event 0xc8ebe3bc6cc2f03e640cefc0f0c872637b7d9492bd5e6318eaba1ab468da9709.
//
// Solidity: event AccessControlUpdated(address indexed accessControl)
func (_Relationship *RelationshipFilterer) WatchAccessControlUpdated(opts *bind.WatchOpts, sink chan<- *RelationshipAccessControlUpdated, accessControl []common.Address) (event.Subscription, error) {

	var accessControlRule []interface{}
	for _, accessControlItem := range accessControl {
		accessControlRule = append(accessControlRule, accessControlItem)
	}

	logs, sub, err := _Relationship.contract.WatchLogs(opts, "AccessControlUpdated", accessControlRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelationshipAccessControlUpdated)
				if err := _Relationship.contract.UnpackLog(event, "AccessControlUpdated", log); err != nil {
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

// ParseAccessControlUpdated is a log parse operation binding the contract event 0xc8ebe3bc6cc2f03e640cefc0f0c872637b7d9492bd5e6318eaba1ab468da9709.
//
// Solidity: event AccessControlUpdated(address indexed accessControl)
func (_Relationship *RelationshipFilterer) ParseAccessControlUpdated(log types.Log) (*RelationshipAccessControlUpdated, error) {
	event := new(RelationshipAccessControlUpdated)
	if err := _Relationship.contract.UnpackLog(event, "AccessControlUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelationshipAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Relationship contract.
type RelationshipAdminChangedIterator struct {
	Event *RelationshipAdminChanged // Event containing the contract specifics and raw log

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
func (it *RelationshipAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelationshipAdminChanged)
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
		it.Event = new(RelationshipAdminChanged)
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
func (it *RelationshipAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelationshipAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelationshipAdminChanged represents a AdminChanged event raised by the Relationship contract.
type RelationshipAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Relationship *RelationshipFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*RelationshipAdminChangedIterator, error) {

	logs, sub, err := _Relationship.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &RelationshipAdminChangedIterator{contract: _Relationship.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Relationship *RelationshipFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *RelationshipAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Relationship.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelationshipAdminChanged)
				if err := _Relationship.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Relationship *RelationshipFilterer) ParseAdminChanged(log types.Log) (*RelationshipAdminChanged, error) {
	event := new(RelationshipAdminChanged)
	if err := _Relationship.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelationshipBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Relationship contract.
type RelationshipBeaconUpgradedIterator struct {
	Event *RelationshipBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *RelationshipBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelationshipBeaconUpgraded)
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
		it.Event = new(RelationshipBeaconUpgraded)
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
func (it *RelationshipBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelationshipBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelationshipBeaconUpgraded represents a BeaconUpgraded event raised by the Relationship contract.
type RelationshipBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Relationship *RelationshipFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*RelationshipBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Relationship.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &RelationshipBeaconUpgradedIterator{contract: _Relationship.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Relationship *RelationshipFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *RelationshipBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Relationship.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelationshipBeaconUpgraded)
				if err := _Relationship.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Relationship *RelationshipFilterer) ParseBeaconUpgraded(log types.Log) (*RelationshipBeaconUpgraded, error) {
	event := new(RelationshipBeaconUpgraded)
	if err := _Relationship.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelationshipInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Relationship contract.
type RelationshipInitializedIterator struct {
	Event *RelationshipInitialized // Event containing the contract specifics and raw log

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
func (it *RelationshipInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelationshipInitialized)
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
		it.Event = new(RelationshipInitialized)
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
func (it *RelationshipInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelationshipInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelationshipInitialized represents a Initialized event raised by the Relationship contract.
type RelationshipInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Relationship *RelationshipFilterer) FilterInitialized(opts *bind.FilterOpts) (*RelationshipInitializedIterator, error) {

	logs, sub, err := _Relationship.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RelationshipInitializedIterator{contract: _Relationship.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Relationship *RelationshipFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RelationshipInitialized) (event.Subscription, error) {

	logs, sub, err := _Relationship.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelationshipInitialized)
				if err := _Relationship.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Relationship *RelationshipFilterer) ParseInitialized(log types.Log) (*RelationshipInitialized, error) {
	event := new(RelationshipInitialized)
	if err := _Relationship.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelationshipRelationPendingProcessorIterator is returned from FilterRelationPendingProcessor and is used to iterate over the raw logs and unpacked data for RelationPendingProcessor events raised by the Relationship contract.
type RelationshipRelationPendingProcessorIterator struct {
	Event *RelationshipRelationPendingProcessor // Event containing the contract specifics and raw log

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
func (it *RelationshipRelationPendingProcessorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelationshipRelationPendingProcessor)
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
		it.Event = new(RelationshipRelationPendingProcessor)
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
func (it *RelationshipRelationPendingProcessorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelationshipRelationPendingProcessorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelationshipRelationPendingProcessor represents a RelationPendingProcessor event raised by the Relationship contract.
type RelationshipRelationPendingProcessor struct {
	SourceContract common.Address
	SourceId       *big.Int
	DestContract   common.Address
	DestId         *big.Int
	RelationshipId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRelationPendingProcessor is a free log retrieval operation binding the contract event 0xff0396d9569545b52312943ba44341eba3800a0d6e83ba9140283c159792839e.
//
// Solidity: event RelationPendingProcessor(address sourceContract, uint256 sourceId, address destContract, uint256 destId, bytes32 indexed relationshipId)
func (_Relationship *RelationshipFilterer) FilterRelationPendingProcessor(opts *bind.FilterOpts, relationshipId [][32]byte) (*RelationshipRelationPendingProcessorIterator, error) {

	var relationshipIdRule []interface{}
	for _, relationshipIdItem := range relationshipId {
		relationshipIdRule = append(relationshipIdRule, relationshipIdItem)
	}

	logs, sub, err := _Relationship.contract.FilterLogs(opts, "RelationPendingProcessor", relationshipIdRule)
	if err != nil {
		return nil, err
	}
	return &RelationshipRelationPendingProcessorIterator{contract: _Relationship.contract, event: "RelationPendingProcessor", logs: logs, sub: sub}, nil
}

// WatchRelationPendingProcessor is a free log subscription operation binding the contract event 0xff0396d9569545b52312943ba44341eba3800a0d6e83ba9140283c159792839e.
//
// Solidity: event RelationPendingProcessor(address sourceContract, uint256 sourceId, address destContract, uint256 destId, bytes32 indexed relationshipId)
func (_Relationship *RelationshipFilterer) WatchRelationPendingProcessor(opts *bind.WatchOpts, sink chan<- *RelationshipRelationPendingProcessor, relationshipId [][32]byte) (event.Subscription, error) {

	var relationshipIdRule []interface{}
	for _, relationshipIdItem := range relationshipId {
		relationshipIdRule = append(relationshipIdRule, relationshipIdItem)
	}

	logs, sub, err := _Relationship.contract.WatchLogs(opts, "RelationPendingProcessor", relationshipIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelationshipRelationPendingProcessor)
				if err := _Relationship.contract.UnpackLog(event, "RelationPendingProcessor", log); err != nil {
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

// ParseRelationPendingProcessor is a log parse operation binding the contract event 0xff0396d9569545b52312943ba44341eba3800a0d6e83ba9140283c159792839e.
//
// Solidity: event RelationPendingProcessor(address sourceContract, uint256 sourceId, address destContract, uint256 destId, bytes32 indexed relationshipId)
func (_Relationship *RelationshipFilterer) ParseRelationPendingProcessor(log types.Log) (*RelationshipRelationPendingProcessor, error) {
	event := new(RelationshipRelationPendingProcessor)
	if err := _Relationship.contract.UnpackLog(event, "RelationPendingProcessor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelationshipRelationSetIterator is returned from FilterRelationSet and is used to iterate over the raw logs and unpacked data for RelationSet events raised by the Relationship contract.
type RelationshipRelationSetIterator struct {
	Event *RelationshipRelationSet // Event containing the contract specifics and raw log

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
func (it *RelationshipRelationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelationshipRelationSet)
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
		it.Event = new(RelationshipRelationSet)
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
func (it *RelationshipRelationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelationshipRelationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelationshipRelationSet represents a RelationSet event raised by the Relationship contract.
type RelationshipRelationSet struct {
	SourceContract common.Address
	SourceId       *big.Int
	DestContract   common.Address
	DestId         *big.Int
	RelationshipId [32]byte
	EndTime        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRelationSet is a free log retrieval operation binding the contract event 0xdac80e4156e67d10c07ce819561c6cd96452ad81db0c68e6a47a8687f3d59271.
//
// Solidity: event RelationSet(address sourceContract, uint256 sourceId, address destContract, uint256 destId, bytes32 indexed relationshipId, uint256 endTime)
func (_Relationship *RelationshipFilterer) FilterRelationSet(opts *bind.FilterOpts, relationshipId [][32]byte) (*RelationshipRelationSetIterator, error) {

	var relationshipIdRule []interface{}
	for _, relationshipIdItem := range relationshipId {
		relationshipIdRule = append(relationshipIdRule, relationshipIdItem)
	}

	logs, sub, err := _Relationship.contract.FilterLogs(opts, "RelationSet", relationshipIdRule)
	if err != nil {
		return nil, err
	}
	return &RelationshipRelationSetIterator{contract: _Relationship.contract, event: "RelationSet", logs: logs, sub: sub}, nil
}

// WatchRelationSet is a free log subscription operation binding the contract event 0xdac80e4156e67d10c07ce819561c6cd96452ad81db0c68e6a47a8687f3d59271.
//
// Solidity: event RelationSet(address sourceContract, uint256 sourceId, address destContract, uint256 destId, bytes32 indexed relationshipId, uint256 endTime)
func (_Relationship *RelationshipFilterer) WatchRelationSet(opts *bind.WatchOpts, sink chan<- *RelationshipRelationSet, relationshipId [][32]byte) (event.Subscription, error) {

	var relationshipIdRule []interface{}
	for _, relationshipIdItem := range relationshipId {
		relationshipIdRule = append(relationshipIdRule, relationshipIdItem)
	}

	logs, sub, err := _Relationship.contract.WatchLogs(opts, "RelationSet", relationshipIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelationshipRelationSet)
				if err := _Relationship.contract.UnpackLog(event, "RelationSet", log); err != nil {
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

// ParseRelationSet is a log parse operation binding the contract event 0xdac80e4156e67d10c07ce819561c6cd96452ad81db0c68e6a47a8687f3d59271.
//
// Solidity: event RelationSet(address sourceContract, uint256 sourceId, address destContract, uint256 destId, bytes32 indexed relationshipId, uint256 endTime)
func (_Relationship *RelationshipFilterer) ParseRelationSet(log types.Log) (*RelationshipRelationSet, error) {
	event := new(RelationshipRelationSet)
	if err := _Relationship.contract.UnpackLog(event, "RelationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelationshipRelationUnsetIterator is returned from FilterRelationUnset and is used to iterate over the raw logs and unpacked data for RelationUnset events raised by the Relationship contract.
type RelationshipRelationUnsetIterator struct {
	Event *RelationshipRelationUnset // Event containing the contract specifics and raw log

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
func (it *RelationshipRelationUnsetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelationshipRelationUnset)
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
		it.Event = new(RelationshipRelationUnset)
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
func (it *RelationshipRelationUnsetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelationshipRelationUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelationshipRelationUnset represents a RelationUnset event raised by the Relationship contract.
type RelationshipRelationUnset struct {
	SourceContract common.Address
	SourceId       *big.Int
	DestContract   common.Address
	DestId         *big.Int
	RelationshipId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRelationUnset is a free log retrieval operation binding the contract event 0x1c597e2ff497f95acc844355a7e0266613bbc54a551e5fa01710e5ad3418e04f.
//
// Solidity: event RelationUnset(address sourceContract, uint256 sourceId, address destContract, uint256 destId, bytes32 indexed relationshipId)
func (_Relationship *RelationshipFilterer) FilterRelationUnset(opts *bind.FilterOpts, relationshipId [][32]byte) (*RelationshipRelationUnsetIterator, error) {

	var relationshipIdRule []interface{}
	for _, relationshipIdItem := range relationshipId {
		relationshipIdRule = append(relationshipIdRule, relationshipIdItem)
	}

	logs, sub, err := _Relationship.contract.FilterLogs(opts, "RelationUnset", relationshipIdRule)
	if err != nil {
		return nil, err
	}
	return &RelationshipRelationUnsetIterator{contract: _Relationship.contract, event: "RelationUnset", logs: logs, sub: sub}, nil
}

// WatchRelationUnset is a free log subscription operation binding the contract event 0x1c597e2ff497f95acc844355a7e0266613bbc54a551e5fa01710e5ad3418e04f.
//
// Solidity: event RelationUnset(address sourceContract, uint256 sourceId, address destContract, uint256 destId, bytes32 indexed relationshipId)
func (_Relationship *RelationshipFilterer) WatchRelationUnset(opts *bind.WatchOpts, sink chan<- *RelationshipRelationUnset, relationshipId [][32]byte) (event.Subscription, error) {

	var relationshipIdRule []interface{}
	for _, relationshipIdItem := range relationshipId {
		relationshipIdRule = append(relationshipIdRule, relationshipIdItem)
	}

	logs, sub, err := _Relationship.contract.WatchLogs(opts, "RelationUnset", relationshipIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelationshipRelationUnset)
				if err := _Relationship.contract.UnpackLog(event, "RelationUnset", log); err != nil {
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

// ParseRelationUnset is a log parse operation binding the contract event 0x1c597e2ff497f95acc844355a7e0266613bbc54a551e5fa01710e5ad3418e04f.
//
// Solidity: event RelationUnset(address sourceContract, uint256 sourceId, address destContract, uint256 destId, bytes32 indexed relationshipId)
func (_Relationship *RelationshipFilterer) ParseRelationUnset(log types.Log) (*RelationshipRelationUnset, error) {
	event := new(RelationshipRelationUnset)
	if err := _Relationship.contract.UnpackLog(event, "RelationUnset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelationshipRelationshipConfigSetIterator is returned from FilterRelationshipConfigSet and is used to iterate over the raw logs and unpacked data for RelationshipConfigSet events raised by the Relationship contract.
type RelationshipRelationshipConfigSetIterator struct {
	Event *RelationshipRelationshipConfigSet // Event containing the contract specifics and raw log

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
func (it *RelationshipRelationshipConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelationshipRelationshipConfigSet)
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
		it.Event = new(RelationshipRelationshipConfigSet)
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
func (it *RelationshipRelationshipConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelationshipRelationshipConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelationshipRelationshipConfigSet represents a RelationshipConfigSet event raised by the Relationship contract.
type RelationshipRelationshipConfigSet struct {
	Name                  string
	RelationshipId        [32]byte
	SourceIPAssetTypeMask *big.Int
	DestIPAssetTypeMask   *big.Int
	OnlySameFranchise     bool
	Processor             common.Address
	MaxTTL                *big.Int
	MinTTL                *big.Int
	Renewable             bool
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRelationshipConfigSet is a free log retrieval operation binding the contract event 0x23d823dca554930dda2141543b0575ecace98e873be3b3fa7dbc7b7c2e8dba6d.
//
// Solidity: event RelationshipConfigSet(string name, bytes32 indexed relationshipId, uint256 sourceIPAssetTypeMask, uint256 destIPAssetTypeMask, bool onlySameFranchise, address processor, uint256 maxTTL, uint256 minTTL, bool renewable)
func (_Relationship *RelationshipFilterer) FilterRelationshipConfigSet(opts *bind.FilterOpts, relationshipId [][32]byte) (*RelationshipRelationshipConfigSetIterator, error) {

	var relationshipIdRule []interface{}
	for _, relationshipIdItem := range relationshipId {
		relationshipIdRule = append(relationshipIdRule, relationshipIdItem)
	}

	logs, sub, err := _Relationship.contract.FilterLogs(opts, "RelationshipConfigSet", relationshipIdRule)
	if err != nil {
		return nil, err
	}
	return &RelationshipRelationshipConfigSetIterator{contract: _Relationship.contract, event: "RelationshipConfigSet", logs: logs, sub: sub}, nil
}

// WatchRelationshipConfigSet is a free log subscription operation binding the contract event 0x23d823dca554930dda2141543b0575ecace98e873be3b3fa7dbc7b7c2e8dba6d.
//
// Solidity: event RelationshipConfigSet(string name, bytes32 indexed relationshipId, uint256 sourceIPAssetTypeMask, uint256 destIPAssetTypeMask, bool onlySameFranchise, address processor, uint256 maxTTL, uint256 minTTL, bool renewable)
func (_Relationship *RelationshipFilterer) WatchRelationshipConfigSet(opts *bind.WatchOpts, sink chan<- *RelationshipRelationshipConfigSet, relationshipId [][32]byte) (event.Subscription, error) {

	var relationshipIdRule []interface{}
	for _, relationshipIdItem := range relationshipId {
		relationshipIdRule = append(relationshipIdRule, relationshipIdItem)
	}

	logs, sub, err := _Relationship.contract.WatchLogs(opts, "RelationshipConfigSet", relationshipIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelationshipRelationshipConfigSet)
				if err := _Relationship.contract.UnpackLog(event, "RelationshipConfigSet", log); err != nil {
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

// ParseRelationshipConfigSet is a log parse operation binding the contract event 0x23d823dca554930dda2141543b0575ecace98e873be3b3fa7dbc7b7c2e8dba6d.
//
// Solidity: event RelationshipConfigSet(string name, bytes32 indexed relationshipId, uint256 sourceIPAssetTypeMask, uint256 destIPAssetTypeMask, bool onlySameFranchise, address processor, uint256 maxTTL, uint256 minTTL, bool renewable)
func (_Relationship *RelationshipFilterer) ParseRelationshipConfigSet(log types.Log) (*RelationshipRelationshipConfigSet, error) {
	event := new(RelationshipRelationshipConfigSet)
	if err := _Relationship.contract.UnpackLog(event, "RelationshipConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelationshipRelationshipConfigUnsetIterator is returned from FilterRelationshipConfigUnset and is used to iterate over the raw logs and unpacked data for RelationshipConfigUnset events raised by the Relationship contract.
type RelationshipRelationshipConfigUnsetIterator struct {
	Event *RelationshipRelationshipConfigUnset // Event containing the contract specifics and raw log

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
func (it *RelationshipRelationshipConfigUnsetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelationshipRelationshipConfigUnset)
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
		it.Event = new(RelationshipRelationshipConfigUnset)
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
func (it *RelationshipRelationshipConfigUnsetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelationshipRelationshipConfigUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelationshipRelationshipConfigUnset represents a RelationshipConfigUnset event raised by the Relationship contract.
type RelationshipRelationshipConfigUnset struct {
	RelationshipId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRelationshipConfigUnset is a free log retrieval operation binding the contract event 0x72ee2f0fb0c2e5516159220d68bb62f4ee27e2eeefb237e60bf0b0e108ce3490.
//
// Solidity: event RelationshipConfigUnset(bytes32 indexed relationshipId)
func (_Relationship *RelationshipFilterer) FilterRelationshipConfigUnset(opts *bind.FilterOpts, relationshipId [][32]byte) (*RelationshipRelationshipConfigUnsetIterator, error) {

	var relationshipIdRule []interface{}
	for _, relationshipIdItem := range relationshipId {
		relationshipIdRule = append(relationshipIdRule, relationshipIdItem)
	}

	logs, sub, err := _Relationship.contract.FilterLogs(opts, "RelationshipConfigUnset", relationshipIdRule)
	if err != nil {
		return nil, err
	}
	return &RelationshipRelationshipConfigUnsetIterator{contract: _Relationship.contract, event: "RelationshipConfigUnset", logs: logs, sub: sub}, nil
}

// WatchRelationshipConfigUnset is a free log subscription operation binding the contract event 0x72ee2f0fb0c2e5516159220d68bb62f4ee27e2eeefb237e60bf0b0e108ce3490.
//
// Solidity: event RelationshipConfigUnset(bytes32 indexed relationshipId)
func (_Relationship *RelationshipFilterer) WatchRelationshipConfigUnset(opts *bind.WatchOpts, sink chan<- *RelationshipRelationshipConfigUnset, relationshipId [][32]byte) (event.Subscription, error) {

	var relationshipIdRule []interface{}
	for _, relationshipIdItem := range relationshipId {
		relationshipIdRule = append(relationshipIdRule, relationshipIdItem)
	}

	logs, sub, err := _Relationship.contract.WatchLogs(opts, "RelationshipConfigUnset", relationshipIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelationshipRelationshipConfigUnset)
				if err := _Relationship.contract.UnpackLog(event, "RelationshipConfigUnset", log); err != nil {
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

// ParseRelationshipConfigUnset is a log parse operation binding the contract event 0x72ee2f0fb0c2e5516159220d68bb62f4ee27e2eeefb237e60bf0b0e108ce3490.
//
// Solidity: event RelationshipConfigUnset(bytes32 indexed relationshipId)
func (_Relationship *RelationshipFilterer) ParseRelationshipConfigUnset(log types.Log) (*RelationshipRelationshipConfigUnset, error) {
	event := new(RelationshipRelationshipConfigUnset)
	if err := _Relationship.contract.UnpackLog(event, "RelationshipConfigUnset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelationshipUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Relationship contract.
type RelationshipUpgradedIterator struct {
	Event *RelationshipUpgraded // Event containing the contract specifics and raw log

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
func (it *RelationshipUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelationshipUpgraded)
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
		it.Event = new(RelationshipUpgraded)
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
func (it *RelationshipUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelationshipUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelationshipUpgraded represents a Upgraded event raised by the Relationship contract.
type RelationshipUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Relationship *RelationshipFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*RelationshipUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Relationship.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &RelationshipUpgradedIterator{contract: _Relationship.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Relationship *RelationshipFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *RelationshipUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Relationship.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelationshipUpgraded)
				if err := _Relationship.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Relationship *RelationshipFilterer) ParseUpgraded(log types.Log) (*RelationshipUpgraded, error) {
	event := new(RelationshipUpgraded)
	if err := _Relationship.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
