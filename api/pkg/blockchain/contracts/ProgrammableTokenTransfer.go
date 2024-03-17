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

// ClientAny2EVMMessage is an auto generated low-level Go binding around an user-defined struct.
type ClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
}

// ClientEVMTokenAmount is an auto generated low-level Go binding around an user-defined struct.
type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

// ProgrammableTokenTransferTransferData is an auto generated low-level Go binding around an user-defined struct.
type ProgrammableTokenTransferTransferData struct {
	Payer       common.Address
	Beneficiary common.Address
}

// ProgrammableTokenTransferMetaData contains all meta data concerning the ProgrammableTokenTransfer contract.
var ProgrammableTokenTransferMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destinationChainSelector\",\"type\":\"uint64\"}],\"name\":\"DestinationChainNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FailedToWithdrawEth\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBeneficiaryAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPayerAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReceiverAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"InvalidRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTransferAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calculatedFees\",\"type\":\"uint256\"}],\"name\":\"NotEnoughBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NothingToWithdraw\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"SourceChainNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structProgrammableTokenTransfer.TransferData\",\"name\":\"data\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destinationChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structProgrammableTokenTransfer.TransferData\",\"name\":\"data\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_destinationChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"allowlistDestinationChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"allowlistSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"allowlistSourceChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"allowlistedDestinationChains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowlistedSenders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"allowlistedSourceChains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_destinationChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"internalType\":\"structProgrammableTokenTransfer.TransferData\",\"name\":\"_data\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferToken\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ProgrammableTokenTransferABI is the input ABI used to generate the binding from.
// Deprecated: Use ProgrammableTokenTransferMetaData.ABI instead.
var ProgrammableTokenTransferABI = ProgrammableTokenTransferMetaData.ABI

// ProgrammableTokenTransfer is an auto generated Go binding around an Ethereum contract.
type ProgrammableTokenTransfer struct {
	ProgrammableTokenTransferCaller     // Read-only binding to the contract
	ProgrammableTokenTransferTransactor // Write-only binding to the contract
	ProgrammableTokenTransferFilterer   // Log filterer for contract events
}

// ProgrammableTokenTransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProgrammableTokenTransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProgrammableTokenTransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProgrammableTokenTransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProgrammableTokenTransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProgrammableTokenTransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProgrammableTokenTransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProgrammableTokenTransferSession struct {
	Contract     *ProgrammableTokenTransfer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ProgrammableTokenTransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProgrammableTokenTransferCallerSession struct {
	Contract *ProgrammableTokenTransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ProgrammableTokenTransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProgrammableTokenTransferTransactorSession struct {
	Contract     *ProgrammableTokenTransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ProgrammableTokenTransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProgrammableTokenTransferRaw struct {
	Contract *ProgrammableTokenTransfer // Generic contract binding to access the raw methods on
}

// ProgrammableTokenTransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProgrammableTokenTransferCallerRaw struct {
	Contract *ProgrammableTokenTransferCaller // Generic read-only contract binding to access the raw methods on
}

// ProgrammableTokenTransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProgrammableTokenTransferTransactorRaw struct {
	Contract *ProgrammableTokenTransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProgrammableTokenTransfer creates a new instance of ProgrammableTokenTransfer, bound to a specific deployed contract.
func NewProgrammableTokenTransfer(address common.Address, backend bind.ContractBackend) (*ProgrammableTokenTransfer, error) {
	contract, err := bindProgrammableTokenTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProgrammableTokenTransfer{ProgrammableTokenTransferCaller: ProgrammableTokenTransferCaller{contract: contract}, ProgrammableTokenTransferTransactor: ProgrammableTokenTransferTransactor{contract: contract}, ProgrammableTokenTransferFilterer: ProgrammableTokenTransferFilterer{contract: contract}}, nil
}

// NewProgrammableTokenTransferCaller creates a new read-only instance of ProgrammableTokenTransfer, bound to a specific deployed contract.
func NewProgrammableTokenTransferCaller(address common.Address, caller bind.ContractCaller) (*ProgrammableTokenTransferCaller, error) {
	contract, err := bindProgrammableTokenTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProgrammableTokenTransferCaller{contract: contract}, nil
}

// NewProgrammableTokenTransferTransactor creates a new write-only instance of ProgrammableTokenTransfer, bound to a specific deployed contract.
func NewProgrammableTokenTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*ProgrammableTokenTransferTransactor, error) {
	contract, err := bindProgrammableTokenTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProgrammableTokenTransferTransactor{contract: contract}, nil
}

