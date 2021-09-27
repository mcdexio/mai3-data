// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reader

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

// ReaderAccountReaderResult is an auto generated low-level Go binding around an user-defined struct.
type ReaderAccountReaderResult struct {
	Cash                    *big.Int
	Position                *big.Int
	AvailableMargin         *big.Int
	Margin                  *big.Int
	SettleableMargin        *big.Int
	IsInitialMarginSafe     bool
	IsMaintenanceMarginSafe bool
	IsMarginSafe            bool
	TargetLeverage          *big.Int
}

// ReaderAccountsResult is an auto generated low-level Go binding around an user-defined struct.
type ReaderAccountsResult struct {
	Account       common.Address
	Position      *big.Int
	Margin        *big.Int
	IsSafe        bool
	AvailableCash *big.Int
}

// ReaderLiquidityPoolReaderResult is an auto generated low-level Go binding around an user-defined struct.
type ReaderLiquidityPoolReaderResult struct {
	IsRunning             bool
	IsFastCreationEnabled bool
	Addresses             [7]common.Address
	IntNums               [5]*big.Int
	UintNums              [6]*big.Int
	Perpetuals            []ReaderPerpetualReaderResult
	IsAMMMaintenanceSafe  bool
}

// ReaderPerpetualReaderResult is an auto generated low-level Go binding around an user-defined struct.
type ReaderPerpetualReaderResult struct {
	State              uint8
	Oracle             common.Address
	Nums               [39]*big.Int
	Symbol             *big.Int
	UnderlyingAsset    string
	IsMarketClosed     bool
	IsTerminated       bool
	AmmCashBalance     *big.Int
	AmmPositionAmount  *big.Int
	IsInversePerpetual bool
}

// ReaderMetaData contains all meta data concerning the Reader contract.
var ReaderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolCreator_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"inverseStateService_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountStorage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSynced\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"cash\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"availableMargin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"margin\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"settleableMargin\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"isInitialMarginSafe\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isMaintenanceMarginSafe\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isMarginSafe\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"targetLeverage\",\"type\":\"int256\"}],\"internalType\":\"structReader.AccountReaderResult\",\"name\":\"accountStorage\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getAccountsInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSynced\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"margin\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"isSafe\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"availableCash\",\"type\":\"int256\"}],\"internalType\":\"structReader.AccountsResult[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"getAccountsInfoByAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSynced\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"position\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"margin\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"isSafe\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"availableCash\",\"type\":\"int256\"}],\"internalType\":\"structReader.AccountsResult[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getL1BlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidityPool\",\"type\":\"address\"}],\"name\":\"getLiquidityPoolStorage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSynced\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isRunning\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isFastCreationEnabled\",\"type\":\"bool\"},{\"internalType\":\"address[7]\",\"name\":\"addresses\",\"type\":\"address[7]\"},{\"internalType\":\"int256[5]\",\"name\":\"intNums\",\"type\":\"int256[5]\"},{\"internalType\":\"uint256[6]\",\"name\":\"uintNums\",\"type\":\"uint256[6]\"},{\"components\":[{\"internalType\":\"enumPerpetualState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"int256[39]\",\"name\":\"nums\",\"type\":\"int256[39]\"},{\"internalType\":\"uint256\",\"name\":\"symbol\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"underlyingAsset\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isMarketClosed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isTerminated\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"ammCashBalance\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"ammPositionAmount\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"isInversePerpetual\",\"type\":\"bool\"}],\"internalType\":\"structReader.PerpetualReaderResult[]\",\"name\":\"perpetuals\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"isAMMMaintenanceSafe\",\"type\":\"bool\"}],\"internalType\":\"structReader.LiquidityPoolReaderResult\",\"name\":\"pool\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidityPool\",\"type\":\"address\"}],\"name\":\"getPoolMargin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSynced\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"poolMargin\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"isSafe\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inverseStateService\",\"outputs\":[{\"internalType\":\"contractIInverseStateService\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidityPool\",\"type\":\"address\"}],\"name\":\"isAMMMaintenanceSafe\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"cashToAdd\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"shareToMint\",\"type\":\"int256\"}],\"name\":\"queryAddLiquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSynced\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"cashToAddResult\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"shareToMintResult\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"shareToRemove\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"cashToReturn\",\"type\":\"int256\"}],\"name\":\"queryRemoveLiquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSynced\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"shareToRemoveResult\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"cashToReturnResult\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"perpetualIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"flags\",\"type\":\"uint32\"}],\"name\":\"queryTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSynced\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"tradePrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"totalFee\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"cost\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"oracles\",\"type\":\"address[]\"}],\"name\":\"readIndexPrices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"isSuccess\",\"type\":\"bool[]\"},{\"internalType\":\"int256[]\",\"name\":\"indexPrices\",\"type\":\"int256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"timestamps\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ReaderABI is the input ABI used to generate the binding from.
// Deprecated: Use ReaderMetaData.ABI instead.
var ReaderABI = ReaderMetaData.ABI

