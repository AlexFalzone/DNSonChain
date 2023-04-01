// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package domainRegistry

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

// DomainRegistryMetaData contains all meta data concerning the DomainRegistry contract.
var DomainRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"certificate\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DomainCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"}],\"name\":\"DomainDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"certificate\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"DomainUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"certificate\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"createDomain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"}],\"name\":\"deleteDomain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"domains\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"certificate\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"}],\"name\":\"getCertificate\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"certificate\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"updateDomain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506120bf806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630eaf97521461005c5780631469e9141461007857806326449235146100945780635195418c146100c8578063ed0f2e75146100e4575b600080fd5b610076600480360381019061007191906112bb565b610115565b005b610092600480360381019061008d919061147f565b610285565b005b6100ae60048036038101906100a99190611542565b6107a5565b6040516100bf95949392919061165a565b60405180910390f35b6100e260048036038101906100dd919061147f565b6109a9565b005b6100fe60048036038101906100f991906112bb565b610bda565b60405161010c929190611717565b60405180910390f35b3373ffffffffffffffffffffffffffffffffffffffff166000838360405161013e92919061177e565b908152602001604051809103902060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146101c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101bd906117e3565b60405180910390fd5b600082826040516101d892919061177e565b9081526020016040518091039020600080820160006101f791906111e5565b60018201600061020791906111e5565b60028201600061021791906111e5565b60038201600090556004820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550507f9b449e8e6e58cfe4813dc8a38f233368018ddd99ddc223e2f1a55eb22eefa2368282604051610279929190611830565b60405180910390a15050565b600086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505090506000805b8251811015610363577f2e0000000000000000000000000000000000000000000000000000000000000083828151811061030f5761030e611854565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19160361035057818061034c906118b2565b9250505b808061035b906118b2565b9150506102d2565b50600081116103a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039e90611946565b60405180910390fd5b600181111561046e5760006103bb8361104a565b90503373ffffffffffffffffffffffffffffffffffffffff166000826040516103e49190611997565b908152602001604051809103902060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461046c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046390611a20565b60405180910390fd5b505b600080898960405161048192919061177e565b9081526020016040518091039020600001805461049d90611a6f565b9050111561058e5742600089896040516104b892919061177e565b908152602001604051809103902060030154111561050b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050290611b12565b60405180910390fd5b6000888860405161051d92919061177e565b90815260200160405180910390206000808201600061053c91906111e5565b60018201600061054c91906111e5565b60028201600061055c91906111e5565b60038201600090556004820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550505b6040518060a0016040528089898080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200187878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020018581526020016000851461064c5784426106479190611b32565b61066e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b81526020013373ffffffffffffffffffffffffffffffffffffffff168152506000898960405161069f92919061177e565b908152602001604051809103902060008201518160000190816106c29190611d12565b5060208201518160010190816106d89190611d12565b5060408201518160020190816106ee9190611d12565b506060820151816003015560808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050503373ffffffffffffffffffffffffffffffffffffffff167fa13d03807a9cdb641092d81b09d6b69e448c291c5e7d4b7a0c57f31a3b72b1d789898989898960405161079396959493929190611de4565b60405180910390a25050505050505050565b6000818051602081018201805184825260208301602085012081835280955050505050506000915090508060000180546107de90611a6f565b80601f016020809104026020016040519081016040528092919081815260200182805461080a90611a6f565b80156108575780601f1061082c57610100808354040283529160200191610857565b820191906000526020600020905b81548152906001019060200180831161083a57829003601f168201915b50505050509080600101805461086c90611a6f565b80601f016020809104026020016040519081016040528092919081815260200182805461089890611a6f565b80156108e55780601f106108ba576101008083540402835291602001916108e5565b820191906000526020600020905b8154815290600101906020018083116108c857829003601f168201915b5050505050908060020180546108fa90611a6f565b80601f016020809104026020016040519081016040528092919081815260200182805461092690611a6f565b80156109735780601f1061094857610100808354040283529160200191610973565b820191906000526020600020905b81548152906001019060200180831161095657829003601f168201915b5050505050908060030154908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905085565b60008087876040516109bc92919061177e565b908152602001604051809103902060000180546109d890611a6f565b905011610a1a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1190611e8e565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff1660008787604051610a4392919061177e565b908152602001604051809103902060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610acb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac2906117e3565b60405180910390fd5b838360008888604051610adf92919061177e565b90815260200160405180910390206001019182610afd929190611eb9565b508160008787604051610b1192919061177e565b90815260200160405180910390206002019081610b2e9190611d12565b5060008114610b48578042610b439190611b32565b610b6a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b60008787604051610b7c92919061177e565b9081526020016040518091039020600301819055507fb8abd760067276c7baaf5f313ea2d97676d7eef530f9c9143c377f4b8b4ca6b5868686868686604051610bca96959493929190611de4565b60405180910390a1505050505050565b6060806000808585604051610bf092919061177e565b908152602001604051809103902090506000816000018054610c1190611a6f565b905011610c53576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c4a90611e8e565b60405180910390fd5b42816003015411610c99576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9090611fd5565b60405180910390fd5b60006002826002018054610cac90611a6f565b9050610cb89190612024565b14610f1d576000816002018054610cce90611a6f565b80601f0160208091040260200160405190810160405280929190818152602001828054610cfa90611a6f565b8015610d475780601f10610d1c57610100808354040283529160200191610d47565b820191906000526020600020905b815481529060010190602001808311610d2a57829003601f168201915b50505050509050600060018251610d5e9190611b32565b67ffffffffffffffff811115610d7757610d7661131e565b5b6040519080825280601f01601f191660200182016040528015610da95781602001600182028036833780820191505090505b509050600060f81b81600081518110610dc557610dc4611854565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060005b8251811015610e7f57828181518110610e1357610e12611854565b5b602001015160f81c60f81b82600183610e2c9190611b32565b81518110610e3d57610e3c611854565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080610e77906118b2565b915050610df7565b508260010181818054610e9190611a6f565b80601f0160208091040260200160405190810160405280929190818152602001828054610ebd90611a6f565b8015610f0a5780601f10610edf57610100808354040283529160200191610f0a565b820191906000526020600020905b815481529060010190602001808311610eed57829003601f168201915b5050505050915094509450505050611043565b8060010181600201818054610f3190611a6f565b80601f0160208091040260200160405190810160405280929190818152602001828054610f5d90611a6f565b8015610faa5780601f10610f7f57610100808354040283529160200191610faa565b820191906000526020600020905b815481529060010190602001808311610f8d57829003601f168201915b50505050509150808054610fbd90611a6f565b80601f0160208091040260200160405190810160405280929190818152602001828054610fe990611a6f565b80156110365780601f1061100b57610100808354040283529160200191611036565b820191906000526020600020905b81548152906001019060200180831161101957829003601f168201915b5050505050905092509250505b9250929050565b60606000805b83518110156110da577f2e0000000000000000000000000000000000000000000000000000000000000084828151811061108d5761108c611854565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916036110c7578091506110da565b80806110d2906118b2565b915050611050565b50600060018285516110ec9190612055565b6110f69190612055565b67ffffffffffffffff81111561110f5761110e61131e565b5b6040519080825280601f01601f1916602001820160405280156111415781602001600182028036833780820191505090505b50905060005b81518110156111da57848160018561115f9190611b32565b6111699190611b32565b8151811061117a57611179611854565b5b602001015160f81c60f81b82828151811061119857611197611854565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806111d2906118b2565b915050611147565b508092505050919050565b5080546111f190611a6f565b6000825580601f106112035750611222565b601f0160209004906000526020600020908101906112219190611225565b5b50565b5b8082111561123e576000816000905550600101611226565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f84011261127b5761127a611256565b5b8235905067ffffffffffffffff8111156112985761129761125b565b5b6020830191508360018202830111156112b4576112b3611260565b5b9250929050565b600080602083850312156112d2576112d161124c565b5b600083013567ffffffffffffffff8111156112f0576112ef611251565b5b6112fc85828601611265565b92509250509250929050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6113568261130d565b810181811067ffffffffffffffff821117156113755761137461131e565b5b80604052505050565b6000611388611242565b9050611394828261134d565b919050565b600067ffffffffffffffff8211156113b4576113b361131e565b5b6113bd8261130d565b9050602081019050919050565b82818337600083830152505050565b60006113ec6113e784611399565b61137e565b90508281526020810184848401111561140857611407611308565b5b6114138482856113ca565b509392505050565b600082601f8301126114305761142f611256565b5b81356114408482602086016113d9565b91505092915050565b6000819050919050565b61145c81611449565b811461146757600080fd5b50565b60008135905061147981611453565b92915050565b6000806000806000806080878903121561149c5761149b61124c565b5b600087013567ffffffffffffffff8111156114ba576114b9611251565b5b6114c689828a01611265565b9650965050602087013567ffffffffffffffff8111156114e9576114e8611251565b5b6114f589828a01611265565b9450945050604087013567ffffffffffffffff81111561151857611517611251565b5b61152489828a0161141b565b925050606061153589828a0161146a565b9150509295509295509295565b6000602082840312156115585761155761124c565b5b600082013567ffffffffffffffff81111561157657611575611251565b5b6115828482850161141b565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156115c55780820151818401526020810190506115aa565b60008484015250505050565b60006115dc8261158b565b6115e68185611596565b93506115f68185602086016115a7565b6115ff8161130d565b840191505092915050565b61161381611449565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061164482611619565b9050919050565b61165481611639565b82525050565b600060a082019050818103600083015261167481886115d1565b9050818103602083015261168881876115d1565b9050818103604083015261169c81866115d1565b90506116ab606083018561160a565b6116b8608083018461164b565b9695505050505050565b600081519050919050565b600082825260208201905092915050565b60006116e9826116c2565b6116f381856116cd565b93506117038185602086016115a7565b61170c8161130d565b840191505092915050565b6000604082019050818103600083015261173181856115d1565b9050818103602083015261174581846116de565b90509392505050565b600081905092915050565b6000611765838561174e565b93506117728385846113ca565b82840190509392505050565b600061178b828486611759565b91508190509392505050565b7f446f6d61696e206e6f74206f776e65642062792073656e646572000000000000600082015250565b60006117cd601a83611596565b91506117d882611797565b602082019050919050565b600060208201905081810360008301526117fc816117c0565b9050919050565b600061180f8385611596565b935061181c8385846113ca565b6118258361130d565b840190509392505050565b6000602082019050818103600083015261184b818486611803565b90509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006118bd82611449565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036118ef576118ee611883565b5b600182019050919050565b7f496e76616c696420686f73746e616d6500000000000000000000000000000000600082015250565b6000611930601083611596565b915061193b826118fa565b602082019050919050565b6000602082019050818103600083015261195f81611923565b9050919050565b60006119718261158b565b61197b818561174e565b935061198b8185602086016115a7565b80840191505092915050565b60006119a38284611966565b915081905092915050565b7f506172656e7420646f6d61696e206e6f742072656769737465726564206f722060008201527f6e6f74206f776e65642062792073656e64657200000000000000000000000000602082015250565b6000611a0a603383611596565b9150611a15826119ae565b604082019050919050565b60006020820190508181036000830152611a39816119fd565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611a8757607f821691505b602082108103611a9a57611a99611a40565b5b50919050565b7f486f73746e616d6520616c7265616479207265676973746572656420616e642060008201527f6e6f742065787069726564000000000000000000000000000000000000000000602082015250565b6000611afc602b83611596565b9150611b0782611aa0565b604082019050919050565b60006020820190508181036000830152611b2b81611aef565b9050919050565b6000611b3d82611449565b9150611b4883611449565b9250828201905080821115611b6057611b5f611883565b5b92915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302611bc87fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611b8b565b611bd28683611b8b565b95508019841693508086168417925050509392505050565b6000819050919050565b6000611c0f611c0a611c0584611449565b611bea565b611449565b9050919050565b6000819050919050565b611c2983611bf4565b611c3d611c3582611c16565b848454611b98565b825550505050565b600090565b611c52611c45565b611c5d818484611c20565b505050565b5b81811015611c8157611c76600082611c4a565b600181019050611c63565b5050565b601f821115611cc657611c9781611b66565b611ca084611b7b565b81016020851015611caf578190505b611cc3611cbb85611b7b565b830182611c62565b50505b505050565b600082821c905092915050565b6000611ce960001984600802611ccb565b1980831691505092915050565b6000611d028383611cd8565b9150826002028217905092915050565b611d1b8261158b565b67ffffffffffffffff811115611d3457611d3361131e565b5b611d3e8254611a6f565b611d49828285611c85565b600060209050601f831160018114611d7c5760008415611d6a578287015190505b611d748582611cf6565b865550611ddc565b601f198416611d8a86611b66565b60005b82811015611db257848901518255600182019150602085019450602081019050611d8d565b86831015611dcf5784890151611dcb601f891682611cd8565b8355505b6001600288020188555050505b505050505050565b60006080820190508181036000830152611dff81888a611803565b90508181036020830152611e14818688611803565b90508181036040830152611e2881856115d1565b9050611e37606083018461160a565b979650505050505050565b7f446f6d61696e206e6f7420666f756e6400000000000000000000000000000000600082015250565b6000611e78601083611596565b9150611e8382611e42565b602082019050919050565b60006020820190508181036000830152611ea781611e6b565b9050919050565b600082905092915050565b611ec38383611eae565b67ffffffffffffffff811115611edc57611edb61131e565b5b611ee68254611a6f565b611ef1828285611c85565b6000601f831160018114611f205760008415611f0e578287013590505b611f188582611cf6565b865550611f80565b601f198416611f2e86611b66565b60005b82811015611f5657848901358255600182019150602085019450602081019050611f31565b86831015611f735784890135611f6f601f891682611cd8565b8355505b6001600288020188555050505b50505050505050565b7f4365727469666963617465206861732065787069726564000000000000000000600082015250565b6000611fbf601783611596565b9150611fca82611f89565b602082019050919050565b60006020820190508181036000830152611fee81611fb2565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061202f82611449565b915061203a83611449565b92508261204a57612049611ff5565b5b828206905092915050565b600061206082611449565b915061206b83611449565b925082820390508181111561208357612082611883565b5b9291505056fea26469706673582212200b66f60f66772a43b38119ee5077a8c51915d34601f645bd2e82b81b8876c5e764736f6c63430008130033",
}

// DomainRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use DomainRegistryMetaData.ABI instead.
var DomainRegistryABI = DomainRegistryMetaData.ABI

// DomainRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DomainRegistryMetaData.Bin instead.
var DomainRegistryBin = DomainRegistryMetaData.Bin

// DeployDomainRegistry deploys a new Ethereum contract, binding an instance of DomainRegistry to it.
func DeployDomainRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DomainRegistry, error) {
	parsed, err := DomainRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DomainRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DomainRegistry{DomainRegistryCaller: DomainRegistryCaller{contract: contract}, DomainRegistryTransactor: DomainRegistryTransactor{contract: contract}, DomainRegistryFilterer: DomainRegistryFilterer{contract: contract}}, nil
}

// DomainRegistry is an auto generated Go binding around an Ethereum contract.
type DomainRegistry struct {
	DomainRegistryCaller     // Read-only binding to the contract
	DomainRegistryTransactor // Write-only binding to the contract
	DomainRegistryFilterer   // Log filterer for contract events
}

// DomainRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DomainRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomainRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DomainRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomainRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DomainRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomainRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DomainRegistrySession struct {
	Contract     *DomainRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DomainRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DomainRegistryCallerSession struct {
	Contract *DomainRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DomainRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DomainRegistryTransactorSession struct {
	Contract     *DomainRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DomainRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DomainRegistryRaw struct {
	Contract *DomainRegistry // Generic contract binding to access the raw methods on
}

// DomainRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DomainRegistryCallerRaw struct {
	Contract *DomainRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// DomainRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DomainRegistryTransactorRaw struct {
	Contract *DomainRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDomainRegistry creates a new instance of DomainRegistry, bound to a specific deployed contract.
func NewDomainRegistry(address common.Address, backend bind.ContractBackend) (*DomainRegistry, error) {
	contract, err := bindDomainRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DomainRegistry{DomainRegistryCaller: DomainRegistryCaller{contract: contract}, DomainRegistryTransactor: DomainRegistryTransactor{contract: contract}, DomainRegistryFilterer: DomainRegistryFilterer{contract: contract}}, nil
}

// NewDomainRegistryCaller creates a new read-only instance of DomainRegistry, bound to a specific deployed contract.
func NewDomainRegistryCaller(address common.Address, caller bind.ContractCaller) (*DomainRegistryCaller, error) {
	contract, err := bindDomainRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DomainRegistryCaller{contract: contract}, nil
}

// NewDomainRegistryTransactor creates a new write-only instance of DomainRegistry, bound to a specific deployed contract.
func NewDomainRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*DomainRegistryTransactor, error) {
	contract, err := bindDomainRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DomainRegistryTransactor{contract: contract}, nil
}

// NewDomainRegistryFilterer creates a new log filterer instance of DomainRegistry, bound to a specific deployed contract.
func NewDomainRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*DomainRegistryFilterer, error) {
	contract, err := bindDomainRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DomainRegistryFilterer{contract: contract}, nil
}

