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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidInput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gettestnum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_code\",\"type\":\"uint256\"}],\"name\":\"isverified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newaddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_code\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_mixed_question\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_mixed_answer\",\"type\":\"string\"}],\"name\":\"resetOrforgetPassword\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_code\",\"type\":\"uint256\"}],\"name\":\"verifycode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200199838038062001998833981810160405281019062000037919062000150565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036200009e576040517fb4fa3fb300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000182565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200011882620000eb565b9050919050565b6200012a816200010b565b81146200013657600080fd5b50565b6000815190506200014a816200011f565b92915050565b600060208284031215620001695762000168620000e6565b5b6000620001798482850162000139565b91505092915050565b61180680620001926000396000f3fe6080604052600436106100705760003560e01c8063814494b11161004e578063814494b11461010c5780638da5cb5b146101355780639e5d4c4914610160578063f6bae2b11461019057610070565b8063264bd688146100755780634f1543fc146100b257806373d4a13a146100dd575b600080fd5b34801561008157600080fd5b5061009c60048036038101906100979190610e5f565b6101b9565b6040516100a99190610ed6565b60405180910390f35b3480156100be57600080fd5b506100c7610200565b6040516100d49190610f00565b60405180910390f35b3480156100e957600080fd5b506100f2610209565b604051610103959493929190610fdb565b60405180910390f35b34801561011857600080fd5b50610133600480360381019061012e9190611076565b610293565b005b34801561014157600080fd5b5061014a61084b565b6040516101579190611157565b60405180910390f35b61017a600480360381019061017591906111d2565b6108b2565b604051610187919061129b565b60405180910390f35b34801561019c57600080fd5b506101b760048036038101906101b29190610e5f565b610936565b005b60006001836040516101cb91906112f9565b9081526020016040518091039020600083815260200190815260200160002060009054906101000a900460ff16905092915050565b60006005905090565b60006060806060806000606067ffffffffffffffff81111561022e5761022d610cfe565b5b6040519080825280601f01601f1916602001820160405280156102605781602001600182028036833780820191505090505b5090506060604d60208301303c8080602001905181019061028191906113be565b95509550955095509550509091929394565b600360159054906101000a900460ff166102d45760008060008060006102b7610209565b945094509450945094506102ce8585858585610a08565b50505050505b60008060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060c001604052908160008201805461032f906114d8565b80601f016020809104026020016040519081016040528092919081815260200182805461035b906114d8565b80156103a85780601f1061037d576101008083540402835291602001916103a8565b820191906000526020600020905b81548152906001019060200180831161038b57829003601f168201915b50505050508152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382018054610421906114d8565b80601f016020809104026020016040519081016040528092919081815260200182805461044d906114d8565b801561049a5780601f1061046f5761010080835404028352916020019161049a565b820191906000526020600020905b81548152906001019060200180831161047d57829003601f168201915b505050505081526020016004820180546104b3906114d8565b80601f01602080910402602001604051908101604052809291908181526020018280546104df906114d8565b801561052c5780601f106105015761010080835404028352916020019161052c565b820191906000526020600020905b81548152906001019060200180831161050f57829003601f168201915b50505050508152602001600582018054610545906114d8565b80601f0160208091040260200160405190810160405280929190818152602001828054610571906114d8565b80156105be5780601f10610593576101008083540402835291602001916105be565b820191906000526020600020905b8154815290600101906020018083116105a157829003601f168201915b50505050508152505090506105d385856101b9565b801561061157506105e261084b565b73ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16145b80156106275750610626858260000151610b5e565b5b801561063d575061063c838260800151610b5e565b5b80156106535750610652828260a00151610b5e565b5b61065c57600080fd5b8381602001818152505085816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050806000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000190816106f291906116b5565b506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301908161075991906116b5565b50608082015181600401908161076f91906116b5565b5060a082015181600501908161078591906116b5565b509050506000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160006107d89190610c72565b60018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556003820160006108179190610c72565b6004820160006108279190610c72565b6005820160006108379190610c72565b505061084286610b86565b50505050505050565b6000600360149054906101000a900460ff161561088c57600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506108af565b600080600080600061089c610209565b9450945094509450945084955050505050505b90565b60606108bc61084b565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610920576040517fea8e4eb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61092c85858585610be5565b9050949350505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109bd576040517fea8e4eb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180836040516109ce91906112f9565b9081526020016040518091039020600083815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60006040518060c00160405280868152602001600081526020018773ffffffffffffffffffffffffffffffffffffffff168152602001858152602001848152602001838152509050806000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000019081610aa491906116b5565b506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003019081610b0b91906116b5565b506080820151816004019081610b2191906116b5565b5060a0820151816005019081610b3791906116b5565b509050506001600360156101000a81548160ff021916908315150217905550505050505050565b600081518351148015610b7e575081805190602001208380519060200120145b905092915050565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600360146101000a81548160ff02191690831515021790555050565b606060008573ffffffffffffffffffffffffffffffffffffffff16858585604051610c119291906117b7565b60006040518083038185875af1925050503d8060008114610c4e576040519150601f19603f3d011682016040523d82523d6000602084013e610c53565b606091505b50809350819250505080610c6957815160208301fd5b50949350505050565b508054610c7e906114d8565b6000825580601f10610c905750610caf565b601f016020900490600052602060002090810190610cae9190610cb2565b5b50565b5b80821115610ccb576000816000905550600101610cb3565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610d3682610ced565b810181811067ffffffffffffffff82111715610d5557610d54610cfe565b5b80604052505050565b6000610d68610ccf565b9050610d748282610d2d565b919050565b600067ffffffffffffffff821115610d9457610d93610cfe565b5b610d9d82610ced565b9050602081019050919050565b82818337600083830152505050565b6000610dcc610dc784610d79565b610d5e565b905082815260208101848484011115610de857610de7610ce8565b5b610df3848285610daa565b509392505050565b600082601f830112610e1057610e0f610ce3565b5b8135610e20848260208601610db9565b91505092915050565b6000819050919050565b610e3c81610e29565b8114610e4757600080fd5b50565b600081359050610e5981610e33565b92915050565b60008060408385031215610e7657610e75610cd9565b5b600083013567ffffffffffffffff811115610e9457610e93610cde565b5b610ea085828601610dfb565b9250506020610eb185828601610e4a565b9150509250929050565b60008115159050919050565b610ed081610ebb565b82525050565b6000602082019050610eeb6000830184610ec7565b92915050565b610efa81610e29565b82525050565b6000602082019050610f156000830184610ef1565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610f4682610f1b565b9050919050565b610f5681610f3b565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610f96578082015181840152602081019050610f7b565b60008484015250505050565b6000610fad82610f5c565b610fb78185610f67565b9350610fc7818560208601610f78565b610fd081610ced565b840191505092915050565b600060a082019050610ff06000830188610f4d565b81810360208301526110028187610fa2565b905081810360408301526110168186610fa2565b9050818103606083015261102a8185610fa2565b9050818103608083015261103e8184610fa2565b90509695505050505050565b61105381610f3b565b811461105e57600080fd5b50565b6000813590506110708161104a565b92915050565b60008060008060008060c0878903121561109357611092610cd9565b5b60006110a189828a01611061565b96505060206110b289828a01611061565b955050604087013567ffffffffffffffff8111156110d3576110d2610cde565b5b6110df89828a01610dfb565b94505060606110f089828a01610e4a565b935050608087013567ffffffffffffffff81111561111157611110610cde565b5b61111d89828a01610dfb565b92505060a087013567ffffffffffffffff81111561113e5761113d610cde565b5b61114a89828a01610dfb565b9150509295509295509295565b600060208201905061116c6000830184610f4d565b92915050565b600080fd5b600080fd5b60008083601f84011261119257611191610ce3565b5b8235905067ffffffffffffffff8111156111af576111ae611172565b5b6020830191508360018202830111156111cb576111ca611177565b5b9250929050565b600080600080606085870312156111ec576111eb610cd9565b5b60006111fa87828801611061565b945050602061120b87828801610e4a565b935050604085013567ffffffffffffffff81111561122c5761122b610cde565b5b6112388782880161117c565b925092505092959194509250565b600081519050919050565b600082825260208201905092915050565b600061126d82611246565b6112778185611251565b9350611287818560208601610f78565b61129081610ced565b840191505092915050565b600060208201905081810360008301526112b58184611262565b905092915050565b600081905092915050565b60006112d382610f5c565b6112dd81856112bd565b93506112ed818560208601610f78565b80840191505092915050565b600061130582846112c8565b915081905092915050565b600061131b82610f1b565b9050919050565b61132b81611310565b811461133657600080fd5b50565b60008151905061134881611322565b92915050565b600061136161135c84610d79565b610d5e565b90508281526020810184848401111561137d5761137c610ce8565b5b611388848285610f78565b509392505050565b600082601f8301126113a5576113a4610ce3565b5b81516113b584826020860161134e565b91505092915050565b600080600080600060a086880312156113da576113d9610cd9565b5b60006113e888828901611339565b955050602086015167ffffffffffffffff81111561140957611408610cde565b5b61141588828901611390565b945050604086015167ffffffffffffffff81111561143657611435610cde565b5b61144288828901611390565b935050606086015167ffffffffffffffff81111561146357611462610cde565b5b61146f88828901611390565b925050608086015167ffffffffffffffff8111156114905761148f610cde565b5b61149c88828901611390565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806114f057607f821691505b602082108103611503576115026114a9565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261156b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261152e565b611575868361152e565b95508019841693508086168417925050509392505050565b6000819050919050565b60006115b26115ad6115a884610e29565b61158d565b610e29565b9050919050565b6000819050919050565b6115cc83611597565b6115e06115d8826115b9565b84845461153b565b825550505050565b600090565b6115f56115e8565b6116008184846115c3565b505050565b5b81811015611624576116196000826115ed565b600181019050611606565b5050565b601f8211156116695761163a81611509565b6116438461151e565b81016020851015611652578190505b61166661165e8561151e565b830182611605565b50505b505050565b600082821c905092915050565b600061168c6000198460080261166e565b1980831691505092915050565b60006116a5838361167b565b9150826002028217905092915050565b6116be82610f5c565b67ffffffffffffffff8111156116d7576116d6610cfe565b5b6116e182546114d8565b6116ec828285611628565b600060209050601f83116001811461171f576000841561170d578287015190505b6117178582611699565b86555061177f565b601f19841661172d86611509565b60005b8281101561175557848901518255600182019150602085019450602081019050611730565b86831015611772578489015161176e601f89168261167b565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b600061179e8385611787565b93506117ab838584610daa565b82840190509392505050565b60006117c4828486611792565b9150819050939250505056fea2646970667358221220ba3a77797581b611810efc1c0a68e68df97d5859cb9e60047ef7785d2ad569d764736f6c63430008120033",
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
// Solidity: function gettestnum() pure returns(uint256)
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
// Solidity: function gettestnum() pure returns(uint256)
func (_Wallet *WalletSession) Gettestnum() (*big.Int, error) {
	return _Wallet.Contract.Gettestnum(&_Wallet.CallOpts)
}

// Gettestnum is a free data retrieval call binding the contract method 0x4f1543fc.
//
// Solidity: function gettestnum() pure returns(uint256)
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