// Reader is an auto generated Go binding around an Ethereum contract.
type Reader struct {
	ReaderCaller     // Read-only binding to the contract
	ReaderTransactor // Write-only binding to the contract
	ReaderFilterer   // Log filterer for contract events
}

// ReaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReaderSession struct {
	Contract     *Reader           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReaderCallerSession struct {
	Contract *ReaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ReaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReaderTransactorSession struct {
	Contract     *ReaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReaderRaw struct {
	Contract *Reader // Generic contract binding to access the raw methods on
}

// ReaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReaderCallerRaw struct {
	Contract *ReaderCaller // Generic read-only contract binding to access the raw methods on
}

// ReaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReaderTransactorRaw struct {
	Contract *ReaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReader creates a new instance of Reader, bound to a specific deployed contract.
func NewReader(address common.Address, backend bind.ContractBackend) (*Reader, error) {
	contract, err := bindReader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reader{ReaderCaller: ReaderCaller{contract: contract}, ReaderTransactor: ReaderTransactor{contract: contract}, ReaderFilterer: ReaderFilterer{contract: contract}}, nil
}

// NewReaderCaller creates a new read-only instance of Reader, bound to a specific deployed contract.
func NewReaderCaller(address common.Address, caller bind.ContractCaller) (*ReaderCaller, error) {
	contract, err := bindReader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReaderCaller{contract: contract}, nil
}

// NewReaderTransactor creates a new write-only instance of Reader, bound to a specific deployed contract.
func NewReaderTransactor(address common.Address, transactor bind.ContractTransactor) (*ReaderTransactor, error) {
	contract, err := bindReader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReaderTransactor{contract: contract}, nil
}

// NewReaderFilterer creates a new log filterer instance of Reader, bound to a specific deployed contract.
func NewReaderFilterer(address common.Address, filterer bind.ContractFilterer) (*ReaderFilterer, error) {
	contract, err := bindReader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReaderFilterer{contract: contract}, nil
}

// bindReader binds a generic wrapper to an already deployed contract.
func bindReader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReaderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reader *ReaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reader.Contract.ReaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reader *ReaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reader.Contract.ReaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reader *ReaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reader.Contract.ReaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reader *ReaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reader *ReaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reader *ReaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reader.Contract.contract.Transact(opts, method, params...)
}