// NewProgrammableTokenTransferFilterer creates a new log filterer instance of ProgrammableTokenTransfer, bound to a specific deployed contract.
func NewProgrammableTokenTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*ProgrammableTokenTransferFilterer, error) {
	contract, err := bindProgrammableTokenTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProgrammableTokenTransferFilterer{contract: contract}, nil
}

// bindProgrammableTokenTransfer binds a generic wrapper to an already deployed contract.
func bindProgrammableTokenTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProgrammableTokenTransferMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProgrammableTokenTransfer.Contract.ProgrammableTokenTransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.ProgrammableTokenTransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.ProgrammableTokenTransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProgrammableTokenTransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.contract.Transact(opts, method, params...)
}

// AllowlistedDestinationChains is a free data retrieval call binding the contract method 0x75c67c66.
//
// Solidity: function allowlistedDestinationChains(uint64 ) view returns(bool)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferCaller) AllowlistedDestinationChains(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _ProgrammableTokenTransfer.contract.Call(opts, &out, "allowlistedDestinationChains", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowlistedDestinationChains is a free data retrieval call binding the contract method 0x75c67c66.
//
// Solidity: function allowlistedDestinationChains(uint64 ) view returns(bool)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferSession) AllowlistedDestinationChains(arg0 uint64) (bool, error) {
	return _ProgrammableTokenTransfer.Contract.AllowlistedDestinationChains(&_ProgrammableTokenTransfer.CallOpts, arg0)
}

// AllowlistedDestinationChains is a free data retrieval call binding the contract method 0x75c67c66.
//
// Solidity: function allowlistedDestinationChains(uint64 ) view returns(bool)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferCallerSession) AllowlistedDestinationChains(arg0 uint64) (bool, error) {
	return _ProgrammableTokenTransfer.Contract.AllowlistedDestinationChains(&_ProgrammableTokenTransfer.CallOpts, arg0)
}

// AllowlistedSenders is a free data retrieval call binding the contract method 0x6159ada1.
//
// Solidity: function allowlistedSenders(address ) view returns(bool)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferCaller) AllowlistedSenders(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ProgrammableTokenTransfer.contract.Call(opts, &out, "allowlistedSenders", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowlistedSenders is a free data retrieval call binding the contract method 0x6159ada1.
//
// Solidity: function allowlistedSenders(address ) view returns(bool)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferSession) AllowlistedSenders(arg0 common.Address) (bool, error) {
	return _ProgrammableTokenTransfer.Contract.AllowlistedSenders(&_ProgrammableTokenTransfer.CallOpts, arg0)
}

// AllowlistedSenders is a free data retrieval call binding the contract method 0x6159ada1.
//
// Solidity: function allowlistedSenders(address ) view returns(bool)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferCallerSession) AllowlistedSenders(arg0 common.Address) (bool, error) {
	return _ProgrammableTokenTransfer.Contract.AllowlistedSenders(&_ProgrammableTokenTransfer.CallOpts, arg0)
}

// AllowlistedSourceChains is a free data retrieval call binding the contract method 0x4030d521.
//
// Solidity: function allowlistedSourceChains(uint64 ) view returns(bool)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferCaller) AllowlistedSourceChains(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _ProgrammableTokenTransfer.contract.Call(opts, &out, "allowlistedSourceChains", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowlistedSourceChains is a free data retrieval call binding the contract method 0x4030d521.
//
// Solidity: function allowlistedSourceChains(uint64 ) view returns(bool)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferSession) AllowlistedSourceChains(arg0 uint64) (bool, error) {
	return _ProgrammableTokenTransfer.Contract.AllowlistedSourceChains(&_ProgrammableTokenTransfer.CallOpts, arg0)
}

