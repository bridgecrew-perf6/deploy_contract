// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TestFilter

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TestFilterABI is the input ABI used to generate the binding from.
const TestFilterABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getAge\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_age\",\"type\":\"uint256\"}],\"name\":\"setAge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name_\",\"type\":\"string\"},{\"name\":\"age_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"newName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TestSetName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAge\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newAge\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TestSetAge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"funcName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"TestGetName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"funcName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"TestGetAge\",\"type\":\"event\"}]"

// TestFilter is an auto generated Go binding around an Ethereum contract.
type TestFilter struct {
	TestFilterCaller     // Read-only binding to the contract
	TestFilterTransactor // Write-only binding to the contract
	TestFilterFilterer   // Log filterer for contract events
}

// TestFilterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestFilterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestFilterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestFilterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestFilterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestFilterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestFilterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestFilterSession struct {
	Contract     *TestFilter       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestFilterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestFilterCallerSession struct {
	Contract *TestFilterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TestFilterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestFilterTransactorSession struct {
	Contract     *TestFilterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TestFilterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestFilterRaw struct {
	Contract *TestFilter // Generic contract binding to access the raw methods on
}

// TestFilterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestFilterCallerRaw struct {
	Contract *TestFilterCaller // Generic read-only contract binding to access the raw methods on
}

// TestFilterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestFilterTransactorRaw struct {
	Contract *TestFilterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestFilter creates a new instance of TestFilter, bound to a specific deployed contract.
func NewTestFilter(address common.Address, backend bind.ContractBackend) (*TestFilter, error) {
	contract, err := bindTestFilter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestFilter{TestFilterCaller: TestFilterCaller{contract: contract}, TestFilterTransactor: TestFilterTransactor{contract: contract}, TestFilterFilterer: TestFilterFilterer{contract: contract}}, nil
}

// NewTestFilterCaller creates a new read-only instance of TestFilter, bound to a specific deployed contract.
func NewTestFilterCaller(address common.Address, caller bind.ContractCaller) (*TestFilterCaller, error) {
	contract, err := bindTestFilter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestFilterCaller{contract: contract}, nil
}

// NewTestFilterTransactor creates a new write-only instance of TestFilter, bound to a specific deployed contract.
func NewTestFilterTransactor(address common.Address, transactor bind.ContractTransactor) (*TestFilterTransactor, error) {
	contract, err := bindTestFilter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestFilterTransactor{contract: contract}, nil
}

// NewTestFilterFilterer creates a new log filterer instance of TestFilter, bound to a specific deployed contract.
func NewTestFilterFilterer(address common.Address, filterer bind.ContractFilterer) (*TestFilterFilterer, error) {
	contract, err := bindTestFilter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestFilterFilterer{contract: contract}, nil
}

// bindTestFilter binds a generic wrapper to an already deployed contract.
func bindTestFilter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestFilterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestFilter *TestFilterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestFilter.Contract.TestFilterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestFilter *TestFilterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestFilter.Contract.TestFilterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestFilter *TestFilterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestFilter.Contract.TestFilterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestFilter *TestFilterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestFilter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestFilter *TestFilterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestFilter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestFilter *TestFilterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestFilter.Contract.contract.Transact(opts, method, params...)
}

// GetAge is a paid mutator transaction binding the contract method 0x967e6e65.
//
// Solidity: function getAge() payable returns(uint256)
func (_TestFilter *TestFilterTransactor) GetAge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestFilter.contract.Transact(opts, "getAge")
}

// GetAge is a paid mutator transaction binding the contract method 0x967e6e65.
//
// Solidity: function getAge() payable returns(uint256)
func (_TestFilter *TestFilterSession) GetAge() (*types.Transaction, error) {
	return _TestFilter.Contract.GetAge(&_TestFilter.TransactOpts)
}

// GetAge is a paid mutator transaction binding the contract method 0x967e6e65.
//
// Solidity: function getAge() payable returns(uint256)
func (_TestFilter *TestFilterTransactorSession) GetAge() (*types.Transaction, error) {
	return _TestFilter.Contract.GetAge(&_TestFilter.TransactOpts)
}

