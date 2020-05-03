// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Deposit

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DepositABI is the input ABI used to generate the binding from.
const DepositABI = "[{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"cashInRequestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"destination\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"cashOutRequestEvent\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"InRequest\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"fiat_uuid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"fiat_address\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"submited_by\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"payload\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"cashInRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"cashInSubmit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"destination\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"cashOutRequest\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Deposit is an auto generated Go binding around an Ethereum contract.
type Deposit struct {
	DepositCaller     // Read-only binding to the contract
	DepositTransactor // Write-only binding to the contract
	DepositFilterer   // Log filterer for contract events
}

// DepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositSession struct {
	Contract     *Deposit          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositCallerSession struct {
	Contract *DepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositTransactorSession struct {
	Contract     *DepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositRaw struct {
	Contract *Deposit // Generic contract binding to access the raw methods on
}

// DepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositCallerRaw struct {
	Contract *DepositCaller // Generic read-only contract binding to access the raw methods on
}

// DepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositTransactorRaw struct {
	Contract *DepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeposit creates a new instance of Deposit, bound to a specific deployed contract.
func NewDeposit(address common.Address, backend bind.ContractBackend) (*Deposit, error) {
	contract, err := bindDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deposit{DepositCaller: DepositCaller{contract: contract}, DepositTransactor: DepositTransactor{contract: contract}, DepositFilterer: DepositFilterer{contract: contract}}, nil
}

// NewDepositCaller creates a new read-only instance of Deposit, bound to a specific deployed contract.
func NewDepositCaller(address common.Address, caller bind.ContractCaller) (*DepositCaller, error) {
	contract, err := bindDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositCaller{contract: contract}, nil
}

// NewDepositTransactor creates a new write-only instance of Deposit, bound to a specific deployed contract.
func NewDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositTransactor, error) {
	contract, err := bindDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositTransactor{contract: contract}, nil
}

// NewDepositFilterer creates a new log filterer instance of Deposit, bound to a specific deployed contract.
func NewDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositFilterer, error) {
	contract, err := bindDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositFilterer{contract: contract}, nil
}

// bindDeposit binds a generic wrapper to an already deployed contract.
func bindDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deposit *DepositRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Deposit.Contract.DepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deposit *DepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deposit.Contract.DepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deposit *DepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deposit.Contract.DepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deposit *DepositCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Deposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deposit *DepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deposit *DepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deposit.Contract.contract.Transact(opts, method, params...)
}

// InRequest is a free data retrieval call binding the contract method 0x3fb4b5de.
//
// Solidity: function InRequest(string ) view returns(string fiat_uuid, uint256 amount, address user, string fiat_address, address submited_by, bool executed, string payload)
func (_Deposit *DepositCaller) InRequest(opts *bind.CallOpts, arg0 string) (struct {
	FiatUuid    string
	Amount      *big.Int
	User        common.Address
	FiatAddress string
	SubmitedBy  common.Address
	Executed    bool
	Payload     string
}, error) {
	ret := new(struct {
		FiatUuid    string
		Amount      *big.Int
		User        common.Address
		FiatAddress string
		SubmitedBy  common.Address
		Executed    bool
		Payload     string
	})
	out := ret
	err := _Deposit.contract.Call(opts, out, "InRequest", arg0)
	return *ret, err
}

// InRequest is a free data retrieval call binding the contract method 0x3fb4b5de.
//
// Solidity: function InRequest(string ) view returns(string fiat_uuid, uint256 amount, address user, string fiat_address, address submited_by, bool executed, string payload)
func (_Deposit *DepositSession) InRequest(arg0 string) (struct {
	FiatUuid    string
	Amount      *big.Int
	User        common.Address
	FiatAddress string
	SubmitedBy  common.Address
	Executed    bool
	Payload     string
}, error) {
	return _Deposit.Contract.InRequest(&_Deposit.CallOpts, arg0)
}

