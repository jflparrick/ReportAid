// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package organisationRecipientBudgets

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

// BudgetsBudget is an auto generated low-level Go binding around an user-defined struct.
type BudgetsBudget struct {
	BudgetType uint8
	BudgetLine [32]byte
	OtherRef   [32]byte
	Finance    BudgetsFinance
}

// BudgetsFinance is an auto generated low-level Go binding around an user-defined struct.
type BudgetsFinance struct {
	Status uint8
	Value  *big.Int
	Start  [32]byte
	End    [32]byte
}

// BudgetsABI is the input ABI used to generate the binding from.
const BudgetsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"budgetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"budgetLine\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"otherRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"start\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"end\",\"type\":\"bytes32\"}],\"internalType\":\"structBudgets.Finance\",\"name\":\"finance\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structBudgets.Budget\",\"name\":\"_budget\",\"type\":\"tuple\"}],\"name\":\"SetBudget\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getBudget\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"budgetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"budgetLine\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"otherRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"start\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"end\",\"type\":\"bytes32\"}],\"internalType\":\"structBudgets.Finance\",\"name\":\"finance\",\"type\":\"tuple\"}],\"internalType\":\"structBudgets.Budget\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getBudgetEnd\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getBudgetLine\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getBudgetReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getBudgetStart\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getBudgetStatus\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getBudgetType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getBudgetValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"}],\"name\":\"getNumBudgets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getOtherRef\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"budgetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"budgetLine\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"otherRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"start\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"end\",\"type\":\"bytes32\"}],\"internalType\":\"structBudgets.Finance\",\"name\":\"finance\",\"type\":\"tuple\"}],\"internalType\":\"structBudgets.Budget\",\"name\":\"_budget\",\"type\":\"tuple\"}],\"name\":\"setBudget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BudgetsFuncSigs maps the 4-byte function signature to its string representation.
var BudgetsFuncSigs = map[string]string{
	"66cf0ca9": "getBudget(uint8,bytes32,bytes32,bytes32)",
	"193a4c6c": "getBudgetEnd(uint8,bytes32,bytes32,bytes32)",
	"54755d3f": "getBudgetLine(uint8,bytes32,bytes32,bytes32)",
	"6be5d013": "getBudgetReference(uint8,bytes32,bytes32,uint256)",
	"4b917df5": "getBudgetStart(uint8,bytes32,bytes32,bytes32)",
	"9541246c": "getBudgetStatus(uint8,bytes32,bytes32,bytes32)",
	"98530d3b": "getBudgetType(uint8,bytes32,bytes32,bytes32)",
	"c3bc3482": "getBudgetValue(uint8,bytes32,bytes32,bytes32)",
	"366b6701": "getNumBudgets(uint8,bytes32,bytes32)",
	"736d5735": "getOtherRef(uint8,bytes32,bytes32,bytes32)",
	"f1e52779": "setBudget(uint8,bytes32,bytes32,bytes32,(uint8,bytes32,bytes32,(uint8,uint256,bytes32,bytes32)))",
}

// Budgets is an auto generated Go binding around an Ethereum contract.
type Budgets struct {
	BudgetsCaller     // Read-only binding to the contract
	BudgetsTransactor // Write-only binding to the contract
	BudgetsFilterer   // Log filterer for contract events
}

// BudgetsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BudgetsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BudgetsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BudgetsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BudgetsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BudgetsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BudgetsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BudgetsSession struct {
	Contract     *Budgets          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BudgetsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BudgetsCallerSession struct {
	Contract *BudgetsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BudgetsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BudgetsTransactorSession struct {
	Contract     *BudgetsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BudgetsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BudgetsRaw struct {
	Contract *Budgets // Generic contract binding to access the raw methods on
}

// BudgetsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BudgetsCallerRaw struct {
	Contract *BudgetsCaller // Generic read-only contract binding to access the raw methods on
}

// BudgetsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BudgetsTransactorRaw struct {
	Contract *BudgetsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBudgets creates a new instance of Budgets, bound to a specific deployed contract.
func NewBudgets(address common.Address, backend bind.ContractBackend) (*Budgets, error) {
	contract, err := bindBudgets(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Budgets{BudgetsCaller: BudgetsCaller{contract: contract}, BudgetsTransactor: BudgetsTransactor{contract: contract}, BudgetsFilterer: BudgetsFilterer{contract: contract}}, nil
}

// NewBudgetsCaller creates a new read-only instance of Budgets, bound to a specific deployed contract.
func NewBudgetsCaller(address common.Address, caller bind.ContractCaller) (*BudgetsCaller, error) {
	contract, err := bindBudgets(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BudgetsCaller{contract: contract}, nil
}

// NewBudgetsTransactor creates a new write-only instance of Budgets, bound to a specific deployed contract.
func NewBudgetsTransactor(address common.Address, transactor bind.ContractTransactor) (*BudgetsTransactor, error) {
	contract, err := bindBudgets(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BudgetsTransactor{contract: contract}, nil
}

// NewBudgetsFilterer creates a new log filterer instance of Budgets, bound to a specific deployed contract.
func NewBudgetsFilterer(address common.Address, filterer bind.ContractFilterer) (*BudgetsFilterer, error) {
	contract, err := bindBudgets(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BudgetsFilterer{contract: contract}, nil
}

// bindBudgets binds a generic wrapper to an already deployed contract.
func bindBudgets(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BudgetsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Budgets *BudgetsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Budgets.Contract.BudgetsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Budgets *BudgetsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Budgets.Contract.BudgetsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Budgets *BudgetsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Budgets.Contract.BudgetsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Budgets *BudgetsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Budgets.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Budgets *BudgetsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Budgets.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Budgets *BudgetsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Budgets.Contract.contract.Transact(opts, method, params...)
}

// GetBudget is a free data retrieval call binding the contract method 0x66cf0ca9.
//
// Solidity: function getBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(BudgetsBudget)
func (_Budgets *BudgetsCaller) GetBudget(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (BudgetsBudget, error) {
	var (
		ret0 = new(BudgetsBudget)
	)
	out := ret0
	err := _Budgets.contract.Call(opts, out, "getBudget", _owner, _firstRef, _secondRef, _budgetRef)
	return *ret0, err
}

// GetBudget is a free data retrieval call binding the contract method 0x66cf0ca9.
//
// Solidity: function getBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(BudgetsBudget)
func (_Budgets *BudgetsSession) GetBudget(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (BudgetsBudget, error) {
	return _Budgets.Contract.GetBudget(&_Budgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudget is a free data retrieval call binding the contract method 0x66cf0ca9.
//
// Solidity: function getBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(BudgetsBudget)
func (_Budgets *BudgetsCallerSession) GetBudget(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (BudgetsBudget, error) {
	return _Budgets.Contract.GetBudget(&_Budgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetEnd is a free data retrieval call binding the contract method 0x193a4c6c.
//
// Solidity: function getBudgetEnd(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_Budgets *BudgetsCaller) GetBudgetEnd(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Budgets.contract.Call(opts, out, "getBudgetEnd", _owner, _firstRef, _secondRef, _budgetRef)
	return *ret0, err
}

// GetBudgetEnd is a free data retrieval call binding the contract method 0x193a4c6c.
//
// Solidity: function getBudgetEnd(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_Budgets *BudgetsSession) GetBudgetEnd(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _Budgets.Contract.GetBudgetEnd(&_Budgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetEnd is a free data retrieval call binding the contract method 0x193a4c6c.
//
// Solidity: function getBudgetEnd(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_Budgets *BudgetsCallerSession) GetBudgetEnd(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _Budgets.Contract.GetBudgetEnd(&_Budgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetLine is a free data retrieval call binding the contract method 0x54755d3f.
//
// Solidity: function getBudgetLine(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_Budgets *BudgetsCaller) GetBudgetLine(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Budgets.contract.Call(opts, out, "getBudgetLine", _owner, _firstRef, _secondRef, _budgetRef)
	return *ret0, err
}

// GetBudgetLine is a free data retrieval call binding the contract method 0x54755d3f.
//
// Solidity: function getBudgetLine(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_Budgets *BudgetsSession) GetBudgetLine(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _Budgets.Contract.GetBudgetLine(&_Budgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetLine is a free data retrieval call binding the contract method 0x54755d3f.
//
// Solidity: function getBudgetLine(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_Budgets *BudgetsCallerSession) GetBudgetLine(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _Budgets.Contract.GetBudgetLine(&_Budgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetReference is a free data retrieval call binding the contract method 0x6be5d013.
//
// Solidity: function getBudgetReference(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, uint256 _index) view returns(bytes32)
func (_Budgets *BudgetsCaller) GetBudgetReference(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Budgets.contract.Call(opts, out, "getBudgetReference", _owner, _firstRef, _secondRef, _index)
	return *ret0, err
}

// GetBudgetReference is a free data retrieval call binding the contract method 0x6be5d013.
//
// Solidity: function getBudgetReference(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, uint256 _index) view returns(bytes32)
func (_Budgets *BudgetsSession) GetBudgetReference(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _Budgets.Contract.GetBudgetReference(&_Budgets.CallOpts, _owner, _firstRef, _secondRef, _index)
}

// GetBudgetReference is a free data retrieval call binding the contract method 0x6be5d013.
//
// Solidity: function getBudgetReference(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, uint256 _index) view returns(bytes32)
func (_Budgets *BudgetsCallerSession) GetBudgetReference(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _Budgets.Contract.GetBudgetReference(&_Budgets.CallOpts, _owner, _firstRef, _secondRef, _index)
}

// GetBudgetStart is a free data retrieval call binding the contract method 0x4b917df5.
//
// Solidity: function getBudgetStart(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_Budgets *BudgetsCaller) GetBudgetStart(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Budgets.contract.Call(opts, out, "getBudgetStart", _owner, _firstRef, _secondRef, _budgetRef)
	return *ret0, err
}

// GetBudgetStart is a free data retrieval call binding the contract method 0x4b917df5.
//
// Solidity: function getBudgetStart(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_Budgets *BudgetsSession) GetBudgetStart(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _Budgets.Contract.GetBudgetStart(&_Budgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetStart is a free data retrieval call binding the contract method 0x4b917df5.
//
// Solidity: function getBudgetStart(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_Budgets *BudgetsCallerSession) GetBudgetStart(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _Budgets.Contract.GetBudgetStart(&_Budgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetStatus is a free data retrieval call binding the contract method 0x9541246c.
//
// Solidity: function getBudgetStatus(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(uint8)
func (_Budgets *BudgetsCaller) GetBudgetStatus(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Budgets.contract.Call(opts, out, "getBudgetStatus", _owner, _firstRef, _secondRef, _budgetRef)
	return *ret0, err
}

// GetBudgetStatus is a free data retrieval call binding the contract method 0x9541246c.
//
// Solidity: function getBudgetStatus(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(uint8)
func (_Budgets *BudgetsSession) GetBudgetStatus(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (uint8, error) {
	return _Budgets.Contract.GetBudgetStatus(&_Budgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetStatus is a free data retrieval call binding the contract method 0x9541246c.
//
// Solidity: function getBudgetStatus(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(uint8)
func (_Budgets *BudgetsCallerSession) GetBudgetStatus(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (uint8, error) {
	return _Budgets.Contract.GetBudgetStatus(&_Budgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetType is a free data retrieval call binding the contract method 0x98530d3b.
//
// Solidity: function getBudgetType(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(uint8)
func (_Budgets *BudgetsCaller) GetBudgetType(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Budgets.contract.Call(opts, out, "getBudgetType", _owner, _firstRef, _secondRef, _budgetRef)
	return *ret0, err
}

// GetBudgetType is a free data retrieval call binding the contract method 0x98530d3b.
//
// Solidity: function getBudgetType(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(uint8)
func (_Budgets *BudgetsSession) GetBudgetType(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (uint8, error) {
	return _Budgets.Contract.GetBudgetType(&_Budgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetType is a free data retrieval call binding the contract method 0x98530d3b.
//
// Solidity: function getBudgetType(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(uint8)
func (_Budgets *BudgetsCallerSession) GetBudgetType(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (uint8, error) {
	return _Budgets.Contract.GetBudgetType(&_Budgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetValue is a free data retrieval call binding the contract method 0xc3bc3482.
//
// Solidity: function getBudgetValue(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(uint256)
func (_Budgets *BudgetsCaller) GetBudgetValue(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Budgets.contract.Call(opts, out, "getBudgetValue", _owner, _firstRef, _secondRef, _budgetRef)
	return *ret0, err
}

// GetBudgetValue is a free data retrieval call binding the contract method 0xc3bc3482.
//
// Solidity: function getBudgetValue(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(uint256)
func (_Budgets *BudgetsSession) GetBudgetValue(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (*big.Int, error) {
	return _Budgets.Contract.GetBudgetValue(&_Budgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetValue is a free data retrieval call binding the contract method 0xc3bc3482.
//
// Solidity: function getBudgetValue(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(uint256)
func (_Budgets *BudgetsCallerSession) GetBudgetValue(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (*big.Int, error) {
	return _Budgets.Contract.GetBudgetValue(&_Budgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetNumBudgets is a free data retrieval call binding the contract method 0x366b6701.
//
// Solidity: function getNumBudgets(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef) view returns(uint256)
func (_Budgets *BudgetsCaller) GetNumBudgets(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Budgets.contract.Call(opts, out, "getNumBudgets", _owner, _firstRef, _secondRef)
	return *ret0, err
}

// GetNumBudgets is a free data retrieval call binding the contract method 0x366b6701.
//
// Solidity: function getNumBudgets(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef) view returns(uint256)
func (_Budgets *BudgetsSession) GetNumBudgets(_owner uint8, _firstRef [32]byte, _secondRef [32]byte) (*big.Int, error) {
	return _Budgets.Contract.GetNumBudgets(&_Budgets.CallOpts, _owner, _firstRef, _secondRef)
}

// GetNumBudgets is a free data retrieval call binding the contract method 0x366b6701.
//
// Solidity: function getNumBudgets(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef) view returns(uint256)
func (_Budgets *BudgetsCallerSession) GetNumBudgets(_owner uint8, _firstRef [32]byte, _secondRef [32]byte) (*big.Int, error) {
	return _Budgets.Contract.GetNumBudgets(&_Budgets.CallOpts, _owner, _firstRef, _secondRef)
}

// GetOtherRef is a free data retrieval call binding the contract method 0x736d5735.
//
// Solidity: function getOtherRef(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_Budgets *BudgetsCaller) GetOtherRef(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Budgets.contract.Call(opts, out, "getOtherRef", _owner, _firstRef, _secondRef, _budgetRef)
	return *ret0, err
}

// GetOtherRef is a free data retrieval call binding the contract method 0x736d5735.
//
// Solidity: function getOtherRef(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_Budgets *BudgetsSession) GetOtherRef(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _Budgets.Contract.GetOtherRef(&_Budgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetOtherRef is a free data retrieval call binding the contract method 0x736d5735.
//
// Solidity: function getOtherRef(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_Budgets *BudgetsCallerSession) GetOtherRef(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _Budgets.Contract.GetOtherRef(&_Budgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// SetBudget is a paid mutator transaction binding the contract method 0xf1e52779.
//
// Solidity: function setBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef, BudgetsBudget _budget) returns()
func (_Budgets *BudgetsTransactor) SetBudget(opts *bind.TransactOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte, _budget BudgetsBudget) (*types.Transaction, error) {
	return _Budgets.contract.Transact(opts, "setBudget", _owner, _firstRef, _secondRef, _budgetRef, _budget)
}

// SetBudget is a paid mutator transaction binding the contract method 0xf1e52779.
//
// Solidity: function setBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef, BudgetsBudget _budget) returns()
func (_Budgets *BudgetsSession) SetBudget(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte, _budget BudgetsBudget) (*types.Transaction, error) {
	return _Budgets.Contract.SetBudget(&_Budgets.TransactOpts, _owner, _firstRef, _secondRef, _budgetRef, _budget)
}

// SetBudget is a paid mutator transaction binding the contract method 0xf1e52779.
//
// Solidity: function setBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef, BudgetsBudget _budget) returns()
func (_Budgets *BudgetsTransactorSession) SetBudget(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte, _budget BudgetsBudget) (*types.Transaction, error) {
	return _Budgets.Contract.SetBudget(&_Budgets.TransactOpts, _owner, _firstRef, _secondRef, _budgetRef, _budget)
}

// BudgetsSetBudgetIterator is returned from FilterSetBudget and is used to iterate over the raw logs and unpacked data for SetBudget events raised by the Budgets contract.
type BudgetsSetBudgetIterator struct {
	Event *BudgetsSetBudget // Event containing the contract specifics and raw log

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
func (it *BudgetsSetBudgetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BudgetsSetBudget)
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
		it.Event = new(BudgetsSetBudget)
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
func (it *BudgetsSetBudgetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BudgetsSetBudgetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BudgetsSetBudget represents a SetBudget event raised by the Budgets contract.
type BudgetsSetBudget struct {
	Owner     uint8
	FirstRef  [32]byte
	SecondRef [32]byte
	BudgetRef [32]byte
	Budget    BudgetsBudget
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetBudget is a free log retrieval operation binding the contract event 0x3a40e9aa34bf78f77f9b1617948f31ae91db81766f528296cc71aa7addab6e7c.
//
// Solidity: event SetBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef, BudgetsBudget _budget)
func (_Budgets *BudgetsFilterer) FilterSetBudget(opts *bind.FilterOpts) (*BudgetsSetBudgetIterator, error) {

	logs, sub, err := _Budgets.contract.FilterLogs(opts, "SetBudget")
	if err != nil {
		return nil, err
	}
	return &BudgetsSetBudgetIterator{contract: _Budgets.contract, event: "SetBudget", logs: logs, sub: sub}, nil
}

// WatchSetBudget is a free log subscription operation binding the contract event 0x3a40e9aa34bf78f77f9b1617948f31ae91db81766f528296cc71aa7addab6e7c.
//
// Solidity: event SetBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef, BudgetsBudget _budget)
func (_Budgets *BudgetsFilterer) WatchSetBudget(opts *bind.WatchOpts, sink chan<- *BudgetsSetBudget) (event.Subscription, error) {

	logs, sub, err := _Budgets.contract.WatchLogs(opts, "SetBudget")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BudgetsSetBudget)
				if err := _Budgets.contract.UnpackLog(event, "SetBudget", log); err != nil {
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

// ParseSetBudget is a log parse operation binding the contract event 0x3a40e9aa34bf78f77f9b1617948f31ae91db81766f528296cc71aa7addab6e7c.
//
// Solidity: event SetBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef, BudgetsBudget _budget)
func (_Budgets *BudgetsFilterer) ParseSetBudget(log types.Log) (*BudgetsSetBudget, error) {
	event := new(BudgetsSetBudget)
	if err := _Budgets.contract.UnpackLog(event, "SetBudget", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IATIBudgetsABI is the input ABI used to generate the binding from.
const IATIBudgetsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"budgetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"budgetLine\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"otherRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"start\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"end\",\"type\":\"bytes32\"}],\"internalType\":\"structBudgets.Finance\",\"name\":\"finance\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structBudgets.Budget\",\"name\":\"_budget\",\"type\":\"tuple\"}],\"name\":\"SetBudget\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getBudget\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"budgetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"budgetLine\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"otherRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"start\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"end\",\"type\":\"bytes32\"}],\"internalType\":\"structBudgets.Finance\",\"name\":\"finance\",\"type\":\"tuple\"}],\"internalType\":\"structBudgets.Budget\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getBudgetEnd\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getBudgetLine\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getBudgetReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getBudgetStart\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getBudgetStatus\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getBudgetType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getBudgetValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"}],\"name\":\"getNumBudgets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getOtherRef\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_owner\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_firstRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_secondRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"budgetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"budgetLine\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"otherRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"start\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"end\",\"type\":\"bytes32\"}],\"internalType\":\"structBudgets.Finance\",\"name\":\"finance\",\"type\":\"tuple\"}],\"internalType\":\"structBudgets.Budget\",\"name\":\"_budget\",\"type\":\"tuple\"}],\"name\":\"setBudget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IATIBudgetsFuncSigs maps the 4-byte function signature to its string representation.
var IATIBudgetsFuncSigs = map[string]string{
	"66cf0ca9": "getBudget(uint8,bytes32,bytes32,bytes32)",
	"193a4c6c": "getBudgetEnd(uint8,bytes32,bytes32,bytes32)",
	"54755d3f": "getBudgetLine(uint8,bytes32,bytes32,bytes32)",
	"6be5d013": "getBudgetReference(uint8,bytes32,bytes32,uint256)",
	"4b917df5": "getBudgetStart(uint8,bytes32,bytes32,bytes32)",
	"9541246c": "getBudgetStatus(uint8,bytes32,bytes32,bytes32)",
	"98530d3b": "getBudgetType(uint8,bytes32,bytes32,bytes32)",
	"c3bc3482": "getBudgetValue(uint8,bytes32,bytes32,bytes32)",
	"366b6701": "getNumBudgets(uint8,bytes32,bytes32)",
	"736d5735": "getOtherRef(uint8,bytes32,bytes32,bytes32)",
	"f1e52779": "setBudget(uint8,bytes32,bytes32,bytes32,(uint8,bytes32,bytes32,(uint8,uint256,bytes32,bytes32)))",
}

// IATIBudgetsBin is the compiled bytecode used for deploying new contracts.
var IATIBudgetsBin = "0x608060405234801561001057600080fd5b50610e7e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80636be5d013116100715780636be5d01314610130578063736d5735146101435780639541246c1461015657806398530d3b14610176578063c3bc348214610189578063f1e527791461019c576100a9565b8063193a4c6c146100ae578063366b6701146100d75780634b917df5146100ea57806354755d3f146100fd57806366cf0ca914610110575b600080fd5b6100c16100bc366004610bee565b6101b1565b6040516100ce9190610d62565b60405180910390f35b6100c16100e5366004610bbb565b61025b565b6100c16100f8366004610bee565b6102dd565b6100c161010b366004610bee565b610387565b61012361011e366004610bee565b61042f565b6040516100ce9190610dbb565b6100c161013e366004610d00565b610534565b6100c1610151366004610bee565b6105fc565b610169610164366004610bee565b6106a6565b6040516100ce9190610dc9565b610169610184366004610bee565b61074e565b6100c1610197366004610bee565b6107f3565b6101af6101aa366004610c28565b61089d565b005b600060ff8516158015906101c85750600760ff8616105b80156101e357508360001a60f81b6001600160f81b03191615155b80156101fe57508260001a60f81b6001600160f81b03191615155b801561021957508160001a60f81b6001600160f81b03191615155b61022257600080fd5b5060ff93909316600090815260016020908152604080832094835293815283822092825291825282812093815292905290206006015490565b600060ff8416158015906102725750600760ff8516105b801561028d57508260001a60f81b6001600160f81b03191615155b80156102a857508160001a60f81b6001600160f81b03191615155b6102b157600080fd5b5060ff831660009081526020818152604080832085845282528083208484529091529020549392505050565b600060ff8516158015906102f45750600760ff8616105b801561030f57508360001a60f81b6001600160f81b03191615155b801561032a57508260001a60f81b6001600160f81b03191615155b801561034557508160001a60f81b6001600160f81b03191615155b61034e57600080fd5b5060ff93909316600090815260016020908152604080832094835293815283822092825291825282812093815292905290206005015490565b600060ff85161580159061039e5750600760ff8616105b80156103b957508360001a60f81b6001600160f81b03191615155b80156103d457508260001a60f81b6001600160f81b03191615155b80156103ef57508160001a60f81b6001600160f81b03191615155b6103f857600080fd5b5060ff9390931660009081526001602081815260408084209584529481528483209383529283528382209482529390915220015490565b610437610b28565b60ff85161580159061044c5750600760ff8616105b801561046757508360001a60f81b6001600160f81b03191615155b801561048257508260001a60f81b6001600160f81b03191615155b801561049d57508160001a60f81b6001600160f81b03191615155b6104a657600080fd5b5060ff9384166000908152600160208181526040808420968452958152858320948352938452848220928252918352839020835160808082018652825487168252928201548185015260028201548186015284519283018552600382015490951682526004810154928201929092526005820154928101929092526006015460608083019190915282015290565b600060ff85161580159061054b5750600760ff8616105b801561056657508360001a60f81b6001600160f81b03191615155b801561058157508260001a60f81b6001600160f81b03191615155b80156105af575060ff8516600090815260208181526040808320878452825280832086845290915290205482105b6105b857600080fd5b60ff8516600090815260208181526040808320878452825280832086845290915290208054839081106105e757fe5b90600052602060002001549050949350505050565b600060ff8516158015906106135750600760ff8616105b801561062e57508360001a60f81b6001600160f81b03191615155b801561064957508260001a60f81b6001600160f81b03191615155b801561066457508160001a60f81b6001600160f81b03191615155b61066d57600080fd5b5060ff93909316600090815260016020908152604080832094835293815283822092825291825282812093815292905290206002015490565b600060ff8516158015906106bd5750600760ff8616105b80156106d857508360001a60f81b6001600160f81b03191615155b80156106f357508260001a60f81b6001600160f81b03191615155b801561070e57508160001a60f81b6001600160f81b03191615155b61071757600080fd5b5060ff9384166000908152600160209081526040808320958352948152848220938252928352838120918152915220600301541690565b600060ff8516158015906107655750600760ff8616105b801561078057508360001a60f81b6001600160f81b03191615155b801561079b57508260001a60f81b6001600160f81b03191615155b80156107b657508160001a60f81b6001600160f81b03191615155b6107bf57600080fd5b5060ff9384166000908152600160209081526040808320958352948152848220938252928352838120918152915220541690565b600060ff85161580159061080a5750600760ff8616105b801561082557508360001a60f81b6001600160f81b03191615155b801561084057508260001a60f81b6001600160f81b03191615155b801561085b57508160001a60f81b6001600160f81b03191615155b61086457600080fd5b5060ff93909316600090815260016020908152604080832094835293815283822092825291825282812093815292905290206004015490565b60ff8516158015906108b25750600760ff8616105b80156108cd57508360001a60f81b6001600160f81b03191615155b80156108e857508260001a60f81b6001600160f81b03191615155b801561090357508160001a60f81b6001600160f81b03191615155b8015610916575060608101515160ff1615155b801561092c5750606081015151600360ff909116105b801561094f575060608101516040015160001a60f81b6001600160f81b03191615155b80156109715750606080820151015160001a60f81b6001600160f81b03191615155b61097a57600080fd5b60ff80861660008181526001602081815260408084208a85528252808420898552825280842088855282528084208751815490881660ff199182161782558389015194820194909455818801516002820155606080890151805160038401805491909a169616959095179097558383015160048083019190915584830151600583015593909601516006909601959095559282528183528382208883528352838220878352909252829020915163745fca7b60e01b815273__$b620240993a55615b607189176fb82e62f$__9263745fca7b92610a5a9287929101610d6b565b60206040518083038186803b158015610a7257600080fd5b505af4158015610a86573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aaa9190610b94565b610ae25760ff851660009081526020818152604080832087845282528083208684528252822080546001810182559083529120018290555b7f3a40e9aa34bf78f77f9b1617948f31ae91db81766f528296cc71aa7addab6e7c8585858585604051610b19959493929190610dd7565b60405180910390a15050505050565b604080516080810182526000808252602082018190529181019190915260608101610b51610b56565b905290565b60408051608081018252600080825260208201819052918101829052606081019190915290565b803560ff81168114610b8e57600080fd5b92915050565b600060208284031215610ba5578081fd5b81518015158114610bb4578182fd5b9392505050565b600080600060608486031215610bcf578182fd5b610bd98585610b7d565b95602085013595506040909401359392505050565b60008060008060808587031215610c03578081fd5b8435610c0e81610e36565b966020860135965060408601359560600135945092505050565b6000806000806000858703610160811215610c41578182fd5b610c4b8888610b7d565b955060208701359450604087013593506060870135925060e0607f1982011215610c73578182fd5b610c7d6080610e0f565b610c8a8960808a01610b7d565b815260a0880135602082015260c08801356040820152608060df1983011215610cb1578283fd5b610cbb6080610e0f565b915060e0880135610ccb81610e36565b825261010088013560208301526101208801356040830152610140909701356060808301919091528701525092959194509290565b60008060008060808587031215610c03578384fd5b60ff81511682526020810151602083015260408101516040830152606081015160ff815116606084015260208101516080840152604081015160a0840152606081015160c0840152505050565b90815260200190565b60006040820184835260206040818501528185548084526060860191508685528285209350845b81811015610dae57845483526001948501949284019201610d92565b5090979650505050505050565b60e08101610b8e8284610d15565b60ff91909116815260200190565b60006101608201905060ff87168252856020830152846040830152836060830152610e056080830184610d15565b9695505050505050565b60405181810167ffffffffffffffff81118282101715610e2e57600080fd5b604052919050565b60ff81168114610e4557600080fd5b5056fea2646970667358221220e28750147182d0e94db906ef5af67662d93c220a2f14a59e23480f809eaf996464736f6c63430006060033"

// DeployIATIBudgets deploys a new Ethereum contract, binding an instance of IATIBudgets to it.
func DeployIATIBudgets(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IATIBudgets, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIBudgetsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	stringsAddr, _, _, _ := DeployStrings(auth, backend)
	IATIBudgetsBin = strings.Replace(IATIBudgetsBin, "__$b620240993a55615b607189176fb82e62f$__", stringsAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IATIBudgetsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IATIBudgets{IATIBudgetsCaller: IATIBudgetsCaller{contract: contract}, IATIBudgetsTransactor: IATIBudgetsTransactor{contract: contract}, IATIBudgetsFilterer: IATIBudgetsFilterer{contract: contract}}, nil
}

// IATIBudgets is an auto generated Go binding around an Ethereum contract.
type IATIBudgets struct {
	IATIBudgetsCaller     // Read-only binding to the contract
	IATIBudgetsTransactor // Write-only binding to the contract
	IATIBudgetsFilterer   // Log filterer for contract events
}

// IATIBudgetsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IATIBudgetsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIBudgetsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IATIBudgetsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIBudgetsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IATIBudgetsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIBudgetsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IATIBudgetsSession struct {
	Contract     *IATIBudgets      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IATIBudgetsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IATIBudgetsCallerSession struct {
	Contract *IATIBudgetsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IATIBudgetsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IATIBudgetsTransactorSession struct {
	Contract     *IATIBudgetsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IATIBudgetsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IATIBudgetsRaw struct {
	Contract *IATIBudgets // Generic contract binding to access the raw methods on
}

// IATIBudgetsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IATIBudgetsCallerRaw struct {
	Contract *IATIBudgetsCaller // Generic read-only contract binding to access the raw methods on
}

// IATIBudgetsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IATIBudgetsTransactorRaw struct {
	Contract *IATIBudgetsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIATIBudgets creates a new instance of IATIBudgets, bound to a specific deployed contract.
func NewIATIBudgets(address common.Address, backend bind.ContractBackend) (*IATIBudgets, error) {
	contract, err := bindIATIBudgets(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IATIBudgets{IATIBudgetsCaller: IATIBudgetsCaller{contract: contract}, IATIBudgetsTransactor: IATIBudgetsTransactor{contract: contract}, IATIBudgetsFilterer: IATIBudgetsFilterer{contract: contract}}, nil
}

// NewIATIBudgetsCaller creates a new read-only instance of IATIBudgets, bound to a specific deployed contract.
func NewIATIBudgetsCaller(address common.Address, caller bind.ContractCaller) (*IATIBudgetsCaller, error) {
	contract, err := bindIATIBudgets(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IATIBudgetsCaller{contract: contract}, nil
}

// NewIATIBudgetsTransactor creates a new write-only instance of IATIBudgets, bound to a specific deployed contract.
func NewIATIBudgetsTransactor(address common.Address, transactor bind.ContractTransactor) (*IATIBudgetsTransactor, error) {
	contract, err := bindIATIBudgets(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IATIBudgetsTransactor{contract: contract}, nil
}

// NewIATIBudgetsFilterer creates a new log filterer instance of IATIBudgets, bound to a specific deployed contract.
func NewIATIBudgetsFilterer(address common.Address, filterer bind.ContractFilterer) (*IATIBudgetsFilterer, error) {
	contract, err := bindIATIBudgets(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IATIBudgetsFilterer{contract: contract}, nil
}

// bindIATIBudgets binds a generic wrapper to an already deployed contract.
func bindIATIBudgets(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIBudgetsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIBudgets *IATIBudgetsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIBudgets.Contract.IATIBudgetsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIBudgets *IATIBudgetsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIBudgets.Contract.IATIBudgetsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIBudgets *IATIBudgetsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIBudgets.Contract.IATIBudgetsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIBudgets *IATIBudgetsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIBudgets.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIBudgets *IATIBudgetsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIBudgets.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIBudgets *IATIBudgetsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIBudgets.Contract.contract.Transact(opts, method, params...)
}

// GetBudget is a free data retrieval call binding the contract method 0x66cf0ca9.
//
// Solidity: function getBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(BudgetsBudget)
func (_IATIBudgets *IATIBudgetsCaller) GetBudget(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (BudgetsBudget, error) {
	var (
		ret0 = new(BudgetsBudget)
	)
	out := ret0
	err := _IATIBudgets.contract.Call(opts, out, "getBudget", _owner, _firstRef, _secondRef, _budgetRef)
	return *ret0, err
}

// GetBudget is a free data retrieval call binding the contract method 0x66cf0ca9.
//
// Solidity: function getBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(BudgetsBudget)
func (_IATIBudgets *IATIBudgetsSession) GetBudget(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (BudgetsBudget, error) {
	return _IATIBudgets.Contract.GetBudget(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudget is a free data retrieval call binding the contract method 0x66cf0ca9.
//
// Solidity: function getBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(BudgetsBudget)
func (_IATIBudgets *IATIBudgetsCallerSession) GetBudget(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (BudgetsBudget, error) {
	return _IATIBudgets.Contract.GetBudget(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetEnd is a free data retrieval call binding the contract method 0x193a4c6c.
//
// Solidity: function getBudgetEnd(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIBudgets *IATIBudgetsCaller) GetBudgetEnd(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIBudgets.contract.Call(opts, out, "getBudgetEnd", _owner, _firstRef, _secondRef, _budgetRef)
	return *ret0, err
}

// GetBudgetEnd is a free data retrieval call binding the contract method 0x193a4c6c.
//
// Solidity: function getBudgetEnd(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIBudgets *IATIBudgetsSession) GetBudgetEnd(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _IATIBudgets.Contract.GetBudgetEnd(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetEnd is a free data retrieval call binding the contract method 0x193a4c6c.
//
// Solidity: function getBudgetEnd(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIBudgets *IATIBudgetsCallerSession) GetBudgetEnd(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _IATIBudgets.Contract.GetBudgetEnd(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetLine is a free data retrieval call binding the contract method 0x54755d3f.
//
// Solidity: function getBudgetLine(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIBudgets *IATIBudgetsCaller) GetBudgetLine(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIBudgets.contract.Call(opts, out, "getBudgetLine", _owner, _firstRef, _secondRef, _budgetRef)
	return *ret0, err
}

// GetBudgetLine is a free data retrieval call binding the contract method 0x54755d3f.
//
// Solidity: function getBudgetLine(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIBudgets *IATIBudgetsSession) GetBudgetLine(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _IATIBudgets.Contract.GetBudgetLine(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetLine is a free data retrieval call binding the contract method 0x54755d3f.
//
// Solidity: function getBudgetLine(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIBudgets *IATIBudgetsCallerSession) GetBudgetLine(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _IATIBudgets.Contract.GetBudgetLine(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetReference is a free data retrieval call binding the contract method 0x6be5d013.
//
// Solidity: function getBudgetReference(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, uint256 _index) view returns(bytes32)
func (_IATIBudgets *IATIBudgetsCaller) GetBudgetReference(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIBudgets.contract.Call(opts, out, "getBudgetReference", _owner, _firstRef, _secondRef, _index)
	return *ret0, err
}

// GetBudgetReference is a free data retrieval call binding the contract method 0x6be5d013.
//
// Solidity: function getBudgetReference(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, uint256 _index) view returns(bytes32)
func (_IATIBudgets *IATIBudgetsSession) GetBudgetReference(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIBudgets.Contract.GetBudgetReference(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef, _index)
}

// GetBudgetReference is a free data retrieval call binding the contract method 0x6be5d013.
//
// Solidity: function getBudgetReference(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, uint256 _index) view returns(bytes32)
func (_IATIBudgets *IATIBudgetsCallerSession) GetBudgetReference(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIBudgets.Contract.GetBudgetReference(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef, _index)
}

// GetBudgetStart is a free data retrieval call binding the contract method 0x4b917df5.
//
// Solidity: function getBudgetStart(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIBudgets *IATIBudgetsCaller) GetBudgetStart(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIBudgets.contract.Call(opts, out, "getBudgetStart", _owner, _firstRef, _secondRef, _budgetRef)
	return *ret0, err
}

// GetBudgetStart is a free data retrieval call binding the contract method 0x4b917df5.
//
// Solidity: function getBudgetStart(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIBudgets *IATIBudgetsSession) GetBudgetStart(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _IATIBudgets.Contract.GetBudgetStart(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetStart is a free data retrieval call binding the contract method 0x4b917df5.
//
// Solidity: function getBudgetStart(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIBudgets *IATIBudgetsCallerSession) GetBudgetStart(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _IATIBudgets.Contract.GetBudgetStart(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetStatus is a free data retrieval call binding the contract method 0x9541246c.
//
// Solidity: function getBudgetStatus(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(uint8)
func (_IATIBudgets *IATIBudgetsCaller) GetBudgetStatus(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIBudgets.contract.Call(opts, out, "getBudgetStatus", _owner, _firstRef, _secondRef, _budgetRef)
	return *ret0, err
}

// GetBudgetStatus is a free data retrieval call binding the contract method 0x9541246c.
//
// Solidity: function getBudgetStatus(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(uint8)
func (_IATIBudgets *IATIBudgetsSession) GetBudgetStatus(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (uint8, error) {
	return _IATIBudgets.Contract.GetBudgetStatus(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetStatus is a free data retrieval call binding the contract method 0x9541246c.
//
// Solidity: function getBudgetStatus(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(uint8)
func (_IATIBudgets *IATIBudgetsCallerSession) GetBudgetStatus(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (uint8, error) {
	return _IATIBudgets.Contract.GetBudgetStatus(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetType is a free data retrieval call binding the contract method 0x98530d3b.
//
// Solidity: function getBudgetType(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(uint8)
func (_IATIBudgets *IATIBudgetsCaller) GetBudgetType(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIBudgets.contract.Call(opts, out, "getBudgetType", _owner, _firstRef, _secondRef, _budgetRef)
	return *ret0, err
}

// GetBudgetType is a free data retrieval call binding the contract method 0x98530d3b.
//
// Solidity: function getBudgetType(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(uint8)
func (_IATIBudgets *IATIBudgetsSession) GetBudgetType(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (uint8, error) {
	return _IATIBudgets.Contract.GetBudgetType(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetType is a free data retrieval call binding the contract method 0x98530d3b.
//
// Solidity: function getBudgetType(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(uint8)
func (_IATIBudgets *IATIBudgetsCallerSession) GetBudgetType(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (uint8, error) {
	return _IATIBudgets.Contract.GetBudgetType(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetValue is a free data retrieval call binding the contract method 0xc3bc3482.
//
// Solidity: function getBudgetValue(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(uint256)
func (_IATIBudgets *IATIBudgetsCaller) GetBudgetValue(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIBudgets.contract.Call(opts, out, "getBudgetValue", _owner, _firstRef, _secondRef, _budgetRef)
	return *ret0, err
}

// GetBudgetValue is a free data retrieval call binding the contract method 0xc3bc3482.
//
// Solidity: function getBudgetValue(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(uint256)
func (_IATIBudgets *IATIBudgetsSession) GetBudgetValue(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (*big.Int, error) {
	return _IATIBudgets.Contract.GetBudgetValue(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetBudgetValue is a free data retrieval call binding the contract method 0xc3bc3482.
//
// Solidity: function getBudgetValue(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(uint256)
func (_IATIBudgets *IATIBudgetsCallerSession) GetBudgetValue(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) (*big.Int, error) {
	return _IATIBudgets.Contract.GetBudgetValue(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetNumBudgets is a free data retrieval call binding the contract method 0x366b6701.
//
// Solidity: function getNumBudgets(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef) view returns(uint256)
func (_IATIBudgets *IATIBudgetsCaller) GetNumBudgets(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIBudgets.contract.Call(opts, out, "getNumBudgets", _owner, _firstRef, _secondRef)
	return *ret0, err
}

// GetNumBudgets is a free data retrieval call binding the contract method 0x366b6701.
//
// Solidity: function getNumBudgets(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef) view returns(uint256)
func (_IATIBudgets *IATIBudgetsSession) GetNumBudgets(_owner uint8, _firstRef [32]byte, _secondRef [32]byte) (*big.Int, error) {
	return _IATIBudgets.Contract.GetNumBudgets(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef)
}

// GetNumBudgets is a free data retrieval call binding the contract method 0x366b6701.
//
// Solidity: function getNumBudgets(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef) view returns(uint256)
func (_IATIBudgets *IATIBudgetsCallerSession) GetNumBudgets(_owner uint8, _firstRef [32]byte, _secondRef [32]byte) (*big.Int, error) {
	return _IATIBudgets.Contract.GetNumBudgets(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef)
}

// GetOtherRef is a free data retrieval call binding the contract method 0x736d5735.
//
// Solidity: function getOtherRef(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIBudgets *IATIBudgetsCaller) GetOtherRef(opts *bind.CallOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIBudgets.contract.Call(opts, out, "getOtherRef", _owner, _firstRef, _secondRef, _budgetRef)
	return *ret0, err
}

// GetOtherRef is a free data retrieval call binding the contract method 0x736d5735.
//
// Solidity: function getOtherRef(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIBudgets *IATIBudgetsSession) GetOtherRef(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _IATIBudgets.Contract.GetOtherRef(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// GetOtherRef is a free data retrieval call binding the contract method 0x736d5735.
//
// Solidity: function getOtherRef(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIBudgets *IATIBudgetsCallerSession) GetOtherRef(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _IATIBudgets.Contract.GetOtherRef(&_IATIBudgets.CallOpts, _owner, _firstRef, _secondRef, _budgetRef)
}

// SetBudget is a paid mutator transaction binding the contract method 0xf1e52779.
//
// Solidity: function setBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef, BudgetsBudget _budget) returns()
func (_IATIBudgets *IATIBudgetsTransactor) SetBudget(opts *bind.TransactOpts, _owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte, _budget BudgetsBudget) (*types.Transaction, error) {
	return _IATIBudgets.contract.Transact(opts, "setBudget", _owner, _firstRef, _secondRef, _budgetRef, _budget)
}

// SetBudget is a paid mutator transaction binding the contract method 0xf1e52779.
//
// Solidity: function setBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef, BudgetsBudget _budget) returns()
func (_IATIBudgets *IATIBudgetsSession) SetBudget(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte, _budget BudgetsBudget) (*types.Transaction, error) {
	return _IATIBudgets.Contract.SetBudget(&_IATIBudgets.TransactOpts, _owner, _firstRef, _secondRef, _budgetRef, _budget)
}

// SetBudget is a paid mutator transaction binding the contract method 0xf1e52779.
//
// Solidity: function setBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef, BudgetsBudget _budget) returns()
func (_IATIBudgets *IATIBudgetsTransactorSession) SetBudget(_owner uint8, _firstRef [32]byte, _secondRef [32]byte, _budgetRef [32]byte, _budget BudgetsBudget) (*types.Transaction, error) {
	return _IATIBudgets.Contract.SetBudget(&_IATIBudgets.TransactOpts, _owner, _firstRef, _secondRef, _budgetRef, _budget)
}

// IATIBudgetsSetBudgetIterator is returned from FilterSetBudget and is used to iterate over the raw logs and unpacked data for SetBudget events raised by the IATIBudgets contract.
type IATIBudgetsSetBudgetIterator struct {
	Event *IATIBudgetsSetBudget // Event containing the contract specifics and raw log

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
func (it *IATIBudgetsSetBudgetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IATIBudgetsSetBudget)
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
		it.Event = new(IATIBudgetsSetBudget)
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
func (it *IATIBudgetsSetBudgetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IATIBudgetsSetBudgetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IATIBudgetsSetBudget represents a SetBudget event raised by the IATIBudgets contract.
type IATIBudgetsSetBudget struct {
	Owner     uint8
	FirstRef  [32]byte
	SecondRef [32]byte
	BudgetRef [32]byte
	Budget    BudgetsBudget
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetBudget is a free log retrieval operation binding the contract event 0x3a40e9aa34bf78f77f9b1617948f31ae91db81766f528296cc71aa7addab6e7c.
//
// Solidity: event SetBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef, BudgetsBudget _budget)
func (_IATIBudgets *IATIBudgetsFilterer) FilterSetBudget(opts *bind.FilterOpts) (*IATIBudgetsSetBudgetIterator, error) {

	logs, sub, err := _IATIBudgets.contract.FilterLogs(opts, "SetBudget")
	if err != nil {
		return nil, err
	}
	return &IATIBudgetsSetBudgetIterator{contract: _IATIBudgets.contract, event: "SetBudget", logs: logs, sub: sub}, nil
}

// WatchSetBudget is a free log subscription operation binding the contract event 0x3a40e9aa34bf78f77f9b1617948f31ae91db81766f528296cc71aa7addab6e7c.
//
// Solidity: event SetBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef, BudgetsBudget _budget)
func (_IATIBudgets *IATIBudgetsFilterer) WatchSetBudget(opts *bind.WatchOpts, sink chan<- *IATIBudgetsSetBudget) (event.Subscription, error) {

	logs, sub, err := _IATIBudgets.contract.WatchLogs(opts, "SetBudget")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IATIBudgetsSetBudget)
				if err := _IATIBudgets.contract.UnpackLog(event, "SetBudget", log); err != nil {
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

// ParseSetBudget is a log parse operation binding the contract event 0x3a40e9aa34bf78f77f9b1617948f31ae91db81766f528296cc71aa7addab6e7c.
//
// Solidity: event SetBudget(uint8 _owner, bytes32 _firstRef, bytes32 _secondRef, bytes32 _budgetRef, BudgetsBudget _budget)
func (_IATIBudgets *IATIBudgetsFilterer) ParseSetBudget(log types.Log) (*IATIBudgetsSetBudget, error) {
	event := new(IATIBudgetsSetBudget)
	if err := _IATIBudgets.contract.UnpackLog(event, "SetBudget", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IATIOrganisationRecipientBudgetsABI is the input ABI used to generate the binding from.
const IATIOrganisationRecipientBudgetsABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_budgets\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getNumRecipientBudgets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getRecipientBudget\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"budgetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"budgetLine\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"otherRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"start\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"end\",\"type\":\"bytes32\"}],\"internalType\":\"structBudgets.Finance\",\"name\":\"finance\",\"type\":\"tuple\"}],\"internalType\":\"structBudgets.Budget\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getRecipientBudgetEnd\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getRecipientBudgetLine\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getRecipientBudgetOrg\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRecipientBudgetReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getRecipientBudgetStart\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getRecipientBudgetStatus\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getRecipientBudgetValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"budgetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"budgetLine\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"otherRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"start\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"end\",\"type\":\"bytes32\"}],\"internalType\":\"structBudgets.Finance\",\"name\":\"finance\",\"type\":\"tuple\"}],\"internalType\":\"structBudgets.Budget\",\"name\":\"_budget\",\"type\":\"tuple\"}],\"name\":\"setRecipientBudget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IATIOrganisationRecipientBudgetsFuncSigs maps the 4-byte function signature to its string representation.
var IATIOrganisationRecipientBudgetsFuncSigs = map[string]string{
	"6bdc1aca": "getNumRecipientBudgets(bytes32,bytes32)",
	"6422072b": "getRecipientBudget(bytes32,bytes32,bytes32)",
	"e3391566": "getRecipientBudgetEnd(bytes32,bytes32,bytes32)",
	"2fc452f8": "getRecipientBudgetLine(bytes32,bytes32,bytes32)",
	"0629561d": "getRecipientBudgetOrg(bytes32,bytes32,bytes32)",
	"41cd0fd0": "getRecipientBudgetReference(bytes32,bytes32,uint256)",
	"999af84d": "getRecipientBudgetStart(bytes32,bytes32,bytes32)",
	"98e2d588": "getRecipientBudgetStatus(bytes32,bytes32,bytes32)",
	"e8ae58f0": "getRecipientBudgetValue(bytes32,bytes32,bytes32)",
	"497c6d69": "setRecipientBudget(bytes32,bytes32,bytes32,(uint8,bytes32,bytes32,(uint8,uint256,bytes32,bytes32)))",
}

// IATIOrganisationRecipientBudgetsBin is the compiled bytecode used for deploying new contracts.
var IATIOrganisationRecipientBudgetsBin = "0x608060405234801561001057600080fd5b506040516108fd3803806108fd83398101604081905261002f91610067565b6001600160a01b03811661004257600080fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055610095565b600060208284031215610078578081fd5b81516001600160a01b038116811461008e578182fd5b9392505050565b610859806100a46000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80636bdc1aca116100665780636bdc1aca1461012757806398e2d5881461013a578063999af84d1461015a578063e33915661461016d578063e8ae58f0146101805761009e565b80630629561d146100a35780632fc452f8146100cc57806341cd0fd0146100df578063497c6d69146100f25780636422072b14610107575b600080fd5b6100b66100b1366004610547565b610193565b6040516100c39190610752565b60405180910390f35b6100b66100da366004610547565b610222565b6100b66100ed366004610639565b61023b565b610105610100366004610572565b610254565b005b61011a610115366004610547565b6102c3565b6040516100c3919061075b565b6100b6610135366004610526565b610350565b61014d610148366004610547565b6103dd565b6040516100c39190610769565b6100b6610168366004610547565b610463565b6100b661017b366004610547565b61047c565b6100b661018e366004610547565b610495565b600080546001600160a01b031663736d573560035b8686866040518563ffffffff1660e01b81526004016101ca9493929190610792565b60206040518083038186803b1580156101e257600080fd5b505afa1580156101f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061021a919061050e565b949350505050565b600080546001600160a01b03166354755d3f60036101a8565b600080546001600160a01b0316636be5d01360036101a8565b6000546001600160a01b031663f1e527796003868686866040518663ffffffff1660e01b815260040161028b9594939291906107b2565b600060405180830381600087803b1580156102a557600080fd5b505af11580156102b9573d6000803e3d6000fd5b5050505050505050565b6102cb6104ae565b6000546001600160a01b03166366cf0ca960038686866040518563ffffffff1660e01b81526004016103009493929190610792565b60e06040518083038186803b15801561031857600080fd5b505afa15801561032c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061021a919061064d565b600080546001600160a01b031663366b6701600385856040518463ffffffff1660e01b815260040161038493929190610777565b60206040518083038186803b15801561039c57600080fd5b505afa1580156103b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103d4919061050e565b90505b92915050565b600080546001600160a01b0316639541246c60038686866040518563ffffffff1660e01b81526004016104139493929190610792565b60206040518083038186803b15801561042b57600080fd5b505afa15801561043f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061021a91906106e2565b600080546001600160a01b0316634b917df560036101a8565b600080546001600160a01b031663193a4c6c60036101a8565b600080546001600160a01b031663c3bc348260036101a8565b6040805160808101825260008082526020820181905291810191909152606081016104d76104dc565b905290565b60408051608081018252600080825260208201819052918101829052606081019190915290565b80516103d781610811565b60006020828403121561051f578081fd5b5051919050565b60008060408385031215610538578081fd5b50508035926020909101359150565b60008060006060848603121561055b578081fd5b505081359360208301359350604090920135919050565b600080600080848603610140811215610589578182fd5b85359450602086013593506040860135925060e0605f19820112156105ac578182fd5b6105b660806107ea565b60608701356105c481610811565b8152608087810135602083015260a0880135604083015260bf19830112156105ea578283fd5b6105f460806107ea565b915060c087013561060481610811565b825260e08701356020830152610100870135604083015261012090960135606080830191909152860152509194909350909190565b60008060006060848603121561055b578283fd5b600081830360e081121561065f578182fd5b61066960806107ea565b835161067481610811565b815260208481015190820152604080850151908201526080605f198301121561069b578283fd5b6106a560806107ea565b91506106b48560608601610503565b82526080840151602083015260a0840151604083015260c09093015160608083019190915283015250919050565b6000602082840312156106f3578081fd5b81516106fe81610811565b9392505050565b60ff81511682526020810151602083015260408101516040830152606081015160ff815116606084015260208101516080840152604081015160a0840152606081015160c0840152505050565b90815260200190565b60e081016103d78284610705565b60ff91909116815260200190565b60ff9390931683526020830191909152604082015260600190565b60ff94909416845260208401929092526040830152606082015260800190565b60006101608201905060ff871682528560208301528460408301528360608301526107e06080830184610705565b9695505050505050565b60405181810167ffffffffffffffff8111828210171561080957600080fd5b604052919050565b60ff8116811461082057600080fd5b5056fea2646970667358221220472a2d1b66c9655d202ace5a115a117841917000a255475d1669c24c60c85b8f64736f6c63430006060033"

// DeployIATIOrganisationRecipientBudgets deploys a new Ethereum contract, binding an instance of IATIOrganisationRecipientBudgets to it.
func DeployIATIOrganisationRecipientBudgets(auth *bind.TransactOpts, backend bind.ContractBackend, _budgets common.Address) (common.Address, *types.Transaction, *IATIOrganisationRecipientBudgets, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIOrganisationRecipientBudgetsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IATIOrganisationRecipientBudgetsBin), backend, _budgets)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IATIOrganisationRecipientBudgets{IATIOrganisationRecipientBudgetsCaller: IATIOrganisationRecipientBudgetsCaller{contract: contract}, IATIOrganisationRecipientBudgetsTransactor: IATIOrganisationRecipientBudgetsTransactor{contract: contract}, IATIOrganisationRecipientBudgetsFilterer: IATIOrganisationRecipientBudgetsFilterer{contract: contract}}, nil
}

// IATIOrganisationRecipientBudgets is an auto generated Go binding around an Ethereum contract.
type IATIOrganisationRecipientBudgets struct {
	IATIOrganisationRecipientBudgetsCaller     // Read-only binding to the contract
	IATIOrganisationRecipientBudgetsTransactor // Write-only binding to the contract
	IATIOrganisationRecipientBudgetsFilterer   // Log filterer for contract events
}

// IATIOrganisationRecipientBudgetsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IATIOrganisationRecipientBudgetsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIOrganisationRecipientBudgetsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IATIOrganisationRecipientBudgetsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIOrganisationRecipientBudgetsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IATIOrganisationRecipientBudgetsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIOrganisationRecipientBudgetsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IATIOrganisationRecipientBudgetsSession struct {
	Contract     *IATIOrganisationRecipientBudgets // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                     // Call options to use throughout this session
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IATIOrganisationRecipientBudgetsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IATIOrganisationRecipientBudgetsCallerSession struct {
	Contract *IATIOrganisationRecipientBudgetsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                           // Call options to use throughout this session
}

// IATIOrganisationRecipientBudgetsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IATIOrganisationRecipientBudgetsTransactorSession struct {
	Contract     *IATIOrganisationRecipientBudgetsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// IATIOrganisationRecipientBudgetsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IATIOrganisationRecipientBudgetsRaw struct {
	Contract *IATIOrganisationRecipientBudgets // Generic contract binding to access the raw methods on
}

// IATIOrganisationRecipientBudgetsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IATIOrganisationRecipientBudgetsCallerRaw struct {
	Contract *IATIOrganisationRecipientBudgetsCaller // Generic read-only contract binding to access the raw methods on
}

// IATIOrganisationRecipientBudgetsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IATIOrganisationRecipientBudgetsTransactorRaw struct {
	Contract *IATIOrganisationRecipientBudgetsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIATIOrganisationRecipientBudgets creates a new instance of IATIOrganisationRecipientBudgets, bound to a specific deployed contract.
func NewIATIOrganisationRecipientBudgets(address common.Address, backend bind.ContractBackend) (*IATIOrganisationRecipientBudgets, error) {
	contract, err := bindIATIOrganisationRecipientBudgets(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IATIOrganisationRecipientBudgets{IATIOrganisationRecipientBudgetsCaller: IATIOrganisationRecipientBudgetsCaller{contract: contract}, IATIOrganisationRecipientBudgetsTransactor: IATIOrganisationRecipientBudgetsTransactor{contract: contract}, IATIOrganisationRecipientBudgetsFilterer: IATIOrganisationRecipientBudgetsFilterer{contract: contract}}, nil
}

// NewIATIOrganisationRecipientBudgetsCaller creates a new read-only instance of IATIOrganisationRecipientBudgets, bound to a specific deployed contract.
func NewIATIOrganisationRecipientBudgetsCaller(address common.Address, caller bind.ContractCaller) (*IATIOrganisationRecipientBudgetsCaller, error) {
	contract, err := bindIATIOrganisationRecipientBudgets(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IATIOrganisationRecipientBudgetsCaller{contract: contract}, nil
}

// NewIATIOrganisationRecipientBudgetsTransactor creates a new write-only instance of IATIOrganisationRecipientBudgets, bound to a specific deployed contract.
func NewIATIOrganisationRecipientBudgetsTransactor(address common.Address, transactor bind.ContractTransactor) (*IATIOrganisationRecipientBudgetsTransactor, error) {
	contract, err := bindIATIOrganisationRecipientBudgets(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IATIOrganisationRecipientBudgetsTransactor{contract: contract}, nil
}

// NewIATIOrganisationRecipientBudgetsFilterer creates a new log filterer instance of IATIOrganisationRecipientBudgets, bound to a specific deployed contract.
func NewIATIOrganisationRecipientBudgetsFilterer(address common.Address, filterer bind.ContractFilterer) (*IATIOrganisationRecipientBudgetsFilterer, error) {
	contract, err := bindIATIOrganisationRecipientBudgets(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IATIOrganisationRecipientBudgetsFilterer{contract: contract}, nil
}

// bindIATIOrganisationRecipientBudgets binds a generic wrapper to an already deployed contract.
func bindIATIOrganisationRecipientBudgets(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIOrganisationRecipientBudgetsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIOrganisationRecipientBudgets.Contract.IATIOrganisationRecipientBudgetsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIOrganisationRecipientBudgets.Contract.IATIOrganisationRecipientBudgetsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIOrganisationRecipientBudgets.Contract.IATIOrganisationRecipientBudgetsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIOrganisationRecipientBudgets.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIOrganisationRecipientBudgets.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIOrganisationRecipientBudgets.Contract.contract.Transact(opts, method, params...)
}

// GetNumRecipientBudgets is a free data retrieval call binding the contract method 0x6bdc1aca.
//
// Solidity: function getNumRecipientBudgets(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(uint256)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCaller) GetNumRecipientBudgets(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIOrganisationRecipientBudgets.contract.Call(opts, out, "getNumRecipientBudgets", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetNumRecipientBudgets is a free data retrieval call binding the contract method 0x6bdc1aca.
//
// Solidity: function getNumRecipientBudgets(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(uint256)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsSession) GetNumRecipientBudgets(_organisationsRef [32]byte, _organisationRef [32]byte) (*big.Int, error) {
	return _IATIOrganisationRecipientBudgets.Contract.GetNumRecipientBudgets(&_IATIOrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef)
}

// GetNumRecipientBudgets is a free data retrieval call binding the contract method 0x6bdc1aca.
//
// Solidity: function getNumRecipientBudgets(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(uint256)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCallerSession) GetNumRecipientBudgets(_organisationsRef [32]byte, _organisationRef [32]byte) (*big.Int, error) {
	return _IATIOrganisationRecipientBudgets.Contract.GetNumRecipientBudgets(&_IATIOrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef)
}

// GetRecipientBudget is a free data retrieval call binding the contract method 0x6422072b.
//
// Solidity: function getRecipientBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(BudgetsBudget)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCaller) GetRecipientBudget(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) (BudgetsBudget, error) {
	var (
		ret0 = new(BudgetsBudget)
	)
	out := ret0
	err := _IATIOrganisationRecipientBudgets.contract.Call(opts, out, "getRecipientBudget", _organisationsRef, _organisationRef, _budgetRef)
	return *ret0, err
}

// GetRecipientBudget is a free data retrieval call binding the contract method 0x6422072b.
//
// Solidity: function getRecipientBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(BudgetsBudget)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsSession) GetRecipientBudget(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) (BudgetsBudget, error) {
	return _IATIOrganisationRecipientBudgets.Contract.GetRecipientBudget(&_IATIOrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudget is a free data retrieval call binding the contract method 0x6422072b.
//
// Solidity: function getRecipientBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(BudgetsBudget)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCallerSession) GetRecipientBudget(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) (BudgetsBudget, error) {
	return _IATIOrganisationRecipientBudgets.Contract.GetRecipientBudget(&_IATIOrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetEnd is a free data retrieval call binding the contract method 0xe3391566.
//
// Solidity: function getRecipientBudgetEnd(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCaller) GetRecipientBudgetEnd(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisationRecipientBudgets.contract.Call(opts, out, "getRecipientBudgetEnd", _organisationsRef, _organisationRef, _budgetRef)
	return *ret0, err
}

// GetRecipientBudgetEnd is a free data retrieval call binding the contract method 0xe3391566.
//
// Solidity: function getRecipientBudgetEnd(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsSession) GetRecipientBudgetEnd(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _IATIOrganisationRecipientBudgets.Contract.GetRecipientBudgetEnd(&_IATIOrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetEnd is a free data retrieval call binding the contract method 0xe3391566.
//
// Solidity: function getRecipientBudgetEnd(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCallerSession) GetRecipientBudgetEnd(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _IATIOrganisationRecipientBudgets.Contract.GetRecipientBudgetEnd(&_IATIOrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetLine is a free data retrieval call binding the contract method 0x2fc452f8.
//
// Solidity: function getRecipientBudgetLine(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCaller) GetRecipientBudgetLine(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisationRecipientBudgets.contract.Call(opts, out, "getRecipientBudgetLine", _organisationsRef, _organisationRef, _budgetRef)
	return *ret0, err
}

// GetRecipientBudgetLine is a free data retrieval call binding the contract method 0x2fc452f8.
//
// Solidity: function getRecipientBudgetLine(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsSession) GetRecipientBudgetLine(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _IATIOrganisationRecipientBudgets.Contract.GetRecipientBudgetLine(&_IATIOrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetLine is a free data retrieval call binding the contract method 0x2fc452f8.
//
// Solidity: function getRecipientBudgetLine(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCallerSession) GetRecipientBudgetLine(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _IATIOrganisationRecipientBudgets.Contract.GetRecipientBudgetLine(&_IATIOrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetOrg is a free data retrieval call binding the contract method 0x0629561d.
//
// Solidity: function getRecipientBudgetOrg(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCaller) GetRecipientBudgetOrg(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisationRecipientBudgets.contract.Call(opts, out, "getRecipientBudgetOrg", _organisationsRef, _organisationRef, _budgetRef)
	return *ret0, err
}

// GetRecipientBudgetOrg is a free data retrieval call binding the contract method 0x0629561d.
//
// Solidity: function getRecipientBudgetOrg(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsSession) GetRecipientBudgetOrg(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _IATIOrganisationRecipientBudgets.Contract.GetRecipientBudgetOrg(&_IATIOrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetOrg is a free data retrieval call binding the contract method 0x0629561d.
//
// Solidity: function getRecipientBudgetOrg(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCallerSession) GetRecipientBudgetOrg(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _IATIOrganisationRecipientBudgets.Contract.GetRecipientBudgetOrg(&_IATIOrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetReference is a free data retrieval call binding the contract method 0x41cd0fd0.
//
// Solidity: function getRecipientBudgetReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) view returns(bytes32)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCaller) GetRecipientBudgetReference(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisationRecipientBudgets.contract.Call(opts, out, "getRecipientBudgetReference", _organisationsRef, _organisationRef, _index)
	return *ret0, err
}

// GetRecipientBudgetReference is a free data retrieval call binding the contract method 0x41cd0fd0.
//
// Solidity: function getRecipientBudgetReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) view returns(bytes32)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsSession) GetRecipientBudgetReference(_organisationsRef [32]byte, _organisationRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIOrganisationRecipientBudgets.Contract.GetRecipientBudgetReference(&_IATIOrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _index)
}

// GetRecipientBudgetReference is a free data retrieval call binding the contract method 0x41cd0fd0.
//
// Solidity: function getRecipientBudgetReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) view returns(bytes32)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCallerSession) GetRecipientBudgetReference(_organisationsRef [32]byte, _organisationRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIOrganisationRecipientBudgets.Contract.GetRecipientBudgetReference(&_IATIOrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _index)
}

// GetRecipientBudgetStart is a free data retrieval call binding the contract method 0x999af84d.
//
// Solidity: function getRecipientBudgetStart(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCaller) GetRecipientBudgetStart(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisationRecipientBudgets.contract.Call(opts, out, "getRecipientBudgetStart", _organisationsRef, _organisationRef, _budgetRef)
	return *ret0, err
}

// GetRecipientBudgetStart is a free data retrieval call binding the contract method 0x999af84d.
//
// Solidity: function getRecipientBudgetStart(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsSession) GetRecipientBudgetStart(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _IATIOrganisationRecipientBudgets.Contract.GetRecipientBudgetStart(&_IATIOrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetStart is a free data retrieval call binding the contract method 0x999af84d.
//
// Solidity: function getRecipientBudgetStart(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCallerSession) GetRecipientBudgetStart(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _IATIOrganisationRecipientBudgets.Contract.GetRecipientBudgetStart(&_IATIOrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetStatus is a free data retrieval call binding the contract method 0x98e2d588.
//
// Solidity: function getRecipientBudgetStatus(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(uint8)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCaller) GetRecipientBudgetStatus(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIOrganisationRecipientBudgets.contract.Call(opts, out, "getRecipientBudgetStatus", _organisationsRef, _organisationRef, _budgetRef)
	return *ret0, err
}

// GetRecipientBudgetStatus is a free data retrieval call binding the contract method 0x98e2d588.
//
// Solidity: function getRecipientBudgetStatus(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(uint8)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsSession) GetRecipientBudgetStatus(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) (uint8, error) {
	return _IATIOrganisationRecipientBudgets.Contract.GetRecipientBudgetStatus(&_IATIOrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetStatus is a free data retrieval call binding the contract method 0x98e2d588.
//
// Solidity: function getRecipientBudgetStatus(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(uint8)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCallerSession) GetRecipientBudgetStatus(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) (uint8, error) {
	return _IATIOrganisationRecipientBudgets.Contract.GetRecipientBudgetStatus(&_IATIOrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetValue is a free data retrieval call binding the contract method 0xe8ae58f0.
//
// Solidity: function getRecipientBudgetValue(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(uint256)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCaller) GetRecipientBudgetValue(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIOrganisationRecipientBudgets.contract.Call(opts, out, "getRecipientBudgetValue", _organisationsRef, _organisationRef, _budgetRef)
	return *ret0, err
}

// GetRecipientBudgetValue is a free data retrieval call binding the contract method 0xe8ae58f0.
//
// Solidity: function getRecipientBudgetValue(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(uint256)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsSession) GetRecipientBudgetValue(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) (*big.Int, error) {
	return _IATIOrganisationRecipientBudgets.Contract.GetRecipientBudgetValue(&_IATIOrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetValue is a free data retrieval call binding the contract method 0xe8ae58f0.
//
// Solidity: function getRecipientBudgetValue(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(uint256)
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsCallerSession) GetRecipientBudgetValue(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) (*big.Int, error) {
	return _IATIOrganisationRecipientBudgets.Contract.GetRecipientBudgetValue(&_IATIOrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// SetRecipientBudget is a paid mutator transaction binding the contract method 0x497c6d69.
//
// Solidity: function setRecipientBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef, BudgetsBudget _budget) returns()
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsTransactor) SetRecipientBudget(opts *bind.TransactOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte, _budget BudgetsBudget) (*types.Transaction, error) {
	return _IATIOrganisationRecipientBudgets.contract.Transact(opts, "setRecipientBudget", _organisationsRef, _organisationRef, _budgetRef, _budget)
}

// SetRecipientBudget is a paid mutator transaction binding the contract method 0x497c6d69.
//
// Solidity: function setRecipientBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef, BudgetsBudget _budget) returns()
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsSession) SetRecipientBudget(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte, _budget BudgetsBudget) (*types.Transaction, error) {
	return _IATIOrganisationRecipientBudgets.Contract.SetRecipientBudget(&_IATIOrganisationRecipientBudgets.TransactOpts, _organisationsRef, _organisationRef, _budgetRef, _budget)
}

// SetRecipientBudget is a paid mutator transaction binding the contract method 0x497c6d69.
//
// Solidity: function setRecipientBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef, BudgetsBudget _budget) returns()
func (_IATIOrganisationRecipientBudgets *IATIOrganisationRecipientBudgetsTransactorSession) SetRecipientBudget(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte, _budget BudgetsBudget) (*types.Transaction, error) {
	return _IATIOrganisationRecipientBudgets.Contract.SetRecipientBudget(&_IATIOrganisationRecipientBudgets.TransactOpts, _organisationsRef, _organisationRef, _budgetRef, _budget)
}

// OrganisationRecipientBudgetsABI is the input ABI used to generate the binding from.
const OrganisationRecipientBudgetsABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getNumRecipientBudgets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getRecipientBudget\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"budgetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"budgetLine\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"otherRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"start\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"end\",\"type\":\"bytes32\"}],\"internalType\":\"structBudgets.Finance\",\"name\":\"finance\",\"type\":\"tuple\"}],\"internalType\":\"structBudgets.Budget\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getRecipientBudgetEnd\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getRecipientBudgetLine\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getRecipientBudgetOrg\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRecipientBudgetReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getRecipientBudgetStart\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getRecipientBudgetStatus\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"}],\"name\":\"getRecipientBudgetValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_budgetRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"budgetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"budgetLine\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"otherRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"start\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"end\",\"type\":\"bytes32\"}],\"internalType\":\"structBudgets.Finance\",\"name\":\"finance\",\"type\":\"tuple\"}],\"internalType\":\"structBudgets.Budget\",\"name\":\"_budget\",\"type\":\"tuple\"}],\"name\":\"setRecipientBudget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OrganisationRecipientBudgetsFuncSigs maps the 4-byte function signature to its string representation.
var OrganisationRecipientBudgetsFuncSigs = map[string]string{
	"6bdc1aca": "getNumRecipientBudgets(bytes32,bytes32)",
	"6422072b": "getRecipientBudget(bytes32,bytes32,bytes32)",
	"e3391566": "getRecipientBudgetEnd(bytes32,bytes32,bytes32)",
	"2fc452f8": "getRecipientBudgetLine(bytes32,bytes32,bytes32)",
	"0629561d": "getRecipientBudgetOrg(bytes32,bytes32,bytes32)",
	"41cd0fd0": "getRecipientBudgetReference(bytes32,bytes32,uint256)",
	"999af84d": "getRecipientBudgetStart(bytes32,bytes32,bytes32)",
	"98e2d588": "getRecipientBudgetStatus(bytes32,bytes32,bytes32)",
	"e8ae58f0": "getRecipientBudgetValue(bytes32,bytes32,bytes32)",
	"497c6d69": "setRecipientBudget(bytes32,bytes32,bytes32,(uint8,bytes32,bytes32,(uint8,uint256,bytes32,bytes32)))",
}

// OrganisationRecipientBudgets is an auto generated Go binding around an Ethereum contract.
type OrganisationRecipientBudgets struct {
	OrganisationRecipientBudgetsCaller     // Read-only binding to the contract
	OrganisationRecipientBudgetsTransactor // Write-only binding to the contract
	OrganisationRecipientBudgetsFilterer   // Log filterer for contract events
}

// OrganisationRecipientBudgetsCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrganisationRecipientBudgetsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganisationRecipientBudgetsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrganisationRecipientBudgetsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganisationRecipientBudgetsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrganisationRecipientBudgetsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganisationRecipientBudgetsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrganisationRecipientBudgetsSession struct {
	Contract     *OrganisationRecipientBudgets // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OrganisationRecipientBudgetsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrganisationRecipientBudgetsCallerSession struct {
	Contract *OrganisationRecipientBudgetsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// OrganisationRecipientBudgetsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrganisationRecipientBudgetsTransactorSession struct {
	Contract     *OrganisationRecipientBudgetsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// OrganisationRecipientBudgetsRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrganisationRecipientBudgetsRaw struct {
	Contract *OrganisationRecipientBudgets // Generic contract binding to access the raw methods on
}

// OrganisationRecipientBudgetsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrganisationRecipientBudgetsCallerRaw struct {
	Contract *OrganisationRecipientBudgetsCaller // Generic read-only contract binding to access the raw methods on
}

// OrganisationRecipientBudgetsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrganisationRecipientBudgetsTransactorRaw struct {
	Contract *OrganisationRecipientBudgetsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrganisationRecipientBudgets creates a new instance of OrganisationRecipientBudgets, bound to a specific deployed contract.
func NewOrganisationRecipientBudgets(address common.Address, backend bind.ContractBackend) (*OrganisationRecipientBudgets, error) {
	contract, err := bindOrganisationRecipientBudgets(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrganisationRecipientBudgets{OrganisationRecipientBudgetsCaller: OrganisationRecipientBudgetsCaller{contract: contract}, OrganisationRecipientBudgetsTransactor: OrganisationRecipientBudgetsTransactor{contract: contract}, OrganisationRecipientBudgetsFilterer: OrganisationRecipientBudgetsFilterer{contract: contract}}, nil
}

// NewOrganisationRecipientBudgetsCaller creates a new read-only instance of OrganisationRecipientBudgets, bound to a specific deployed contract.
func NewOrganisationRecipientBudgetsCaller(address common.Address, caller bind.ContractCaller) (*OrganisationRecipientBudgetsCaller, error) {
	contract, err := bindOrganisationRecipientBudgets(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrganisationRecipientBudgetsCaller{contract: contract}, nil
}

// NewOrganisationRecipientBudgetsTransactor creates a new write-only instance of OrganisationRecipientBudgets, bound to a specific deployed contract.
func NewOrganisationRecipientBudgetsTransactor(address common.Address, transactor bind.ContractTransactor) (*OrganisationRecipientBudgetsTransactor, error) {
	contract, err := bindOrganisationRecipientBudgets(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrganisationRecipientBudgetsTransactor{contract: contract}, nil
}

// NewOrganisationRecipientBudgetsFilterer creates a new log filterer instance of OrganisationRecipientBudgets, bound to a specific deployed contract.
func NewOrganisationRecipientBudgetsFilterer(address common.Address, filterer bind.ContractFilterer) (*OrganisationRecipientBudgetsFilterer, error) {
	contract, err := bindOrganisationRecipientBudgets(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrganisationRecipientBudgetsFilterer{contract: contract}, nil
}

// bindOrganisationRecipientBudgets binds a generic wrapper to an already deployed contract.
func bindOrganisationRecipientBudgets(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrganisationRecipientBudgetsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OrganisationRecipientBudgets.Contract.OrganisationRecipientBudgetsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrganisationRecipientBudgets.Contract.OrganisationRecipientBudgetsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrganisationRecipientBudgets.Contract.OrganisationRecipientBudgetsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OrganisationRecipientBudgets.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrganisationRecipientBudgets.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrganisationRecipientBudgets.Contract.contract.Transact(opts, method, params...)
}

// GetNumRecipientBudgets is a free data retrieval call binding the contract method 0x6bdc1aca.
//
// Solidity: function getNumRecipientBudgets(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(uint256)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCaller) GetNumRecipientBudgets(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OrganisationRecipientBudgets.contract.Call(opts, out, "getNumRecipientBudgets", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetNumRecipientBudgets is a free data retrieval call binding the contract method 0x6bdc1aca.
//
// Solidity: function getNumRecipientBudgets(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(uint256)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsSession) GetNumRecipientBudgets(_organisationsRef [32]byte, _organisationRef [32]byte) (*big.Int, error) {
	return _OrganisationRecipientBudgets.Contract.GetNumRecipientBudgets(&_OrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef)
}

// GetNumRecipientBudgets is a free data retrieval call binding the contract method 0x6bdc1aca.
//
// Solidity: function getNumRecipientBudgets(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(uint256)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCallerSession) GetNumRecipientBudgets(_organisationsRef [32]byte, _organisationRef [32]byte) (*big.Int, error) {
	return _OrganisationRecipientBudgets.Contract.GetNumRecipientBudgets(&_OrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef)
}

// GetRecipientBudget is a free data retrieval call binding the contract method 0x6422072b.
//
// Solidity: function getRecipientBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(BudgetsBudget)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCaller) GetRecipientBudget(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) (BudgetsBudget, error) {
	var (
		ret0 = new(BudgetsBudget)
	)
	out := ret0
	err := _OrganisationRecipientBudgets.contract.Call(opts, out, "getRecipientBudget", _organisationsRef, _organisationRef, _budgetRef)
	return *ret0, err
}

// GetRecipientBudget is a free data retrieval call binding the contract method 0x6422072b.
//
// Solidity: function getRecipientBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(BudgetsBudget)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsSession) GetRecipientBudget(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) (BudgetsBudget, error) {
	return _OrganisationRecipientBudgets.Contract.GetRecipientBudget(&_OrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudget is a free data retrieval call binding the contract method 0x6422072b.
//
// Solidity: function getRecipientBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(BudgetsBudget)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCallerSession) GetRecipientBudget(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) (BudgetsBudget, error) {
	return _OrganisationRecipientBudgets.Contract.GetRecipientBudget(&_OrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetEnd is a free data retrieval call binding the contract method 0xe3391566.
//
// Solidity: function getRecipientBudgetEnd(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCaller) GetRecipientBudgetEnd(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OrganisationRecipientBudgets.contract.Call(opts, out, "getRecipientBudgetEnd", _organisationsRef, _organisationRef, _budgetRef)
	return *ret0, err
}

// GetRecipientBudgetEnd is a free data retrieval call binding the contract method 0xe3391566.
//
// Solidity: function getRecipientBudgetEnd(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsSession) GetRecipientBudgetEnd(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _OrganisationRecipientBudgets.Contract.GetRecipientBudgetEnd(&_OrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetEnd is a free data retrieval call binding the contract method 0xe3391566.
//
// Solidity: function getRecipientBudgetEnd(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCallerSession) GetRecipientBudgetEnd(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _OrganisationRecipientBudgets.Contract.GetRecipientBudgetEnd(&_OrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetLine is a free data retrieval call binding the contract method 0x2fc452f8.
//
// Solidity: function getRecipientBudgetLine(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCaller) GetRecipientBudgetLine(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OrganisationRecipientBudgets.contract.Call(opts, out, "getRecipientBudgetLine", _organisationsRef, _organisationRef, _budgetRef)
	return *ret0, err
}

// GetRecipientBudgetLine is a free data retrieval call binding the contract method 0x2fc452f8.
//
// Solidity: function getRecipientBudgetLine(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsSession) GetRecipientBudgetLine(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _OrganisationRecipientBudgets.Contract.GetRecipientBudgetLine(&_OrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetLine is a free data retrieval call binding the contract method 0x2fc452f8.
//
// Solidity: function getRecipientBudgetLine(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCallerSession) GetRecipientBudgetLine(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _OrganisationRecipientBudgets.Contract.GetRecipientBudgetLine(&_OrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetOrg is a free data retrieval call binding the contract method 0x0629561d.
//
// Solidity: function getRecipientBudgetOrg(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCaller) GetRecipientBudgetOrg(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OrganisationRecipientBudgets.contract.Call(opts, out, "getRecipientBudgetOrg", _organisationsRef, _organisationRef, _budgetRef)
	return *ret0, err
}

// GetRecipientBudgetOrg is a free data retrieval call binding the contract method 0x0629561d.
//
// Solidity: function getRecipientBudgetOrg(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsSession) GetRecipientBudgetOrg(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _OrganisationRecipientBudgets.Contract.GetRecipientBudgetOrg(&_OrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetOrg is a free data retrieval call binding the contract method 0x0629561d.
//
// Solidity: function getRecipientBudgetOrg(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCallerSession) GetRecipientBudgetOrg(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _OrganisationRecipientBudgets.Contract.GetRecipientBudgetOrg(&_OrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetReference is a free data retrieval call binding the contract method 0x41cd0fd0.
//
// Solidity: function getRecipientBudgetReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) view returns(bytes32)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCaller) GetRecipientBudgetReference(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OrganisationRecipientBudgets.contract.Call(opts, out, "getRecipientBudgetReference", _organisationsRef, _organisationRef, _index)
	return *ret0, err
}

// GetRecipientBudgetReference is a free data retrieval call binding the contract method 0x41cd0fd0.
//
// Solidity: function getRecipientBudgetReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) view returns(bytes32)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsSession) GetRecipientBudgetReference(_organisationsRef [32]byte, _organisationRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _OrganisationRecipientBudgets.Contract.GetRecipientBudgetReference(&_OrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _index)
}

// GetRecipientBudgetReference is a free data retrieval call binding the contract method 0x41cd0fd0.
//
// Solidity: function getRecipientBudgetReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) view returns(bytes32)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCallerSession) GetRecipientBudgetReference(_organisationsRef [32]byte, _organisationRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _OrganisationRecipientBudgets.Contract.GetRecipientBudgetReference(&_OrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _index)
}

// GetRecipientBudgetStart is a free data retrieval call binding the contract method 0x999af84d.
//
// Solidity: function getRecipientBudgetStart(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCaller) GetRecipientBudgetStart(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OrganisationRecipientBudgets.contract.Call(opts, out, "getRecipientBudgetStart", _organisationsRef, _organisationRef, _budgetRef)
	return *ret0, err
}

// GetRecipientBudgetStart is a free data retrieval call binding the contract method 0x999af84d.
//
// Solidity: function getRecipientBudgetStart(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsSession) GetRecipientBudgetStart(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _OrganisationRecipientBudgets.Contract.GetRecipientBudgetStart(&_OrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetStart is a free data retrieval call binding the contract method 0x999af84d.
//
// Solidity: function getRecipientBudgetStart(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(bytes32)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCallerSession) GetRecipientBudgetStart(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) ([32]byte, error) {
	return _OrganisationRecipientBudgets.Contract.GetRecipientBudgetStart(&_OrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetStatus is a free data retrieval call binding the contract method 0x98e2d588.
//
// Solidity: function getRecipientBudgetStatus(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(uint8)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCaller) GetRecipientBudgetStatus(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _OrganisationRecipientBudgets.contract.Call(opts, out, "getRecipientBudgetStatus", _organisationsRef, _organisationRef, _budgetRef)
	return *ret0, err
}

// GetRecipientBudgetStatus is a free data retrieval call binding the contract method 0x98e2d588.
//
// Solidity: function getRecipientBudgetStatus(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(uint8)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsSession) GetRecipientBudgetStatus(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) (uint8, error) {
	return _OrganisationRecipientBudgets.Contract.GetRecipientBudgetStatus(&_OrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetStatus is a free data retrieval call binding the contract method 0x98e2d588.
//
// Solidity: function getRecipientBudgetStatus(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(uint8)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCallerSession) GetRecipientBudgetStatus(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) (uint8, error) {
	return _OrganisationRecipientBudgets.Contract.GetRecipientBudgetStatus(&_OrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetValue is a free data retrieval call binding the contract method 0xe8ae58f0.
//
// Solidity: function getRecipientBudgetValue(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(uint256)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCaller) GetRecipientBudgetValue(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OrganisationRecipientBudgets.contract.Call(opts, out, "getRecipientBudgetValue", _organisationsRef, _organisationRef, _budgetRef)
	return *ret0, err
}

// GetRecipientBudgetValue is a free data retrieval call binding the contract method 0xe8ae58f0.
//
// Solidity: function getRecipientBudgetValue(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(uint256)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsSession) GetRecipientBudgetValue(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) (*big.Int, error) {
	return _OrganisationRecipientBudgets.Contract.GetRecipientBudgetValue(&_OrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// GetRecipientBudgetValue is a free data retrieval call binding the contract method 0xe8ae58f0.
//
// Solidity: function getRecipientBudgetValue(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) view returns(uint256)
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsCallerSession) GetRecipientBudgetValue(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte) (*big.Int, error) {
	return _OrganisationRecipientBudgets.Contract.GetRecipientBudgetValue(&_OrganisationRecipientBudgets.CallOpts, _organisationsRef, _organisationRef, _budgetRef)
}

// SetRecipientBudget is a paid mutator transaction binding the contract method 0x497c6d69.
//
// Solidity: function setRecipientBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef, BudgetsBudget _budget) returns()
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsTransactor) SetRecipientBudget(opts *bind.TransactOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte, _budget BudgetsBudget) (*types.Transaction, error) {
	return _OrganisationRecipientBudgets.contract.Transact(opts, "setRecipientBudget", _organisationsRef, _organisationRef, _budgetRef, _budget)
}

// SetRecipientBudget is a paid mutator transaction binding the contract method 0x497c6d69.
//
// Solidity: function setRecipientBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef, BudgetsBudget _budget) returns()
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsSession) SetRecipientBudget(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte, _budget BudgetsBudget) (*types.Transaction, error) {
	return _OrganisationRecipientBudgets.Contract.SetRecipientBudget(&_OrganisationRecipientBudgets.TransactOpts, _organisationsRef, _organisationRef, _budgetRef, _budget)
}

// SetRecipientBudget is a paid mutator transaction binding the contract method 0x497c6d69.
//
// Solidity: function setRecipientBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef, BudgetsBudget _budget) returns()
func (_OrganisationRecipientBudgets *OrganisationRecipientBudgetsTransactorSession) SetRecipientBudget(_organisationsRef [32]byte, _organisationRef [32]byte, _budgetRef [32]byte, _budget BudgetsBudget) (*types.Transaction, error) {
	return _OrganisationRecipientBudgets.Contract.SetRecipientBudget(&_OrganisationRecipientBudgets.TransactOpts, _organisationsRef, _organisationRef, _budgetRef, _budget)
}

// StringsABI is the input ABI used to generate the binding from.
const StringsABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_x\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToStr\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_a\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_b\",\"type\":\"bytes32\"}],\"name\":\"compare\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_a\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_b\",\"type\":\"string\"}],\"name\":\"compare\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_xs\",\"type\":\"bytes32[]\"}],\"name\":\"getExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_a\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_xs\",\"type\":\"string[]\"}],\"name\":\"getIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_xs\",\"type\":\"bytes32[]\"}],\"name\":\"getIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// StringsFuncSigs maps the 4-byte function signature to its string representation.
var StringsFuncSigs = map[string]string{
	"ef0b2368": "bytes32ToStr(bytes32)",
	"28602ab2": "compare(bytes32,bytes32)",
	"3a96fdd7": "compare(string,string)",
	"745fca7b": "getExists(bytes32,bytes32[])",
	"c55f207f": "getIndex(bytes32,bytes32[])",
	"02a33464": "getIndex(string,string[])",
}

// StringsBin is the compiled bytecode used for deploying new contracts.
var StringsBin = "0x610662610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061006c5760003560e01c806302a334641461007157806328602ab21461009a5780633a96fdd7146100ba578063745fca7b146100cd578063c55f207f146100e0578063ef0b2368146100f3575b600080fd5b61008461007f36600461046d565b610113565b60405161009191906105dc565b60405180910390f35b6100ad6100a836600461044c565b610180565b604051610091919061057e565b6100ad6100c836600461051d565b6101bd565b6100ad6100db3660046103a8565b61020b565b6100846100ee3660046103a8565b61024b565b610106610101366004610390565b6102b5565b6040516100919190610589565b6000808351118015610126575060008251115b61012f57600080fd5b815160005b835181101561017657606084828151811061014b57fe5b6020026020010151905061015f86826101bd565b1561016d5781925050610176565b50600101610134565b5090505b92915050565b600082811a60f81b6001600160f81b031916158015906101af57508160001a60f81b6001600160f81b03191615155b6101b857600080fd5b501490565b60008083511180156101d0575060008251115b6101d957600080fd5b8151835184918491146101f15760009250505061017a565b80805190602001208280519060200120149250505061017a565b600082811a60f81b6001600160f81b03191661022657600080fd5b81516000901561024457600061023c858561024b565b845114159150505b9392505050565b600082811a60f81b6001600160f81b0319161580159061026c575060008251115b61027557600080fd5b815160005b8351811015610176576102a08585838151811061029357fe5b6020026020010151610180565b156102ad57809150610176565b60010161027a565b60408051602080825281830190925260609182919060208201818036833701905050905060005b6020811015610320578381602081106102f157fe5b1a60f81b82828151811061030157fe5b60200101906001600160f81b031916908160001a9053506001016102dc565b5092915050565b600082601f830112610337578081fd5b813567ffffffffffffffff81111561034d578182fd5b610360601f8201601f19166020016105e5565b915080825283602082850101111561037757600080fd5b8060208401602084013760009082016020015292915050565b6000602082840312156103a1578081fd5b5035919050565b600080604083850312156103ba578081fd5b8235915060208084013567ffffffffffffffff8111156103d8578283fd5b80850186601f8201126103e9578384fd5b803591506103fe6103f98361060c565b6105e5565b82815283810190828501858502840186018a101561041a578687fd5b8693505b8484101561043c57803583526001939093019291850191850161041e565b5080955050505050509250929050565b6000806040838503121561045e578182fd5b50508035926020909101359150565b6000806040838503121561047f578182fd5b823567ffffffffffffffff80821115610496578384fd5b6104a286838701610327565b93506020915081850135818111156104b8578384fd5b85019050601f810186136104ca578283fd5b80356104d86103f98261060c565b81815283810190838501865b8481101561050d576104fb8b888435890101610327565b845292860192908601906001016104e4565b5096999098509650505050505050565b6000806040838503121561052f578182fd5b823567ffffffffffffffff80821115610546578384fd5b61055286838701610327565b93506020850135915080821115610567578283fd5b5061057485828601610327565b9150509250929050565b901515815260200190565b6000602080835283518082850152825b818110156105b557858101830151858201604001528201610599565b818111156105c65783604083870101525b50601f01601f1916929092016040019392505050565b90815260200190565b60405181810167ffffffffffffffff8111828210171561060457600080fd5b604052919050565b600067ffffffffffffffff821115610622578081fd5b506020908102019056fea2646970667358221220afe80f7b79e21558518987b567fad792efadb78d8c923a4d1825b3998ada86f464736f6c63430006060033"

// DeployStrings deploys a new Ethereum contract, binding an instance of Strings to it.
func DeployStrings(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Strings, error) {
	parsed, err := abi.JSON(strings.NewReader(StringsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StringsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// Strings is an auto generated Go binding around an Ethereum contract.
type Strings struct {
	StringsCaller     // Read-only binding to the contract
	StringsTransactor // Write-only binding to the contract
	StringsFilterer   // Log filterer for contract events
}

// StringsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StringsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StringsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StringsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StringsSession struct {
	Contract     *Strings          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StringsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StringsCallerSession struct {
	Contract *StringsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StringsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StringsTransactorSession struct {
	Contract     *StringsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StringsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StringsRaw struct {
	Contract *Strings // Generic contract binding to access the raw methods on
}

// StringsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StringsCallerRaw struct {
	Contract *StringsCaller // Generic read-only contract binding to access the raw methods on
}

// StringsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StringsTransactorRaw struct {
	Contract *StringsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrings creates a new instance of Strings, bound to a specific deployed contract.
func NewStrings(address common.Address, backend bind.ContractBackend) (*Strings, error) {
	contract, err := bindStrings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// NewStringsCaller creates a new read-only instance of Strings, bound to a specific deployed contract.
func NewStringsCaller(address common.Address, caller bind.ContractCaller) (*StringsCaller, error) {
	contract, err := bindStrings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StringsCaller{contract: contract}, nil
}

// NewStringsTransactor creates a new write-only instance of Strings, bound to a specific deployed contract.
func NewStringsTransactor(address common.Address, transactor bind.ContractTransactor) (*StringsTransactor, error) {
	contract, err := bindStrings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StringsTransactor{contract: contract}, nil
}

// NewStringsFilterer creates a new log filterer instance of Strings, bound to a specific deployed contract.
func NewStringsFilterer(address common.Address, filterer bind.ContractFilterer) (*StringsFilterer, error) {
	contract, err := bindStrings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StringsFilterer{contract: contract}, nil
}

// bindStrings binds a generic wrapper to an already deployed contract.
func bindStrings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StringsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.StringsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transact(opts, method, params...)
}

// Bytes32ToStr is a free data retrieval call binding the contract method 0xef0b2368.
//
// Solidity: function bytes32ToStr(bytes32 _x) pure returns(string)
func (_Strings *StringsCaller) Bytes32ToStr(opts *bind.CallOpts, _x [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Strings.contract.Call(opts, out, "bytes32ToStr", _x)
	return *ret0, err
}

// Bytes32ToStr is a free data retrieval call binding the contract method 0xef0b2368.
//
// Solidity: function bytes32ToStr(bytes32 _x) pure returns(string)
func (_Strings *StringsSession) Bytes32ToStr(_x [32]byte) (string, error) {
	return _Strings.Contract.Bytes32ToStr(&_Strings.CallOpts, _x)
}

// Bytes32ToStr is a free data retrieval call binding the contract method 0xef0b2368.
//
// Solidity: function bytes32ToStr(bytes32 _x) pure returns(string)
func (_Strings *StringsCallerSession) Bytes32ToStr(_x [32]byte) (string, error) {
	return _Strings.Contract.Bytes32ToStr(&_Strings.CallOpts, _x)
}

// Compare is a free data retrieval call binding the contract method 0x28602ab2.
//
// Solidity: function compare(bytes32 _a, bytes32 _b) pure returns(bool)
func (_Strings *StringsCaller) Compare(opts *bind.CallOpts, _a [32]byte, _b [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Strings.contract.Call(opts, out, "compare", _a, _b)
	return *ret0, err
}

// Compare is a free data retrieval call binding the contract method 0x28602ab2.
//
// Solidity: function compare(bytes32 _a, bytes32 _b) pure returns(bool)
func (_Strings *StringsSession) Compare(_a [32]byte, _b [32]byte) (bool, error) {
	return _Strings.Contract.Compare(&_Strings.CallOpts, _a, _b)
}

// Compare is a free data retrieval call binding the contract method 0x28602ab2.
//
// Solidity: function compare(bytes32 _a, bytes32 _b) pure returns(bool)
func (_Strings *StringsCallerSession) Compare(_a [32]byte, _b [32]byte) (bool, error) {
	return _Strings.Contract.Compare(&_Strings.CallOpts, _a, _b)
}

// Compare0 is a free data retrieval call binding the contract method 0x3a96fdd7.
//
// Solidity: function compare(string _a, string _b) pure returns(bool)
func (_Strings *StringsCaller) Compare0(opts *bind.CallOpts, _a string, _b string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Strings.contract.Call(opts, out, "compare0", _a, _b)
	return *ret0, err
}

// Compare0 is a free data retrieval call binding the contract method 0x3a96fdd7.
//
// Solidity: function compare(string _a, string _b) pure returns(bool)
func (_Strings *StringsSession) Compare0(_a string, _b string) (bool, error) {
	return _Strings.Contract.Compare0(&_Strings.CallOpts, _a, _b)
}

// Compare0 is a free data retrieval call binding the contract method 0x3a96fdd7.
//
// Solidity: function compare(string _a, string _b) pure returns(bool)
func (_Strings *StringsCallerSession) Compare0(_a string, _b string) (bool, error) {
	return _Strings.Contract.Compare0(&_Strings.CallOpts, _a, _b)
}

// GetExists is a free data retrieval call binding the contract method 0x745fca7b.
//
// Solidity: function getExists(bytes32 _x, bytes32[] _xs) pure returns(bool)
func (_Strings *StringsCaller) GetExists(opts *bind.CallOpts, _x [32]byte, _xs [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Strings.contract.Call(opts, out, "getExists", _x, _xs)
	return *ret0, err
}

// GetExists is a free data retrieval call binding the contract method 0x745fca7b.
//
// Solidity: function getExists(bytes32 _x, bytes32[] _xs) pure returns(bool)
func (_Strings *StringsSession) GetExists(_x [32]byte, _xs [][32]byte) (bool, error) {
	return _Strings.Contract.GetExists(&_Strings.CallOpts, _x, _xs)
}

// GetExists is a free data retrieval call binding the contract method 0x745fca7b.
//
// Solidity: function getExists(bytes32 _x, bytes32[] _xs) pure returns(bool)
func (_Strings *StringsCallerSession) GetExists(_x [32]byte, _xs [][32]byte) (bool, error) {
	return _Strings.Contract.GetExists(&_Strings.CallOpts, _x, _xs)
}

// GetIndex is a free data retrieval call binding the contract method 0x02a33464.
//
// Solidity: function getIndex(string _a, string[] _xs) pure returns(uint256)
func (_Strings *StringsCaller) GetIndex(opts *bind.CallOpts, _a string, _xs []string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Strings.contract.Call(opts, out, "getIndex", _a, _xs)
	return *ret0, err
}

// GetIndex is a free data retrieval call binding the contract method 0x02a33464.
//
// Solidity: function getIndex(string _a, string[] _xs) pure returns(uint256)
func (_Strings *StringsSession) GetIndex(_a string, _xs []string) (*big.Int, error) {
	return _Strings.Contract.GetIndex(&_Strings.CallOpts, _a, _xs)
}

// GetIndex is a free data retrieval call binding the contract method 0x02a33464.
//
// Solidity: function getIndex(string _a, string[] _xs) pure returns(uint256)
func (_Strings *StringsCallerSession) GetIndex(_a string, _xs []string) (*big.Int, error) {
	return _Strings.Contract.GetIndex(&_Strings.CallOpts, _a, _xs)
}

// GetIndex0 is a free data retrieval call binding the contract method 0xc55f207f.
//
// Solidity: function getIndex(bytes32 _x, bytes32[] _xs) pure returns(uint256)
func (_Strings *StringsCaller) GetIndex0(opts *bind.CallOpts, _x [32]byte, _xs [][32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Strings.contract.Call(opts, out, "getIndex0", _x, _xs)
	return *ret0, err
}

// GetIndex0 is a free data retrieval call binding the contract method 0xc55f207f.
//
// Solidity: function getIndex(bytes32 _x, bytes32[] _xs) pure returns(uint256)
func (_Strings *StringsSession) GetIndex0(_x [32]byte, _xs [][32]byte) (*big.Int, error) {
	return _Strings.Contract.GetIndex0(&_Strings.CallOpts, _x, _xs)
}

// GetIndex0 is a free data retrieval call binding the contract method 0xc55f207f.
//
// Solidity: function getIndex(bytes32 _x, bytes32[] _xs) pure returns(uint256)
func (_Strings *StringsCallerSession) GetIndex0(_x [32]byte, _xs [][32]byte) (*big.Int, error) {
	return _Strings.Contract.GetIndex0(&_Strings.CallOpts, _x, _xs)
}