// AllowlistedSourceChains is a free data retrieval call binding the contract method 0x4030d521.
//
// Solidity: function allowlistedSourceChains(uint64 ) view returns(bool)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferCallerSession) AllowlistedSourceChains(arg0 uint64) (bool, error) {
	return _ProgrammableTokenTransfer.Contract.AllowlistedSourceChains(&_ProgrammableTokenTransfer.CallOpts, arg0)
}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProgrammableTokenTransfer.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferSession) GetRouter() (common.Address, error) {
	return _ProgrammableTokenTransfer.Contract.GetRouter(&_ProgrammableTokenTransfer.CallOpts)
}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferCallerSession) GetRouter() (common.Address, error) {
	return _ProgrammableTokenTransfer.Contract.GetRouter(&_ProgrammableTokenTransfer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProgrammableTokenTransfer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferSession) Owner() (common.Address, error) {
	return _ProgrammableTokenTransfer.Contract.Owner(&_ProgrammableTokenTransfer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferCallerSession) Owner() (common.Address, error) {
	return _ProgrammableTokenTransfer.Contract.Owner(&_ProgrammableTokenTransfer.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ProgrammableTokenTransfer.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ProgrammableTokenTransfer.Contract.SupportsInterface(&_ProgrammableTokenTransfer.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ProgrammableTokenTransfer.Contract.SupportsInterface(&_ProgrammableTokenTransfer.CallOpts, interfaceId)
}

// AllowlistDestinationChain is a paid mutator transaction binding the contract method 0x96d3b83d.
//
// Solidity: function allowlistDestinationChain(uint64 _destinationChainSelector, bool allowed) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactor) AllowlistDestinationChain(opts *bind.TransactOpts, _destinationChainSelector uint64, allowed bool) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.contract.Transact(opts, "allowlistDestinationChain", _destinationChainSelector, allowed)
}

// AllowlistDestinationChain is a paid mutator transaction binding the contract method 0x96d3b83d.
//
// Solidity: function allowlistDestinationChain(uint64 _destinationChainSelector, bool allowed) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferSession) AllowlistDestinationChain(_destinationChainSelector uint64, allowed bool) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.AllowlistDestinationChain(&_ProgrammableTokenTransfer.TransactOpts, _destinationChainSelector, allowed)
}

// AllowlistDestinationChain is a paid mutator transaction binding the contract method 0x96d3b83d.
//
// Solidity: function allowlistDestinationChain(uint64 _destinationChainSelector, bool allowed) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactorSession) AllowlistDestinationChain(_destinationChainSelector uint64, allowed bool) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.AllowlistDestinationChain(&_ProgrammableTokenTransfer.TransactOpts, _destinationChainSelector, allowed)
}

// AllowlistSender is a paid mutator transaction binding the contract method 0xeab5b02c.
//
// Solidity: function allowlistSender(address _sender, bool allowed) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactor) AllowlistSender(opts *bind.TransactOpts, _sender common.Address, allowed bool) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.contract.Transact(opts, "allowlistSender", _sender, allowed)
}

// AllowlistSender is a paid mutator transaction binding the contract method 0xeab5b02c.
//
// Solidity: function allowlistSender(address _sender, bool allowed) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferSession) AllowlistSender(_sender common.Address, allowed bool) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.AllowlistSender(&_ProgrammableTokenTransfer.TransactOpts, _sender, allowed)
}

// AllowlistSender is a paid mutator transaction binding the contract method 0xeab5b02c.
//
// Solidity: function allowlistSender(address _sender, bool allowed) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactorSession) AllowlistSender(_sender common.Address, allowed bool) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.AllowlistSender(&_ProgrammableTokenTransfer.TransactOpts, _sender, allowed)
}

// AllowlistSourceChain is a paid mutator transaction binding the contract method 0xdb04fa49.
//
// Solidity: function allowlistSourceChain(uint64 _sourceChainSelector, bool allowed) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactor) AllowlistSourceChain(opts *bind.TransactOpts, _sourceChainSelector uint64, allowed bool) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.contract.Transact(opts, "allowlistSourceChain", _sourceChainSelector, allowed)
}

