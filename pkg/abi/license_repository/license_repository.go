// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package license_repository

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

// LicenseRepositoryLicense is an auto generated low-level Go binding around an user-defined struct.
type LicenseRepositoryLicense struct {
	Rights   uint8
	TermsURI string
}

// LicenseRepositoryMetaData contains all meta data concerning the LicenseRepository contract.
var LicenseRepositoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"termsURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumLicenseRepository.Rights\",\"name\":\"rights\",\"type\":\"uint8\"}],\"name\":\"LicenseTemplateCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"enumLicenseRepository.Rights\",\"name\":\"rights\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"termsUri\",\"type\":\"string\"}],\"name\":\"addTemplate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"licenseTemplateAt\",\"outputs\":[{\"components\":[{\"internalType\":\"enumLicenseRepository.Rights\",\"name\":\"rights\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"termsURI\",\"type\":\"string\"}],\"internalType\":\"structLicenseRepository.License\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"licenseTemplates\",\"outputs\":[{\"internalType\":\"enumLicenseRepository.Rights\",\"name\":\"rights\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"termsURI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLicenseTemplates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LicenseRepositoryABI is the input ABI used to generate the binding from.
// Deprecated: Use LicenseRepositoryMetaData.ABI instead.
var LicenseRepositoryABI = LicenseRepositoryMetaData.ABI

// LicenseRepository is an auto generated Go binding around an Ethereum contract.
type LicenseRepository struct {
	LicenseRepositoryCaller     // Read-only binding to the contract
	LicenseRepositoryTransactor // Write-only binding to the contract
	LicenseRepositoryFilterer   // Log filterer for contract events
}

// LicenseRepositoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LicenseRepositoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LicenseRepositoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LicenseRepositoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LicenseRepositoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LicenseRepositoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LicenseRepositorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LicenseRepositorySession struct {
	Contract     *LicenseRepository // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LicenseRepositoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LicenseRepositoryCallerSession struct {
	Contract *LicenseRepositoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LicenseRepositoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LicenseRepositoryTransactorSession struct {
	Contract     *LicenseRepositoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LicenseRepositoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LicenseRepositoryRaw struct {
	Contract *LicenseRepository // Generic contract binding to access the raw methods on
}

// LicenseRepositoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LicenseRepositoryCallerRaw struct {
	Contract *LicenseRepositoryCaller // Generic read-only contract binding to access the raw methods on
}

// LicenseRepositoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LicenseRepositoryTransactorRaw struct {
	Contract *LicenseRepositoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLicenseRepository creates a new instance of LicenseRepository, bound to a specific deployed contract.
