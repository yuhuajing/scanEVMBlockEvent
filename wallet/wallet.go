// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wallet

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

// WalletMetaData contains all meta data concerning the Wallet contract.
var WalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidInput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gettestnum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_code\",\"type\":\"uint256\"}],\"name\":\"isverified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newaddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_code\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_mixed_question\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_mixed_answer\",\"type\":\"string\"}],\"name\":\"resetOrforgetPassword\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"settestnum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_code\",\"type\":\"uint256\"}],\"name\":\"verifycode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001a0438038062001a04833981810160405281019062000037919062000150565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036200009e576040517fb4fa3fb300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000182565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200011882620000eb565b9050919050565b6200012a816200010b565b81146200013657600080fd5b50565b6000815190506200014a816200011f565b92915050565b600060208284031215620001695762000168620000e6565b5b6000620001798482850162000139565b91505092915050565b61187280620001926000396000f3fe60806040526004361061007b5760003560e01c80638da5cb5b1161004e5780638da5cb5b146101405780639e5d4c491461016b578063e2da61611461019b578063f6bae2b1146101c45761007b565b8063264bd688146100805780634f1543fc146100bd57806373d4a13a146100e8578063814494b114610117575b600080fd5b34801561008c57600080fd5b506100a760048036038101906100a29190610e9e565b6101ed565b6040516100b49190610f15565b60405180910390f35b3480156100c957600080fd5b506100d2610234565b6040516100df9190610f3f565b60405180910390f35b3480156100f457600080fd5b506100fd61023e565b60405161010e95949392919061101a565b60405180910390f35b34801561012357600080fd5b5061013e600480360381019061013991906110b5565b6102c8565b005b34801561014c57600080fd5b50610155610880565b6040516101629190611196565b60405180910390f35b61018560048036038101906101809190611211565b6108e7565b60405161019291906112da565b60405180910390f35b3480156101a757600080fd5b506101c260048036038101906101bd91906112fc565b61096b565b005b3480156101d057600080fd5b506101eb60048036038101906101e69190610e9e565b610975565b005b60006001836040516101ff9190611365565b9081526020016040518091039020600083815260200190815260200160002060009054906101000a900460ff16905092915050565b6000600454905090565b60006060806060806000606067ffffffffffffffff81111561026357610262610d3d565b5b6040519080825280601f01601f1916602001820160405280156102955781602001600182028036833780820191505090505b5090506060604d60208301303c808060200190518101906102b6919061142a565b95509550955095509550509091929394565b600360159054906101000a900460ff166103095760008060008060006102ec61023e565b945094509450945094506103038585858585610a47565b50505050505b60008060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060c001604052908160008201805461036490611544565b80601f016020809104026020016040519081016040528092919081815260200182805461039090611544565b80156103dd5780601f106103b2576101008083540402835291602001916103dd565b820191906000526020600020905b8154815290600101906020018083116103c057829003601f168201915b50505050508152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201805461045690611544565b80601f016020809104026020016040519081016040528092919081815260200182805461048290611544565b80156104cf5780601f106104a4576101008083540402835291602001916104cf565b820191906000526020600020905b8154815290600101906020018083116104b257829003601f168201915b505050505081526020016004820180546104e890611544565b80601f016020809104026020016040519081016040528092919081815260200182805461051490611544565b80156105615780601f1061053657610100808354040283529160200191610561565b820191906000526020600020905b81548152906001019060200180831161054457829003601f168201915b5050505050815260200160058201805461057a90611544565b80601f01602080910402602001604051908101604052809291908181526020018280546105a690611544565b80156105f35780601f106105c8576101008083540402835291602001916105f3565b820191906000526020600020905b8154815290600101906020018083116105d657829003601f168201915b505050505081525050905061060885856101ed565b80156106465750610617610880565b73ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16145b801561065c575061065b858260000151610b9d565b5b80156106725750610671838260800151610b9d565b5b80156106885750610687828260a00151610b9d565b5b61069157600080fd5b8381602001818152505085816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050806000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000190816107279190611721565b506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301908161078e9190611721565b5060808201518160040190816107a49190611721565b5060a08201518160050190816107ba9190611721565b509050506000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000808201600061080d9190610cb1565b60018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560038201600061084c9190610cb1565b60048201600061085c9190610cb1565b60058201600061086c9190610cb1565b505061087786610bc5565b50505050505050565b6000600360149054906101000a900460ff16156108c157600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506108e4565b60008060008060006108d161023e565b9450945094509450945084955050505050505b90565b60606108f1610880565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610955576040517fea8e4eb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61096185858585610c24565b9050949350505050565b8060048190555050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109fc576040517fea8e4eb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018083604051610a0d9190611365565b9081526020016040518091039020600083815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60006040518060c00160405280868152602001600081526020018773ffffffffffffffffffffffffffffffffffffffff168152602001858152602001848152602001838152509050806000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000019081610ae39190611721565b506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003019081610b4a9190611721565b506080820151816004019081610b609190611721565b5060a0820151816005019081610b769190611721565b509050506001600360156101000a81548160ff021916908315150217905550505050505050565b600081518351148015610bbd575081805190602001208380519060200120145b905092915050565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600360146101000a81548160ff02191690831515021790555050565b606060008573ffffffffffffffffffffffffffffffffffffffff16858585604051610c50929190611823565b60006040518083038185875af1925050503d8060008114610c8d576040519150601f19603f3d011682016040523d82523d6000602084013e610c92565b606091505b50809350819250505080610ca857815160208301fd5b50949350505050565b508054610cbd90611544565b6000825580601f10610ccf5750610cee565b601f016020900490600052602060002090810190610ced9190610cf1565b5b50565b5b80821115610d0a576000816000905550600101610cf2565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610d7582610d2c565b810181811067ffffffffffffffff82111715610d9457610d93610d3d565b5b80604052505050565b6000610da7610d0e565b9050610db38282610d6c565b919050565b600067ffffffffffffffff821115610dd357610dd2610d3d565b5b610ddc82610d2c565b9050602081019050919050565b82818337600083830152505050565b6000610e0b610e0684610db8565b610d9d565b905082815260208101848484011115610e2757610e26610d27565b5b610e32848285610de9565b509392505050565b600082601f830112610e4f57610e4e610d22565b5b8135610e5f848260208601610df8565b91505092915050565b6000819050919050565b610e7b81610e68565b8114610e8657600080fd5b50565b600081359050610e9881610e72565b92915050565b60008060408385031215610eb557610eb4610d18565b5b600083013567ffffffffffffffff811115610ed357610ed2610d1d565b5b610edf85828601610e3a565b9250506020610ef085828601610e89565b9150509250929050565b60008115159050919050565b610f0f81610efa565b82525050565b6000602082019050610f2a6000830184610f06565b92915050565b610f3981610e68565b82525050565b6000602082019050610f546000830184610f30565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610f8582610f5a565b9050919050565b610f9581610f7a565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610fd5578082015181840152602081019050610fba565b60008484015250505050565b6000610fec82610f9b565b610ff68185610fa6565b9350611006818560208601610fb7565b61100f81610d2c565b840191505092915050565b600060a08201905061102f6000830188610f8c565b81810360208301526110418187610fe1565b905081810360408301526110558186610fe1565b905081810360608301526110698185610fe1565b9050818103608083015261107d8184610fe1565b90509695505050505050565b61109281610f7a565b811461109d57600080fd5b50565b6000813590506110af81611089565b92915050565b60008060008060008060c087890312156110d2576110d1610d18565b5b60006110e089828a016110a0565b96505060206110f189828a016110a0565b955050604087013567ffffffffffffffff81111561111257611111610d1d565b5b61111e89828a01610e3a565b945050606061112f89828a01610e89565b935050608087013567ffffffffffffffff8111156111505761114f610d1d565b5b61115c89828a01610e3a565b92505060a087013567ffffffffffffffff81111561117d5761117c610d1d565b5b61118989828a01610e3a565b9150509295509295509295565b60006020820190506111ab6000830184610f8c565b92915050565b600080fd5b600080fd5b60008083601f8401126111d1576111d0610d22565b5b8235905067ffffffffffffffff8111156111ee576111ed6111b1565b5b60208301915083600182028301111561120a576112096111b6565b5b9250929050565b6000806000806060858703121561122b5761122a610d18565b5b6000611239878288016110a0565b945050602061124a87828801610e89565b935050604085013567ffffffffffffffff81111561126b5761126a610d1d565b5b611277878288016111bb565b925092505092959194509250565b600081519050919050565b600082825260208201905092915050565b60006112ac82611285565b6112b68185611290565b93506112c6818560208601610fb7565b6112cf81610d2c565b840191505092915050565b600060208201905081810360008301526112f481846112a1565b905092915050565b60006020828403121561131257611311610d18565b5b600061132084828501610e89565b91505092915050565b600081905092915050565b600061133f82610f9b565b6113498185611329565b9350611359818560208601610fb7565b80840191505092915050565b60006113718284611334565b915081905092915050565b600061138782610f5a565b9050919050565b6113978161137c565b81146113a257600080fd5b50565b6000815190506113b48161138e565b92915050565b60006113cd6113c884610db8565b610d9d565b9050828152602081018484840111156113e9576113e8610d27565b5b6113f4848285610fb7565b509392505050565b600082601f83011261141157611410610d22565b5b81516114218482602086016113ba565b91505092915050565b600080600080600060a0868803121561144657611445610d18565b5b6000611454888289016113a5565b955050602086015167ffffffffffffffff81111561147557611474610d1d565b5b611481888289016113fc565b945050604086015167ffffffffffffffff8111156114a2576114a1610d1d565b5b6114ae888289016113fc565b935050606086015167ffffffffffffffff8111156114cf576114ce610d1d565b5b6114db888289016113fc565b925050608086015167ffffffffffffffff8111156114fc576114fb610d1d565b5b611508888289016113fc565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061155c57607f821691505b60208210810361156f5761156e611515565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026115d77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261159a565b6115e1868361159a565b95508019841693508086168417925050509392505050565b6000819050919050565b600061161e61161961161484610e68565b6115f9565b610e68565b9050919050565b6000819050919050565b61163883611603565b61164c61164482611625565b8484546115a7565b825550505050565b600090565b611661611654565b61166c81848461162f565b505050565b5b8181101561169057611685600082611659565b600181019050611672565b5050565b601f8211156116d5576116a681611575565b6116af8461158a565b810160208510156116be578190505b6116d26116ca8561158a565b830182611671565b50505b505050565b600082821c905092915050565b60006116f8600019846008026116da565b1980831691505092915050565b600061171183836116e7565b9150826002028217905092915050565b61172a82610f9b565b67ffffffffffffffff81111561174357611742610d3d565b5b61174d8254611544565b611758828285611694565b600060209050601f83116001811461178b5760008415611779578287015190505b6117838582611705565b8655506117eb565b601f19841661179986611575565b60005b828110156117c15784890151825560018201915060208501945060208101905061179c565b868310156117de57848901516117da601f8916826116e7565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b600061180a83856117f3565b9350611817838584610de9565b82840190509392505050565b60006118308284866117fe565b9150819050939250505056fea264697066735822122089362cfdc216fbd8aae5357d083d569f0b7178f857d6e4d1685e626ecc52576664736f6c63430008120033",
}

// WalletABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletMetaData.ABI instead.
var WalletABI = WalletMetaData.ABI

// WalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletMetaData.Bin instead.
var WalletBin = WalletMetaData.Bin

// DeployWallet deploys a new Ethereum contract, binding an instance of Wallet to it.
func DeployWallet(auth *bind.TransactOpts, backend bind.ContractBackend, _manager common.Address) (common.Address, *types.Transaction, *Wallet, error) {
	parsed, err := WalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletBin), backend, _manager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}, WalletFilterer: WalletFilterer{contract: contract}}, nil
}

// Wallet is an auto generated Go binding around an Ethereum contract.
type Wallet struct {
	WalletCaller     // Read-only binding to the contract
	WalletTransactor // Write-only binding to the contract
	WalletFilterer   // Log filterer for contract events
}

// WalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletSession struct {
	Contract     *Wallet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletCallerSession struct {
	Contract *WalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletTransactorSession struct {
	Contract     *WalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletRaw struct {
	Contract *Wallet // Generic contract binding to access the raw methods on
}

// WalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletCallerRaw struct {
	Contract *WalletCaller // Generic read-only contract binding to access the raw methods on
}

// WalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletTransactorRaw struct {
	Contract *WalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWallet creates a new instance of Wallet, bound to a specific deployed contract.
func NewWallet(address common.Address, backend bind.ContractBackend) (*Wallet, error) {
	contract, err := bindWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}, WalletFilterer: WalletFilterer{contract: contract}}, nil
}

