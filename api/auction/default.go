// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"highestBider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"}],\"name\":\"End\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"Start\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allBids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"end\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ended\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBider\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_openingBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"start\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"started\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561000f575f80fd5b5060405161141c38038061141c83398181016040528101906100319190610139565b3373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508060c081815250505050610177565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100d5826100ac565b9050919050565b6100e5816100cb565b81146100ef575f80fd5b50565b5f81519050610100816100dc565b92915050565b5f819050919050565b61011881610106565b8114610122575f80fd5b50565b5f815190506101338161010f565b92915050565b5f806040838503121561014f5761014e6100a8565b5b5f61015c858286016100f2565b925050602061016d85828601610125565b9150509250929050565b60805160a05160c0516112436101d95f395f81816107850152818161085e0152610a6f01525f81816105c6015281816107270152610a1001525f818161060f0152818161063301528181610763015281816108880152610adb01526112435ff3fe6080604052600436106100c1575f3560e01c80636b57707e1161007e578063bd20160711610058578063bd20160714610209578063c6bc518214610245578063d57bde791461026f578063efbe1c1c14610299576100c1565b80636b57707e1461018d5780638da5cb5b146101b75780638fb4b573146101e1576100c1565b806312fa6feb146100c55780631998aeef146100ef5780631f2698ab146100f957806337970be6146101235780633ccfd60b1461014d57806347ccca0214610163575b5f80fd5b3480156100d0575f80fd5b506100d96102af565b6040516100e69190610bb7565b60405180910390f35b6100f76102c1565b005b348015610104575f80fd5b5061010d6104a2565b60405161011a9190610bb7565b60405180910390f35b34801561012e575f80fd5b506101376104b2565b6040516101449190610be8565b60405180910390f35b348015610158575f80fd5b506101616104b8565b005b34801561016e575f80fd5b506101776105c4565b6040516101849190610c7b565b60405180910390f35b348015610198575f80fd5b506101a16105e8565b6040516101ae9190610cb4565b60405180910390f35b3480156101c2575f80fd5b506101cb61060d565b6040516101d89190610ced565b60405180910390f35b3480156101ec575f80fd5b5061020760048036038101906102029190610d34565b610631565b005b348015610214575f80fd5b5061022f600480360381019061022a9190610d9c565b610847565b60405161023c9190610be8565b60405180910390f35b348015610250575f80fd5b5061025961085c565b6040516102669190610be8565b60405180910390f35b34801561027a575f80fd5b50610283610880565b6040516102909190610be8565b60405180910390f35b3480156102a4575f80fd5b506102ad610886565b005b5f60019054906101000a900460ff1681565b5f8054906101000a900460ff1661030d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030490610e21565b60405180910390fd5b6002543411610351576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034890610e89565b60405180910390fd5b6001544210610395576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161038c90610ef1565b60405180910390fd5b60025460045f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546104049190610f3c565b92505081905550346002819055503360035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167fe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2346040516104989190610be8565b60405180910390a2565b5f8054906101000a900460ff1681565b60015481565b5f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505f811115610588573373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015610586573d5f803e3d5ffd5b505b7f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a942436433826040516105b9929190610f6f565b60405180910390a150565b7f000000000000000000000000000000000000000000000000000000000000000081565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b690611006565b60405180910390fd5b5f8054906101000a900460ff161561070c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107039061106e565b60405180910390fd5b81600281905550804261071f9190610f3c565b6001819055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd7f0000000000000000000000000000000000000000000000000000000000000000307f00000000000000000000000000000000000000000000000000000000000000006040518463ffffffff1660e01b81526004016107c2939291906110ac565b5f604051808303815f87803b1580156107d9575f80fd5b505af11580156107eb573d5f803e3d5ffd5b5050505060015f806101000a81548160ff0219169083151502179055507f5a7d5f50ab70a39d193bf53f0fb995377776df93bb6fdcf1cd1868b7e0dd44ee4260015460405161083b9291906110e1565b60405180910390a15050565b6004602052805f5260405f205f915090505481565b7f000000000000000000000000000000000000000000000000000000000000000081565b60025481565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610914576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161090b90611006565b60405180910390fd5b5f8054906101000a900460ff16610960576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095790610e21565b60405180910390fd5b6001544210156109a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099c90611152565b60405180910390fd5b5f60019054906101000a900460ff16156109f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109eb906111ba565b60405180910390fd5b60015f60016101000a81548160ff0219169083151502179055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd3060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff167f00000000000000000000000000000000000000000000000000000000000000006040518463ffffffff1660e01b8152600401610aac939291906111d8565b5f604051808303815f87803b158015610ac3575f80fd5b505af1158015610ad5573d5f803e3d5ffd5b505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166108fc60025490811502906040515f60405180830381858888f19350505050158015610b3e573d5f803e3d5ffd5b507f7d7570b046e5ead3a4f3fbc9fa2445113625a1e79823776146316bafed6b8cc560035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600254604051610b93929190610f6f565b60405180910390a1565b5f8115159050919050565b610bb181610b9d565b82525050565b5f602082019050610bca5f830184610ba8565b92915050565b5f819050919050565b610be281610bd0565b82525050565b5f602082019050610bfb5f830184610bd9565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610c43610c3e610c3984610c01565b610c20565b610c01565b9050919050565b5f610c5482610c29565b9050919050565b5f610c6582610c4a565b9050919050565b610c7581610c5b565b82525050565b5f602082019050610c8e5f830184610c6c565b92915050565b5f610c9e82610c01565b9050919050565b610cae81610c94565b82525050565b5f602082019050610cc75f830184610ca5565b92915050565b5f610cd782610c01565b9050919050565b610ce781610ccd565b82525050565b5f602082019050610d005f830184610cde565b92915050565b5f80fd5b610d1381610bd0565b8114610d1d575f80fd5b50565b5f81359050610d2e81610d0a565b92915050565b5f8060408385031215610d4a57610d49610d06565b5b5f610d5785828601610d20565b9250506020610d6885828601610d20565b9150509250929050565b610d7b81610c94565b8114610d85575f80fd5b50565b5f81359050610d9681610d72565b92915050565b5f60208284031215610db157610db0610d06565b5b5f610dbe84828501610d88565b91505092915050565b5f82825260208201905092915050565b7f41756374696f6e206e6f742073746172746564000000000000000000000000005f82015250565b5f610e0b601383610dc7565b9150610e1682610dd7565b602082019050919050565b5f6020820190508181035f830152610e3881610dff565b9050919050565b7f42696420746f6f206c6f770000000000000000000000000000000000000000005f82015250565b5f610e73600b83610dc7565b9150610e7e82610e3f565b602082019050919050565b5f6020820190508181035f830152610ea081610e67565b9050919050565b7f41756374696f6e20656e646564000000000000000000000000000000000000005f82015250565b5f610edb600d83610dc7565b9150610ee682610ea7565b602082019050919050565b5f6020820190508181035f830152610f0881610ecf565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610f4682610bd0565b9150610f5183610bd0565b9250828201905080821115610f6957610f68610f0f565b5b92915050565b5f604082019050610f825f830185610ca5565b610f8f6020830184610bd9565b9392505050565b7f4f6e6c79206f776e65722063616e2063616c6c20746869732066756e6374696f5f8201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b5f610ff0602183610dc7565b9150610ffb82610f96565b604082019050919050565b5f6020820190508181035f83015261101d81610fe4565b9050919050565b7f41756374696f6e20616c726561647920737461727465640000000000000000005f82015250565b5f611058601783610dc7565b915061106382611024565b602082019050919050565b5f6020820190508181035f8301526110858161104c565b9050919050565b5f61109682610c4a565b9050919050565b6110a68161108c565b82525050565b5f6060820190506110bf5f83018661109d565b6110cc6020830185610ca5565b6110d96040830184610bd9565b949350505050565b5f6040820190506110f45f830185610bd9565b6111016020830184610bd9565b9392505050565b7f41756374696f6e206e6f7420656e6465640000000000000000000000000000005f82015250565b5f61113c601183610dc7565b915061114782611108565b602082019050919050565b5f6020820190508181035f83015261116981611130565b9050919050565b7f41756374696f6e20616c726561647920656e64656400000000000000000000005f82015250565b5f6111a4601583610dc7565b91506111af82611170565b602082019050919050565b5f6020820190508181035f8301526111d181611198565b9050919050565b5f6060820190506111eb5f830186610ca5565b6111f86020830185610ca5565b6112056040830184610bd9565b94935050505056fea2646970667358221220d6adf730384add0e35cbad1b8584e8333c90bfed79ab05efbaeb5f020b14dc6064736f6c63430008190033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _nft common.Address, _nftId *big.Int) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend, _nft, _nftId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// AllBids is a free data retrieval call binding the contract method 0xbd201607.
//
// Solidity: function allBids(address ) view returns(uint256)
func (_Contract *ContractCaller) AllBids(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allBids", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllBids is a free data retrieval call binding the contract method 0xbd201607.
//
// Solidity: function allBids(address ) view returns(uint256)
func (_Contract *ContractSession) AllBids(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.AllBids(&_Contract.CallOpts, arg0)
}

// AllBids is a free data retrieval call binding the contract method 0xbd201607.
//
// Solidity: function allBids(address ) view returns(uint256)
func (_Contract *ContractCallerSession) AllBids(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.AllBids(&_Contract.CallOpts, arg0)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Contract *ContractCaller) Ended(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ended")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Contract *ContractSession) Ended() (bool, error) {
	return _Contract.Contract.Ended(&_Contract.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Contract *ContractCallerSession) Ended() (bool, error) {
	return _Contract.Contract.Ended(&_Contract.CallOpts)
}

// EndedTime is a free data retrieval call binding the contract method 0x37970be6.
//
// Solidity: function endedTime() view returns(uint256)
func (_Contract *ContractCaller) EndedTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "endedTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndedTime is a free data retrieval call binding the contract method 0x37970be6.
//
// Solidity: function endedTime() view returns(uint256)
func (_Contract *ContractSession) EndedTime() (*big.Int, error) {
	return _Contract.Contract.EndedTime(&_Contract.CallOpts)
}

// EndedTime is a free data retrieval call binding the contract method 0x37970be6.
//
// Solidity: function endedTime() view returns(uint256)
func (_Contract *ContractCallerSession) EndedTime() (*big.Int, error) {
	return _Contract.Contract.EndedTime(&_Contract.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Contract *ContractCaller) HighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "highestBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Contract *ContractSession) HighestBid() (*big.Int, error) {
	return _Contract.Contract.HighestBid(&_Contract.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Contract *ContractCallerSession) HighestBid() (*big.Int, error) {
	return _Contract.Contract.HighestBid(&_Contract.CallOpts)
}

// HighestBider is a free data retrieval call binding the contract method 0x6b57707e.
//
// Solidity: function highestBider() view returns(address)
func (_Contract *ContractCaller) HighestBider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "highestBider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBider is a free data retrieval call binding the contract method 0x6b57707e.
//
// Solidity: function highestBider() view returns(address)
func (_Contract *ContractSession) HighestBider() (common.Address, error) {
	return _Contract.Contract.HighestBider(&_Contract.CallOpts)
}

// HighestBider is a free data retrieval call binding the contract method 0x6b57707e.
//
// Solidity: function highestBider() view returns(address)
func (_Contract *ContractCallerSession) HighestBider() (common.Address, error) {
	return _Contract.Contract.HighestBider(&_Contract.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Contract *ContractCaller) Nft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Contract *ContractSession) Nft() (common.Address, error) {
	return _Contract.Contract.Nft(&_Contract.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Contract *ContractCallerSession) Nft() (common.Address, error) {
	return _Contract.Contract.Nft(&_Contract.CallOpts)
}

// NftId is a free data retrieval call binding the contract method 0xc6bc5182.
//
// Solidity: function nftId() view returns(uint256)
func (_Contract *ContractCaller) NftId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nftId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftId is a free data retrieval call binding the contract method 0xc6bc5182.
//
// Solidity: function nftId() view returns(uint256)
func (_Contract *ContractSession) NftId() (*big.Int, error) {
	return _Contract.Contract.NftId(&_Contract.CallOpts)
}

// NftId is a free data retrieval call binding the contract method 0xc6bc5182.
//
// Solidity: function nftId() view returns(uint256)
func (_Contract *ContractCallerSession) NftId() (*big.Int, error) {
	return _Contract.Contract.NftId(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Contract *ContractCaller) Started(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "started")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Contract *ContractSession) Started() (bool, error) {
	return _Contract.Contract.Started(&_Contract.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Contract *ContractCallerSession) Started() (bool, error) {
	return _Contract.Contract.Started(&_Contract.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Contract *ContractTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Contract *ContractSession) Bid() (*types.Transaction, error) {
	return _Contract.Contract.Bid(&_Contract.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Contract *ContractTransactorSession) Bid() (*types.Transaction, error) {
	return _Contract.Contract.Bid(&_Contract.TransactOpts)
}

// End is a paid mutator transaction binding the contract method 0xefbe1c1c.
//
// Solidity: function end() returns()
func (_Contract *ContractTransactor) End(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "end")
}

// End is a paid mutator transaction binding the contract method 0xefbe1c1c.
//
// Solidity: function end() returns()
func (_Contract *ContractSession) End() (*types.Transaction, error) {
	return _Contract.Contract.End(&_Contract.TransactOpts)
}

// End is a paid mutator transaction binding the contract method 0xefbe1c1c.
//
// Solidity: function end() returns()
func (_Contract *ContractTransactorSession) End() (*types.Transaction, error) {
	return _Contract.Contract.End(&_Contract.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0x8fb4b573.
//
// Solidity: function start(uint256 _openingBid, uint256 _duration) returns()
func (_Contract *ContractTransactor) Start(opts *bind.TransactOpts, _openingBid *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "start", _openingBid, _duration)
}

// Start is a paid mutator transaction binding the contract method 0x8fb4b573.
//
// Solidity: function start(uint256 _openingBid, uint256 _duration) returns()
func (_Contract *ContractSession) Start(_openingBid *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Start(&_Contract.TransactOpts, _openingBid, _duration)
}

// Start is a paid mutator transaction binding the contract method 0x8fb4b573.
//
// Solidity: function start(uint256 _openingBid, uint256 _duration) returns()
func (_Contract *ContractTransactorSession) Start(_openingBid *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Start(&_Contract.TransactOpts, _openingBid, _duration)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// ContractBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the Contract contract.
type ContractBidIterator struct {
	Event *ContractBid // Event containing the contract specifics and raw log

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
func (it *ContractBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBid)
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
		it.Event = new(ContractBid)
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
func (it *ContractBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBid represents a Bid event raised by the Contract contract.
type ContractBid struct {
	Bidder common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 value)
func (_Contract *ContractFilterer) FilterBid(opts *bind.FilterOpts, bidder []common.Address) (*ContractBidIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Bid", bidderRule)
	if err != nil {
		return nil, err
	}
	return &ContractBidIterator{contract: _Contract.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 value)
func (_Contract *ContractFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *ContractBid, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Bid", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBid)
				if err := _Contract.contract.UnpackLog(event, "Bid", log); err != nil {
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

// ParseBid is a log parse operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 value)
func (_Contract *ContractFilterer) ParseBid(log types.Log) (*ContractBid, error) {
	event := new(ContractBid)
	if err := _Contract.contract.UnpackLog(event, "Bid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEndIterator is returned from FilterEnd and is used to iterate over the raw logs and unpacked data for End events raised by the Contract contract.
type ContractEndIterator struct {
	Event *ContractEnd // Event containing the contract specifics and raw log

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
func (it *ContractEndIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEnd)
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
		it.Event = new(ContractEnd)
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
func (it *ContractEndIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEndIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEnd represents a End event raised by the Contract contract.
type ContractEnd struct {
	HighestBider common.Address
	HighestBid   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEnd is a free log retrieval operation binding the contract event 0x7d7570b046e5ead3a4f3fbc9fa2445113625a1e79823776146316bafed6b8cc5.
//
// Solidity: event End(address highestBider, uint256 highestBid)
func (_Contract *ContractFilterer) FilterEnd(opts *bind.FilterOpts) (*ContractEndIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "End")
	if err != nil {
		return nil, err
	}
	return &ContractEndIterator{contract: _Contract.contract, event: "End", logs: logs, sub: sub}, nil
}

// WatchEnd is a free log subscription operation binding the contract event 0x7d7570b046e5ead3a4f3fbc9fa2445113625a1e79823776146316bafed6b8cc5.
//
// Solidity: event End(address highestBider, uint256 highestBid)
func (_Contract *ContractFilterer) WatchEnd(opts *bind.WatchOpts, sink chan<- *ContractEnd) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "End")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEnd)
				if err := _Contract.contract.UnpackLog(event, "End", log); err != nil {
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

// ParseEnd is a log parse operation binding the contract event 0x7d7570b046e5ead3a4f3fbc9fa2445113625a1e79823776146316bafed6b8cc5.
//
// Solidity: event End(address highestBider, uint256 highestBid)
func (_Contract *ContractFilterer) ParseEnd(log types.Log) (*ContractEnd, error) {
	event := new(ContractEnd)
	if err := _Contract.contract.UnpackLog(event, "End", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStartIterator is returned from FilterStart and is used to iterate over the raw logs and unpacked data for Start events raised by the Contract contract.
type ContractStartIterator struct {
	Event *ContractStart // Event containing the contract specifics and raw log

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
func (it *ContractStartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStart)
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
		it.Event = new(ContractStart)
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
func (it *ContractStartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStart represents a Start event raised by the Contract contract.
type ContractStart struct {
	StartTime *big.Int
	EndTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStart is a free log retrieval operation binding the contract event 0x5a7d5f50ab70a39d193bf53f0fb995377776df93bb6fdcf1cd1868b7e0dd44ee.
//
// Solidity: event Start(uint256 startTime, uint256 endTime)
func (_Contract *ContractFilterer) FilterStart(opts *bind.FilterOpts) (*ContractStartIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Start")
	if err != nil {
		return nil, err
	}
	return &ContractStartIterator{contract: _Contract.contract, event: "Start", logs: logs, sub: sub}, nil
}

// WatchStart is a free log subscription operation binding the contract event 0x5a7d5f50ab70a39d193bf53f0fb995377776df93bb6fdcf1cd1868b7e0dd44ee.
//
// Solidity: event Start(uint256 startTime, uint256 endTime)
func (_Contract *ContractFilterer) WatchStart(opts *bind.WatchOpts, sink chan<- *ContractStart) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Start")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStart)
				if err := _Contract.contract.UnpackLog(event, "Start", log); err != nil {
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

// ParseStart is a log parse operation binding the contract event 0x5a7d5f50ab70a39d193bf53f0fb995377776df93bb6fdcf1cd1868b7e0dd44ee.
//
// Solidity: event Start(uint256 startTime, uint256 endTime)
func (_Contract *ContractFilterer) ParseStart(log types.Log) (*ContractStart, error) {
	event := new(ContractStart)
	if err := _Contract.contract.UnpackLog(event, "Start", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Contract contract.
type ContractWithdrawIterator struct {
	Event *ContractWithdraw // Event containing the contract specifics and raw log

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
func (it *ContractWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWithdraw)
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
		it.Event = new(ContractWithdraw)
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
func (it *ContractWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWithdraw represents a Withdraw event raised by the Contract contract.
type ContractWithdraw struct {
	Bidder common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address bidder, uint256 value)
func (_Contract *ContractFilterer) FilterWithdraw(opts *bind.FilterOpts) (*ContractWithdrawIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &ContractWithdrawIterator{contract: _Contract.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address bidder, uint256 value)
func (_Contract *ContractFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ContractWithdraw) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWithdraw)
				if err := _Contract.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address bidder, uint256 value)
func (_Contract *ContractFilterer) ParseWithdraw(log types.Log) (*ContractWithdraw, error) {
	event := new(ContractWithdraw)
	if err := _Contract.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
