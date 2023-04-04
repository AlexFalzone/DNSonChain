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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"certificate\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DomainCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"}],\"name\":\"DomainDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"certificate\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiryDate\",\"type\":\"uint256\"}],\"name\":\"DomainUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"certificate\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expiryDate\",\"type\":\"uint256\"}],\"name\":\"createDomain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"}],\"name\":\"deleteDomain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"domains\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"certificate\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expiryDate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"}],\"name\":\"getCertificate\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hostname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"certificate\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expiryDate\",\"type\":\"uint256\"}],\"name\":\"updateDomain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611f87806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630eaf97521461005c5780631469e9141461007857806326449235146100945780635195418c146100c8578063ed0f2e75146100e4575b600080fd5b6100766004803603810190610071919061116a565b610115565b005b610092600480360381019061008d919061132e565b61033a565b005b6100ae60048036038101906100a991906113f1565b61085a565b6040516100bf959493929190611509565b60405180910390f35b6100e260048036038101906100dd919061132e565b610a5e565b005b6100fe60048036038101906100f9919061116a565b610cd3565b60405161010c929190611571565b60405180910390f35b81818080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050503373ffffffffffffffffffffffffffffffffffffffff1660008260405161018191906115e4565b908152602001604051809103902060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610209576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161020090611647565b60405180910390fd5b600080848460405161021c92919061168c565b90815260200160405180910390206000018054610238906116d4565b90501161027a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027190611751565b60405180910390fd5b6000838360405161028c92919061168c565b9081526020016040518091039020600080820160006102ab9190611094565b6001820160006102bb9190611094565b6002820160006102cb9190611094565b60038201600090556004820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550507f9b449e8e6e58cfe4813dc8a38f233368018ddd99ddc223e2f1a55eb22eefa236838360405161032d92919061179e565b60405180910390a1505050565b600086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505090506000805b8251811015610418577f2e000000000000000000000000000000000000000000000000000000000000008382815181106103c4576103c36117c2565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19160361040557818061040190611820565b9250505b808061041090611820565b915050610387565b506000811161045c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610453906118b4565b60405180910390fd5b600181111561052357600061047083610ef9565b90503373ffffffffffffffffffffffffffffffffffffffff1660008260405161049991906115e4565b908152602001604051809103902060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610521576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161051890611946565b60405180910390fd5b505b600080898960405161053692919061168c565b90815260200160405180910390206000018054610552906116d4565b9050111561064357426000898960405161056d92919061168c565b90815260200160405180910390206003015411156105c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b7906119d8565b60405180910390fd5b600088886040516105d292919061168c565b9081526020016040518091039020600080820160006105f19190611094565b6001820160006106019190611094565b6002820160006106119190611094565b60038201600090556004820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550505b6040518060a0016040528089898080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200187878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152602001858152602001600085146107015784426106fc91906119f8565b610723565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b81526020013373ffffffffffffffffffffffffffffffffffffffff168152506000898960405161075492919061168c565b908152602001604051809103902060008201518160000190816107779190611bd8565b50602082015181600101908161078d9190611bd8565b5060408201518160020190816107a39190611bd8565b506060820151816003015560808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050503373ffffffffffffffffffffffffffffffffffffffff167fa13d03807a9cdb641092d81b09d6b69e448c291c5e7d4b7a0c57f31a3b72b1d789898989898960405161084896959493929190611caa565b60405180910390a25050505050505050565b600081805160208101820180518482526020830160208501208183528095505050505050600091509050806000018054610893906116d4565b80601f01602080910402602001604051908101604052809291908181526020018280546108bf906116d4565b801561090c5780601f106108e15761010080835404028352916020019161090c565b820191906000526020600020905b8154815290600101906020018083116108ef57829003601f168201915b505050505090806001018054610921906116d4565b80601f016020809104026020016040519081016040528092919081815260200182805461094d906116d4565b801561099a5780601f1061096f5761010080835404028352916020019161099a565b820191906000526020600020905b81548152906001019060200180831161097d57829003601f168201915b5050505050908060020180546109af906116d4565b80601f01602080910402602001604051908101604052809291908181526020018280546109db906116d4565b8015610a285780601f106109fd57610100808354040283529160200191610a28565b820191906000526020600020905b815481529060010190602001808311610a0b57829003601f168201915b5050505050908060030154908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905085565b85858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050503373ffffffffffffffffffffffffffffffffffffffff16600082604051610aca91906115e4565b908152602001604051809103902060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610b52576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4990611647565b60405180910390fd5b6000808888604051610b6592919061168c565b90815260200160405180910390206000018054610b81906116d4565b905011610bc3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bba90611751565b60405180910390fd5b848460008989604051610bd792919061168c565b90815260200160405180910390206001019182610bf5929190611d13565b508260008888604051610c0992919061168c565b90815260200160405180910390206002019081610c269190611bd8565b5060008214610c40578142610c3b91906119f8565b610c62565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b60008888604051610c7492919061168c565b9081526020016040518091039020600301819055507fb8abd760067276c7baaf5f313ea2d97676d7eef530f9c9143c377f4b8b4ca6b5878787878787604051610cc296959493929190611caa565b60405180910390a150505050505050565b6060806000808585604051610ce992919061168c565b908152602001604051809103902090506000816000018054610d0a906116d4565b905011610d4c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d4390611751565b60405180910390fd5b42816003015411610d92576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d8990611e2f565b60405180910390fd5b6000816002018054610da3906116d4565b80601f0160208091040260200160405190810160405280929190818152602001828054610dcf906116d4565b8015610e1c5780601f10610df157610100808354040283529160200191610e1c565b820191906000526020600020905b815481529060010190602001808311610dff57829003601f168201915b50505050509050600060028251610e339190611e7e565b14610e5b5780604051602001610e499190611efb565b60405160208183030381529060405290505b8160010181818054610e6c906116d4565b80601f0160208091040260200160405190810160405280929190818152602001828054610e98906116d4565b8015610ee55780601f10610eba57610100808354040283529160200191610ee5565b820191906000526020600020905b815481529060010190602001808311610ec857829003601f168201915b505050505091509350935050509250929050565b60606000805b8351811015610f89577f2e00000000000000000000000000000000000000000000000000000000000000848281518110610f3c57610f3b6117c2565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191603610f7657809150610f89565b8080610f8190611820565b915050610eff565b5060006001828551610f9b9190611f1d565b610fa59190611f1d565b67ffffffffffffffff811115610fbe57610fbd6111cd565b5b6040519080825280601f01601f191660200182016040528015610ff05781602001600182028036833780820191505090505b50905060005b815181101561108957848160018561100e91906119f8565b61101891906119f8565b81518110611029576110286117c2565b5b602001015160f81c60f81b828281518110611047576110466117c2565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350808061108190611820565b915050610ff6565b508092505050919050565b5080546110a0906116d4565b6000825580601f106110b257506110d1565b601f0160209004906000526020600020908101906110d091906110d4565b5b50565b5b808211156110ed5760008160009055506001016110d5565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f84011261112a57611129611105565b5b8235905067ffffffffffffffff8111156111475761114661110a565b5b6020830191508360018202830111156111635761116261110f565b5b9250929050565b60008060208385031215611181576111806110fb565b5b600083013567ffffffffffffffff81111561119f5761119e611100565b5b6111ab85828601611114565b92509250509250929050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611205826111bc565b810181811067ffffffffffffffff82111715611224576112236111cd565b5b80604052505050565b60006112376110f1565b905061124382826111fc565b919050565b600067ffffffffffffffff821115611263576112626111cd565b5b61126c826111bc565b9050602081019050919050565b82818337600083830152505050565b600061129b61129684611248565b61122d565b9050828152602081018484840111156112b7576112b66111b7565b5b6112c2848285611279565b509392505050565b600082601f8301126112df576112de611105565b5b81356112ef848260208601611288565b91505092915050565b6000819050919050565b61130b816112f8565b811461131657600080fd5b50565b60008135905061132881611302565b92915050565b6000806000806000806080878903121561134b5761134a6110fb565b5b600087013567ffffffffffffffff81111561136957611368611100565b5b61137589828a01611114565b9650965050602087013567ffffffffffffffff81111561139857611397611100565b5b6113a489828a01611114565b9450945050604087013567ffffffffffffffff8111156113c7576113c6611100565b5b6113d389828a016112ca565b92505060606113e489828a01611319565b9150509295509295509295565b600060208284031215611407576114066110fb565b5b600082013567ffffffffffffffff81111561142557611424611100565b5b611431848285016112ca565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015611474578082015181840152602081019050611459565b60008484015250505050565b600061148b8261143a565b6114958185611445565b93506114a5818560208601611456565b6114ae816111bc565b840191505092915050565b6114c2816112f8565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006114f3826114c8565b9050919050565b611503816114e8565b82525050565b600060a08201905081810360008301526115238188611480565b905081810360208301526115378187611480565b9050818103604083015261154b8186611480565b905061155a60608301856114b9565b61156760808301846114fa565b9695505050505050565b6000604082019050818103600083015261158b8185611480565b9050818103602083015261159f8184611480565b90509392505050565b600081905092915050565b60006115be8261143a565b6115c881856115a8565b93506115d8818560208601611456565b80840191505092915050565b60006115f082846115b3565b915081905092915050565b7f446f6d61696e206e6f74206f776e65642062792073656e646572000000000000600082015250565b6000611631601a83611445565b915061163c826115fb565b602082019050919050565b6000602082019050818103600083015261166081611624565b9050919050565b600061167383856115a8565b9350611680838584611279565b82840190509392505050565b6000611699828486611667565b91508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806116ec57607f821691505b6020821081036116ff576116fe6116a5565b5b50919050565b7f446f6d61696e206e6f7420666f756e6400000000000000000000000000000000600082015250565b600061173b601083611445565b915061174682611705565b602082019050919050565b6000602082019050818103600083015261176a8161172e565b9050919050565b600061177d8385611445565b935061178a838584611279565b611793836111bc565b840190509392505050565b600060208201905081810360008301526117b9818486611771565b90509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061182b826112f8565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361185d5761185c6117f1565b5b600182019050919050565b7f496e76616c696420686f73746e616d6500000000000000000000000000000000600082015250565b600061189e601083611445565b91506118a982611868565b602082019050919050565b600060208201905081810360008301526118cd81611891565b9050919050565b7f506172656e7420646f6d61696e206e6f742072656769737465726564206f722060008201527f6e6f74206f776e65642062792073656e64657200000000000000000000000000602082015250565b6000611930603383611445565b915061193b826118d4565b604082019050919050565b6000602082019050818103600083015261195f81611923565b9050919050565b7f486f73746e616d6520616c7265616479207265676973746572656420616e642060008201527f6e6f742065787069726564000000000000000000000000000000000000000000602082015250565b60006119c2602b83611445565b91506119cd82611966565b604082019050919050565b600060208201905081810360008301526119f1816119b5565b9050919050565b6000611a03826112f8565b9150611a0e836112f8565b9250828201905080821115611a2657611a256117f1565b5b92915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302611a8e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611a51565b611a988683611a51565b95508019841693508086168417925050509392505050565b6000819050919050565b6000611ad5611ad0611acb846112f8565b611ab0565b6112f8565b9050919050565b6000819050919050565b611aef83611aba565b611b03611afb82611adc565b848454611a5e565b825550505050565b600090565b611b18611b0b565b611b23818484611ae6565b505050565b5b81811015611b4757611b3c600082611b10565b600181019050611b29565b5050565b601f821115611b8c57611b5d81611a2c565b611b6684611a41565b81016020851015611b75578190505b611b89611b8185611a41565b830182611b28565b50505b505050565b600082821c905092915050565b6000611baf60001984600802611b91565b1980831691505092915050565b6000611bc88383611b9e565b9150826002028217905092915050565b611be18261143a565b67ffffffffffffffff811115611bfa57611bf96111cd565b5b611c0482546116d4565b611c0f828285611b4b565b600060209050601f831160018114611c425760008415611c30578287015190505b611c3a8582611bbc565b865550611ca2565b601f198416611c5086611a2c565b60005b82811015611c7857848901518255600182019150602085019450602081019050611c53565b86831015611c955784890151611c91601f891682611b9e565b8355505b6001600288020188555050505b505050505050565b60006080820190508181036000830152611cc581888a611771565b90508181036020830152611cda818688611771565b90508181036040830152611cee8185611480565b9050611cfd60608301846114b9565b979650505050505050565b600082905092915050565b611d1d8383611d08565b67ffffffffffffffff811115611d3657611d356111cd565b5b611d4082546116d4565b611d4b828285611b4b565b6000601f831160018114611d7a5760008415611d68578287013590505b611d728582611bbc565b865550611dda565b601f198416611d8886611a2c565b60005b82811015611db057848901358255600182019150602085019450602081019050611d8b565b86831015611dcd5784890135611dc9601f891682611b9e565b8355505b6001600288020188555050505b50505050505050565b7f4365727469666963617465206861732065787069726564000000000000000000600082015250565b6000611e19601783611445565b9150611e2482611de3565b602082019050919050565b60006020820190508181036000830152611e4881611e0c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611e89826112f8565b9150611e94836112f8565b925082611ea457611ea3611e4f565b5b828206905092915050565b7f3000000000000000000000000000000000000000000000000000000000000000600082015250565b6000611ee56001836115a8565b9150611ef082611eaf565b600182019050919050565b6000611f0682611ed8565b9150611f1282846115b3565b915081905092915050565b6000611f28826112f8565b9150611f33836112f8565b9250828203905081811115611f4b57611f4a6117f1565b5b9291505056fea2646970667358221220e0bc540b2aade1395a910452f14cf0e1b1a6621369c0b181a09741fc2c5dbfb164736f6c63430008130033",
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
// Solidity: function domains(string ) view returns(string hostname, string ip, string certificate, uint256 expiryDate, address owner)
func (_DomainRegistry *DomainRegistryCaller) Domains(opts *bind.CallOpts, arg0 string) (struct {
	Hostname    string
	Ip          string
	Certificate string
	ExpiryDate  *big.Int
	Owner       common.Address
}, error) {
	var out []interface{}
	err := _DomainRegistry.contract.Call(opts, &out, "domains", arg0)

	outstruct := new(struct {
		Hostname    string
		Ip          string
		Certificate string
		ExpiryDate  *big.Int
		Owner       common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hostname = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Ip = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Certificate = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ExpiryDate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Domains is a free data retrieval call binding the contract method 0x26449235.
//
// Solidity: function domains(string ) view returns(string hostname, string ip, string certificate, uint256 expiryDate, address owner)
func (_DomainRegistry *DomainRegistrySession) Domains(arg0 string) (struct {
	Hostname    string
	Ip          string
	Certificate string
	ExpiryDate  *big.Int
	Owner       common.Address
}, error) {
	return _DomainRegistry.Contract.Domains(&_DomainRegistry.CallOpts, arg0)
}

// Domains is a free data retrieval call binding the contract method 0x26449235.
//
// Solidity: function domains(string ) view returns(string hostname, string ip, string certificate, uint256 expiryDate, address owner)
func (_DomainRegistry *DomainRegistryCallerSession) Domains(arg0 string) (struct {
	Hostname    string
	Ip          string
	Certificate string
	ExpiryDate  *big.Int
	Owner       common.Address
}, error) {
	return _DomainRegistry.Contract.Domains(&_DomainRegistry.CallOpts, arg0)
}

// GetCertificate is a free data retrieval call binding the contract method 0xed0f2e75.
//
// Solidity: function getCertificate(string hostname) view returns(string, string)
func (_DomainRegistry *DomainRegistryCaller) GetCertificate(opts *bind.CallOpts, hostname string) (string, string, error) {
	var out []interface{}
	err := _DomainRegistry.contract.Call(opts, &out, "getCertificate", hostname)

	if err != nil {
		return *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetCertificate is a free data retrieval call binding the contract method 0xed0f2e75.
//
// Solidity: function getCertificate(string hostname) view returns(string, string)
func (_DomainRegistry *DomainRegistrySession) GetCertificate(hostname string) (string, string, error) {
	return _DomainRegistry.Contract.GetCertificate(&_DomainRegistry.CallOpts, hostname)
}

// GetCertificate is a free data retrieval call binding the contract method 0xed0f2e75.
//
// Solidity: function getCertificate(string hostname) view returns(string, string)
func (_DomainRegistry *DomainRegistryCallerSession) GetCertificate(hostname string) (string, string, error) {
	return _DomainRegistry.Contract.GetCertificate(&_DomainRegistry.CallOpts, hostname)
}

// CreateDomain is a paid mutator transaction binding the contract method 0x1469e914.
//
// Solidity: function createDomain(string hostname, string ip, string certificate, uint256 expiryDate) returns()
func (_DomainRegistry *DomainRegistryTransactor) CreateDomain(opts *bind.TransactOpts, hostname string, ip string, certificate string, expiryDate *big.Int) (*types.Transaction, error) {
	return _DomainRegistry.contract.Transact(opts, "createDomain", hostname, ip, certificate, expiryDate)
}

// CreateDomain is a paid mutator transaction binding the contract method 0x1469e914.
//
// Solidity: function createDomain(string hostname, string ip, string certificate, uint256 expiryDate) returns()
func (_DomainRegistry *DomainRegistrySession) CreateDomain(hostname string, ip string, certificate string, expiryDate *big.Int) (*types.Transaction, error) {
	return _DomainRegistry.Contract.CreateDomain(&_DomainRegistry.TransactOpts, hostname, ip, certificate, expiryDate)
}

// CreateDomain is a paid mutator transaction binding the contract method 0x1469e914.
//
// Solidity: function createDomain(string hostname, string ip, string certificate, uint256 expiryDate) returns()
func (_DomainRegistry *DomainRegistryTransactorSession) CreateDomain(hostname string, ip string, certificate string, expiryDate *big.Int) (*types.Transaction, error) {
	return _DomainRegistry.Contract.CreateDomain(&_DomainRegistry.TransactOpts, hostname, ip, certificate, expiryDate)
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
// Solidity: function updateDomain(string hostname, string ip, string certificate, uint256 expiryDate) returns()
func (_DomainRegistry *DomainRegistryTransactor) UpdateDomain(opts *bind.TransactOpts, hostname string, ip string, certificate string, expiryDate *big.Int) (*types.Transaction, error) {
	return _DomainRegistry.contract.Transact(opts, "updateDomain", hostname, ip, certificate, expiryDate)
}

// UpdateDomain is a paid mutator transaction binding the contract method 0x5195418c.
//
// Solidity: function updateDomain(string hostname, string ip, string certificate, uint256 expiryDate) returns()
func (_DomainRegistry *DomainRegistrySession) UpdateDomain(hostname string, ip string, certificate string, expiryDate *big.Int) (*types.Transaction, error) {
	return _DomainRegistry.Contract.UpdateDomain(&_DomainRegistry.TransactOpts, hostname, ip, certificate, expiryDate)
}

// UpdateDomain is a paid mutator transaction binding the contract method 0x5195418c.
//
// Solidity: function updateDomain(string hostname, string ip, string certificate, uint256 expiryDate) returns()
func (_DomainRegistry *DomainRegistryTransactorSession) UpdateDomain(hostname string, ip string, certificate string, expiryDate *big.Int) (*types.Transaction, error) {
	return _DomainRegistry.Contract.UpdateDomain(&_DomainRegistry.TransactOpts, hostname, ip, certificate, expiryDate)
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
	ExpiryDate  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDomainUpdated is a free log retrieval operation binding the contract event 0xb8abd760067276c7baaf5f313ea2d97676d7eef530f9c9143c377f4b8b4ca6b5.
//
// Solidity: event DomainUpdated(string hostname, string ip, string certificate, uint256 expiryDate)
func (_DomainRegistry *DomainRegistryFilterer) FilterDomainUpdated(opts *bind.FilterOpts) (*DomainRegistryDomainUpdatedIterator, error) {

	logs, sub, err := _DomainRegistry.contract.FilterLogs(opts, "DomainUpdated")
	if err != nil {
		return nil, err
	}
	return &DomainRegistryDomainUpdatedIterator{contract: _DomainRegistry.contract, event: "DomainUpdated", logs: logs, sub: sub}, nil
}

// WatchDomainUpdated is a free log subscription operation binding the contract event 0xb8abd760067276c7baaf5f313ea2d97676d7eef530f9c9143c377f4b8b4ca6b5.
//
// Solidity: event DomainUpdated(string hostname, string ip, string certificate, uint256 expiryDate)
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
// Solidity: event DomainUpdated(string hostname, string ip, string certificate, uint256 expiryDate)
func (_DomainRegistry *DomainRegistryFilterer) ParseDomainUpdated(log types.Log) (*DomainRegistryDomainUpdated, error) {
	event := new(DomainRegistryDomainUpdated)
	if err := _DomainRegistry.contract.UnpackLog(event, "DomainUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
