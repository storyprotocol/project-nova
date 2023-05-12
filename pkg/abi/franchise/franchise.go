// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package franchise

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

// IFranchiseCharacterCollection is an auto generated low-level Go binding around an user-defined struct.
type IFranchiseCharacterCollection struct {
	Collection common.Address
	Name       string
	Symbol     string
}

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

// FranchiseMetaData contains all meta data concerning the Franchise contract.
var FranchiseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_characterRegistryFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CharacterAlreadyCreated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidName\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVaultAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryNotCreated\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"canonCollection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fanficCollection\",\"type\":\"address\"}],\"name\":\"CharactersCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"canonCollection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fanficCollection\",\"type\":\"address\"}],\"name\":\"StoriesCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_CANON_COLLECTION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_FANFIC_COLLECTION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structICharacterRegistry.Author\",\"name\":\"author\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"licenseModule\",\"type\":\"address\"},{\"internalType\":\"enumICharacterRegistry.CharacterType\",\"name\":\"charType\",\"type\":\"uint8\"}],\"internalType\":\"structICharacterRegistry.CharacterInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"externalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"}],\"name\":\"addCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"author\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structNftToken[]\",\"name\":\"characters\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"autographModule\",\"type\":\"address\"}],\"internalType\":\"structStoryInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"addFanFiction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"author\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structNftToken[]\",\"name\":\"characters\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"autographModule\",\"type\":\"address\"}],\"internalType\":\"structStoryInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"addStory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"characterRegistry\",\"outputs\":[{\"internalType\":\"contractCharacterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"collectCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"collectStory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getCharacterAutograph\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getCharacterCollectors\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getStoryAutograph\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getStoryCollectors\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structIFranchise.CharacterCollection\",\"name\":\"canon\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structIFranchise.CharacterCollection\",\"name\":\"fanfic\",\"type\":\"tuple\"}],\"name\":\"initCharacters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"licenseModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"licenseNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"promoteFanFiction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"licenseData\",\"type\":\"bytes\"}],\"name\":\"requestLicense\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"setLicenseModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStoryRegistry\",\"name\":\"storyRegistry_\",\"type\":\"address\"}],\"name\":\"setStoryRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storyRegistry\",\"outputs\":[{\"internalType\":\"contractIStoryRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FranchiseABI is the input ABI used to generate the binding from.
// Deprecated: Use FranchiseMetaData.ABI instead.
var FranchiseABI = FranchiseMetaData.ABI

// Franchise is an auto generated Go binding around an Ethereum contract.
type Franchise struct {
	FranchiseCaller     // Read-only binding to the contract
	FranchiseTransactor // Write-only binding to the contract
	FranchiseFilterer   // Log filterer for contract events
}