// bindDomainRegistry binds a generic wrapper to an already deployed contract.
func bindDomainRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DomainRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DomainRegistry *DomainRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DomainRegistry.Contract.DomainRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DomainRegistry *DomainRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DomainRegistry.Contract.DomainRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DomainRegistry *DomainRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DomainRegistry.Contract.DomainRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DomainRegistry *DomainRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DomainRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DomainRegistry *DomainRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DomainRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DomainRegistry *DomainRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DomainRegistry.Contract.contract.Transact(opts, method, params...)
}

// Domains is a free data retrieval call binding the contract method 0x26449235.
//
// Solidity: function domains(string ) view returns(string hostname, string ip, string certificate, uint256 data, address owner)
func (_DomainRegistry *DomainRegistryCaller) Domains(opts *bind.CallOpts, arg0 string) (struct {
	Hostname    string
	Ip          string
	Certificate string
	Data        *big.Int
	Owner       common.Address
}, error) {
	var out []interface{}
	err := _DomainRegistry.contract.Call(opts, &out, "domains", arg0)

	outstruct := new(struct {
		Hostname    string
		Ip          string
		Certificate string
		Data        *big.Int
		Owner       common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hostname = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Ip = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Certificate = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Data = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Domains is a free data retrieval call binding the contract method 0x26449235.
//
// Solidity: function domains(string ) view returns(string hostname, string ip, string certificate, uint256 data, address owner)
func (_DomainRegistry *DomainRegistrySession) Domains(arg0 string) (struct {
	Hostname    string
	Ip          string
	Certificate string
	Data        *big.Int
	Owner       common.Address
}, error) {
	return _DomainRegistry.Contract.Domains(&_DomainRegistry.CallOpts, arg0)
}

// Domains is a free data retrieval call binding the contract method 0x26449235.
//
// Solidity: function domains(string ) view returns(string hostname, string ip, string certificate, uint256 data, address owner)
func (_DomainRegistry *DomainRegistryCallerSession) Domains(arg0 string) (struct {
	Hostname    string
	Ip          string
	Certificate string
	Data        *big.Int
	Owner       common.Address
}, error) {
	return _DomainRegistry.Contract.Domains(&_DomainRegistry.CallOpts, arg0)
}

// GetCertificate is a free data retrieval call binding the contract method 0xed0f2e75.
//
// Solidity: function getCertificate(string hostname) view returns(string, bytes)
func (_DomainRegistry *DomainRegistryCaller) GetCertificate(opts *bind.CallOpts, hostname string) (string, []byte, error) {
	var out []interface{}
	err := _DomainRegistry.contract.Call(opts, &out, "getCertificate", hostname)

	if err != nil {
		return *new(string), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// GetCertificate is a free data retrieval call binding the contract method 0xed0f2e75.
//
// Solidity: function getCertificate(string hostname) view returns(string, bytes)
func (_DomainRegistry *DomainRegistrySession) GetCertificate(hostname string) (string, []byte, error) {
	return _DomainRegistry.Contract.GetCertificate(&_DomainRegistry.CallOpts, hostname)
}

// GetCertificate is a free data retrieval call binding the contract method 0xed0f2e75.
//
// Solidity: function getCertificate(string hostname) view returns(string, bytes)
func (_DomainRegistry *DomainRegistryCallerSession) GetCertificate(hostname string) (string, []byte, error) {
	return _DomainRegistry.Contract.GetCertificate(&_DomainRegistry.CallOpts, hostname)
}

// CreateDomain is a paid mutator transaction binding the contract method 0x1469e914.
//
// Solidity: function createDomain(string hostname, string ip, string certificate, uint256 data) returns()
func (_DomainRegistry *DomainRegistryTransactor) CreateDomain(opts *bind.TransactOpts, hostname string, ip string, certificate string, data *big.Int) (*types.Transaction, error) {
	return _DomainRegistry.contract.Transact(opts, "createDomain", hostname, ip, certificate, data)
}

// CreateDomain is a paid mutator transaction binding the contract method 0x1469e914.
//
// Solidity: function createDomain(string hostname, string ip, string certificate, uint256 data) returns()
func (_DomainRegistry *DomainRegistrySession) CreateDomain(hostname string, ip string, certificate string, data *big.Int) (*types.Transaction, error) {
	return _DomainRegistry.Contract.CreateDomain(&_DomainRegistry.TransactOpts, hostname, ip, certificate, data)
}

// CreateDomain is a paid mutator transaction binding the contract method 0x1469e914.
//
// Solidity: function createDomain(string hostname, string ip, string certificate, uint256 data) returns()
func (_DomainRegistry *DomainRegistryTransactorSession) CreateDomain(hostname string, ip string, certificate string, data *big.Int) (*types.Transaction, error) {
	return _DomainRegistry.Contract.CreateDomain(&_DomainRegistry.TransactOpts, hostname, ip, certificate, data)
}

// DeleteDomain is a paid mutator transaction binding the contract method 0x0eaf9752.
//
// Solidity: function deleteDomain(string hostname) returns()
func (_DomainRegistry *DomainRegistryTransactor) DeleteDomain(opts *bind.TransactOpts, hostname string) (*types.Transaction, error) {
	return _DomainRegistry.contract.Transact(opts, "deleteDomain", hostname)
}

// DeleteDomain is a paid mutator transaction binding the contract method 0x0eaf9752.
//
// Solidity: function deleteDomain(string hostname) returns()
func (_DomainRegistry *DomainRegistrySession) DeleteDomain(hostname string) (*types.Transaction, error) {
	return _DomainRegistry.Contract.DeleteDomain(&_DomainRegistry.TransactOpts, hostname)
}

// DeleteDomain is a paid mutator transaction binding the contract method 0x0eaf9752.
//
// Solidity: function deleteDomain(string hostname) returns()
func (_DomainRegistry *DomainRegistryTransactorSession) DeleteDomain(hostname string) (*types.Transaction, error) {
	return _DomainRegistry.Contract.DeleteDomain(&_DomainRegistry.TransactOpts, hostname)
}

// UpdateDomain is a paid mutator transaction binding the contract method 0x5195418c.
//
// Solidity: function updateDomain(string hostname, string ip, string certificate, uint256 data) returns()
func (_DomainRegistry *DomainRegistryTransactor) UpdateDomain(opts *bind.TransactOpts, hostname string, ip string, certificate string, data *big.Int) (*types.Transaction, error) {
	return _DomainRegistry.contract.Transact(opts, "updateDomain", hostname, ip, certificate, data)
}

// UpdateDomain is a paid mutator transaction binding the contract method 0x5195418c.
//
// Solidity: function updateDomain(string hostname, string ip, string certificate, uint256 data) returns()
func (_DomainRegistry *DomainRegistrySession) UpdateDomain(hostname string, ip string, certificate string, data *big.Int) (*types.Transaction, error) {
	return _DomainRegistry.Contract.UpdateDomain(&_DomainRegistry.TransactOpts, hostname, ip, certificate, data)
}

// UpdateDomain is a paid mutator transaction binding the contract method 0x5195418c.
//
// Solidity: function updateDomain(string hostname, string ip, string certificate, uint256 data) returns()
func (_DomainRegistry *DomainRegistryTransactorSession) UpdateDomain(hostname string, ip string, certificate string, data *big.Int) (*types.Transaction, error) {
	return _DomainRegistry.Contract.UpdateDomain(&_DomainRegistry.TransactOpts, hostname, ip, certificate, data)
}

// DomainRegistryDomainCreatedIterator is returned from FilterDomainCreated and is used to iterate over the raw logs and unpacked data for DomainCreated events raised by the DomainRegistry contract.
type DomainRegistryDomainCreatedIterator struct {
	Event *DomainRegistryDomainCreated // Event containing the contract specifics and raw log

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
func (it *DomainRegistryDomainCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomainRegistryDomainCreated)
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
		it.Event = new(DomainRegistryDomainCreated)
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
func (it *DomainRegistryDomainCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomainRegistryDomainCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomainRegistryDomainCreated represents a DomainCreated event raised by the DomainRegistry contract.
type DomainRegistryDomainCreated struct {
	Hostname    string
	Ip          string
	Certificate string
	Data        *big.Int
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDomainCreated is a free log retrieval operation binding the contract event 0xa13d03807a9cdb641092d81b09d6b69e448c291c5e7d4b7a0c57f31a3b72b1d7.
//
// Solidity: event DomainCreated(string hostname, string ip, string certificate, uint256 data, address indexed owner)
func (_DomainRegistry *DomainRegistryFilterer) FilterDomainCreated(opts *bind.FilterOpts, owner []common.Address) (*DomainRegistryDomainCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _DomainRegistry.contract.FilterLogs(opts, "DomainCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &DomainRegistryDomainCreatedIterator{contract: _DomainRegistry.contract, event: "DomainCreated", logs: logs, sub: sub}, nil
}

// WatchDomainCreated is a free log subscription operation binding the contract event 0xa13d03807a9cdb641092d81b09d6b69e448c291c5e7d4b7a0c57f31a3b72b1d7.
//
// Solidity: event DomainCreated(string hostname, string ip, string certificate, uint256 data, address indexed owner)
func (_DomainRegistry *DomainRegistryFilterer) WatchDomainCreated(opts *bind.WatchOpts, sink chan<- *DomainRegistryDomainCreated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _DomainRegistry.contract.WatchLogs(opts, "DomainCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomainRegistryDomainCreated)
				if err := _DomainRegistry.contract.UnpackLog(event, "DomainCreated", log); err != nil {
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

// ParseDomainCreated is a log parse operation binding the contract event 0xa13d03807a9cdb641092d81b09d6b69e448c291c5e7d4b7a0c57f31a3b72b1d7.
//
// Solidity: event DomainCreated(string hostname, string ip, string certificate, uint256 data, address indexed owner)
func (_DomainRegistry *DomainRegistryFilterer) ParseDomainCreated(log types.Log) (*DomainRegistryDomainCreated, error) {
	event := new(DomainRegistryDomainCreated)
	if err := _DomainRegistry.contract.UnpackLog(event, "DomainCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DomainRegistryDomainDeletedIterator is returned from FilterDomainDeleted and is used to iterate over the raw logs and unpacked data for DomainDeleted events raised by the DomainRegistry contract.
type DomainRegistryDomainDeletedIterator struct {
	Event *DomainRegistryDomainDeleted // Event containing the contract specifics and raw log

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
func (it *DomainRegistryDomainDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomainRegistryDomainDeleted)
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
		it.Event = new(DomainRegistryDomainDeleted)
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
func (it *DomainRegistryDomainDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomainRegistryDomainDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomainRegistryDomainDeleted represents a DomainDeleted event raised by the DomainRegistry contract.
type DomainRegistryDomainDeleted struct {
	Hostname string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDomainDeleted is a free log retrieval operation binding the contract event 0x9b449e8e6e58cfe4813dc8a38f233368018ddd99ddc223e2f1a55eb22eefa236.
//
// Solidity: event DomainDeleted(string hostname)
func (_DomainRegistry *DomainRegistryFilterer) FilterDomainDeleted(opts *bind.FilterOpts) (*DomainRegistryDomainDeletedIterator, error) {

	logs, sub, err := _DomainRegistry.contract.FilterLogs(opts, "DomainDeleted")
	if err != nil {
		return nil, err
	}
	return &DomainRegistryDomainDeletedIterator{contract: _DomainRegistry.contract, event: "DomainDeleted", logs: logs, sub: sub}, nil
}

// WatchDomainDeleted is a free log subscription operation binding the contract event 0x9b449e8e6e58cfe4813dc8a38f233368018ddd99ddc223e2f1a55eb22eefa236.
//
// Solidity: event DomainDeleted(string hostname)
func (_DomainRegistry *DomainRegistryFilterer) WatchDomainDeleted(opts *bind.WatchOpts, sink chan<- *DomainRegistryDomainDeleted) (event.Subscription, error) {

	logs, sub, err := _DomainRegistry.contract.WatchLogs(opts, "DomainDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomainRegistryDomainDeleted)
				if err := _DomainRegistry.contract.UnpackLog(event, "DomainDeleted", log); err != nil {
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

// ParseDomainDeleted is a log parse operation binding the contract event 0x9b449e8e6e58cfe4813dc8a38f233368018ddd99ddc223e2f1a55eb22eefa236.
//
// Solidity: event DomainDeleted(string hostname)
func (_DomainRegistry *DomainRegistryFilterer) ParseDomainDeleted(log types.Log) (*DomainRegistryDomainDeleted, error) {
	event := new(DomainRegistryDomainDeleted)
	if err := _DomainRegistry.contract.UnpackLog(event, "DomainDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DomainRegistryDomainUpdatedIterator is returned from FilterDomainUpdated and is used to iterate over the raw logs and unpacked data for DomainUpdated events raised by the DomainRegistry contract.
type DomainRegistryDomainUpdatedIterator struct {
	Event *DomainRegistryDomainUpdated // Event containing the contract specifics and raw log

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
func (it *DomainRegistryDomainUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DomainRegistryDomainUpdated)
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
		it.Event = new(DomainRegistryDomainUpdated)
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
func (it *DomainRegistryDomainUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DomainRegistryDomainUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DomainRegistryDomainUpdated represents a DomainUpdated event raised by the DomainRegistry contract.
type DomainRegistryDomainUpdated struct {
	Hostname    string
	Ip          string
	Certificate string
	Data        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDomainUpdated is a free log retrieval operation binding the contract event 0xb8abd760067276c7baaf5f313ea2d97676d7eef530f9c9143c377f4b8b4ca6b5.
//
// Solidity: event DomainUpdated(string hostname, string ip, string certificate, uint256 data)
func (_DomainRegistry *DomainRegistryFilterer) FilterDomainUpdated(opts *bind.FilterOpts) (*DomainRegistryDomainUpdatedIterator, error) {

	logs, sub, err := _DomainRegistry.contract.FilterLogs(opts, "DomainUpdated")
	if err != nil {
		return nil, err
	}
	return &DomainRegistryDomainUpdatedIterator{contract: _DomainRegistry.contract, event: "DomainUpdated", logs: logs, sub: sub}, nil
}

// WatchDomainUpdated is a free log subscription operation binding the contract event 0xb8abd760067276c7baaf5f313ea2d97676d7eef530f9c9143c377f4b8b4ca6b5.
//
// Solidity: event DomainUpdated(string hostname, string ip, string certificate, uint256 data)
func (_DomainRegistry *DomainRegistryFilterer) WatchDomainUpdated(opts *bind.WatchOpts, sink chan<- *DomainRegistryDomainUpdated) (event.Subscription, error) {

	logs, sub, err := _DomainRegistry.contract.WatchLogs(opts, "DomainUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DomainRegistryDomainUpdated)
				if err := _DomainRegistry.contract.UnpackLog(event, "DomainUpdated", log); err != nil {
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

// ParseDomainUpdated is a log parse operation binding the contract event 0xb8abd760067276c7baaf5f313ea2d97676d7eef530f9c9143c377f4b8b4ca6b5.
//
// Solidity: event DomainUpdated(string hostname, string ip, string certificate, uint256 data)
func (_DomainRegistry *DomainRegistryFilterer) ParseDomainUpdated(log types.Log) (*DomainRegistryDomainUpdated, error) {
	event := new(DomainRegistryDomainUpdated)
	if err := _DomainRegistry.contract.UnpackLog(event, "DomainUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