// GetName is a paid mutator transaction binding the contract method 0x17d7de7c.
//
// Solidity: function getName() payable returns(string)
func (_TestFilter *TestFilterTransactor) GetName(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestFilter.contract.Transact(opts, "getName")
}

// GetName is a paid mutator transaction binding the contract method 0x17d7de7c.
//
// Solidity: function getName() payable returns(string)
func (_TestFilter *TestFilterSession) GetName() (*types.Transaction, error) {
	return _TestFilter.Contract.GetName(&_TestFilter.TransactOpts)
}

// GetName is a paid mutator transaction binding the contract method 0x17d7de7c.
//
// Solidity: function getName() payable returns(string)
func (_TestFilter *TestFilterTransactorSession) GetName() (*types.Transaction, error) {
	return _TestFilter.Contract.GetName(&_TestFilter.TransactOpts)
}

// SetAge is a paid mutator transaction binding the contract method 0xd5dcf127.
//
// Solidity: function setAge(uint256 _age) returns()
func (_TestFilter *TestFilterTransactor) SetAge(opts *bind.TransactOpts, _age *big.Int) (*types.Transaction, error) {
	return _TestFilter.contract.Transact(opts, "setAge", _age)
}

// SetAge is a paid mutator transaction binding the contract method 0xd5dcf127.
//
// Solidity: function setAge(uint256 _age) returns()
func (_TestFilter *TestFilterSession) SetAge(_age *big.Int) (*types.Transaction, error) {
	return _TestFilter.Contract.SetAge(&_TestFilter.TransactOpts, _age)
}