// InRequest is a free data retrieval call binding the contract method 0x3fb4b5de.
//
// Solidity: function InRequest(string ) view returns(string fiat_uuid, uint256 amount, address user, string fiat_address, address submited_by, bool executed, string payload)
func (_Deposit *DepositCallerSession) InRequest(arg0 string) (struct {
	FiatUuid    string
	Amount      *big.Int
	User        common.Address
	FiatAddress string
	SubmitedBy  common.Address
	Executed    bool
	Payload     string
}, error) {
	return _Deposit.Contract.InRequest(&_Deposit.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Deposit *DepositCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Deposit.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Deposit *DepositSession) IsOwner() (bool, error) {
	return _Deposit.Contract.IsOwner(&_Deposit.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Deposit *DepositCallerSession) IsOwner() (bool, error) {
	return _Deposit.Contract.IsOwner(&_Deposit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Deposit *DepositCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Deposit.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Deposit *DepositSession) Owner() (common.Address, error) {
	return _Deposit.Contract.Owner(&_Deposit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Deposit *DepositCallerSession) Owner() (common.Address, error) {
	return _Deposit.Contract.Owner(&_Deposit.CallOpts)
}

// CashInRequest is a paid mutator transaction binding the contract method 0x9b65f4e3.
//
// Solidity: function cashInRequest(string from, address user, string uuid, uint256 amount) returns()
func (_Deposit *DepositTransactor) CashInRequest(opts *bind.TransactOpts, from string, user common.Address, uuid string, amount *big.Int) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "cashInRequest", from, user, uuid, amount)
}

// CashInRequest is a paid mutator transaction binding the contract method 0x9b65f4e3.
//
// Solidity: function cashInRequest(string from, address user, string uuid, uint256 amount) returns()
func (_Deposit *DepositSession) CashInRequest(from string, user common.Address, uuid string, amount *big.Int) (*types.Transaction, error) {
	return _Deposit.Contract.CashInRequest(&_Deposit.TransactOpts, from, user, uuid, amount)
}

// CashInRequest is a paid mutator transaction binding the contract method 0x9b65f4e3.
//
// Solidity: function cashInRequest(string from, address user, string uuid, uint256 amount) returns()
func (_Deposit *DepositTransactorSession) CashInRequest(from string, user common.Address, uuid string, amount *big.Int) (*types.Transaction, error) {
	return _Deposit.Contract.CashInRequest(&_Deposit.TransactOpts, from, user, uuid, amount)
}

// CashInSubmit is a paid mutator transaction binding the contract method 0x8a3cac3c.
//
// Solidity: function cashInSubmit(string uuid) returns()
func (_Deposit *DepositTransactor) CashInSubmit(opts *bind.TransactOpts, uuid string) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "cashInSubmit", uuid)
}

// CashInSubmit is a paid mutator transaction binding the contract method 0x8a3cac3c.
//
// Solidity: function cashInSubmit(string uuid) returns()
func (_Deposit *DepositSession) CashInSubmit(uuid string) (*types.Transaction, error) {
	return _Deposit.Contract.CashInSubmit(&_Deposit.TransactOpts, uuid)
}

// CashInSubmit is a paid mutator transaction binding the contract method 0x8a3cac3c.
//
// Solidity: function cashInSubmit(string uuid) returns()
func (_Deposit *DepositTransactorSession) CashInSubmit(uuid string) (*types.Transaction, error) {
	return _Deposit.Contract.CashInSubmit(&_Deposit.TransactOpts, uuid)
}

// CashOutRequest is a paid mutator transaction binding the contract method 0xd91f86b2.
//
// Solidity: function cashOutRequest(string destination, address user) payable returns()
func (_Deposit *DepositTransactor) CashOutRequest(opts *bind.TransactOpts, destination string, user common.Address) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "cashOutRequest", destination, user)
}

