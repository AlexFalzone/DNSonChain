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
	Bin: "0x608060405234801561001057600080fd5b5061229f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630eaf97521461005c57806326449235146100785780632ad09666146100ac578063e532fb98146100c8578063ed0f2e75146100e4575b600080fd5b610076600480360381019061007191906112cd565b610115565b005b610092600480360381019061008d919061145b565b610285565b6040516100a39594939291906115d2565b60405180910390f35b6100c660048036038101906100c19190611707565b610489565b005b6100e260048036038101906100dd9190611707565b6109a9565b005b6100fe60048036038101906100f991906112cd565b610bda565b60405161010c9291906117ca565b60405180910390f35b3373ffffffffffffffffffffffffffffffffffffffff166000838360405161013e929190611831565b908152602001604051809103902060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146101c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101bd90611896565b60405180910390fd5b600082826040516101d8929190611831565b9081526020016040518091039020600080820160006101f791906111b7565b60018201600061020791906111b7565b60028201600061021791906111f7565b60038201600090556004820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550507f9b449e8e6e58cfe4813dc8a38f233368018ddd99ddc223e2f1a55eb22eefa23682826040516102799291906118e3565b60405180910390a15050565b6000818051602081018201805184825260208301602085012081835280955050505050506000915090508060000180546102be90611936565b80601f01602080910402602001604051908101604052809291908181526020018280546102ea90611936565b80156103375780601f1061030c57610100808354040283529160200191610337565b820191906000526020600020905b81548152906001019060200180831161031a57829003601f168201915b50505050509080600101805461034c90611936565b80601f016020809104026020016040519081016040528092919081815260200182805461037890611936565b80156103c55780601f1061039a576101008083540402835291602001916103c5565b820191906000526020600020905b8154815290600101906020018083116103a857829003601f168201915b5050505050908060020180546103da90611936565b80601f016020809104026020016040519081016040528092919081815260200182805461040690611936565b80156104535780601f1061042857610100808354040283529160200191610453565b820191906000526020600020905b81548152906001019060200180831161043657829003601f168201915b5050505050908060030154908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905085565b600086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505090506000805b8251811015610567577f2e0000000000000000000000000000000000000000000000000000000000000083828151811061051357610512611967565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191603610554578180610550906119c5565b9250505b808061055f906119c5565b9150506104d6565b50600081116105ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a290611a59565b60405180910390fd5b60018111156106725760006105bf8361101c565b90503373ffffffffffffffffffffffffffffffffffffffff166000826040516105e89190611aaa565b908152602001604051809103902060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610670576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066790611b33565b60405180910390fd5b505b6000808989604051610685929190611831565b908152602001604051809103902060000180546106a190611936565b905011156107925742600089896040516106bc929190611831565b908152602001604051809103902060030154111561070f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070690611bc5565b60405180910390fd5b60008888604051610721929190611831565b90815260200160405180910390206000808201600061074091906111b7565b60018201600061075091906111b7565b60028201600061076091906111f7565b60038201600090556004820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550505b6040518060a0016040528089898080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200187878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020018581526020016000851461085057844261084b9190611be5565b610872565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b81526020013373ffffffffffffffffffffffffffffffffffffffff16815250600089896040516108a3929190611831565b908152602001604051809103902060008201518160000190816108c69190611dc5565b5060208201518160010190816108dc9190611dc5565b5060408201518160020190816108f29190611ef2565b506060820151816003015560808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050503373ffffffffffffffffffffffffffffffffffffffff167fd694233903b144a99ef1896eb8f5ebff710c263412bfcff1b8d37d5c574e1d1089898989898960405161099796959493929190611fc4565b60405180910390a25050505050505050565b60008087876040516109bc929190611831565b908152602001604051809103902060000180546109d890611936565b905011610a1a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a119061206e565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff1660008787604051610a43929190611831565b908152602001604051809103902060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610acb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac290611896565b60405180910390fd5b838360008888604051610adf929190611831565b90815260200160405180910390206001019182610afd929190612099565b508160008787604051610b11929190611831565b90815260200160405180910390206002019081610b2e9190611ef2565b5060008114610b48578042610b439190611be5565b610b6a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b60008787604051610b7c929190611831565b9081526020016040518091039020600301819055507f6ee95bc81bd859a73439a27fc26eb56779fa812f28d8271df16444b27b0e02fd868686868686604051610bca96959493929190611fc4565b60405180910390a1505050505050565b6060806000808585604051610bf0929190611831565b908152602001604051809103902090506000816000018054610c1190611936565b905011610c53576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c4a9061206e565b60405180910390fd5b42816003015411610c99576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c90906121b5565b60405180910390fd5b60006002826002018054610cac90611936565b9050610cb89190612204565b14610eef5760006001826002018054610cd090611936565b9050610cdc9190611be5565b67ffffffffffffffff811115610cf557610cf4611330565b5b6040519080825280601f01601f191660200182016040528015610d275781602001600182028036833780820191505090505b509050600060f81b81600081518110610d4357610d42611967565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060005b826002018054610d8490611936565b9050811015610e525782600201818154610d9d90611936565b8110610dac57610dab611967565b5b815460011615610dcb5790600052602060002090602091828204019190065b9054901a7f01000000000000000000000000000000000000000000000000000000000000000282600183610dff9190611be5565b81518110610e1057610e0f611967565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080610e4a906119c5565b915050610d75565b508160010181818054610e6490611936565b80601f0160208091040260200160405190810160405280929190818152602001828054610e9090611936565b8015610edd5780601f10610eb257610100808354040283529160200191610edd565b820191906000526020600020905b815481529060010190602001808311610ec057829003601f168201915b50505050509150935093505050611015565b8060010181600201818054610f0390611936565b80601f0160208091040260200160405190810160405280929190818152602001828054610f2f90611936565b8015610f7c5780601f10610f5157610100808354040283529160200191610f7c565b820191906000526020600020905b815481529060010190602001808311610f5f57829003601f168201915b50505050509150808054610f8f90611936565b80601f0160208091040260200160405190810160405280929190818152602001828054610fbb90611936565b80156110085780601f10610fdd57610100808354040283529160200191611008565b820191906000526020600020905b815481529060010190602001808311610feb57829003601f168201915b5050505050905092509250505b9250929050565b60606000805b83518110156110ac577f2e0000000000000000000000000000000000000000000000000000000000000084828151811061105f5761105e611967565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191603611099578091506110ac565b80806110a4906119c5565b915050611022565b50600060018285516110be9190612235565b6110c89190612235565b67ffffffffffffffff8111156110e1576110e0611330565b5b6040519080825280601f01601f1916602001820160405280156111135781602001600182028036833780820191505090505b50905060005b81518110156111ac5784816001856111319190611be5565b61113b9190611be5565b8151811061114c5761114b611967565b5b602001015160f81c60f81b82828151811061116a57611169611967565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806111a4906119c5565b915050611119565b508092505050919050565b5080546111c390611936565b6000825580601f106111d557506111f4565b601f0160209004906000526020600020908101906111f39190611237565b5b50565b50805461120390611936565b6000825580601f106112155750611234565b601f0160209004906000526020600020908101906112339190611237565b5b50565b5b80821115611250576000816000905550600101611238565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f84011261128d5761128c611268565b5b8235905067ffffffffffffffff8111156112aa576112a961126d565b5b6020830191508360018202830111156112c6576112c5611272565b5b9250929050565b600080602083850312156112e4576112e361125e565b5b600083013567ffffffffffffffff81111561130257611301611263565b5b61130e85828601611277565b92509250509250929050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6113688261131f565b810181811067ffffffffffffffff8211171561138757611386611330565b5b80604052505050565b600061139a611254565b90506113a6828261135f565b919050565b600067ffffffffffffffff8211156113c6576113c5611330565b5b6113cf8261131f565b9050602081019050919050565b82818337600083830152505050565b60006113fe6113f9846113ab565b611390565b90508281526020810184848401111561141a5761141961131a565b5b6114258482856113dc565b509392505050565b600082601f83011261144257611441611268565b5b81356114528482602086016113eb565b91505092915050565b6000602082840312156114715761147061125e565b5b600082013567ffffffffffffffff81111561148f5761148e611263565b5b61149b8482850161142d565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156114de5780820151818401526020810190506114c3565b60008484015250505050565b60006114f5826114a4565b6114ff81856114af565b935061150f8185602086016114c0565b6115188161131f565b840191505092915050565b600081519050919050565b600082825260208201905092915050565b600061154a82611523565b611554818561152e565b93506115648185602086016114c0565b61156d8161131f565b840191505092915050565b6000819050919050565b61158b81611578565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006115bc82611591565b9050919050565b6115cc816115b1565b82525050565b600060a08201905081810360008301526115ec81886114ea565b9050818103602083015261160081876114ea565b90508181036040830152611614818661153f565b90506116236060830185611582565b61163060808301846115c3565b9695505050505050565b600067ffffffffffffffff82111561165557611654611330565b5b61165e8261131f565b9050602081019050919050565b600061167e6116798461163a565b611390565b90508281526020810184848401111561169a5761169961131a565b5b6116a58482856113dc565b509392505050565b600082601f8301126116c2576116c1611268565b5b81356116d284826020860161166b565b91505092915050565b6116e481611578565b81146116ef57600080fd5b50565b600081359050611701816116db565b92915050565b600080600080600080608087890312156117245761172361125e565b5b600087013567ffffffffffffffff81111561174257611741611263565b5b61174e89828a01611277565b9650965050602087013567ffffffffffffffff81111561177157611770611263565b5b61177d89828a01611277565b9450945050604087013567ffffffffffffffff8111156117a05761179f611263565b5b6117ac89828a016116ad565b92505060606117bd89828a016116f2565b9150509295509295509295565b600060408201905081810360008301526117e481856114ea565b905081810360208301526117f8818461153f565b90509392505050565b600081905092915050565b60006118188385611801565b93506118258385846113dc565b82840190509392505050565b600061183e82848661180c565b91508190509392505050565b7f446f6d61696e206e6f74206f776e65642062792073656e646572000000000000600082015250565b6000611880601a836114af565b915061188b8261184a565b602082019050919050565b600060208201905081810360008301526118af81611873565b9050919050565b60006118c283856114af565b93506118cf8385846113dc565b6118d88361131f565b840190509392505050565b600060208201905081810360008301526118fe8184866118b6565b90509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061194e57607f821691505b60208210810361196157611960611907565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006119d082611578565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611a0257611a01611996565b5b600182019050919050565b7f496e76616c696420686f73746e616d6500000000000000000000000000000000600082015250565b6000611a436010836114af565b9150611a4e82611a0d565b602082019050919050565b60006020820190508181036000830152611a7281611a36565b9050919050565b6000611a84826114a4565b611a8e8185611801565b9350611a9e8185602086016114c0565b80840191505092915050565b6000611ab68284611a79565b915081905092915050565b7f506172656e7420646f6d61696e206e6f742072656769737465726564206f722060008201527f6e6f74206f776e65642062792073656e64657200000000000000000000000000602082015250565b6000611b1d6033836114af565b9150611b2882611ac1565b604082019050919050565b60006020820190508181036000830152611b4c81611b10565b9050919050565b7f486f73746e616d6520616c7265616479207265676973746572656420616e642060008201527f6e6f742065787069726564000000000000000000000000000000000000000000602082015250565b6000611baf602b836114af565b9150611bba82611b53565b604082019050919050565b60006020820190508181036000830152611bde81611ba2565b9050919050565b6000611bf082611578565b9150611bfb83611578565b9250828201905080821115611c1357611c12611996565b5b92915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302611c7b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611c3e565b611c858683611c3e565b95508019841693508086168417925050509392505050565b6000819050919050565b6000611cc2611cbd611cb884611578565b611c9d565b611578565b9050919050565b6000819050919050565b611cdc83611ca7565b611cf0611ce882611cc9565b848454611c4b565b825550505050565b600090565b611d05611cf8565b611d10818484611cd3565b505050565b5b81811015611d3457611d29600082611cfd565b600181019050611d16565b5050565b601f821115611d7957611d4a81611c19565b611d5384611c2e565b81016020851015611d62578190505b611d76611d6e85611c2e565b830182611d15565b50505b505050565b600082821c905092915050565b6000611d9c60001984600802611d7e565b1980831691505092915050565b6000611db58383611d8b565b9150826002028217905092915050565b611dce826114a4565b67ffffffffffffffff811115611de757611de6611330565b5b611df18254611936565b611dfc828285611d38565b600060209050601f831160018114611e2f5760008415611e1d578287015190505b611e278582611da9565b865550611e8f565b601f198416611e3d86611c19565b60005b82811015611e6557848901518255600182019150602085019450602081019050611e40565b86831015611e825784890151611e7e601f891682611d8b565b8355505b6001600288020188555050505b505050505050565b60008190508160005260206000209050919050565b601f821115611eed57611ebe81611e97565b611ec784611c2e565b81016020851015611ed6578190505b611eea611ee285611c2e565b830182611d15565b50505b505050565b611efb82611523565b67ffffffffffffffff811115611f1457611f13611330565b5b611f1e8254611936565b611f29828285611eac565b600060209050601f831160018114611f5c5760008415611f4a578287015190505b611f548582611da9565b865550611fbc565b601f198416611f6a86611e97565b60005b82811015611f9257848901518255600182019150602085019450602081019050611f6d565b86831015611faf5784890151611fab601f891682611d8b565b8355505b6001600288020188555050505b505050505050565b60006080820190508181036000830152611fdf81888a6118b6565b90508181036020830152611ff48186886118b6565b90508181036040830152612008818561153f565b90506120176060830184611582565b979650505050505050565b7f446f6d61696e206e6f7420666f756e6400000000000000000000000000000000600082015250565b60006120586010836114af565b915061206382612022565b602082019050919050565b600060208201905081810360008301526120878161204b565b9050919050565b600082905092915050565b6120a3838361208e565b67ffffffffffffffff8111156120bc576120bb611330565b5b6120c68254611936565b6120d1828285611d38565b6000601f83116001811461210057600084156120ee578287013590505b6120f88582611da9565b865550612160565b601f19841661210e86611c19565b60005b8281101561213657848901358255600182019150602085019450602081019050612111565b86831015612153578489013561214f601f891682611d8b565b8355505b6001600288020188555050505b50505050505050565b7f4365727469666963617465206861732065787069726564000000000000000000600082015250565b600061219f6017836114af565b91506121aa82612169565b602082019050919050565b600060208201905081810360008301526121ce81612192565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061220f82611578565b915061221a83611578565b92508261222a576122296121d5565b5b828206905092915050565b600061224082611578565b915061224b83611578565b925082820390508181111561226357612262611996565b5b9291505056fea26469706673582212209f84ac6de8d1623e891921aa768ea281f4af5a392b5d9cb16f4a2e6d11559a7264736f6c63430008130033",
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