// AllowlistSourceChain is a paid mutator transaction binding the contract method 0xdb04fa49.
//
// Solidity: function allowlistSourceChain(uint64 _sourceChainSelector, bool allowed) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferSession) AllowlistSourceChain(_sourceChainSelector uint64, allowed bool) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.AllowlistSourceChain(&_ProgrammableTokenTransfer.TransactOpts, _sourceChainSelector, allowed)
}

// AllowlistSourceChain is a paid mutator transaction binding the contract method 0xdb04fa49.
//
// Solidity: function allowlistSourceChain(uint64 _sourceChainSelector, bool allowed) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactorSession) AllowlistSourceChain(_sourceChainSelector uint64, allowed bool) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.AllowlistSourceChain(&_ProgrammableTokenTransfer.TransactOpts, _sourceChainSelector, allowed)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.contract.Transact(opts, "ccipReceive", message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.CcipReceive(&_ProgrammableTokenTransfer.TransactOpts, message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.CcipReceive(&_ProgrammableTokenTransfer.TransactOpts, message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.RenounceOwnership(&_ProgrammableTokenTransfer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.RenounceOwnership(&_ProgrammableTokenTransfer.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.TransferOwnership(&_ProgrammableTokenTransfer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.TransferOwnership(&_ProgrammableTokenTransfer.TransactOpts, newOwner)
}

// TransferToken is a paid mutator transaction binding the contract method 0xb892eaff.
//
// Solidity: function transferToken(uint64 _destinationChainSelector, address _receiver, (address,address) _data, address _token, uint256 _amount) returns(bytes32 messageId)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactor) TransferToken(opts *bind.TransactOpts, _destinationChainSelector uint64, _receiver common.Address, _data ProgrammableTokenTransferTransferData, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.contract.Transact(opts, "transferToken", _destinationChainSelector, _receiver, _data, _token, _amount)
}

// TransferToken is a paid mutator transaction binding the contract method 0xb892eaff.
//
// Solidity: function transferToken(uint64 _destinationChainSelector, address _receiver, (address,address) _data, address _token, uint256 _amount) returns(bytes32 messageId)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferSession) TransferToken(_destinationChainSelector uint64, _receiver common.Address, _data ProgrammableTokenTransferTransferData, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.TransferToken(&_ProgrammableTokenTransfer.TransactOpts, _destinationChainSelector, _receiver, _data, _token, _amount)
}

// TransferToken is a paid mutator transaction binding the contract method 0xb892eaff.
//
// Solidity: function transferToken(uint64 _destinationChainSelector, address _receiver, (address,address) _data, address _token, uint256 _amount) returns(bytes32 messageId)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactorSession) TransferToken(_destinationChainSelector uint64, _receiver common.Address, _data ProgrammableTokenTransferTransferData, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.TransferToken(&_ProgrammableTokenTransfer.TransactOpts, _destinationChainSelector, _receiver, _data, _token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _beneficiary) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactor) Withdraw(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.contract.Transact(opts, "withdraw", _beneficiary)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _beneficiary) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferSession) Withdraw(_beneficiary common.Address) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.Withdraw(&_ProgrammableTokenTransfer.TransactOpts, _beneficiary)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _beneficiary) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactorSession) Withdraw(_beneficiary common.Address) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.Withdraw(&_ProgrammableTokenTransfer.TransactOpts, _beneficiary)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3aeac4e1.
//
// Solidity: function withdrawToken(address _beneficiary, address _token) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactor) WithdrawToken(opts *bind.TransactOpts, _beneficiary common.Address, _token common.Address) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.contract.Transact(opts, "withdrawToken", _beneficiary, _token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3aeac4e1.
//
// Solidity: function withdrawToken(address _beneficiary, address _token) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferSession) WithdrawToken(_beneficiary common.Address, _token common.Address) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.WithdrawToken(&_ProgrammableTokenTransfer.TransactOpts, _beneficiary, _token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3aeac4e1.
//
// Solidity: function withdrawToken(address _beneficiary, address _token) returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactorSession) WithdrawToken(_beneficiary common.Address, _token common.Address) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.WithdrawToken(&_ProgrammableTokenTransfer.TransactOpts, _beneficiary, _token)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferSession) Receive() (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.Receive(&_ProgrammableTokenTransfer.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferTransactorSession) Receive() (*types.Transaction, error) {
	return _ProgrammableTokenTransfer.Contract.Receive(&_ProgrammableTokenTransfer.TransactOpts)
}