// CashOutRequest is a paid mutator transaction binding the contract method 0xd91f86b2.
//
// Solidity: function cashOutRequest(string destination, address user) payable returns()
func (_Deposit *DepositSession) CashOutRequest(destination string, user common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.CashOutRequest(&_Deposit.TransactOpts, destination, user)
}

// CashOutRequest is a paid mutator transaction binding the contract method 0xd91f86b2.
//
// Solidity: function cashOutRequest(string destination, address user) payable returns()
func (_Deposit *DepositTransactorSession) CashOutRequest(destination string, user common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.CashOutRequest(&_Deposit.TransactOpts, destination, user)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Deposit *DepositTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Deposit *DepositSession) RenounceOwnership() (*types.Transaction, error) {
	return _Deposit.Contract.RenounceOwnership(&_Deposit.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Deposit *DepositTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Deposit.Contract.RenounceOwnership(&_Deposit.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Deposit *DepositTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Deposit *DepositSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.TransferOwnership(&_Deposit.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Deposit *DepositTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.TransferOwnership(&_Deposit.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Deposit *DepositTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Deposit.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Deposit *DepositSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Deposit.Contract.Fallback(&_Deposit.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Deposit *DepositTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Deposit.Contract.Fallback(&_Deposit.TransactOpts, calldata)
}

// DepositOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Deposit contract.
type DepositOwnershipTransferredIterator struct {
	Event *DepositOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DepositOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositOwnershipTransferred)
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
		it.Event = new(DepositOwnershipTransferred)
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
func (it *DepositOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositOwnershipTransferred represents a OwnershipTransferred event raised by the Deposit contract.
type DepositOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Deposit *DepositFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DepositOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Deposit.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DepositOwnershipTransferredIterator{contract: _Deposit.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Deposit *DepositFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DepositOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Deposit.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositOwnershipTransferred)
				if err := _Deposit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Deposit *DepositFilterer) ParseOwnershipTransferred(log types.Log) (*DepositOwnershipTransferred, error) {
	event := new(DepositOwnershipTransferred)
	if err := _Deposit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DepositCashInRequestEventIterator is returned from FilterCashInRequestEvent and is used to iterate over the raw logs and unpacked data for CashInRequestEvent events raised by the Deposit contract.
type DepositCashInRequestEventIterator struct {
	Event *DepositCashInRequestEvent // Event containing the contract specifics and raw log

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
func (it *DepositCashInRequestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositCashInRequestEvent)
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
		it.Event = new(DepositCashInRequestEvent)
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
func (it *DepositCashInRequestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositCashInRequestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositCashInRequestEvent represents a CashInRequestEvent event raised by the Deposit contract.
type DepositCashInRequestEvent struct {
	From   common.Hash
	User   common.Address
	Amount *big.Int
	Uuid   common.Hash
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCashInRequestEvent is a free log retrieval operation binding the contract event 0xc7257fe765b6dd4946dfa7e3829b142c46f4566767caec90f8396dcaf7e625c1.
//
// Solidity: event cashInRequestEvent(string indexed from, address indexed user, uint256 amount, string indexed uuid)
func (_Deposit *DepositFilterer) FilterCashInRequestEvent(opts *bind.FilterOpts, from []string, user []common.Address, uuid []string) (*DepositCashInRequestEventIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var uuidRule []interface{}
	for _, uuidItem := range uuid {
		uuidRule = append(uuidRule, uuidItem)
	}

	logs, sub, err := _Deposit.contract.FilterLogs(opts, "cashInRequestEvent", fromRule, userRule, uuidRule)
	if err != nil {
		return nil, err
	}
	return &DepositCashInRequestEventIterator{contract: _Deposit.contract, event: "cashInRequestEvent", logs: logs, sub: sub}, nil
}

// WatchCashInRequestEvent is a free log subscription operation binding the contract event 0xc7257fe765b6dd4946dfa7e3829b142c46f4566767caec90f8396dcaf7e625c1.
//
// Solidity: event cashInRequestEvent(string indexed from, address indexed user, uint256 amount, string indexed uuid)
func (_Deposit *DepositFilterer) WatchCashInRequestEvent(opts *bind.WatchOpts, sink chan<- *DepositCashInRequestEvent, from []string, user []common.Address, uuid []string) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var uuidRule []interface{}
	for _, uuidItem := range uuid {
		uuidRule = append(uuidRule, uuidItem)
	}

	logs, sub, err := _Deposit.contract.WatchLogs(opts, "cashInRequestEvent", fromRule, userRule, uuidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositCashInRequestEvent)
				if err := _Deposit.contract.UnpackLog(event, "cashInRequestEvent", log); err != nil {
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

// ParseCashInRequestEvent is a log parse operation binding the contract event 0xc7257fe765b6dd4946dfa7e3829b142c46f4566767caec90f8396dcaf7e625c1.
//
// Solidity: event cashInRequestEvent(string indexed from, address indexed user, uint256 amount, string indexed uuid)
func (_Deposit *DepositFilterer) ParseCashInRequestEvent(log types.Log) (*DepositCashInRequestEvent, error) {
	event := new(DepositCashInRequestEvent)
	if err := _Deposit.contract.UnpackLog(event, "cashInRequestEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DepositCashOutRequestEventIterator is returned from FilterCashOutRequestEvent and is used to iterate over the raw logs and unpacked data for CashOutRequestEvent events raised by the Deposit contract.
type DepositCashOutRequestEventIterator struct {
	Event *DepositCashOutRequestEvent // Event containing the contract specifics and raw log

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
func (it *DepositCashOutRequestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositCashOutRequestEvent)
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
		it.Event = new(DepositCashOutRequestEvent)
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
func (it *DepositCashOutRequestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositCashOutRequestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositCashOutRequestEvent represents a CashOutRequestEvent event raised by the Deposit contract.
type DepositCashOutRequestEvent struct {
	Destination common.Hash
	User        common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCashOutRequestEvent is a free log retrieval operation binding the contract event 0xe3cb954797a95fbb1caebb51fad66117ac891d940890eb9c2d8c4e3051761288.
//
// Solidity: event cashOutRequestEvent(string indexed destination, address indexed user, uint256 amount)
func (_Deposit *DepositFilterer) FilterCashOutRequestEvent(opts *bind.FilterOpts, destination []string, user []common.Address) (*DepositCashOutRequestEventIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Deposit.contract.FilterLogs(opts, "cashOutRequestEvent", destinationRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DepositCashOutRequestEventIterator{contract: _Deposit.contract, event: "cashOutRequestEvent", logs: logs, sub: sub}, nil
}

// WatchCashOutRequestEvent is a free log subscription operation binding the contract event 0xe3cb954797a95fbb1caebb51fad66117ac891d940890eb9c2d8c4e3051761288.
//
// Solidity: event cashOutRequestEvent(string indexed destination, address indexed user, uint256 amount)
func (_Deposit *DepositFilterer) WatchCashOutRequestEvent(opts *bind.WatchOpts, sink chan<- *DepositCashOutRequestEvent, destination []string, user []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Deposit.contract.WatchLogs(opts, "cashOutRequestEvent", destinationRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositCashOutRequestEvent)
				if err := _Deposit.contract.UnpackLog(event, "cashOutRequestEvent", log); err != nil {
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

// ParseCashOutRequestEvent is a log parse operation binding the contract event 0xe3cb954797a95fbb1caebb51fad66117ac891d940890eb9c2d8c4e3051761288.
//
// Solidity: event cashOutRequestEvent(string indexed destination, address indexed user, uint256 amount)
func (_Deposit *DepositFilterer) ParseCashOutRequestEvent(log types.Log) (*DepositCashOutRequestEvent, error) {
	event := new(DepositCashOutRequestEvent)
	if err := _Deposit.contract.UnpackLog(event, "cashOutRequestEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}