// GetAccountStorage is a free data retrieval call binding the contract method 0xeb16510d.
//
// Solidity: function getAccountStorage(address liquidityPool, uint256 perpetualIndex, address account) view returns(bool isSynced, (int256,int256,int256,int256,int256,bool,bool,bool,int256) accountStorage)
func (_Reader *ReaderCaller) GetAccountStorage(opts *bind.CallOpts, liquidityPool common.Address, perpetualIndex *big.Int, account common.Address) (struct {
	IsSynced       bool
	AccountStorage ReaderAccountReaderResult
}, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getAccountStorage", liquidityPool, perpetualIndex, account)

	outstruct := new(struct {
		IsSynced       bool
		AccountStorage ReaderAccountReaderResult
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsSynced = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.AccountStorage = *abi.ConvertType(out[1], new(ReaderAccountReaderResult)).(*ReaderAccountReaderResult)

	return *outstruct, err

}

// GetAccountStorage is a free data retrieval call binding the contract method 0xeb16510d.
//
// Solidity: function getAccountStorage(address liquidityPool, uint256 perpetualIndex, address account) view returns(bool isSynced, (int256,int256,int256,int256,int256,bool,bool,bool,int256) accountStorage)
func (_Reader *ReaderSession) GetAccountStorage(liquidityPool common.Address, perpetualIndex *big.Int, account common.Address) (struct {
	IsSynced       bool
	AccountStorage ReaderAccountReaderResult
}, error) {
	return _Reader.Contract.GetAccountStorage(&_Reader.CallOpts, liquidityPool, perpetualIndex, account)
}

// GetAccountStorage is a free data retrieval call binding the contract method 0xeb16510d.
//
// Solidity: function getAccountStorage(address liquidityPool, uint256 perpetualIndex, address account) view returns(bool isSynced, (int256,int256,int256,int256,int256,bool,bool,bool,int256) accountStorage)
func (_Reader *ReaderCallerSession) GetAccountStorage(liquidityPool common.Address, perpetualIndex *big.Int, account common.Address) (struct {
	IsSynced       bool
	AccountStorage ReaderAccountReaderResult
}, error) {
	return _Reader.Contract.GetAccountStorage(&_Reader.CallOpts, liquidityPool, perpetualIndex, account)
}

// GetImplementation is a free data retrieval call binding the contract method 0x15ac72ca.
//
// Solidity: function getImplementation(address proxy) view returns(address)
func (_Reader *ReaderCaller) GetImplementation(opts *bind.CallOpts, proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getImplementation", proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0x15ac72ca.
//
// Solidity: function getImplementation(address proxy) view returns(address)
func (_Reader *ReaderSession) GetImplementation(proxy common.Address) (common.Address, error) {
	return _Reader.Contract.GetImplementation(&_Reader.CallOpts, proxy)
}

// GetImplementation is a free data retrieval call binding the contract method 0x15ac72ca.
//
// Solidity: function getImplementation(address proxy) view returns(address)
func (_Reader *ReaderCallerSession) GetImplementation(proxy common.Address) (common.Address, error) {
	return _Reader.Contract.GetImplementation(&_Reader.CallOpts, proxy)
}

// GetL1BlockNumber is a free data retrieval call binding the contract method 0xb9b3efe9.
//
// Solidity: function getL1BlockNumber() view returns(uint256)
func (_Reader *ReaderCaller) GetL1BlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getL1BlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetL1BlockNumber is a free data retrieval call binding the contract method 0xb9b3efe9.
//
// Solidity: function getL1BlockNumber() view returns(uint256)
func (_Reader *ReaderSession) GetL1BlockNumber() (*big.Int, error) {
	return _Reader.Contract.GetL1BlockNumber(&_Reader.CallOpts)
}

// GetL1BlockNumber is a free data retrieval call binding the contract method 0xb9b3efe9.
//
// Solidity: function getL1BlockNumber() view returns(uint256)
func (_Reader *ReaderCallerSession) GetL1BlockNumber() (*big.Int, error) {
	return _Reader.Contract.GetL1BlockNumber(&_Reader.CallOpts)
}

// GetLiquidityPoolStorage is a free data retrieval call binding the contract method 0x574408c1.
//
// Solidity: function getLiquidityPoolStorage(address liquidityPool) view returns(bool isSynced, (bool,bool,address[7],int256[5],uint256[6],(uint8,address,int256[39],uint256,string,bool,bool,int256,int256,bool)[],bool) pool)
func (_Reader *ReaderCaller) GetLiquidityPoolStorage(opts *bind.CallOpts, liquidityPool common.Address) (struct {
	IsSynced bool
	Pool     ReaderLiquidityPoolReaderResult
}, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getLiquidityPoolStorage", liquidityPool)

	outstruct := new(struct {
		IsSynced bool
		Pool     ReaderLiquidityPoolReaderResult
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsSynced = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Pool = *abi.ConvertType(out[1], new(ReaderLiquidityPoolReaderResult)).(*ReaderLiquidityPoolReaderResult)

	return *outstruct, err

}

// GetLiquidityPoolStorage is a free data retrieval call binding the contract method 0x574408c1.
//
// Solidity: function getLiquidityPoolStorage(address liquidityPool) view returns(bool isSynced, (bool,bool,address[7],int256[5],uint256[6],(uint8,address,int256[39],uint256,string,bool,bool,int256,int256,bool)[],bool) pool)
func (_Reader *ReaderSession) GetLiquidityPoolStorage(liquidityPool common.Address) (struct {
	IsSynced bool
	Pool     ReaderLiquidityPoolReaderResult
}, error) {
	return _Reader.Contract.GetLiquidityPoolStorage(&_Reader.CallOpts, liquidityPool)
}

// GetLiquidityPoolStorage is a free data retrieval call binding the contract method 0x574408c1.
//
// Solidity: function getLiquidityPoolStorage(address liquidityPool) view returns(bool isSynced, (bool,bool,address[7],int256[5],uint256[6],(uint8,address,int256[39],uint256,string,bool,bool,int256,int256,bool)[],bool) pool)
func (_Reader *ReaderCallerSession) GetLiquidityPoolStorage(liquidityPool common.Address) (struct {
	IsSynced bool
	Pool     ReaderLiquidityPoolReaderResult
}, error) {
	return _Reader.Contract.GetLiquidityPoolStorage(&_Reader.CallOpts, liquidityPool)
}

// GetPoolMargin is a free data retrieval call binding the contract method 0x41de031a.
//
// Solidity: function getPoolMargin(address liquidityPool) view returns(bool isSynced, int256 poolMargin, bool isSafe)
func (_Reader *ReaderCaller) GetPoolMargin(opts *bind.CallOpts, liquidityPool common.Address) (struct {
	IsSynced   bool
	PoolMargin *big.Int
	IsSafe     bool
}, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getPoolMargin", liquidityPool)

	outstruct := new(struct {
		IsSynced   bool
		PoolMargin *big.Int
		IsSafe     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsSynced = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PoolMargin = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsSafe = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetPoolMargin is a free data retrieval call binding the contract method 0x41de031a.
//
// Solidity: function getPoolMargin(address liquidityPool) view returns(bool isSynced, int256 poolMargin, bool isSafe)
func (_Reader *ReaderSession) GetPoolMargin(liquidityPool common.Address) (struct {
	IsSynced   bool
	PoolMargin *big.Int
	IsSafe     bool
}, error) {
	return _Reader.Contract.GetPoolMargin(&_Reader.CallOpts, liquidityPool)
}

// GetPoolMargin is a free data retrieval call binding the contract method 0x41de031a.
//
// Solidity: function getPoolMargin(address liquidityPool) view returns(bool isSynced, int256 poolMargin, bool isSafe)
func (_Reader *ReaderCallerSession) GetPoolMargin(liquidityPool common.Address) (struct {
	IsSynced   bool
	PoolMargin *big.Int
	IsSafe     bool
}, error) {
	return _Reader.Contract.GetPoolMargin(&_Reader.CallOpts, liquidityPool)
}

// InverseStateService is a free data retrieval call binding the contract method 0x34a3cf33.
//
// Solidity: function inverseStateService() view returns(address)
func (_Reader *ReaderCaller) InverseStateService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "inverseStateService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InverseStateService is a free data retrieval call binding the contract method 0x34a3cf33.
//
// Solidity: function inverseStateService() view returns(address)
func (_Reader *ReaderSession) InverseStateService() (common.Address, error) {
	return _Reader.Contract.InverseStateService(&_Reader.CallOpts)
}

// InverseStateService is a free data retrieval call binding the contract method 0x34a3cf33.
//
// Solidity: function inverseStateService() view returns(address)
func (_Reader *ReaderCallerSession) InverseStateService() (common.Address, error) {
	return _Reader.Contract.InverseStateService(&_Reader.CallOpts)
}

// PoolCreator is a free data retrieval call binding the contract method 0xc6c1decd.
//
// Solidity: function poolCreator() view returns(address)
func (_Reader *ReaderCaller) PoolCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "poolCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolCreator is a free data retrieval call binding the contract method 0xc6c1decd.
//
// Solidity: function poolCreator() view returns(address)
func (_Reader *ReaderSession) PoolCreator() (common.Address, error) {
	return _Reader.Contract.PoolCreator(&_Reader.CallOpts)
}

// PoolCreator is a free data retrieval call binding the contract method 0xc6c1decd.
//
// Solidity: function poolCreator() view returns(address)
func (_Reader *ReaderCallerSession) PoolCreator() (common.Address, error) {
	return _Reader.Contract.PoolCreator(&_Reader.CallOpts)
}

// QueryAddLiquidity is a free data retrieval call binding the contract method 0x7e4b4e45.
//
// Solidity: function queryAddLiquidity(address liquidityPool, int256 cashToAdd, int256 shareToMint) view returns(bool isSynced, int256 cashToAddResult, int256 shareToMintResult)
func (_Reader *ReaderCaller) QueryAddLiquidity(opts *bind.CallOpts, liquidityPool common.Address, cashToAdd *big.Int, shareToMint *big.Int) (struct {
	IsSynced          bool
	CashToAddResult   *big.Int
	ShareToMintResult *big.Int
}, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "queryAddLiquidity", liquidityPool, cashToAdd, shareToMint)

	outstruct := new(struct {
		IsSynced          bool
		CashToAddResult   *big.Int
		ShareToMintResult *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsSynced = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CashToAddResult = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ShareToMintResult = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QueryAddLiquidity is a free data retrieval call binding the contract method 0x7e4b4e45.
//
// Solidity: function queryAddLiquidity(address liquidityPool, int256 cashToAdd, int256 shareToMint) view returns(bool isSynced, int256 cashToAddResult, int256 shareToMintResult)
func (_Reader *ReaderSession) QueryAddLiquidity(liquidityPool common.Address, cashToAdd *big.Int, shareToMint *big.Int) (struct {
	IsSynced          bool
	CashToAddResult   *big.Int
	ShareToMintResult *big.Int
}, error) {
	return _Reader.Contract.QueryAddLiquidity(&_Reader.CallOpts, liquidityPool, cashToAdd, shareToMint)
}

// QueryAddLiquidity is a free data retrieval call binding the contract method 0x7e4b4e45.
//
// Solidity: function queryAddLiquidity(address liquidityPool, int256 cashToAdd, int256 shareToMint) view returns(bool isSynced, int256 cashToAddResult, int256 shareToMintResult)
func (_Reader *ReaderCallerSession) QueryAddLiquidity(liquidityPool common.Address, cashToAdd *big.Int, shareToMint *big.Int) (struct {
	IsSynced          bool
	CashToAddResult   *big.Int
	ShareToMintResult *big.Int
}, error) {
	return _Reader.Contract.QueryAddLiquidity(&_Reader.CallOpts, liquidityPool, cashToAdd, shareToMint)
}

// QueryRemoveLiquidity is a free data retrieval call binding the contract method 0x3c070544.
//
// Solidity: function queryRemoveLiquidity(address liquidityPool, int256 shareToRemove, int256 cashToReturn) view returns(bool isSynced, int256 shareToRemoveResult, int256 cashToReturnResult)
func (_Reader *ReaderCaller) QueryRemoveLiquidity(opts *bind.CallOpts, liquidityPool common.Address, shareToRemove *big.Int, cashToReturn *big.Int) (struct {
	IsSynced            bool
	ShareToRemoveResult *big.Int
	CashToReturnResult  *big.Int
}, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "queryRemoveLiquidity", liquidityPool, shareToRemove, cashToReturn)

	outstruct := new(struct {
		IsSynced            bool
		ShareToRemoveResult *big.Int
		CashToReturnResult  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsSynced = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ShareToRemoveResult = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CashToReturnResult = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QueryRemoveLiquidity is a free data retrieval call binding the contract method 0x3c070544.
//
// Solidity: function queryRemoveLiquidity(address liquidityPool, int256 shareToRemove, int256 cashToReturn) view returns(bool isSynced, int256 shareToRemoveResult, int256 cashToReturnResult)
func (_Reader *ReaderSession) QueryRemoveLiquidity(liquidityPool common.Address, shareToRemove *big.Int, cashToReturn *big.Int) (struct {
	IsSynced            bool
	ShareToRemoveResult *big.Int
	CashToReturnResult  *big.Int
}, error) {
	return _Reader.Contract.QueryRemoveLiquidity(&_Reader.CallOpts, liquidityPool, shareToRemove, cashToReturn)
}

// QueryRemoveLiquidity is a free data retrieval call binding the contract method 0x3c070544.
//
// Solidity: function queryRemoveLiquidity(address liquidityPool, int256 shareToRemove, int256 cashToReturn) view returns(bool isSynced, int256 shareToRemoveResult, int256 cashToReturnResult)
func (_Reader *ReaderCallerSession) QueryRemoveLiquidity(liquidityPool common.Address, shareToRemove *big.Int, cashToReturn *big.Int) (struct {
	IsSynced            bool
	ShareToRemoveResult *big.Int
	CashToReturnResult  *big.Int
}, error) {
	return _Reader.Contract.QueryRemoveLiquidity(&_Reader.CallOpts, liquidityPool, shareToRemove, cashToReturn)
}

// QueryTrade is a free data retrieval call binding the contract method 0x04a6c736.
//
// Solidity: function queryTrade(address liquidityPool, uint256 perpetualIndex, address trader, int256 amount, address referrer, uint32 flags) view returns(bool isSynced, int256 tradePrice, int256 totalFee, int256 cost)
func (_Reader *ReaderCaller) QueryTrade(opts *bind.CallOpts, liquidityPool common.Address, perpetualIndex *big.Int, trader common.Address, amount *big.Int, referrer common.Address, flags uint32) (struct {
	IsSynced   bool
	TradePrice *big.Int
	TotalFee   *big.Int
	Cost       *big.Int
}, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "queryTrade", liquidityPool, perpetualIndex, trader, amount, referrer, flags)

	outstruct := new(struct {
		IsSynced   bool
		TradePrice *big.Int
		TotalFee   *big.Int
		Cost       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsSynced = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.TradePrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Cost = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QueryTrade is a free data retrieval call binding the contract method 0x04a6c736.
//
// Solidity: function queryTrade(address liquidityPool, uint256 perpetualIndex, address trader, int256 amount, address referrer, uint32 flags) view returns(bool isSynced, int256 tradePrice, int256 totalFee, int256 cost)
func (_Reader *ReaderSession) QueryTrade(liquidityPool common.Address, perpetualIndex *big.Int, trader common.Address, amount *big.Int, referrer common.Address, flags uint32) (struct {
	IsSynced   bool
	TradePrice *big.Int
	TotalFee   *big.Int
	Cost       *big.Int
}, error) {
	return _Reader.Contract.QueryTrade(&_Reader.CallOpts, liquidityPool, perpetualIndex, trader, amount, referrer, flags)
}

// QueryTrade is a free data retrieval call binding the contract method 0x04a6c736.
//
// Solidity: function queryTrade(address liquidityPool, uint256 perpetualIndex, address trader, int256 amount, address referrer, uint32 flags) view returns(bool isSynced, int256 tradePrice, int256 totalFee, int256 cost)
func (_Reader *ReaderCallerSession) QueryTrade(liquidityPool common.Address, perpetualIndex *big.Int, trader common.Address, amount *big.Int, referrer common.Address, flags uint32) (struct {
	IsSynced   bool
	TradePrice *big.Int
	TotalFee   *big.Int
	Cost       *big.Int
}, error) {
	return _Reader.Contract.QueryTrade(&_Reader.CallOpts, liquidityPool, perpetualIndex, trader, amount, referrer, flags)
}

// ReadIndexPrices is a free data retrieval call binding the contract method 0xa90dc662.
//
// Solidity: function readIndexPrices(address[] oracles) view returns(bool[] isSuccess, int256[] indexPrices, uint256[] timestamps)
func (_Reader *ReaderCaller) ReadIndexPrices(opts *bind.CallOpts, oracles []common.Address) (struct {
	IsSuccess   []bool
	IndexPrices []*big.Int
	Timestamps  []*big.Int
}, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "readIndexPrices", oracles)

	outstruct := new(struct {
		IsSuccess   []bool
		IndexPrices []*big.Int
		Timestamps  []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsSuccess = *abi.ConvertType(out[0], new([]bool)).(*[]bool)
	outstruct.IndexPrices = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Timestamps = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// ReadIndexPrices is a free data retrieval call binding the contract method 0xa90dc662.
//
// Solidity: function readIndexPrices(address[] oracles) view returns(bool[] isSuccess, int256[] indexPrices, uint256[] timestamps)
func (_Reader *ReaderSession) ReadIndexPrices(oracles []common.Address) (struct {
	IsSuccess   []bool
	IndexPrices []*big.Int
	Timestamps  []*big.Int
}, error) {
	return _Reader.Contract.ReadIndexPrices(&_Reader.CallOpts, oracles)
}

// ReadIndexPrices is a free data retrieval call binding the contract method 0xa90dc662.
//
// Solidity: function readIndexPrices(address[] oracles) view returns(bool[] isSuccess, int256[] indexPrices, uint256[] timestamps)
func (_Reader *ReaderCallerSession) ReadIndexPrices(oracles []common.Address) (struct {
	IsSuccess   []bool
	IndexPrices []*big.Int
	Timestamps  []*big.Int
}, error) {
	return _Reader.Contract.ReadIndexPrices(&_Reader.CallOpts, oracles)
}

// GetAccountsInfo is a paid mutator transaction binding the contract method 0x77ee51ee.
//
// Solidity: function getAccountsInfo(address liquidityPool, uint256 perpetualIndex, uint256 begin, uint256 end) returns(bool isSynced, (address,int256,int256,bool,int256)[] result)
func (_Reader *ReaderTransactor) GetAccountsInfo(opts *bind.TransactOpts, liquidityPool common.Address, perpetualIndex *big.Int, begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Reader.contract.Transact(opts, "getAccountsInfo", liquidityPool, perpetualIndex, begin, end)
}

// GetAccountsInfo is a paid mutator transaction binding the contract method 0x77ee51ee.
//
// Solidity: function getAccountsInfo(address liquidityPool, uint256 perpetualIndex, uint256 begin, uint256 end) returns(bool isSynced, (address,int256,int256,bool,int256)[] result)
func (_Reader *ReaderSession) GetAccountsInfo(liquidityPool common.Address, perpetualIndex *big.Int, begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Reader.Contract.GetAccountsInfo(&_Reader.TransactOpts, liquidityPool, perpetualIndex, begin, end)
}

// GetAccountsInfo is a paid mutator transaction binding the contract method 0x77ee51ee.
//
// Solidity: function getAccountsInfo(address liquidityPool, uint256 perpetualIndex, uint256 begin, uint256 end) returns(bool isSynced, (address,int256,int256,bool,int256)[] result)
func (_Reader *ReaderTransactorSession) GetAccountsInfo(liquidityPool common.Address, perpetualIndex *big.Int, begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Reader.Contract.GetAccountsInfo(&_Reader.TransactOpts, liquidityPool, perpetualIndex, begin, end)
}

// GetAccountsInfoByAddress is a paid mutator transaction binding the contract method 0xbc0a2603.
//
// Solidity: function getAccountsInfoByAddress(address liquidityPool, uint256 perpetualIndex, address[] accounts) returns(bool isSynced, (address,int256,int256,bool,int256)[] result)
func (_Reader *ReaderTransactor) GetAccountsInfoByAddress(opts *bind.TransactOpts, liquidityPool common.Address, perpetualIndex *big.Int, accounts []common.Address) (*types.Transaction, error) {
	return _Reader.contract.Transact(opts, "getAccountsInfoByAddress", liquidityPool, perpetualIndex, accounts)
}

// GetAccountsInfoByAddress is a paid mutator transaction binding the contract method 0xbc0a2603.
//
// Solidity: function getAccountsInfoByAddress(address liquidityPool, uint256 perpetualIndex, address[] accounts) returns(bool isSynced, (address,int256,int256,bool,int256)[] result)
func (_Reader *ReaderSession) GetAccountsInfoByAddress(liquidityPool common.Address, perpetualIndex *big.Int, accounts []common.Address) (*types.Transaction, error) {
	return _Reader.Contract.GetAccountsInfoByAddress(&_Reader.TransactOpts, liquidityPool, perpetualIndex, accounts)
}

// GetAccountsInfoByAddress is a paid mutator transaction binding the contract method 0xbc0a2603.
//
// Solidity: function getAccountsInfoByAddress(address liquidityPool, uint256 perpetualIndex, address[] accounts) returns(bool isSynced, (address,int256,int256,bool,int256)[] result)
func (_Reader *ReaderTransactorSession) GetAccountsInfoByAddress(liquidityPool common.Address, perpetualIndex *big.Int, accounts []common.Address) (*types.Transaction, error) {
	return _Reader.Contract.GetAccountsInfoByAddress(&_Reader.TransactOpts, liquidityPool, perpetualIndex, accounts)
}

// IsAMMMaintenanceSafe is a paid mutator transaction binding the contract method 0x46a83af8.
//
// Solidity: function isAMMMaintenanceSafe(address liquidityPool) returns(bool)
func (_Reader *ReaderTransactor) IsAMMMaintenanceSafe(opts *bind.TransactOpts, liquidityPool common.Address) (*types.Transaction, error) {
	return _Reader.contract.Transact(opts, "isAMMMaintenanceSafe", liquidityPool)
}

// IsAMMMaintenanceSafe is a paid mutator transaction binding the contract method 0x46a83af8.
//
// Solidity: function isAMMMaintenanceSafe(address liquidityPool) returns(bool)
func (_Reader *ReaderSession) IsAMMMaintenanceSafe(liquidityPool common.Address) (*types.Transaction, error) {
	return _Reader.Contract.IsAMMMaintenanceSafe(&_Reader.TransactOpts, liquidityPool)
}

// IsAMMMaintenanceSafe is a paid mutator transaction binding the contract method 0x46a83af8.
//
// Solidity: function isAMMMaintenanceSafe(address liquidityPool) returns(bool)
func (_Reader *ReaderTransactorSession) IsAMMMaintenanceSafe(liquidityPool common.Address) (*types.Transaction, error) {
	return _Reader.Contract.IsAMMMaintenanceSafe(&_Reader.TransactOpts, liquidityPool)
}