func NewLicenseRepository(address common.Address, backend bind.ContractBackend) (*LicenseRepository, error) {
	contract, err := bindLicenseRepository(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LicenseRepository{LicenseRepositoryCaller: LicenseRepositoryCaller{contract: contract}, LicenseRepositoryTransactor: LicenseRepositoryTransactor{contract: contract}, LicenseRepositoryFilterer: LicenseRepositoryFilterer{contract: contract}}, nil
}

// NewLicenseRepositoryCaller creates a new read-only instance of LicenseRepository, bound to a specific deployed contract.
func NewLicenseRepositoryCaller(address common.Address, caller bind.ContractCaller) (*LicenseRepositoryCaller, error) {
	contract, err := bindLicenseRepository(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LicenseRepositoryCaller{contract: contract}, nil
}

// NewLicenseRepositoryTransactor creates a new write-only instance of LicenseRepository, bound to a specific deployed contract.
func NewLicenseRepositoryTransactor(address common.Address, transactor bind.ContractTransactor) (*LicenseRepositoryTransactor, error) {
	contract, err := bindLicenseRepository(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LicenseRepositoryTransactor{contract: contract}, nil
}

// NewLicenseRepositoryFilterer creates a new log filterer instance of LicenseRepository, bound to a specific deployed contract.
func NewLicenseRepositoryFilterer(address common.Address, filterer bind.ContractFilterer) (*LicenseRepositoryFilterer, error) {
	contract, err := bindLicenseRepository(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LicenseRepositoryFilterer{contract: contract}, nil
}

// bindLicenseRepository binds a generic wrapper to an already deployed contract.
func bindLicenseRepository(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LicenseRepositoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LicenseRepository *LicenseRepositoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LicenseRepository.Contract.LicenseRepositoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LicenseRepository *LicenseRepositoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LicenseRepository.Contract.LicenseRepositoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LicenseRepository *LicenseRepositoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LicenseRepository.Contract.LicenseRepositoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LicenseRepository *LicenseRepositoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LicenseRepository.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LicenseRepository *LicenseRepositoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LicenseRepository.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LicenseRepository *LicenseRepositoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LicenseRepository.Contract.contract.Transact(opts, method, params...)
}

// LicenseTemplateAt is a free data retrieval call binding the contract method 0x0bfdebb5.
//
// Solidity: function licenseTemplateAt(uint256 index) view returns((uint8,string))
func (_LicenseRepository *LicenseRepositoryCaller) LicenseTemplateAt(opts *bind.CallOpts, index *big.Int) (LicenseRepositoryLicense, error) {
	var out []interface{}
	err := _LicenseRepository.contract.Call(opts, &out, "licenseTemplateAt", index)

	if err != nil {
		return *new(LicenseRepositoryLicense), err
	}

	out0 := *abi.ConvertType(out[0], new(LicenseRepositoryLicense)).(*LicenseRepositoryLicense)

	return out0, err

}

// LicenseTemplateAt is a free data retrieval call binding the contract method 0x0bfdebb5.
//
// Solidity: function licenseTemplateAt(uint256 index) view returns((uint8,string))
func (_LicenseRepository *LicenseRepositorySession) LicenseTemplateAt(index *big.Int) (LicenseRepositoryLicense, error) {
	return _LicenseRepository.Contract.LicenseTemplateAt(&_LicenseRepository.CallOpts, index)
}

// LicenseTemplateAt is a free data retrieval call binding the contract method 0x0bfdebb5.
//
// Solidity: function licenseTemplateAt(uint256 index) view returns((uint8,string))
func (_LicenseRepository *LicenseRepositoryCallerSession) LicenseTemplateAt(index *big.Int) (LicenseRepositoryLicense, error) {
	return _LicenseRepository.Contract.LicenseTemplateAt(&_LicenseRepository.CallOpts, index)
}

// LicenseTemplates is a free data retrieval call binding the contract method 0xba232648.
//
// Solidity: function licenseTemplates(uint256 ) view returns(uint8 rights, string termsURI)
func (_LicenseRepository *LicenseRepositoryCaller) LicenseTemplates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Rights   uint8
	TermsURI string
}, error) {
	var out []interface{}
	err := _LicenseRepository.contract.Call(opts, &out, "licenseTemplates", arg0)

	outstruct := new(struct {
		Rights   uint8
		TermsURI string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Rights = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.TermsURI = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// LicenseTemplates is a free data retrieval call binding the contract method 0xba232648.
//
// Solidity: function licenseTemplates(uint256 ) view returns(uint8 rights, string termsURI)
func (_LicenseRepository *LicenseRepositorySession) LicenseTemplates(arg0 *big.Int) (struct {
	Rights   uint8
	TermsURI string
}, error) {
	return _LicenseRepository.Contract.LicenseTemplates(&_LicenseRepository.CallOpts, arg0)
}

// LicenseTemplates is a free data retrieval call binding the contract method 0xba232648.
//
// Solidity: function licenseTemplates(uint256 ) view returns(uint8 rights, string termsURI)
func (_LicenseRepository *LicenseRepositoryCallerSession) LicenseTemplates(arg0 *big.Int) (struct {
	Rights   uint8
	TermsURI string
}, error) {
	return _LicenseRepository.Contract.LicenseTemplates(&_LicenseRepository.CallOpts, arg0)
}

// TotalLicenseTemplates is a free data retrieval call binding the contract method 0x577a5f5d.
//
// Solidity: function totalLicenseTemplates() view returns(uint256)
func (_LicenseRepository *LicenseRepositoryCaller) TotalLicenseTemplates(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LicenseRepository.contract.Call(opts, &out, "totalLicenseTemplates")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLicenseTemplates is a free data retrieval call binding the contract method 0x577a5f5d.
//
// Solidity: function totalLicenseTemplates() view returns(uint256)
func (_LicenseRepository *LicenseRepositorySession) TotalLicenseTemplates() (*big.Int, error) {
	return _LicenseRepository.Contract.TotalLicenseTemplates(&_LicenseRepository.CallOpts)
}

// TotalLicenseTemplates is a free data retrieval call binding the contract method 0x577a5f5d.
//
// Solidity: function totalLicenseTemplates() view returns(uint256)
func (_LicenseRepository *LicenseRepositoryCallerSession) TotalLicenseTemplates() (*big.Int, error) {
	return _LicenseRepository.Contract.TotalLicenseTemplates(&_LicenseRepository.CallOpts)
}

// AddTemplate is a paid mutator transaction binding the contract method 0xd0e85ba7.
//
// Solidity: function addTemplate(uint8 rights, string termsUri) returns(uint256)
func (_LicenseRepository *LicenseRepositoryTransactor) AddTemplate(opts *bind.TransactOpts, rights uint8, termsUri string) (*types.Transaction, error) {
	return _LicenseRepository.contract.Transact(opts, "addTemplate", rights, termsUri)
}

// AddTemplate is a paid mutator transaction binding the contract method 0xd0e85ba7.
//
// Solidity: function addTemplate(uint8 rights, string termsUri) returns(uint256)
func (_LicenseRepository *LicenseRepositorySession) AddTemplate(rights uint8, termsUri string) (*types.Transaction, error) {
	return _LicenseRepository.Contract.AddTemplate(&_LicenseRepository.TransactOpts, rights, termsUri)
}

// AddTemplate is a paid mutator transaction binding the contract method 0xd0e85ba7.
//
// Solidity: function addTemplate(uint8 rights, string termsUri) returns(uint256)
func (_LicenseRepository *LicenseRepositoryTransactorSession) AddTemplate(rights uint8, termsUri string) (*types.Transaction, error) {
	return _LicenseRepository.Contract.AddTemplate(&_LicenseRepository.TransactOpts, rights, termsUri)
}

// LicenseRepositoryLicenseTemplateCreatedIterator is returned from FilterLicenseTemplateCreated and is used to iterate over the raw logs and unpacked data for LicenseTemplateCreated events raised by the LicenseRepository contract.
type LicenseRepositoryLicenseTemplateCreatedIterator struct {
	Event *LicenseRepositoryLicenseTemplateCreated // Event containing the contract specifics and raw log

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
func (it *LicenseRepositoryLicenseTemplateCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenseRepositoryLicenseTemplateCreated)
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
		it.Event = new(LicenseRepositoryLicenseTemplateCreated)
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
func (it *LicenseRepositoryLicenseTemplateCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenseRepositoryLicenseTemplateCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenseRepositoryLicenseTemplateCreated represents a LicenseTemplateCreated event raised by the LicenseRepository contract.
type LicenseRepositoryLicenseTemplateCreated struct {
	LicenseId *big.Int
	TermsURI  string
	Rights    uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLicenseTemplateCreated is a free log retrieval operation binding the contract event 0xdf9796cf7d77cd0946d474f80bce815c431a58786366c89da9c7c4d5b91a6eee.
//
// Solidity: event LicenseTemplateCreated(uint256 indexed licenseId, string termsURI, uint8 rights)
func (_LicenseRepository *LicenseRepositoryFilterer) FilterLicenseTemplateCreated(opts *bind.FilterOpts, licenseId []*big.Int) (*LicenseRepositoryLicenseTemplateCreatedIterator, error) {

	var licenseIdRule []interface{}
	for _, licenseIdItem := range licenseId {
		licenseIdRule = append(licenseIdRule, licenseIdItem)
	}

	logs, sub, err := _LicenseRepository.contract.FilterLogs(opts, "LicenseTemplateCreated", licenseIdRule)
	if err != nil {
		return nil, err
	}
	return &LicenseRepositoryLicenseTemplateCreatedIterator{contract: _LicenseRepository.contract, event: "LicenseTemplateCreated", logs: logs, sub: sub}, nil
}

// WatchLicenseTemplateCreated is a free log subscription operation binding the contract event 0xdf9796cf7d77cd0946d474f80bce815c431a58786366c89da9c7c4d5b91a6eee.
//
// Solidity: event LicenseTemplateCreated(uint256 indexed licenseId, string termsURI, uint8 rights)
func (_LicenseRepository *LicenseRepositoryFilterer) WatchLicenseTemplateCreated(opts *bind.WatchOpts, sink chan<- *LicenseRepositoryLicenseTemplateCreated, licenseId []*big.Int) (event.Subscription, error) {

	var licenseIdRule []interface{}
	for _, licenseIdItem := range licenseId {
		licenseIdRule = append(licenseIdRule, licenseIdItem)
	}

	logs, sub, err := _LicenseRepository.contract.WatchLogs(opts, "LicenseTemplateCreated", licenseIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenseRepositoryLicenseTemplateCreated)
				if err := _LicenseRepository.contract.UnpackLog(event, "LicenseTemplateCreated", log); err != nil {
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

// ParseLicenseTemplateCreated is a log parse operation binding the contract event 0xdf9796cf7d77cd0946d474f80bce815c431a58786366c89da9c7c4d5b91a6eee.
//
// Solidity: event LicenseTemplateCreated(uint256 indexed licenseId, string termsURI, uint8 rights)
func (_LicenseRepository *LicenseRepositoryFilterer) ParseLicenseTemplateCreated(log types.Log) (*LicenseRepositoryLicenseTemplateCreated, error) {
	event := new(LicenseRepositoryLicenseTemplateCreated)
	if err := _LicenseRepository.contract.UnpackLog(event, "LicenseTemplateCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
