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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"certificate\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DomainCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"}],\"name\":\"DomainDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"certificate\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"DomainUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"certificate\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"createDomain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"}],\"name\":\"deleteDomain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"domains\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"certificate\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"}],\"name\":\"getCertificate\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"certificate\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"updateDomain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611ed6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630eaf97521461005c57806326449235146100785780632ad09666146100ac578063e532fb98146100c8578063ed0f2e75146100e4575b600080fd5b61007660048036038101906100719190610f31565b610115565b005b610092600480360381019061008d91906110bf565b610285565b6040516100a3959493929190611236565b60405180910390f35b6100c660048036038101906100c19190611320565b610489565b005b6100e260048036038101906100dd9190611320565b6108d0565b005b6100fe60048036038101906100f99190610f31565b610a95565b60405161010c9291906113e9565b60405180910390f35b3373ffffffffffffffffffffffffffffffffffffffff166000838360405161013e929190611450565b908152602001604051809103902060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146101c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101bd906114b5565b60405180910390fd5b600082826040516101d8929190611450565b9081526020016040518091039020600080820160006101f79190610e1b565b6001820160006102079190610e1b565b6002820160006102179190610e5b565b60038201600090556004820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550507f9b449e8e6e58cfe4813dc8a38f233368018ddd99ddc223e2f1a55eb22eefa2368282604051610279929190611502565b60405180910390a15050565b6000818051602081018201805184825260208301602085012081835280955050505050506000915090508060000180546102be90611555565b80601f01602080910402602001604051908101604052809291908181526020018280546102ea90611555565b80156103375780601f1061030c57610100808354040283529160200191610337565b820191906000526020600020905b81548152906001019060200180831161031a57829003601f168201915b50505050509080600101805461034c90611555565b80601f016020809104026020016040519081016040528092919081815260200182805461037890611555565b80156103c55780601f1061039a576101008083540402835291602001916103c5565b820191906000526020600020905b8154815290600101906020018083116103a857829003601f168201915b5050505050908060020180546103da90611555565b80601f016020809104026020016040519081016040528092919081815260200182805461040690611555565b80156104535780601f1061042857610100808354040283529160200191610453565b820191906000526020600020905b81548152906001019060200180831161043657829003601f168201915b5050505050908060030154908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905085565b600087878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505090506000805b8251811015610567577f2e0000000000000000000000000000000000000000000000000000000000000083828151811061051357610512611586565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191603610554578180610550906115e4565b9250505b808061055f906115e4565b9150506104d6565b50600081116105ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a290611678565b60405180910390fd5b60018111156106725760006105bf83610c80565b90503373ffffffffffffffffffffffffffffffffffffffff166000826040516105e891906116c9565b908152602001604051809103902060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610670576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066790611752565b60405180910390fd5b505b6040518060a001604052808a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200188888080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200186868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020016000851461077457844261076f9190611772565b610796565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b81526020013373ffffffffffffffffffffffffffffffffffffffff1681525060008a8a6040516107c7929190611450565b908152602001604051809103902060008201518160000190816107ea9190611952565b5060208201518160010190816108009190611952565b5060408201518160020190816108169190611a7f565b506060820151816003015560808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050503373ffffffffffffffffffffffffffffffffffffffff167fd694233903b144a99ef1896eb8f5ebff710c263412bfcff1b8d37d5c574e1d108a8a8a8a8a8a8a6040516108bd9796959493929190611b7e565b60405180910390a2505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff16600088886040516108f9929190611450565b908152602001604051809103902060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610981576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610978906114b5565b60405180910390fd5b848460008989604051610995929190611450565b908152602001604051809103902060010191826109b3929190611be9565b508282600089896040516109c8929190611450565b908152602001604051809103902060020191826109e6929190611cc4565b5060008114610a005780426109fb9190611772565b610a22565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b60008888604051610a34929190611450565b9081526020016040518091039020600301819055507f6ee95bc81bd859a73439a27fc26eb56779fa812f28d8271df16444b27b0e02fd87878787878787604051610a849796959493929190611b7e565b60405180910390a150505050505050565b6060806000808585604051610aab929190611450565b908152602001604051809103902090506000816000018054610acc90611555565b905011610b0e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0590611de0565b60405180910390fd5b42816003015411610b54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4b90611e4c565b60405180910390fd5b8060010181600201818054610b6890611555565b80601f0160208091040260200160405190810160405280929190818152602001828054610b9490611555565b8015610be15780601f10610bb657610100808354040283529160200191610be1565b820191906000526020600020905b815481529060010190602001808311610bc457829003601f168201915b50505050509150808054610bf490611555565b80601f0160208091040260200160405190810160405280929190818152602001828054610c2090611555565b8015610c6d5780601f10610c4257610100808354040283529160200191610c6d565b820191906000526020600020905b815481529060010190602001808311610c5057829003601f168201915b5050505050905092509250509250929050565b60606000805b8351811015610d10577f2e00000000000000000000000000000000000000000000000000000000000000848281518110610cc357610cc2611586565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191603610cfd57809150610d10565b8080610d08906115e4565b915050610c86565b5060006001828551610d229190611e6c565b610d2c9190611e6c565b67ffffffffffffffff811115610d4557610d44610f94565b5b6040519080825280601f01601f191660200182016040528015610d775781602001600182028036833780820191505090505b50905060005b8151811015610e10578481600185610d959190611772565b610d9f9190611772565b81518110610db057610daf611586565b5b602001015160f81c60f81b828281518110610dce57610dcd611586565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080610e08906115e4565b915050610d7d565b508092505050919050565b508054610e2790611555565b6000825580601f10610e395750610e58565b601f016020900490600052602060002090810190610e579190610e9b565b5b50565b508054610e6790611555565b6000825580601f10610e795750610e98565b601f016020900490600052602060002090810190610e979190610e9b565b5b50565b5b80821115610eb4576000816000905550600101610e9c565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f840112610ef157610ef0610ecc565b5b8235905067ffffffffffffffff811115610f0e57610f0d610ed1565b5b602083019150836001820283011115610f2a57610f29610ed6565b5b9250929050565b60008060208385031215610f4857610f47610ec2565b5b600083013567ffffffffffffffff811115610f6657610f65610ec7565b5b610f7285828601610edb565b92509250509250929050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610fcc82610f83565b810181811067ffffffffffffffff82111715610feb57610fea610f94565b5b80604052505050565b6000610ffe610eb8565b905061100a8282610fc3565b919050565b600067ffffffffffffffff82111561102a57611029610f94565b5b61103382610f83565b9050602081019050919050565b82818337600083830152505050565b600061106261105d8461100f565b610ff4565b90508281526020810184848401111561107e5761107d610f7e565b5b611089848285611040565b509392505050565b600082601f8301126110a6576110a5610ecc565b5b81356110b684826020860161104f565b91505092915050565b6000602082840312156110d5576110d4610ec2565b5b600082013567ffffffffffffffff8111156110f3576110f2610ec7565b5b6110ff84828501611091565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015611142578082015181840152602081019050611127565b60008484015250505050565b600061115982611108565b6111638185611113565b9350611173818560208601611124565b61117c81610f83565b840191505092915050565b600081519050919050565b600082825260208201905092915050565b60006111ae82611187565b6111b88185611192565b93506111c8818560208601611124565b6111d181610f83565b840191505092915050565b6000819050919050565b6111ef816111dc565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611220826111f5565b9050919050565b61123081611215565b82525050565b600060a0820190508181036000830152611250818861114e565b90508181036020830152611264818761114e565b9050818103604083015261127881866111a3565b905061128760608301856111e6565b6112946080830184611227565b9695505050505050565b60008083601f8401126112b4576112b3610ecc565b5b8235905067ffffffffffffffff8111156112d1576112d0610ed1565b5b6020830191508360018202830111156112ed576112ec610ed6565b5b9250929050565b6112fd816111dc565b811461130857600080fd5b50565b60008135905061131a816112f4565b92915050565b60008060008060008060006080888a03121561133f5761133e610ec2565b5b600088013567ffffffffffffffff81111561135d5761135c610ec7565b5b6113698a828b01610edb565b9750975050602088013567ffffffffffffffff81111561138c5761138b610ec7565b5b6113988a828b01610edb565b9550955050604088013567ffffffffffffffff8111156113bb576113ba610ec7565b5b6113c78a828b0161129e565b935093505060606113da8a828b0161130b565b91505092959891949750929550565b60006040820190508181036000830152611403818561114e565b9050818103602083015261141781846111a3565b90509392505050565b600081905092915050565b60006114378385611420565b9350611444838584611040565b82840190509392505050565b600061145d82848661142b565b91508190509392505050565b7f446f6d61696e206e6f74206f776e65642062792073656e646572000000000000600082015250565b600061149f601a83611113565b91506114aa82611469565b602082019050919050565b600060208201905081810360008301526114ce81611492565b9050919050565b60006114e18385611113565b93506114ee838584611040565b6114f783610f83565b840190509392505050565b6000602082019050818103600083015261151d8184866114d5565b90509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061156d57607f821691505b6020821081036115805761157f611526565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006115ef826111dc565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611621576116206115b5565b5b600182019050919050565b7f496e76616c696420686f73746e616d6500000000000000000000000000000000600082015250565b6000611662601083611113565b915061166d8261162c565b602082019050919050565b6000602082019050818103600083015261169181611655565b9050919050565b60006116a382611108565b6116ad8185611420565b93506116bd818560208601611124565b80840191505092915050565b60006116d58284611698565b915081905092915050565b7f506172656e7420646f6d61696e206e6f742072656769737465726564206f722060008201527f6e6f74206f776e65642062792073656e64657200000000000000000000000000602082015250565b600061173c603383611113565b9150611747826116e0565b604082019050919050565b6000602082019050818103600083015261176b8161172f565b9050919050565b600061177d826111dc565b9150611788836111dc565b92508282019050808211156117a05761179f6115b5565b5b92915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026118087fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826117cb565b61181286836117cb565b95508019841693508086168417925050509392505050565b6000819050919050565b600061184f61184a611845846111dc565b61182a565b6111dc565b9050919050565b6000819050919050565b61186983611834565b61187d61187582611856565b8484546117d8565b825550505050565b600090565b611892611885565b61189d818484611860565b505050565b5b818110156118c1576118b660008261188a565b6001810190506118a3565b5050565b601f821115611906576118d7816117a6565b6118e0846117bb565b810160208510156118ef578190505b6119036118fb856117bb565b8301826118a2565b50505b505050565b600082821c905092915050565b60006119296000198460080261190b565b1980831691505092915050565b60006119428383611918565b9150826002028217905092915050565b61195b82611108565b67ffffffffffffffff81111561197457611973610f94565b5b61197e8254611555565b6119898282856118c5565b600060209050601f8311600181146119bc57600084156119aa578287015190505b6119b48582611936565b865550611a1c565b601f1984166119ca866117a6565b60005b828110156119f2578489015182556001820191506020850194506020810190506119cd565b86831015611a0f5784890151611a0b601f891682611918565b8355505b6001600288020188555050505b505050505050565b60008190508160005260206000209050919050565b601f821115611a7a57611a4b81611a24565b611a54846117bb565b81016020851015611a63578190505b611a77611a6f856117bb565b8301826118a2565b50505b505050565b611a8882611187565b67ffffffffffffffff811115611aa157611aa0610f94565b5b611aab8254611555565b611ab6828285611a39565b600060209050601f831160018114611ae95760008415611ad7578287015190505b611ae18582611936565b865550611b49565b601f198416611af786611a24565b60005b82811015611b1f57848901518255600182019150602085019450602081019050611afa565b86831015611b3c5784890151611b38601f891682611918565b8355505b6001600288020188555050505b505050505050565b6000611b5d8385611192565b9350611b6a838584611040565b611b7383610f83565b840190509392505050565b60006080820190508181036000830152611b9981898b6114d5565b90508181036020830152611bae8187896114d5565b90508181036040830152611bc3818587611b51565b9050611bd260608301846111e6565b98975050505050505050565b600082905092915050565b611bf38383611bde565b67ffffffffffffffff811115611c0c57611c0b610f94565b5b611c168254611555565b611c218282856118c5565b6000601f831160018114611c505760008415611c3e578287013590505b611c488582611936565b865550611cb0565b601f198416611c5e866117a6565b60005b82811015611c8657848901358255600182019150602085019450602081019050611c61565b86831015611ca35784890135611c9f601f891682611918565b8355505b6001600288020188555050505b50505050505050565b600082905092915050565b611cce8383611cb9565b67ffffffffffffffff811115611ce757611ce6610f94565b5b611cf18254611555565b611cfc828285611a39565b6000601f831160018114611d2b5760008415611d19578287013590505b611d238582611936565b865550611d8b565b601f198416611d3986611a24565b60005b82811015611d6157848901358255600182019150602085019450602081019050611d3c565b86831015611d7e5784890135611d7a601f891682611918565b8355505b6001600288020188555050505b50505050505050565b7f446f6d61696e206e6f7420666f756e6400000000000000000000000000000000600082015250565b6000611dca601083611113565b9150611dd582611d94565b602082019050919050565b60006020820190508181036000830152611df981611dbd565b9050919050565b7f4365727469666963617465206861732065787069726564000000000000000000600082015250565b6000611e36601783611113565b9150611e4182611e00565b602082019050919050565b60006020820190508181036000830152611e6581611e29565b9050919050565b6000611e77826111dc565b9150611e82836111dc565b9250828203905081811115611e9a57611e996115b5565b5b9291505056fea264697066735822122002ed9da84137e22051db0b1e0000330c98f1e4e100614dfc73301ff5b1a7bde164736f6c63430008130033",
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
// Solidity: function domains(string ) view returns(string hostname, string ip, bytes certificate, uint256 data, address owner)
func (_DomainRegistry *DomainRegistryCaller) Domains(opts *bind.CallOpts, arg0 string) (struct {
	Hostname    string
	Ip          string
	Certificate []byte
	Data        *big.Int
	Owner       common.Address
}, error) {
	var out []interface{}
	err := _DomainRegistry.contract.Call(opts, &out, "domains", arg0)

	outstruct := new(struct {
		Hostname    string
		Ip          string
		Certificate []byte
		Data        *big.Int
		Owner       common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hostname = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Ip = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Certificate = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Data = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Domains is a free data retrieval call binding the contract method 0x26449235.
//
// Solidity: function domains(string ) view returns(string hostname, string ip, bytes certificate, uint256 data, address owner)
func (_DomainRegistry *DomainRegistrySession) Domains(arg0 string) (struct {
	Hostname    string
	Ip          string
	Certificate []byte
	Data        *big.Int
	Owner       common.Address
}, error) {
	return _DomainRegistry.Contract.Domains(&_DomainRegistry.CallOpts, arg0)
}

// Domains is a free data retrieval call binding the contract method 0x26449235.
//
// Solidity: function domains(string ) view returns(string hostname, string ip, bytes certificate, uint256 data, address owner)
func (_DomainRegistry *DomainRegistryCallerSession) Domains(arg0 string) (struct {
	Hostname    string
	Ip          string
	Certificate []byte
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

// CreateDomain is a paid mutator transaction binding the contract method 0x2ad09666.
//
// Solidity: function createDomain(string hostname, string ip, bytes certificate, uint256 data) returns()
func (_DomainRegistry *DomainRegistryTransactor) CreateDomain(opts *bind.TransactOpts, hostname string, ip string, certificate []byte, data *big.Int) (*types.Transaction, error) {
	return _DomainRegistry.contract.Transact(opts, "createDomain", hostname, ip, certificate, data)
}

// CreateDomain is a paid mutator transaction binding the contract method 0x2ad09666.
//
// Solidity: function createDomain(string hostname, string ip, bytes certificate, uint256 data) returns()
func (_DomainRegistry *DomainRegistrySession) CreateDomain(hostname string, ip string, certificate []byte, data *big.Int) (*types.Transaction, error) {
	return _DomainRegistry.Contract.CreateDomain(&_DomainRegistry.TransactOpts, hostname, ip, certificate, data)
}

// CreateDomain is a paid mutator transaction binding the contract method 0x2ad09666.
//
// Solidity: function createDomain(string hostname, string ip, bytes certificate, uint256 data) returns()
func (_DomainRegistry *DomainRegistryTransactorSession) CreateDomain(hostname string, ip string, certificate []byte, data *big.Int) (*types.Transaction, error) {
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

// UpdateDomain is a paid mutator transaction binding the contract method 0xe532fb98.
//
// Solidity: function updateDomain(string hostname, string ip, bytes certificate, uint256 data) returns()
func (_DomainRegistry *DomainRegistryTransactor) UpdateDomain(opts *bind.TransactOpts, hostname string, ip string, certificate []byte, data *big.Int) (*types.Transaction, error) {
	return _DomainRegistry.contract.Transact(opts, "updateDomain", hostname, ip, certificate, data)
}

// UpdateDomain is a paid mutator transaction binding the contract method 0xe532fb98.
//
// Solidity: function updateDomain(string hostname, string ip, bytes certificate, uint256 data) returns()
func (_DomainRegistry *DomainRegistrySession) UpdateDomain(hostname string, ip string, certificate []byte, data *big.Int) (*types.Transaction, error) {
	return _DomainRegistry.Contract.UpdateDomain(&_DomainRegistry.TransactOpts, hostname, ip, certificate, data)
}

// UpdateDomain is a paid mutator transaction binding the contract method 0xe532fb98.
//
// Solidity: function updateDomain(string hostname, string ip, bytes certificate, uint256 data) returns()
func (_DomainRegistry *DomainRegistryTransactorSession) UpdateDomain(hostname string, ip string, certificate []byte, data *big.Int) (*types.Transaction, error) {
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
	Certificate []byte
	Data        *big.Int
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDomainCreated is a free log retrieval operation binding the contract event 0xd694233903b144a99ef1896eb8f5ebff710c263412bfcff1b8d37d5c574e1d10.
//
// Solidity: event DomainCreated(string hostname, string ip, bytes certificate, uint256 data, address indexed owner)
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

// WatchDomainCreated is a free log subscription operation binding the contract event 0xd694233903b144a99ef1896eb8f5ebff710c263412bfcff1b8d37d5c574e1d10.
//
// Solidity: event DomainCreated(string hostname, string ip, bytes certificate, uint256 data, address indexed owner)
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

// ParseDomainCreated is a log parse operation binding the contract event 0xd694233903b144a99ef1896eb8f5ebff710c263412bfcff1b8d37d5c574e1d10.
//
// Solidity: event DomainCreated(string hostname, string ip, bytes certificate, uint256 data, address indexed owner)
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
	Certificate []byte
	Data        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDomainUpdated is a free log retrieval operation binding the contract event 0x6ee95bc81bd859a73439a27fc26eb56779fa812f28d8271df16444b27b0e02fd.
//
// Solidity: event DomainUpdated(string hostname, string ip, bytes certificate, uint256 data)
func (_DomainRegistry *DomainRegistryFilterer) FilterDomainUpdated(opts *bind.FilterOpts) (*DomainRegistryDomainUpdatedIterator, error) {

	logs, sub, err := _DomainRegistry.contract.FilterLogs(opts, "DomainUpdated")
	if err != nil {
		return nil, err
	}
	return &DomainRegistryDomainUpdatedIterator{contract: _DomainRegistry.contract, event: "DomainUpdated", logs: logs, sub: sub}, nil
}

// WatchDomainUpdated is a free log subscription operation binding the contract event 0x6ee95bc81bd859a73439a27fc26eb56779fa812f28d8271df16444b27b0e02fd.
//
// Solidity: event DomainUpdated(string hostname, string ip, bytes certificate, uint256 data)
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

// ParseDomainUpdated is a log parse operation binding the contract event 0x6ee95bc81bd859a73439a27fc26eb56779fa812f28d8271df16444b27b0e02fd.
//
// Solidity: event DomainUpdated(string hostname, string ip, bytes certificate, uint256 data)
func (_DomainRegistry *DomainRegistryFilterer) ParseDomainUpdated(log types.Log) (*DomainRegistryDomainUpdated, error) {
	event := new(DomainRegistryDomainUpdated)
	if err := _DomainRegistry.contract.UnpackLog(event, "DomainUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