// SetAge is a paid mutator transaction binding the contract method 0xd5dcf127.
//
// Solidity: function setAge(uint256 _age) returns()
func (_TestFilter *TestFilterTransactorSession) SetAge(_age *big.Int) (*types.Transaction, error) {
	return _TestFilter.Contract.SetAge(&_TestFilter.TransactOpts, _age)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_TestFilter *TestFilterTransactor) SetName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _TestFilter.contract.Transact(opts, "setName", _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_TestFilter *TestFilterSession) SetName(_name string) (*types.Transaction, error) {
	return _TestFilter.Contract.SetName(&_TestFilter.TransactOpts, _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_TestFilter *TestFilterTransactorSession) SetName(_name string) (*types.Transaction, error) {
	return _TestFilter.Contract.SetName(&_TestFilter.TransactOpts, _name)
}

// TestFilterTestGetAgeIterator is returned from FilterTestGetAge and is used to iterate over the raw logs and unpacked data for TestGetAge events raised by the TestFilter contract.
type TestFilterTestGetAgeIterator struct {
	Event *TestFilterTestGetAge // Event containing the contract specifics and raw log

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
func (it *TestFilterTestGetAgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestFilterTestGetAge)
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
		it.Event = new(TestFilterTestGetAge)
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
func (it *TestFilterTestGetAgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestFilterTestGetAgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestFilterTestGetAge represents a TestGetAge event raised by the TestFilter contract.
type TestFilterTestGetAge struct {
	FuncName    string
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTestGetAge is a free log retrieval operation binding the contract event 0x8cc3a86be5eb071b844e63d24f8bc49c8f9ed32ee6d51c9d27799b096a3b1a5f.
//
// Solidity: event TestGetAge(string funcName, uint256 blockNumber)
func (_TestFilter *TestFilterFilterer) FilterTestGetAge(opts *bind.FilterOpts) (*TestFilterTestGetAgeIterator, error) {

	logs, sub, err := _TestFilter.contract.FilterLogs(opts, "TestGetAge")
	if err != nil {
		return nil, err
	}
	return &TestFilterTestGetAgeIterator{contract: _TestFilter.contract, event: "TestGetAge", logs: logs, sub: sub}, nil
}

// WatchTestGetAge is a free log subscription operation binding the contract event 0x8cc3a86be5eb071b844e63d24f8bc49c8f9ed32ee6d51c9d27799b096a3b1a5f.
//
// Solidity: event TestGetAge(string funcName, uint256 blockNumber)
func (_TestFilter *TestFilterFilterer) WatchTestGetAge(opts *bind.WatchOpts, sink chan<- *TestFilterTestGetAge) (event.Subscription, error) {

	logs, sub, err := _TestFilter.contract.WatchLogs(opts, "TestGetAge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestFilterTestGetAge)
				if err := _TestFilter.contract.UnpackLog(event, "TestGetAge", log); err != nil {
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

// ParseTestGetAge is a log parse operation binding the contract event 0x8cc3a86be5eb071b844e63d24f8bc49c8f9ed32ee6d51c9d27799b096a3b1a5f.
//
// Solidity: event TestGetAge(string funcName, uint256 blockNumber)
func (_TestFilter *TestFilterFilterer) ParseTestGetAge(log types.Log) (*TestFilterTestGetAge, error) {
	event := new(TestFilterTestGetAge)
	if err := _TestFilter.contract.UnpackLog(event, "TestGetAge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestFilterTestGetNameIterator is returned from FilterTestGetName and is used to iterate over the raw logs and unpacked data for TestGetName events raised by the TestFilter contract.
type TestFilterTestGetNameIterator struct {
	Event *TestFilterTestGetName // Event containing the contract specifics and raw log

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
func (it *TestFilterTestGetNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestFilterTestGetName)
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
		it.Event = new(TestFilterTestGetName)
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
func (it *TestFilterTestGetNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestFilterTestGetNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestFilterTestGetName represents a TestGetName event raised by the TestFilter contract.
type TestFilterTestGetName struct {
	FuncName    string
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTestGetName is a free log retrieval operation binding the contract event 0x8f314ef2ed508967ae72c0a8ea5b69d8c5c4c14b88111b93d68ad8e6c640632e.
//
// Solidity: event TestGetName(string funcName, uint256 blockNumber)
func (_TestFilter *TestFilterFilterer) FilterTestGetName(opts *bind.FilterOpts) (*TestFilterTestGetNameIterator, error) {

	logs, sub, err := _TestFilter.contract.FilterLogs(opts, "TestGetName")
	if err != nil {
		return nil, err
	}
	return &TestFilterTestGetNameIterator{contract: _TestFilter.contract, event: "TestGetName", logs: logs, sub: sub}, nil
}

// WatchTestGetName is a free log subscription operation binding the contract event 0x8f314ef2ed508967ae72c0a8ea5b69d8c5c4c14b88111b93d68ad8e6c640632e.
//
// Solidity: event TestGetName(string funcName, uint256 blockNumber)
func (_TestFilter *TestFilterFilterer) WatchTestGetName(opts *bind.WatchOpts, sink chan<- *TestFilterTestGetName) (event.Subscription, error) {

	logs, sub, err := _TestFilter.contract.WatchLogs(opts, "TestGetName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestFilterTestGetName)
				if err := _TestFilter.contract.UnpackLog(event, "TestGetName", log); err != nil {
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

// ParseTestGetName is a log parse operation binding the contract event 0x8f314ef2ed508967ae72c0a8ea5b69d8c5c4c14b88111b93d68ad8e6c640632e.
//
// Solidity: event TestGetName(string funcName, uint256 blockNumber)
func (_TestFilter *TestFilterFilterer) ParseTestGetName(log types.Log) (*TestFilterTestGetName, error) {
	event := new(TestFilterTestGetName)
	if err := _TestFilter.contract.UnpackLog(event, "TestGetName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestFilterTestSetAgeIterator is returned from FilterTestSetAge and is used to iterate over the raw logs and unpacked data for TestSetAge events raised by the TestFilter contract.
type TestFilterTestSetAgeIterator struct {
	Event *TestFilterTestSetAge // Event containing the contract specifics and raw log

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
func (it *TestFilterTestSetAgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestFilterTestSetAge)
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
		it.Event = new(TestFilterTestSetAge)
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
func (it *TestFilterTestSetAgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestFilterTestSetAgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestFilterTestSetAge represents a TestSetAge event raised by the TestFilter contract.
type TestFilterTestSetAge struct {
	OldAge    *big.Int
	NewAge    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTestSetAge is a free log retrieval operation binding the contract event 0xf9905ed439c974a0f8c3a6dceec1e59ff1282efae84554b2e43ace1c0213b566.
//
// Solidity: event TestSetAge(uint256 oldAge, uint256 newAge, uint256 timestamp)
func (_TestFilter *TestFilterFilterer) FilterTestSetAge(opts *bind.FilterOpts) (*TestFilterTestSetAgeIterator, error) {

	logs, sub, err := _TestFilter.contract.FilterLogs(opts, "TestSetAge")
	if err != nil {
		return nil, err
	}
	return &TestFilterTestSetAgeIterator{contract: _TestFilter.contract, event: "TestSetAge", logs: logs, sub: sub}, nil
}

// WatchTestSetAge is a free log subscription operation binding the contract event 0xf9905ed439c974a0f8c3a6dceec1e59ff1282efae84554b2e43ace1c0213b566.
//
// Solidity: event TestSetAge(uint256 oldAge, uint256 newAge, uint256 timestamp)
func (_TestFilter *TestFilterFilterer) WatchTestSetAge(opts *bind.WatchOpts, sink chan<- *TestFilterTestSetAge) (event.Subscription, error) {

	logs, sub, err := _TestFilter.contract.WatchLogs(opts, "TestSetAge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestFilterTestSetAge)
				if err := _TestFilter.contract.UnpackLog(event, "TestSetAge", log); err != nil {
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

// ParseTestSetAge is a log parse operation binding the contract event 0xf9905ed439c974a0f8c3a6dceec1e59ff1282efae84554b2e43ace1c0213b566.
//
// Solidity: event TestSetAge(uint256 oldAge, uint256 newAge, uint256 timestamp)
func (_TestFilter *TestFilterFilterer) ParseTestSetAge(log types.Log) (*TestFilterTestSetAge, error) {
	event := new(TestFilterTestSetAge)
	if err := _TestFilter.contract.UnpackLog(event, "TestSetAge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestFilterTestSetNameIterator is returned from FilterTestSetName and is used to iterate over the raw logs and unpacked data for TestSetName events raised by the TestFilter contract.
type TestFilterTestSetNameIterator struct {
	Event *TestFilterTestSetName // Event containing the contract specifics and raw log

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
func (it *TestFilterTestSetNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestFilterTestSetName)
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
		it.Event = new(TestFilterTestSetName)
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
func (it *TestFilterTestSetNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestFilterTestSetNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestFilterTestSetName represents a TestSetName event raised by the TestFilter contract.
type TestFilterTestSetName struct {
	OldName   string
	NewName   string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTestSetName is a free log retrieval operation binding the contract event 0x9aa97cf26d841d503a2834c9a4e166b8c17f3cd001238a2bfd891c0138b3a80f.
//
// Solidity: event TestSetName(string oldName, string newName, uint256 timestamp)
func (_TestFilter *TestFilterFilterer) FilterTestSetName(opts *bind.FilterOpts) (*TestFilterTestSetNameIterator, error) {

	logs, sub, err := _TestFilter.contract.FilterLogs(opts, "TestSetName")
	if err != nil {
		return nil, err
	}
	return &TestFilterTestSetNameIterator{contract: _TestFilter.contract, event: "TestSetName", logs: logs, sub: sub}, nil
}

// WatchTestSetName is a free log subscription operation binding the contract event 0x9aa97cf26d841d503a2834c9a4e166b8c17f3cd001238a2bfd891c0138b3a80f.
//
// Solidity: event TestSetName(string oldName, string newName, uint256 timestamp)
func (_TestFilter *TestFilterFilterer) WatchTestSetName(opts *bind.WatchOpts, sink chan<- *TestFilterTestSetName) (event.Subscription, error) {

	logs, sub, err := _TestFilter.contract.WatchLogs(opts, "TestSetName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestFilterTestSetName)
				if err := _TestFilter.contract.UnpackLog(event, "TestSetName", log); err != nil {
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

// ParseTestSetName is a log parse operation binding the contract event 0x9aa97cf26d841d503a2834c9a4e166b8c17f3cd001238a2bfd891c0138b3a80f.
//
// Solidity: event TestSetName(string oldName, string newName, uint256 timestamp)
func (_TestFilter *TestFilterFilterer) ParseTestSetName(log types.Log) (*TestFilterTestSetName, error) {
	event := new(TestFilterTestSetName)
	if err := _TestFilter.contract.UnpackLog(event, "TestSetName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