// FranchiseCaller is an auto generated read-only Go binding around an Ethereum contract.
type FranchiseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FranchiseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FranchiseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FranchiseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FranchiseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FranchiseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FranchiseSession struct {
	Contract     *Franchise        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FranchiseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FranchiseCallerSession struct {
	Contract *FranchiseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FranchiseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FranchiseTransactorSession struct {
	Contract     *FranchiseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FranchiseRaw is an auto generated low-level Go binding around an Ethereum contract.
type FranchiseRaw struct {
	Contract *Franchise // Generic contract binding to access the raw methods on
}

// FranchiseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FranchiseCallerRaw struct {
	Contract *FranchiseCaller // Generic read-only contract binding to access the raw methods on
}

// FranchiseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FranchiseTransactorRaw struct {
	Contract *FranchiseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFranchise creates a new instance of Franchise, bound to a specific deployed contract.
func NewFranchise(address common.Address, backend bind.ContractBackend) (*Franchise, error) {
	contract, err := bindFranchise(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Franchise{FranchiseCaller: FranchiseCaller{contract: contract}, FranchiseTransactor: FranchiseTransactor{contract: contract}, FranchiseFilterer: FranchiseFilterer{contract: contract}}, nil
}

// NewFranchiseCaller creates a new read-only instance of Franchise, bound to a specific deployed contract.
func NewFranchiseCaller(address common.Address, caller bind.ContractCaller) (*FranchiseCaller, error) {
	contract, err := bindFranchise(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FranchiseCaller{contract: contract}, nil
}

// NewFranchiseTransactor creates a new write-only instance of Franchise, bound to a specific deployed contract.
func NewFranchiseTransactor(address common.Address, transactor bind.ContractTransactor) (*FranchiseTransactor, error) {
	contract, err := bindFranchise(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FranchiseTransactor{contract: contract}, nil
}

// NewFranchiseFilterer creates a new log filterer instance of Franchise, bound to a specific deployed contract.
func NewFranchiseFilterer(address common.Address, filterer bind.ContractFilterer) (*FranchiseFilterer, error) {
	contract, err := bindFranchise(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FranchiseFilterer{contract: contract}, nil
}

// bindFranchise binds a generic wrapper to an already deployed contract.
func bindFranchise(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FranchiseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Franchise *FranchiseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Franchise.Contract.FranchiseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Franchise *FranchiseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Franchise.Contract.FranchiseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Franchise *FranchiseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Franchise.Contract.FranchiseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Franchise *FranchiseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Franchise.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Franchise *FranchiseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Franchise.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Franchise *FranchiseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Franchise.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTCANONCOLLECTION is a free data retrieval call binding the contract method 0x37042dde.
//
// Solidity: function DEFAULT_CANON_COLLECTION() view returns(address)
func (_Franchise *FranchiseCaller) DEFAULTCANONCOLLECTION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Franchise.contract.Call(opts, &out, "DEFAULT_CANON_COLLECTION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEFAULTCANONCOLLECTION is a free data retrieval call binding the contract method 0x37042dde.
//
// Solidity: function DEFAULT_CANON_COLLECTION() view returns(address)
func (_Franchise *FranchiseSession) DEFAULTCANONCOLLECTION() (common.Address, error) {
	return _Franchise.Contract.DEFAULTCANONCOLLECTION(&_Franchise.CallOpts)
}

// DEFAULTCANONCOLLECTION is a free data retrieval call binding the contract method 0x37042dde.
//
// Solidity: function DEFAULT_CANON_COLLECTION() view returns(address)
func (_Franchise *FranchiseCallerSession) DEFAULTCANONCOLLECTION() (common.Address, error) {
	return _Franchise.Contract.DEFAULTCANONCOLLECTION(&_Franchise.CallOpts)
}

// DEFAULTFANFICCOLLECTION is a free data retrieval call binding the contract method 0x84cbbab5.
//
// Solidity: function DEFAULT_FANFIC_COLLECTION() view returns(address)
func (_Franchise *FranchiseCaller) DEFAULTFANFICCOLLECTION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Franchise.contract.Call(opts, &out, "DEFAULT_FANFIC_COLLECTION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEFAULTFANFICCOLLECTION is a free data retrieval call binding the contract method 0x84cbbab5.
//
// Solidity: function DEFAULT_FANFIC_COLLECTION() view returns(address)
func (_Franchise *FranchiseSession) DEFAULTFANFICCOLLECTION() (common.Address, error) {
	return _Franchise.Contract.DEFAULTFANFICCOLLECTION(&_Franchise.CallOpts)
}

// DEFAULTFANFICCOLLECTION is a free data retrieval call binding the contract method 0x84cbbab5.
//
// Solidity: function DEFAULT_FANFIC_COLLECTION() view returns(address)
func (_Franchise *FranchiseCallerSession) DEFAULTFANFICCOLLECTION() (common.Address, error) {
	return _Franchise.Contract.DEFAULTFANFICCOLLECTION(&_Franchise.CallOpts)
}

// CharacterRegistry is a free data retrieval call binding the contract method 0x82a7e977.
//
// Solidity: function characterRegistry() view returns(address)
func (_Franchise *FranchiseCaller) CharacterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Franchise.contract.Call(opts, &out, "characterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CharacterRegistry is a free data retrieval call binding the contract method 0x82a7e977.
//
// Solidity: function characterRegistry() view returns(address)
func (_Franchise *FranchiseSession) CharacterRegistry() (common.Address, error) {
	return _Franchise.Contract.CharacterRegistry(&_Franchise.CallOpts)
}

// CharacterRegistry is a free data retrieval call binding the contract method 0x82a7e977.
//
// Solidity: function characterRegistry() view returns(address)
func (_Franchise *FranchiseCallerSession) CharacterRegistry() (common.Address, error) {
	return _Franchise.Contract.CharacterRegistry(&_Franchise.CallOpts)
}

// GetCharacterAutograph is a free data retrieval call binding the contract method 0x4595e554.
//
// Solidity: function getCharacterAutograph(address collection, uint256 id) view returns(address, uint256)
func (_Franchise *FranchiseCaller) GetCharacterAutograph(opts *bind.CallOpts, collection common.Address, id *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Franchise.contract.Call(opts, &out, "getCharacterAutograph", collection, id)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetCharacterAutograph is a free data retrieval call binding the contract method 0x4595e554.
//
// Solidity: function getCharacterAutograph(address collection, uint256 id) view returns(address, uint256)
func (_Franchise *FranchiseSession) GetCharacterAutograph(collection common.Address, id *big.Int) (common.Address, *big.Int, error) {
	return _Franchise.Contract.GetCharacterAutograph(&_Franchise.CallOpts, collection, id)
}

// GetCharacterAutograph is a free data retrieval call binding the contract method 0x4595e554.
//
// Solidity: function getCharacterAutograph(address collection, uint256 id) view returns(address, uint256)
func (_Franchise *FranchiseCallerSession) GetCharacterAutograph(collection common.Address, id *big.Int) (common.Address, *big.Int, error) {
	return _Franchise.Contract.GetCharacterAutograph(&_Franchise.CallOpts, collection, id)
}

// GetCharacterCollectors is a free data retrieval call binding the contract method 0x52430b6a.
//
// Solidity: function getCharacterCollectors(address collection, uint256 id) view returns(address[])
func (_Franchise *FranchiseCaller) GetCharacterCollectors(opts *bind.CallOpts, collection common.Address, id *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Franchise.contract.Call(opts, &out, "getCharacterCollectors", collection, id)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCharacterCollectors is a free data retrieval call binding the contract method 0x52430b6a.
//
// Solidity: function getCharacterCollectors(address collection, uint256 id) view returns(address[])
func (_Franchise *FranchiseSession) GetCharacterCollectors(collection common.Address, id *big.Int) ([]common.Address, error) {
	return _Franchise.Contract.GetCharacterCollectors(&_Franchise.CallOpts, collection, id)
}

// GetCharacterCollectors is a free data retrieval call binding the contract method 0x52430b6a.
//
// Solidity: function getCharacterCollectors(address collection, uint256 id) view returns(address[])
func (_Franchise *FranchiseCallerSession) GetCharacterCollectors(collection common.Address, id *big.Int) ([]common.Address, error) {
	return _Franchise.Contract.GetCharacterCollectors(&_Franchise.CallOpts, collection, id)
}

// GetStoryAutograph is a free data retrieval call binding the contract method 0x78fce0ae.
//
// Solidity: function getStoryAutograph(address collection, uint256 id) view returns(address, uint256)
func (_Franchise *FranchiseCaller) GetStoryAutograph(opts *bind.CallOpts, collection common.Address, id *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Franchise.contract.Call(opts, &out, "getStoryAutograph", collection, id)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetStoryAutograph is a free data retrieval call binding the contract method 0x78fce0ae.
//
// Solidity: function getStoryAutograph(address collection, uint256 id) view returns(address, uint256)
func (_Franchise *FranchiseSession) GetStoryAutograph(collection common.Address, id *big.Int) (common.Address, *big.Int, error) {
	return _Franchise.Contract.GetStoryAutograph(&_Franchise.CallOpts, collection, id)
}

// GetStoryAutograph is a free data retrieval call binding the contract method 0x78fce0ae.
//
// Solidity: function getStoryAutograph(address collection, uint256 id) view returns(address, uint256)
func (_Franchise *FranchiseCallerSession) GetStoryAutograph(collection common.Address, id *big.Int) (common.Address, *big.Int, error) {
	return _Franchise.Contract.GetStoryAutograph(&_Franchise.CallOpts, collection, id)
}

// GetStoryCollectors is a free data retrieval call binding the contract method 0x78016c0e.
//
// Solidity: function getStoryCollectors(address collection, uint256 id) view returns(address[])
func (_Franchise *FranchiseCaller) GetStoryCollectors(opts *bind.CallOpts, collection common.Address, id *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Franchise.contract.Call(opts, &out, "getStoryCollectors", collection, id)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStoryCollectors is a free data retrieval call binding the contract method 0x78016c0e.
//
// Solidity: function getStoryCollectors(address collection, uint256 id) view returns(address[])
func (_Franchise *FranchiseSession) GetStoryCollectors(collection common.Address, id *big.Int) ([]common.Address, error) {
	return _Franchise.Contract.GetStoryCollectors(&_Franchise.CallOpts, collection, id)
}

// GetStoryCollectors is a free data retrieval call binding the contract method 0x78016c0e.
//
// Solidity: function getStoryCollectors(address collection, uint256 id) view returns(address[])
func (_Franchise *FranchiseCallerSession) GetStoryCollectors(collection common.Address, id *big.Int) ([]common.Address, error) {
	return _Franchise.Contract.GetStoryCollectors(&_Franchise.CallOpts, collection, id)
}

// LicenseModule is a free data retrieval call binding the contract method 0xe5bbfcf2.
//
// Solidity: function licenseModule() view returns(address)
func (_Franchise *FranchiseCaller) LicenseModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Franchise.contract.Call(opts, &out, "licenseModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LicenseModule is a free data retrieval call binding the contract method 0xe5bbfcf2.
//
// Solidity: function licenseModule() view returns(address)
func (_Franchise *FranchiseSession) LicenseModule() (common.Address, error) {
	return _Franchise.Contract.LicenseModule(&_Franchise.CallOpts)
}

// LicenseModule is a free data retrieval call binding the contract method 0xe5bbfcf2.
//
// Solidity: function licenseModule() view returns(address)
func (_Franchise *FranchiseCallerSession) LicenseModule() (common.Address, error) {
	return _Franchise.Contract.LicenseModule(&_Franchise.CallOpts)
}

// LicenseNFT is a free data retrieval call binding the contract method 0xda68d47d.
//
// Solidity: function licenseNFT() view returns(address)
func (_Franchise *FranchiseCaller) LicenseNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Franchise.contract.Call(opts, &out, "licenseNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LicenseNFT is a free data retrieval call binding the contract method 0xda68d47d.
//
// Solidity: function licenseNFT() view returns(address)
func (_Franchise *FranchiseSession) LicenseNFT() (common.Address, error) {
	return _Franchise.Contract.LicenseNFT(&_Franchise.CallOpts)
}

// LicenseNFT is a free data retrieval call binding the contract method 0xda68d47d.
//
// Solidity: function licenseNFT() view returns(address)
func (_Franchise *FranchiseCallerSession) LicenseNFT() (common.Address, error) {
	return _Franchise.Contract.LicenseNFT(&_Franchise.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Franchise *FranchiseCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Franchise.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Franchise *FranchiseSession) Name() (string, error) {
	return _Franchise.Contract.Name(&_Franchise.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Franchise *FranchiseCallerSession) Name() (string, error) {
	return _Franchise.Contract.Name(&_Franchise.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Franchise *FranchiseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Franchise.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Franchise *FranchiseSession) Owner() (common.Address, error) {
	return _Franchise.Contract.Owner(&_Franchise.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Franchise *FranchiseCallerSession) Owner() (common.Address, error) {
	return _Franchise.Contract.Owner(&_Franchise.CallOpts)
}

// StoryRegistry is a free data retrieval call binding the contract method 0xc4e6ce9b.
//
// Solidity: function storyRegistry() view returns(address)
func (_Franchise *FranchiseCaller) StoryRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Franchise.contract.Call(opts, &out, "storyRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StoryRegistry is a free data retrieval call binding the contract method 0xc4e6ce9b.
//
// Solidity: function storyRegistry() view returns(address)
func (_Franchise *FranchiseSession) StoryRegistry() (common.Address, error) {
	return _Franchise.Contract.StoryRegistry(&_Franchise.CallOpts)
}

// StoryRegistry is a free data retrieval call binding the contract method 0xc4e6ce9b.
//
// Solidity: function storyRegistry() view returns(address)
func (_Franchise *FranchiseCallerSession) StoryRegistry() (common.Address, error) {
	return _Franchise.Contract.StoryRegistry(&_Franchise.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Franchise *FranchiseCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Franchise.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Franchise *FranchiseSession) Vault() (common.Address, error) {
	return _Franchise.Contract.Vault(&_Franchise.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Franchise *FranchiseCallerSession) Vault() (common.Address, error) {
	return _Franchise.Contract.Vault(&_Franchise.CallOpts)
}

// AddCharacter is a paid mutator transaction binding the contract method 0xced1dd88.
//
// Solidity: function addCharacter(address collection, (string,string,(address,string),address,uint8) info, uint256 externalId, address owner, string imageURI) returns()
func (_Franchise *FranchiseTransactor) AddCharacter(opts *bind.TransactOpts, collection common.Address, info ICharacterRegistryCharacterInfo, externalId *big.Int, owner common.Address, imageURI string) (*types.Transaction, error) {
	return _Franchise.contract.Transact(opts, "addCharacter", collection, info, externalId, owner, imageURI)
}

// AddCharacter is a paid mutator transaction binding the contract method 0xced1dd88.
//
// Solidity: function addCharacter(address collection, (string,string,(address,string),address,uint8) info, uint256 externalId, address owner, string imageURI) returns()
func (_Franchise *FranchiseSession) AddCharacter(collection common.Address, info ICharacterRegistryCharacterInfo, externalId *big.Int, owner common.Address, imageURI string) (*types.Transaction, error) {
	return _Franchise.Contract.AddCharacter(&_Franchise.TransactOpts, collection, info, externalId, owner, imageURI)
}

// AddCharacter is a paid mutator transaction binding the contract method 0xced1dd88.
//
// Solidity: function addCharacter(address collection, (string,string,(address,string),address,uint8) info, uint256 externalId, address owner, string imageURI) returns()
func (_Franchise *FranchiseTransactorSession) AddCharacter(collection common.Address, info ICharacterRegistryCharacterInfo, externalId *big.Int, owner common.Address, imageURI string) (*types.Transaction, error) {
	return _Franchise.Contract.AddCharacter(&_Franchise.TransactOpts, collection, info, externalId, owner, imageURI)
}

// AddFanFiction is a paid mutator transaction binding the contract method 0xf064212b.
//
// Solidity: function addFanFiction((address[],string,(address,uint256)[],address) info) returns()
func (_Franchise *FranchiseTransactor) AddFanFiction(opts *bind.TransactOpts, info StoryInfo) (*types.Transaction, error) {
	return _Franchise.contract.Transact(opts, "addFanFiction", info)
}

// AddFanFiction is a paid mutator transaction binding the contract method 0xf064212b.
//
// Solidity: function addFanFiction((address[],string,(address,uint256)[],address) info) returns()
func (_Franchise *FranchiseSession) AddFanFiction(info StoryInfo) (*types.Transaction, error) {
	return _Franchise.Contract.AddFanFiction(&_Franchise.TransactOpts, info)
}

// AddFanFiction is a paid mutator transaction binding the contract method 0xf064212b.
//
// Solidity: function addFanFiction((address[],string,(address,uint256)[],address) info) returns()
func (_Franchise *FranchiseTransactorSession) AddFanFiction(info StoryInfo) (*types.Transaction, error) {
	return _Franchise.Contract.AddFanFiction(&_Franchise.TransactOpts, info)
}

// AddStory is a paid mutator transaction binding the contract method 0x8463f073.
//
// Solidity: function addStory((address[],string,(address,uint256)[],address) info) returns()
func (_Franchise *FranchiseTransactor) AddStory(opts *bind.TransactOpts, info StoryInfo) (*types.Transaction, error) {
	return _Franchise.contract.Transact(opts, "addStory", info)
}

// AddStory is a paid mutator transaction binding the contract method 0x8463f073.
//
// Solidity: function addStory((address[],string,(address,uint256)[],address) info) returns()
func (_Franchise *FranchiseSession) AddStory(info StoryInfo) (*types.Transaction, error) {
	return _Franchise.Contract.AddStory(&_Franchise.TransactOpts, info)
}

// AddStory is a paid mutator transaction binding the contract method 0x8463f073.
//
// Solidity: function addStory((address[],string,(address,uint256)[],address) info) returns()
func (_Franchise *FranchiseTransactorSession) AddStory(info StoryInfo) (*types.Transaction, error) {
	return _Franchise.Contract.AddStory(&_Franchise.TransactOpts, info)
}

// CollectCharacter is a paid mutator transaction binding the contract method 0xc92ec7e8.
//
// Solidity: function collectCharacter(address collection, uint256 id) returns()
func (_Franchise *FranchiseTransactor) CollectCharacter(opts *bind.TransactOpts, collection common.Address, id *big.Int) (*types.Transaction, error) {
	return _Franchise.contract.Transact(opts, "collectCharacter", collection, id)
}

// CollectCharacter is a paid mutator transaction binding the contract method 0xc92ec7e8.
//
// Solidity: function collectCharacter(address collection, uint256 id) returns()
func (_Franchise *FranchiseSession) CollectCharacter(collection common.Address, id *big.Int) (*types.Transaction, error) {
	return _Franchise.Contract.CollectCharacter(&_Franchise.TransactOpts, collection, id)
}

// CollectCharacter is a paid mutator transaction binding the contract method 0xc92ec7e8.
//
// Solidity: function collectCharacter(address collection, uint256 id) returns()
func (_Franchise *FranchiseTransactorSession) CollectCharacter(collection common.Address, id *big.Int) (*types.Transaction, error) {
	return _Franchise.Contract.CollectCharacter(&_Franchise.TransactOpts, collection, id)
}

// CollectStory is a paid mutator transaction binding the contract method 0x0364c358.
//
// Solidity: function collectStory(address collection, uint256 id) returns()
func (_Franchise *FranchiseTransactor) CollectStory(opts *bind.TransactOpts, collection common.Address, id *big.Int) (*types.Transaction, error) {
	return _Franchise.contract.Transact(opts, "collectStory", collection, id)
}

// CollectStory is a paid mutator transaction binding the contract method 0x0364c358.
//
// Solidity: function collectStory(address collection, uint256 id) returns()
func (_Franchise *FranchiseSession) CollectStory(collection common.Address, id *big.Int) (*types.Transaction, error) {
	return _Franchise.Contract.CollectStory(&_Franchise.TransactOpts, collection, id)
}

// CollectStory is a paid mutator transaction binding the contract method 0x0364c358.
//
// Solidity: function collectStory(address collection, uint256 id) returns()
func (_Franchise *FranchiseTransactorSession) CollectStory(collection common.Address, id *big.Int) (*types.Transaction, error) {
	return _Franchise.Contract.CollectStory(&_Franchise.TransactOpts, collection, id)
}

// InitCharacters is a paid mutator transaction binding the contract method 0xc7433788.
//
// Solidity: function initCharacters((address,string,string) canon, (address,string,string) fanfic) returns()
func (_Franchise *FranchiseTransactor) InitCharacters(opts *bind.TransactOpts, canon IFranchiseCharacterCollection, fanfic IFranchiseCharacterCollection) (*types.Transaction, error) {
	return _Franchise.contract.Transact(opts, "initCharacters", canon, fanfic)
}

// InitCharacters is a paid mutator transaction binding the contract method 0xc7433788.
//
// Solidity: function initCharacters((address,string,string) canon, (address,string,string) fanfic) returns()
func (_Franchise *FranchiseSession) InitCharacters(canon IFranchiseCharacterCollection, fanfic IFranchiseCharacterCollection) (*types.Transaction, error) {
	return _Franchise.Contract.InitCharacters(&_Franchise.TransactOpts, canon, fanfic)
}

// InitCharacters is a paid mutator transaction binding the contract method 0xc7433788.
//
// Solidity: function initCharacters((address,string,string) canon, (address,string,string) fanfic) returns()
func (_Franchise *FranchiseTransactorSession) InitCharacters(canon IFranchiseCharacterCollection, fanfic IFranchiseCharacterCollection) (*types.Transaction, error) {
	return _Franchise.Contract.InitCharacters(&_Franchise.TransactOpts, canon, fanfic)
}

// PromoteFanFiction is a paid mutator transaction binding the contract method 0xe59cac57.
//
// Solidity: function promoteFanFiction(uint256 tokenId) returns()
func (_Franchise *FranchiseTransactor) PromoteFanFiction(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Franchise.contract.Transact(opts, "promoteFanFiction", tokenId)
}

// PromoteFanFiction is a paid mutator transaction binding the contract method 0xe59cac57.
//
// Solidity: function promoteFanFiction(uint256 tokenId) returns()
func (_Franchise *FranchiseSession) PromoteFanFiction(tokenId *big.Int) (*types.Transaction, error) {
	return _Franchise.Contract.PromoteFanFiction(&_Franchise.TransactOpts, tokenId)
}

// PromoteFanFiction is a paid mutator transaction binding the contract method 0xe59cac57.
//
// Solidity: function promoteFanFiction(uint256 tokenId) returns()
func (_Franchise *FranchiseTransactorSession) PromoteFanFiction(tokenId *big.Int) (*types.Transaction, error) {
	return _Franchise.Contract.PromoteFanFiction(&_Franchise.TransactOpts, tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Franchise *FranchiseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Franchise.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Franchise *FranchiseSession) RenounceOwnership() (*types.Transaction, error) {
	return _Franchise.Contract.RenounceOwnership(&_Franchise.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Franchise *FranchiseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Franchise.Contract.RenounceOwnership(&_Franchise.TransactOpts)
}

// RequestLicense is a paid mutator transaction binding the contract method 0xe572b949.
//
// Solidity: function requestLicense(address collection, uint256 tokenID, bytes licenseData) returns()
func (_Franchise *FranchiseTransactor) RequestLicense(opts *bind.TransactOpts, collection common.Address, tokenID *big.Int, licenseData []byte) (*types.Transaction, error) {
	return _Franchise.contract.Transact(opts, "requestLicense", collection, tokenID, licenseData)
}

// RequestLicense is a paid mutator transaction binding the contract method 0xe572b949.
//
// Solidity: function requestLicense(address collection, uint256 tokenID, bytes licenseData) returns()
func (_Franchise *FranchiseSession) RequestLicense(collection common.Address, tokenID *big.Int, licenseData []byte) (*types.Transaction, error) {
	return _Franchise.Contract.RequestLicense(&_Franchise.TransactOpts, collection, tokenID, licenseData)
}

// RequestLicense is a paid mutator transaction binding the contract method 0xe572b949.
//
// Solidity: function requestLicense(address collection, uint256 tokenID, bytes licenseData) returns()
func (_Franchise *FranchiseTransactorSession) RequestLicense(collection common.Address, tokenID *big.Int, licenseData []byte) (*types.Transaction, error) {
	return _Franchise.Contract.RequestLicense(&_Franchise.TransactOpts, collection, tokenID, licenseData)
}

// SetLicenseModule is a paid mutator transaction binding the contract method 0xa0bc44eb.
//
// Solidity: function setLicenseModule(address module) returns()
func (_Franchise *FranchiseTransactor) SetLicenseModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _Franchise.contract.Transact(opts, "setLicenseModule", module)
}

// SetLicenseModule is a paid mutator transaction binding the contract method 0xa0bc44eb.
//
// Solidity: function setLicenseModule(address module) returns()
func (_Franchise *FranchiseSession) SetLicenseModule(module common.Address) (*types.Transaction, error) {
	return _Franchise.Contract.SetLicenseModule(&_Franchise.TransactOpts, module)
}

// SetLicenseModule is a paid mutator transaction binding the contract method 0xa0bc44eb.
//
// Solidity: function setLicenseModule(address module) returns()
func (_Franchise *FranchiseTransactorSession) SetLicenseModule(module common.Address) (*types.Transaction, error) {
	return _Franchise.Contract.SetLicenseModule(&_Franchise.TransactOpts, module)
}

// SetStoryRegistry is a paid mutator transaction binding the contract method 0x8d6c058b.
//
// Solidity: function setStoryRegistry(address storyRegistry_) returns()
func (_Franchise *FranchiseTransactor) SetStoryRegistry(opts *bind.TransactOpts, storyRegistry_ common.Address) (*types.Transaction, error) {
	return _Franchise.contract.Transact(opts, "setStoryRegistry", storyRegistry_)
}

// SetStoryRegistry is a paid mutator transaction binding the contract method 0x8d6c058b.
//
// Solidity: function setStoryRegistry(address storyRegistry_) returns()
func (_Franchise *FranchiseSession) SetStoryRegistry(storyRegistry_ common.Address) (*types.Transaction, error) {
	return _Franchise.Contract.SetStoryRegistry(&_Franchise.TransactOpts, storyRegistry_)
}

// SetStoryRegistry is a paid mutator transaction binding the contract method 0x8d6c058b.
//
// Solidity: function setStoryRegistry(address storyRegistry_) returns()
func (_Franchise *FranchiseTransactorSession) SetStoryRegistry(storyRegistry_ common.Address) (*types.Transaction, error) {
	return _Franchise.Contract.SetStoryRegistry(&_Franchise.TransactOpts, storyRegistry_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Franchise *FranchiseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Franchise.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Franchise *FranchiseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Franchise.Contract.TransferOwnership(&_Franchise.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Franchise *FranchiseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Franchise.Contract.TransferOwnership(&_Franchise.TransactOpts, newOwner)
}

// FranchiseCharactersCreatedIterator is returned from FilterCharactersCreated and is used to iterate over the raw logs and unpacked data for CharactersCreated events raised by the Franchise contract.
type FranchiseCharactersCreatedIterator struct {
	Event *FranchiseCharactersCreated // Event containing the contract specifics and raw log

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
func (it *FranchiseCharactersCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FranchiseCharactersCreated)
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
		it.Event = new(FranchiseCharactersCreated)
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
func (it *FranchiseCharactersCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FranchiseCharactersCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FranchiseCharactersCreated represents a CharactersCreated event raised by the Franchise contract.
type FranchiseCharactersCreated struct {
	Registry         common.Address
	CanonCollection  common.Address
	FanficCollection common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCharactersCreated is a free log retrieval operation binding the contract event 0x0926c24f70a117b6b0ef1f9ff269362edc450c1cc5ef61e30d503a437e7ec48f.
//
// Solidity: event CharactersCreated(address registry, address canonCollection, address fanficCollection)
func (_Franchise *FranchiseFilterer) FilterCharactersCreated(opts *bind.FilterOpts) (*FranchiseCharactersCreatedIterator, error) {

	logs, sub, err := _Franchise.contract.FilterLogs(opts, "CharactersCreated")
	if err != nil {
		return nil, err
	}
	return &FranchiseCharactersCreatedIterator{contract: _Franchise.contract, event: "CharactersCreated", logs: logs, sub: sub}, nil
}

// WatchCharactersCreated is a free log subscription operation binding the contract event 0x0926c24f70a117b6b0ef1f9ff269362edc450c1cc5ef61e30d503a437e7ec48f.
//
// Solidity: event CharactersCreated(address registry, address canonCollection, address fanficCollection)
func (_Franchise *FranchiseFilterer) WatchCharactersCreated(opts *bind.WatchOpts, sink chan<- *FranchiseCharactersCreated) (event.Subscription, error) {

	logs, sub, err := _Franchise.contract.WatchLogs(opts, "CharactersCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FranchiseCharactersCreated)
				if err := _Franchise.contract.UnpackLog(event, "CharactersCreated", log); err != nil {
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

// ParseCharactersCreated is a log parse operation binding the contract event 0x0926c24f70a117b6b0ef1f9ff269362edc450c1cc5ef61e30d503a437e7ec48f.
//
// Solidity: event CharactersCreated(address registry, address canonCollection, address fanficCollection)
func (_Franchise *FranchiseFilterer) ParseCharactersCreated(log types.Log) (*FranchiseCharactersCreated, error) {
	event := new(FranchiseCharactersCreated)
	if err := _Franchise.contract.UnpackLog(event, "CharactersCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FranchiseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Franchise contract.
type FranchiseOwnershipTransferredIterator struct {
	Event *FranchiseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FranchiseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FranchiseOwnershipTransferred)
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
		it.Event = new(FranchiseOwnershipTransferred)
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
func (it *FranchiseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FranchiseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FranchiseOwnershipTransferred represents a OwnershipTransferred event raised by the Franchise contract.
type FranchiseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Franchise *FranchiseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FranchiseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Franchise.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FranchiseOwnershipTransferredIterator{contract: _Franchise.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Franchise *FranchiseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FranchiseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Franchise.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FranchiseOwnershipTransferred)
				if err := _Franchise.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Franchise *FranchiseFilterer) ParseOwnershipTransferred(log types.Log) (*FranchiseOwnershipTransferred, error) {
	event := new(FranchiseOwnershipTransferred)
	if err := _Franchise.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FranchiseStoriesCreatedIterator is returned from FilterStoriesCreated and is used to iterate over the raw logs and unpacked data for StoriesCreated events raised by the Franchise contract.
type FranchiseStoriesCreatedIterator struct {
	Event *FranchiseStoriesCreated // Event containing the contract specifics and raw log

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
func (it *FranchiseStoriesCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FranchiseStoriesCreated)
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
		it.Event = new(FranchiseStoriesCreated)
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
func (it *FranchiseStoriesCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FranchiseStoriesCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FranchiseStoriesCreated represents a StoriesCreated event raised by the Franchise contract.
type FranchiseStoriesCreated struct {
	Registry         common.Address
	CanonCollection  common.Address
	FanficCollection common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterStoriesCreated is a free log retrieval operation binding the contract event 0xd4a5ffcdc7537c2393e2f0a4b947c163d50c2ac75cc5e783117b8a39ee6e4230.
//
// Solidity: event StoriesCreated(address registry, address canonCollection, address fanficCollection)
func (_Franchise *FranchiseFilterer) FilterStoriesCreated(opts *bind.FilterOpts) (*FranchiseStoriesCreatedIterator, error) {

	logs, sub, err := _Franchise.contract.FilterLogs(opts, "StoriesCreated")
	if err != nil {
		return nil, err
	}
	return &FranchiseStoriesCreatedIterator{contract: _Franchise.contract, event: "StoriesCreated", logs: logs, sub: sub}, nil
}

// WatchStoriesCreated is a free log subscription operation binding the contract event 0xd4a5ffcdc7537c2393e2f0a4b947c163d50c2ac75cc5e783117b8a39ee6e4230.
//
// Solidity: event StoriesCreated(address registry, address canonCollection, address fanficCollection)
func (_Franchise *FranchiseFilterer) WatchStoriesCreated(opts *bind.WatchOpts, sink chan<- *FranchiseStoriesCreated) (event.Subscription, error) {

	logs, sub, err := _Franchise.contract.WatchLogs(opts, "StoriesCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FranchiseStoriesCreated)
				if err := _Franchise.contract.UnpackLog(event, "StoriesCreated", log); err != nil {
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

// ParseStoriesCreated is a log parse operation binding the contract event 0xd4a5ffcdc7537c2393e2f0a4b947c163d50c2ac75cc5e783117b8a39ee6e4230.
//
// Solidity: event StoriesCreated(address registry, address canonCollection, address fanficCollection)
func (_Franchise *FranchiseFilterer) ParseStoriesCreated(log types.Log) (*FranchiseStoriesCreated, error) {
	event := new(FranchiseStoriesCreated)
	if err := _Franchise.contract.UnpackLog(event, "StoriesCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
