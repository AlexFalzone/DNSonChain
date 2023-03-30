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
	Bin: "0x608060405234801561001057600080fd5b50612088806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630eaf97521461005c57806326449235146100785780632ad09666146100ac578063e532fb98146100c8578063ed0f2e75146100e4575b600080fd5b61007660048036038101906100719190611051565b610115565b005b610092600480360381019061008d91906111df565b610285565b6040516100a3959493929190611356565b60405180910390f35b6100c660048036038101906100c19190611440565b610489565b005b6100e260048036038101906100dd9190611440565b6109f0565b005b6100fe60048036038101906100f99190611051565b610bb5565b60405161010c929190611509565b60405180910390f35b3373ffffffffffffffffffffffffffffffffffffffff166000838360405161013e929190611570565b908152602001604051809103902060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146101c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101bd906115d5565b60405180910390fd5b600082826040516101d8929190611570565b9081526020016040518091039020600080820160006101f79190610f3b565b6001820160006102079190610f3b565b6002820160006102179190610f7b565b60038201600090556004820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550507f9b449e8e6e58cfe4813dc8a38f233368018ddd99ddc223e2f1a55eb22eefa2368282604051610279929190611622565b60405180910390a15050565b6000818051602081018201805184825260208301602085012081835280955050505050506000915090508060000180546102be90611675565b80601f01602080910402602001604051908101604052809291908181526020018280546102ea90611675565b80156103375780601f1061030c57610100808354040283529160200191610337565b820191906000526020600020905b81548152906001019060200180831161031a57829003601f168201915b50505050509080600101805461034c90611675565b80601f016020809104026020016040519081016040528092919081815260200182805461037890611675565b80156103c55780601f1061039a576101008083540402835291602001916103c5565b820191906000526020600020905b8154815290600101906020018083116103a857829003601f168201915b5050505050908060020180546103da90611675565b80601f016020809104026020016040519081016040528092919081815260200182805461040690611675565b80156104535780601f1061042857610100808354040283529160200191610453565b820191906000526020600020905b81548152906001019060200180831161043657829003601f168201915b5050505050908060030154908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905085565b600087878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505090506000805b8251811015610567577f2e00000000000000000000000000000000000000000000000000000000000000838281518110610513576105126116a6565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19160361055457818061055090611704565b9250505b808061055f90611704565b9150506104d6565b50600081116105ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a290611798565b60405180910390fd5b60018111156106725760006105bf83610da0565b90503373ffffffffffffffffffffffffffffffffffffffff166000826040516105e891906117e9565b908152602001604051809103902060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610670576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066790611872565b60405180910390fd5b505b6000808a8a604051610685929190611570565b908152602001604051809103902060000180546106a190611675565b90501115610792574260008a8a6040516106bc929190611570565b908152602001604051809103902060030154111561070f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070690611904565b60405180910390fd5b60008989604051610721929190611570565b9081526020016040518091039020600080820160006107409190610f3b565b6001820160006107509190610f3b565b6002820160006107609190610f7b565b60038201600090556004820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550505b6040518060a001604052808a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200188888080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200186868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020016000851461089457844261088f9190611924565b6108b6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b81526020013373ffffffffffffffffffffffffffffffffffffffff1681525060008a8a6040516108e7929190611570565b9081526020016040518091039020600082015181600001908161090a9190611b04565b5060208201518160010190816109209190611b04565b5060408201518160020190816109369190611c31565b506060820151816003015560808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050503373ffffffffffffffffffffffffffffffffffffffff167fd694233903b144a99ef1896eb8f5ebff710c263412bfcff1b8d37d5c574e1d108a8a8a8a8a8a8a6040516109dd9796959493929190611d30565b60405180910390a2505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff1660008888604051610a19929190611570565b908152602001604051809103902060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610aa1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a98906115d5565b60405180910390fd5b848460008989604051610ab5929190611570565b90815260200160405180910390206001019182610ad3929190611d9b565b50828260008989604051610ae8929190611570565b90815260200160405180910390206002019182610b06929190611e76565b5060008114610b20578042610b1b9190611924565b610b42565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b60008888604051610b54929190611570565b9081526020016040518091039020600301819055507f6ee95bc81bd859a73439a27fc26eb56779fa812f28d8271df16444b27b0e02fd87878787878787604051610ba49796959493929190611d30565b60405180910390a150505050505050565b6060806000808585604051610bcb929190611570565b908152602001604051809103902090506000816000018054610bec90611675565b905011610c2e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c2590611f92565b60405180910390fd5b42816003015411610c74576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6b90611ffe565b60405180910390fd5b8060010181600201818054610c8890611675565b80601f0160208091040260200160405190810160405280929190818152602001828054610cb490611675565b8015610d015780601f10610cd657610100808354040283529160200191610d01565b820191906000526020600020905b815481529060010190602001808311610ce457829003601f168201915b50505050509150808054610d1490611675565b80601f0160208091040260200160405190810160405280929190818152602001828054610d4090611675565b8015610d8d5780601f10610d6257610100808354040283529160200191610d8d565b820191906000526020600020905b815481529060010190602001808311610d7057829003601f168201915b5050505050905092509250509250929050565b60606000805b8351811015610e30577f2e00000000000000000000000000000000000000000000000000000000000000848281518110610de357610de26116a6565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191603610e1d57809150610e30565b8080610e2890611704565b915050610da6565b5060006001828551610e42919061201e565b610e4c919061201e565b67ffffffffffffffff811115610e6557610e646110b4565b5b6040519080825280601f01601f191660200182016040528015610e975781602001600182028036833780820191505090505b50905060005b8151811015610f30578481600185610eb59190611924565b610ebf9190611924565b81518110610ed057610ecf6116a6565b5b602001015160f81c60f81b828281518110610eee57610eed6116a6565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080610f2890611704565b915050610e9d565b508092505050919050565b508054610f4790611675565b6000825580601f10610f595750610f78565b601f016020900490600052602060002090810190610f779190610fbb565b5b50565b508054610f8790611675565b6000825580601f10610f995750610fb8565b601f016020900490600052602060002090810190610fb79190610fbb565b5b50565b5b80821115610fd4576000816000905550600101610fbc565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f84011261101157611010610fec565b5b8235905067ffffffffffffffff81111561102e5761102d610ff1565b5b60208301915083600182028301111561104a57611049610ff6565b5b9250929050565b6000806020838503121561106857611067610fe2565b5b600083013567ffffffffffffffff81111561108657611085610fe7565b5b61109285828601610ffb565b92509250509250929050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6110ec826110a3565b810181811067ffffffffffffffff8211171561110b5761110a6110b4565b5b80604052505050565b600061111e610fd8565b905061112a82826110e3565b919050565b600067ffffffffffffffff82111561114a576111496110b4565b5b611153826110a3565b9050602081019050919050565b82818337600083830152505050565b600061118261117d8461112f565b611114565b90508281526020810184848401111561119e5761119d61109e565b5b6111a9848285611160565b509392505050565b600082601f8301126111c6576111c5610fec565b5b81356111d684826020860161116f565b91505092915050565b6000602082840312156111f5576111f4610fe2565b5b600082013567ffffffffffffffff81111561121357611212610fe7565b5b61121f848285016111b1565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015611262578082015181840152602081019050611247565b60008484015250505050565b600061127982611228565b6112838185611233565b9350611293818560208601611244565b61129c816110a3565b840191505092915050565b600081519050919050565b600082825260208201905092915050565b60006112ce826112a7565b6112d881856112b2565b93506112e8818560208601611244565b6112f1816110a3565b840191505092915050565b6000819050919050565b61130f816112fc565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061134082611315565b9050919050565b61135081611335565b82525050565b600060a0820190508181036000830152611370818861126e565b90508181036020830152611384818761126e565b9050818103604083015261139881866112c3565b90506113a76060830185611306565b6113b46080830184611347565b9695505050505050565b60008083601f8401126113d4576113d3610fec565b5b8235905067ffffffffffffffff8111156113f1576113f0610ff1565b5b60208301915083600182028301111561140d5761140c610ff6565b5b9250929050565b61141d816112fc565b811461142857600080fd5b50565b60008135905061143a81611414565b92915050565b60008060008060008060006080888a03121561145f5761145e610fe2565b5b600088013567ffffffffffffffff81111561147d5761147c610fe7565b5b6114898a828b01610ffb565b9750975050602088013567ffffffffffffffff8111156114ac576114ab610fe7565b5b6114b88a828b01610ffb565b9550955050604088013567ffffffffffffffff8111156114db576114da610fe7565b5b6114e78a828b016113be565b935093505060606114fa8a828b0161142b565b91505092959891949750929550565b60006040820190508181036000830152611523818561126e565b9050818103602083015261153781846112c3565b90509392505050565b600081905092915050565b60006115578385611540565b9350611564838584611160565b82840190509392505050565b600061157d82848661154b565b91508190509392505050565b7f446f6d61696e206e6f74206f776e65642062792073656e646572000000000000600082015250565b60006115bf601a83611233565b91506115ca82611589565b602082019050919050565b600060208201905081810360008301526115ee816115b2565b9050919050565b60006116018385611233565b935061160e838584611160565b611617836110a3565b840190509392505050565b6000602082019050818103600083015261163d8184866115f5565b90509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061168d57607f821691505b6020821081036116a05761169f611646565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061170f826112fc565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611741576117406116d5565b5b600182019050919050565b7f496e76616c696420686f73746e616d6500000000000000000000000000000000600082015250565b6000611782601083611233565b915061178d8261174c565b602082019050919050565b600060208201905081810360008301526117b181611775565b9050919050565b60006117c382611228565b6117cd8185611540565b93506117dd818560208601611244565b80840191505092915050565b60006117f582846117b8565b915081905092915050565b7f506172656e7420646f6d61696e206e6f742072656769737465726564206f722060008201527f6e6f74206f776e65642062792073656e64657200000000000000000000000000602082015250565b600061185c603383611233565b915061186782611800565b604082019050919050565b6000602082019050818103600083015261188b8161184f565b9050919050565b7f486f73746e616d6520616c7265616479207265676973746572656420616e642060008201527f6e6f742065787069726564000000000000000000000000000000000000000000602082015250565b60006118ee602b83611233565b91506118f982611892565b604082019050919050565b6000602082019050818103600083015261191d816118e1565b9050919050565b600061192f826112fc565b915061193a836112fc565b9250828201905080821115611952576119516116d5565b5b92915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026119ba7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261197d565b6119c4868361197d565b95508019841693508086168417925050509392505050565b6000819050919050565b6000611a016119fc6119f7846112fc565b6119dc565b6112fc565b9050919050565b6000819050919050565b611a1b836119e6565b611a2f611a2782611a08565b84845461198a565b825550505050565b600090565b611a44611a37565b611a4f818484611a12565b505050565b5b81811015611a7357611a68600082611a3c565b600181019050611a55565b5050565b601f821115611ab857611a8981611958565b611a928461196d565b81016020851015611aa1578190505b611ab5611aad8561196d565b830182611a54565b50505b505050565b600082821c905092915050565b6000611adb60001984600802611abd565b1980831691505092915050565b6000611af48383611aca565b9150826002028217905092915050565b611b0d82611228565b67ffffffffffffffff811115611b2657611b256110b4565b5b611b308254611675565b611b3b828285611a77565b600060209050601f831160018114611b6e5760008415611b5c578287015190505b611b668582611ae8565b865550611bce565b601f198416611b7c86611958565b60005b82811015611ba457848901518255600182019150602085019450602081019050611b7f565b86831015611bc15784890151611bbd601f891682611aca565b8355505b6001600288020188555050505b505050505050565b60008190508160005260206000209050919050565b601f821115611c2c57611bfd81611bd6565b611c068461196d565b81016020851015611c15578190505b611c29611c218561196d565b830182611a54565b50505b505050565b611c3a826112a7565b67ffffffffffffffff811115611c5357611c526110b4565b5b611c5d8254611675565b611c68828285611beb565b600060209050601f831160018114611c9b5760008415611c89578287015190505b611c938582611ae8565b865550611cfb565b601f198416611ca986611bd6565b60005b82811015611cd157848901518255600182019150602085019450602081019050611cac565b86831015611cee5784890151611cea601f891682611aca565b8355505b6001600288020188555050505b505050505050565b6000611d0f83856112b2565b9350611d1c838584611160565b611d25836110a3565b840190509392505050565b60006080820190508181036000830152611d4b81898b6115f5565b90508181036020830152611d608187896115f5565b90508181036040830152611d75818587611d03565b9050611d846060830184611306565b98975050505050505050565b600082905092915050565b611da58383611d90565b67ffffffffffffffff811115611dbe57611dbd6110b4565b5b611dc88254611675565b611dd3828285611a77565b6000601f831160018114611e025760008415611df0578287013590505b611dfa8582611ae8565b865550611e62565b601f198416611e1086611958565b60005b82811015611e3857848901358255600182019150602085019450602081019050611e13565b86831015611e555784890135611e51601f891682611aca565b8355505b6001600288020188555050505b50505050505050565b600082905092915050565b611e808383611e6b565b67ffffffffffffffff811115611e9957611e986110b4565b5b611ea38254611675565b611eae828285611beb565b6000601f831160018114611edd5760008415611ecb578287013590505b611ed58582611ae8565b865550611f3d565b601f198416611eeb86611bd6565b60005b82811015611f1357848901358255600182019150602085019450602081019050611eee565b86831015611f305784890135611f2c601f891682611aca565b8355505b6001600288020188555050505b50505050505050565b7f446f6d61696e206e6f7420666f756e6400000000000000000000000000000000600082015250565b6000611f7c601083611233565b9150611f8782611f46565b602082019050919050565b60006020820190508181036000830152611fab81611f6f565b9050919050565b7f4365727469666963617465206861732065787069726564000000000000000000600082015250565b6000611fe8601783611233565b9150611ff382611fb2565b602082019050919050565b6000602082019050818103600083015261201781611fdb565b9050919050565b6000612029826112fc565b9150612034836112fc565b925082820390508181111561204c5761204b6116d5565b5b9291505056fea2646970667358221220d9adfc4541ff4aa4ad19ffa8fcf8b59ff89c705f78a7db00c2c472d4313f9da364736f6c63430008130033",
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