// ProgrammableTokenTransferMessageReceivedIterator is returned from FilterMessageReceived and is used to iterate over the raw logs and unpacked data for MessageReceived events raised by the ProgrammableTokenTransfer contract.
type ProgrammableTokenTransferMessageReceivedIterator struct {
	Event *ProgrammableTokenTransferMessageReceived // Event containing the contract specifics and raw log

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
func (it *ProgrammableTokenTransferMessageReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProgrammableTokenTransferMessageReceived)
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
		it.Event = new(ProgrammableTokenTransferMessageReceived)
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
func (it *ProgrammableTokenTransferMessageReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProgrammableTokenTransferMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProgrammableTokenTransferMessageReceived represents a MessageReceived event raised by the ProgrammableTokenTransfer contract.
type ProgrammableTokenTransferMessageReceived struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              common.Address
	Data                ProgrammableTokenTransferTransferData
	Token               common.Address
	TokenAmount         *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMessageReceived is a free log retrieval operation binding the contract event 0x3512ebab48659a32bc61ed8699c211efcc8688c7a778dcb6fb9968e325e0445e.
//
// Solidity: event MessageReceived(bytes32 indexed messageId, uint64 indexed sourceChainSelector, address sender, (address,address) data, address token, uint256 tokenAmount)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferFilterer) FilterMessageReceived(opts *bind.FilterOpts, messageId [][32]byte, sourceChainSelector []uint64) (*ProgrammableTokenTransferMessageReceivedIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}
	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _ProgrammableTokenTransfer.contract.FilterLogs(opts, "MessageReceived", messageIdRule, sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &ProgrammableTokenTransferMessageReceivedIterator{contract: _ProgrammableTokenTransfer.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

// WatchMessageReceived is a free log subscription operation binding the contract event 0x3512ebab48659a32bc61ed8699c211efcc8688c7a778dcb6fb9968e325e0445e.
//
// Solidity: event MessageReceived(bytes32 indexed messageId, uint64 indexed sourceChainSelector, address sender, (address,address) data, address token, uint256 tokenAmount)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *ProgrammableTokenTransferMessageReceived, messageId [][32]byte, sourceChainSelector []uint64) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}
	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _ProgrammableTokenTransfer.contract.WatchLogs(opts, "MessageReceived", messageIdRule, sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProgrammableTokenTransferMessageReceived)
				if err := _ProgrammableTokenTransfer.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

// ParseMessageReceived is a log parse operation binding the contract event 0x3512ebab48659a32bc61ed8699c211efcc8688c7a778dcb6fb9968e325e0445e.
//
// Solidity: event MessageReceived(bytes32 indexed messageId, uint64 indexed sourceChainSelector, address sender, (address,address) data, address token, uint256 tokenAmount)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferFilterer) ParseMessageReceived(log types.Log) (*ProgrammableTokenTransferMessageReceived, error) {
	event := new(ProgrammableTokenTransferMessageReceived)
	if err := _ProgrammableTokenTransfer.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProgrammableTokenTransferMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the ProgrammableTokenTransfer contract.
type ProgrammableTokenTransferMessageSentIterator struct {
	Event *ProgrammableTokenTransferMessageSent // Event containing the contract specifics and raw log

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
func (it *ProgrammableTokenTransferMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProgrammableTokenTransferMessageSent)
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
		it.Event = new(ProgrammableTokenTransferMessageSent)
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
func (it *ProgrammableTokenTransferMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProgrammableTokenTransferMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProgrammableTokenTransferMessageSent represents a MessageSent event raised by the ProgrammableTokenTransfer contract.
type ProgrammableTokenTransferMessageSent struct {
	MessageId                [32]byte
	DestinationChainSelector uint64
	Receiver                 common.Address
	Data                     ProgrammableTokenTransferTransferData
	Token                    common.Address
	TokenAmount              *big.Int
	Fees                     *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0x41b3f2aa0d49d1cad6d633fd1e8852ea8dd967316e53acbe70b7e0a633aee85c.
//
// Solidity: event MessageSent(bytes32 indexed messageId, uint64 indexed destinationChainSelector, address receiver, (address,address) data, address token, uint256 tokenAmount, uint256 fees)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferFilterer) FilterMessageSent(opts *bind.FilterOpts, messageId [][32]byte, destinationChainSelector []uint64) (*ProgrammableTokenTransferMessageSentIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}
	var destinationChainSelectorRule []interface{}
	for _, destinationChainSelectorItem := range destinationChainSelector {
		destinationChainSelectorRule = append(destinationChainSelectorRule, destinationChainSelectorItem)
	}

	logs, sub, err := _ProgrammableTokenTransfer.contract.FilterLogs(opts, "MessageSent", messageIdRule, destinationChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &ProgrammableTokenTransferMessageSentIterator{contract: _ProgrammableTokenTransfer.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0x41b3f2aa0d49d1cad6d633fd1e8852ea8dd967316e53acbe70b7e0a633aee85c.
//
// Solidity: event MessageSent(bytes32 indexed messageId, uint64 indexed destinationChainSelector, address receiver, (address,address) data, address token, uint256 tokenAmount, uint256 fees)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *ProgrammableTokenTransferMessageSent, messageId [][32]byte, destinationChainSelector []uint64) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}
	var destinationChainSelectorRule []interface{}
	for _, destinationChainSelectorItem := range destinationChainSelector {
		destinationChainSelectorRule = append(destinationChainSelectorRule, destinationChainSelectorItem)
	}

	logs, sub, err := _ProgrammableTokenTransfer.contract.WatchLogs(opts, "MessageSent", messageIdRule, destinationChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProgrammableTokenTransferMessageSent)
				if err := _ProgrammableTokenTransfer.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0x41b3f2aa0d49d1cad6d633fd1e8852ea8dd967316e53acbe70b7e0a633aee85c.
//
// Solidity: event MessageSent(bytes32 indexed messageId, uint64 indexed destinationChainSelector, address receiver, (address,address) data, address token, uint256 tokenAmount, uint256 fees)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferFilterer) ParseMessageSent(log types.Log) (*ProgrammableTokenTransferMessageSent, error) {
	event := new(ProgrammableTokenTransferMessageSent)
	if err := _ProgrammableTokenTransfer.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProgrammableTokenTransferOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProgrammableTokenTransfer contract.
type ProgrammableTokenTransferOwnershipTransferredIterator struct {
	Event *ProgrammableTokenTransferOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProgrammableTokenTransferOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProgrammableTokenTransferOwnershipTransferred)
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
		it.Event = new(ProgrammableTokenTransferOwnershipTransferred)
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
func (it *ProgrammableTokenTransferOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProgrammableTokenTransferOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProgrammableTokenTransferOwnershipTransferred represents a OwnershipTransferred event raised by the ProgrammableTokenTransfer contract.
type ProgrammableTokenTransferOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProgrammableTokenTransferOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProgrammableTokenTransfer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProgrammableTokenTransferOwnershipTransferredIterator{contract: _ProgrammableTokenTransfer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProgrammableTokenTransferOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProgrammableTokenTransfer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProgrammableTokenTransferOwnershipTransferred)
				if err := _ProgrammableTokenTransfer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ProgrammableTokenTransfer *ProgrammableTokenTransferFilterer) ParseOwnershipTransferred(log types.Log) (*ProgrammableTokenTransferOwnershipTransferred, error) {
	event := new(ProgrammableTokenTransferOwnershipTransferred)
	if err := _ProgrammableTokenTransfer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