// NewWalletCaller creates a new read-only instance of Wallet, bound to a specific deployed contract.
func NewWalletCaller(address common.Address, caller bind.ContractCaller) (*WalletCaller, error) {
	contract, err := bindWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletCaller{contract: contract}, nil
}

// NewWalletTransactor creates a new write-only instance of Wallet, bound to a specific deployed contract.
func NewWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletTransactor, error) {
	contract, err := bindWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletTransactor{contract: contract}, nil
}

// NewWalletFilterer creates a new log filterer instance of Wallet, bound to a specific deployed contract.
func NewWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletFilterer, error) {
	contract, err := bindWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletFilterer{contract: contract}, nil
}

// bindWallet binds a generic wrapper to an already deployed contract.
func bindWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.WalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transact(opts, method, params...)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(address, string, string, string, string)
func (_Wallet *WalletCaller) Data(opts *bind.CallOpts) (common.Address, string, string, string, string, error) {
	var out []interface{}
	err := _Wallet.contract.Call(opts, &out, "data")

	if err != nil {
		return *new(common.Address), *new(string), *new(string), *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(string)).(*string)

	return out0, out1, out2, out3, out4, err

}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(address, string, string, string, string)
func (_Wallet *WalletSession) Data() (common.Address, string, string, string, string, error) {
	return _Wallet.Contract.Data(&_Wallet.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(address, string, string, string, string)
func (_Wallet *WalletCallerSession) Data() (common.Address, string, string, string, string, error) {
	return _Wallet.Contract.Data(&_Wallet.CallOpts)
}

// Gettestnum is a free data retrieval call binding the contract method 0x4f1543fc.
//
// Solidity: function gettestnum() view returns(uint256)
func (_Wallet *WalletCaller) Gettestnum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Wallet.contract.Call(opts, &out, "gettestnum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gettestnum is a free data retrieval call binding the contract method 0x4f1543fc.
//
// Solidity: function gettestnum() view returns(uint256)
func (_Wallet *WalletSession) Gettestnum() (*big.Int, error) {
	return _Wallet.Contract.Gettestnum(&_Wallet.CallOpts)
}

// Gettestnum is a free data retrieval call binding the contract method 0x4f1543fc.
//
// Solidity: function gettestnum() view returns(uint256)
func (_Wallet *WalletCallerSession) Gettestnum() (*big.Int, error) {
	return _Wallet.Contract.Gettestnum(&_Wallet.CallOpts)
}

// Isverified is a free data retrieval call binding the contract method 0x264bd688.
//
// Solidity: function isverified(string _email, uint256 _code) view returns(bool)
func (_Wallet *WalletCaller) Isverified(opts *bind.CallOpts, _email string, _code *big.Int) (bool, error) {
	var out []interface{}
	err := _Wallet.contract.Call(opts, &out, "isverified", _email, _code)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Isverified is a free data retrieval call binding the contract method 0x264bd688.
//
// Solidity: function isverified(string _email, uint256 _code) view returns(bool)
func (_Wallet *WalletSession) Isverified(_email string, _code *big.Int) (bool, error) {
	return _Wallet.Contract.Isverified(&_Wallet.CallOpts, _email, _code)
}

// Isverified is a free data retrieval call binding the contract method 0x264bd688.
//
// Solidity: function isverified(string _email, uint256 _code) view returns(bool)
func (_Wallet *WalletCallerSession) Isverified(_email string, _code *big.Int) (bool, error) {
	return _Wallet.Contract.Isverified(&_Wallet.CallOpts, _email, _code)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wallet *WalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wallet *WalletSession) Owner() (common.Address, error) {
	return _Wallet.Contract.Owner(&_Wallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wallet *WalletCallerSession) Owner() (common.Address, error) {
	return _Wallet.Contract.Owner(&_Wallet.CallOpts)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes _calldata) payable returns(bytes)
func (_Wallet *WalletTransactor) ExecuteCall(opts *bind.TransactOpts, to common.Address, value *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "executeCall", to, value, _calldata)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes _calldata) payable returns(bytes)
func (_Wallet *WalletSession) ExecuteCall(to common.Address, value *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _Wallet.Contract.ExecuteCall(&_Wallet.TransactOpts, to, value, _calldata)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes _calldata) payable returns(bytes)
func (_Wallet *WalletTransactorSession) ExecuteCall(to common.Address, value *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _Wallet.Contract.ExecuteCall(&_Wallet.TransactOpts, to, value, _calldata)
}

// ResetOrforgetPassword is a paid mutator transaction binding the contract method 0x814494b1.
//
// Solidity: function resetOrforgetPassword(address _address, address _newaddress, string _email, uint256 _code, string _mixed_question, string _mixed_answer) returns()
func (_Wallet *WalletTransactor) ResetOrforgetPassword(opts *bind.TransactOpts, _address common.Address, _newaddress common.Address, _email string, _code *big.Int, _mixed_question string, _mixed_answer string) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "resetOrforgetPassword", _address, _newaddress, _email, _code, _mixed_question, _mixed_answer)
}

// ResetOrforgetPassword is a paid mutator transaction binding the contract method 0x814494b1.
//
// Solidity: function resetOrforgetPassword(address _address, address _newaddress, string _email, uint256 _code, string _mixed_question, string _mixed_answer) returns()
func (_Wallet *WalletSession) ResetOrforgetPassword(_address common.Address, _newaddress common.Address, _email string, _code *big.Int, _mixed_question string, _mixed_answer string) (*types.Transaction, error) {
	return _Wallet.Contract.ResetOrforgetPassword(&_Wallet.TransactOpts, _address, _newaddress, _email, _code, _mixed_question, _mixed_answer)
}

// ResetOrforgetPassword is a paid mutator transaction binding the contract method 0x814494b1.
//
// Solidity: function resetOrforgetPassword(address _address, address _newaddress, string _email, uint256 _code, string _mixed_question, string _mixed_answer) returns()
func (_Wallet *WalletTransactorSession) ResetOrforgetPassword(_address common.Address, _newaddress common.Address, _email string, _code *big.Int, _mixed_question string, _mixed_answer string) (*types.Transaction, error) {
	return _Wallet.Contract.ResetOrforgetPassword(&_Wallet.TransactOpts, _address, _newaddress, _email, _code, _mixed_question, _mixed_answer)
}

// Settestnum is a paid mutator transaction binding the contract method 0xe2da6161.
//
// Solidity: function settestnum(uint256 _number) returns()
func (_Wallet *WalletTransactor) Settestnum(opts *bind.TransactOpts, _number *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "settestnum", _number)
}

// Settestnum is a paid mutator transaction binding the contract method 0xe2da6161.
//
// Solidity: function settestnum(uint256 _number) returns()
func (_Wallet *WalletSession) Settestnum(_number *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Settestnum(&_Wallet.TransactOpts, _number)
}

// Settestnum is a paid mutator transaction binding the contract method 0xe2da6161.
//
// Solidity: function settestnum(uint256 _number) returns()
func (_Wallet *WalletTransactorSession) Settestnum(_number *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Settestnum(&_Wallet.TransactOpts, _number)
}

// Verifycode is a paid mutator transaction binding the contract method 0xf6bae2b1.
//
// Solidity: function verifycode(string _email, uint256 _code) returns()
func (_Wallet *WalletTransactor) Verifycode(opts *bind.TransactOpts, _email string, _code *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "verifycode", _email, _code)
}

// Verifycode is a paid mutator transaction binding the contract method 0xf6bae2b1.
//
// Solidity: function verifycode(string _email, uint256 _code) returns()
func (_Wallet *WalletSession) Verifycode(_email string, _code *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Verifycode(&_Wallet.TransactOpts, _email, _code)
}

// Verifycode is a paid mutator transaction binding the contract method 0xf6bae2b1.
//
// Solidity: function verifycode(string _email, uint256 _code) returns()
func (_Wallet *WalletTransactorSession) Verifycode(_email string, _code *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Verifycode(&_Wallet.TransactOpts, _email, _code)
}
