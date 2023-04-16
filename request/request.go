package request

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"
	"test/domainRegistry"
	"test/util"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	client *ethclient.Client
	once   sync.Once
	err    error
)

// custom error type to handle dial client error
type DialClientError struct {
	err error
}

func (e *DialClientError) Error() string {
	return fmt.Sprintf("Dial Client Error: %v", e.err)
}

const (
	keyfile        = "private_key.key"
	contractAdress = "0x987a844139EC14B97e91E9e524dc86A24DC8D8bF"
)

// DialClient returns a singleton client instance to interact with the Ethereum blockchain.
// If the client is not yet initialized, it tries to establish a connection to the specified URL.
// It returns an error if the connection fails.
// Parameters:
//   - url: The Ethereum node URL as a string.
//
// Returns:
//   - *ethclient.Client: The client instance to interact with the Ethereum blockchain.
//   - error: An error, if any occurs during the connection process.
func DialClient(url string) (*ethclient.Client, error) {
	once.Do(func() {
		client, err = ethclient.Dial(url)
	})

	if client != nil {
		return client, nil
	} else {
		return nil, &DialClientError{err: err}
	}
}

// createDomain creates a new domain in the domain registry contract.
// It takes the transaction options, the contract instance, the domain name,
// and the certificate string as input. It sends the transaction and logs the result.
//
// Parameters:
//   - auth: *bind.TransactOpts representing the transaction options.
//   - instance: *domainRegistry.DomainRegistry instance representing the contract.
//   - domain: The domain name as a string.
//   - certificate: The certificate as a PEM-encoded string.
func createDomain(auth *bind.TransactOpts, instance *domainRegistry.DomainRegistry, domain string, certificate string, expirationTimestamp int64) {

	tx, err := instance.CreateDomain(auth, domain, util.GetIP(), certificate, big.NewInt(expirationTimestamp))
	if err != nil {
		fmt.Println("Impossibile creare il dominio: ", err)
		return
	}
	fmt.Printf("Domain created: %s", tx.Hash().Hex())
}

// updateDomain updates an existing domain in the domain registry contract.
// It takes the transaction options, the contract instance, the domain name,
// and the certificate string as input. It sends the transaction and logs the result.
//
// Parameters:
//   - auth: *bind.TransactOpts representing the transaction options.
//   - instance: *domainRegistry.DomainRegistry instance representing the contract.
//   - domain: The domain name as a string.
//   - certificate: The certificate as a PEM-encoded string.
func updateDomain(auth *bind.TransactOpts, instance *domainRegistry.DomainRegistry, domain string, certificate string, expirationTimestamp int64) {

	tx, err := instance.UpdateDomain(auth, domain, util.GetIP(), certificate, big.NewInt(expirationTimestamp))
	if err != nil {
		fmt.Println("Impossibile aggiornare il dominio: ", err)
		return
	}
	fmt.Printf("Domain updated: %s", tx.Hash().Hex())
}

// deleteDomain revokes a domain in the domain registry contract.
// It takes the transaction options and the contract instance as input,
// and sends the transaction to revoke the specified domain. It logs the result.
//
// Parameters:
//   - auth: *bind.TransactOpts representing the transaction options.
//   - instance: *domainRegistry.DomainRegistry instance representing the contract.
//   - domain: The domain name as a string.
func deleteDomain(auth *bind.TransactOpts, instance *domainRegistry.DomainRegistry, domain string) {
	tx, err := instance.DeleteDomain(auth, domain)
	if err != nil {
		fmt.Println("Impossibile revocare il dominio: ", err)
		return
	}
	fmt.Printf("Domain revoked: %s", tx.Hash().Hex())
}

// getCertificate retrieves the certificate for a domain from the domain registry contract.
// It takes the transaction options, the contract instance, and the domain name as input.
// It returns the IP address, the certificate as a PEM-encoded string, and an error if any occurs.
//
// Parameters:
//   - auth: *bind.TransactOpts representing the transaction options.
//   - instance: *domainRegistry.DomainRegistry instance representing the contract.
//   - domain: The domain name as a string.
//
// Returns:
//   - ip: The IP address of the domain as a string.
//   - certificate: The certificate of the domain as a PEM-encoded string.
//   - error: An error, if any occurs during the execution of the function.
func getCertificate(auth *bind.TransactOpts, instance *domainRegistry.DomainRegistry, domain string) (string, string, error) {
	ip, certificate, err := instance.GetCertificate(&bind.CallOpts{}, domain)
	if err != nil {
		fmt.Println("Impossibile recuperare il certificato: ", err)
		return "", "", err
	}
	return ip, certificate, err
}

// estimateGas estimates the gas required to perform a specific action on the domain registry contract.
// It takes the transaction options, the Ethereum client, the contract address, the function name,
// and a variadic list of arguments for the function. It returns the estimated gas as a uint64 and an error, if any occurs.
//
// Parameters:
//   - auth: *bind.TransactOpts representing the transaction options.
//   - client: *ethclient.Client instance representing the Ethereum client.
//   - address: common.Address representing the contract address.
//   - functionName: The name of the contract function to estimate gas for.
//   - args: A variadic list of arguments for the specified function.
//
// Returns:
//   - uint64: The estimated gas for the specified function call.
//   - error: An error, if any occurs during the gas estimation process.
func estimateGas(auth *bind.TransactOpts, client *ethclient.Client, address common.Address, functionName string, args ...interface{}) (uint64, error) {
	gasAuth := *auth
	gasAuth.GasLimit = 0

	_, err := domainRegistry.NewDomainRegistry(address, client)
	if err != nil {
		return 0, err
	}

	var gasEstimate uint64
	var input abi.Method

	contractABI, err := abi.JSON(strings.NewReader(domainRegistry.DomainRegistryABI))
	if err != nil {
		return 0, err
	}

	found := false
	for name, method := range contractABI.Methods {
		if name == functionName {
			input = method
			found = true
			break
		}
	}
	if !found {
		return 0, fmt.Errorf("unknown function name: %s", functionName)
	}

	encodedArgs, err := input.Inputs.Pack(args...)
	if err != nil {
		return 0, err
	}

	callMsg := ethereum.CallMsg{
		From: gasAuth.From,
		To:   &address,
		Data: append(input.ID, encodedArgs...),
	}
	gasEstimate, err = client.EstimateGas(context.Background(), callMsg)
	if err != nil {
		return 0, err
	}

	return gasEstimate, nil
}

// Request performs various actions on the domain registry contract based on the user's choice.
// It takes an ethclient.Client, an integer representing the action to perform, the domain name,
// and the certificate string as input. The actions are as follows:
// 1 - Create a new domain.
// 2 - Revoke a domain.
// 3 - Update an existing domain.
// 4 - Get the certificate for a domain.
// The function returns the IP address and certificate of the domain as strings, and an error if any occurs.
//
// Parameters:
//   - client: *ethclient.Client instance to interact with the Ethereum blockchain.
//   - choice: Integer representing the action to perform (1: Create, 2: Revoke, 3: Update, 4: Get certificate).
//   - name: The domain name as a string.
//   - cert: The certificate as a PEM-encoded string.
//
// Returns:
//   - ip: The IP address of the domain as a string.
//   - certificate: The certificate of the domain as a PEM-encoded string.
//   - error: An error, if any occurs during the execution of the function.
func Request(client *ethclient.Client, choice int, name string, cert string, expirationTimestamp int64) (string, string, error) {
	privateKey, err := ReadOrCreatePrivateKey(keyfile)
	if err != nil {
		log.Fatal(err)
	}

	// We then need to extract the Ethereum address from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasPrice = gasPrice

	address := common.HexToAddress(contractAdress)
	instance, err := domainRegistry.NewDomainRegistry(address, client)
	if err != nil {
		log.Fatal(err)
	}

	if choice == 1 { //create certificate
		estimatedGas, err := estimateGas(auth, client, address, "createDomain", name, util.GetIP(), cert, big.NewInt(expirationTimestamp))
		if err != nil {
			log.Fatalf("Failed to estimate gas: %v", err)
		}
		auth.GasLimit = estimatedGas // Use the estimated gas limit

		createDomain(auth, instance, name, cert, expirationTimestamp)
	} else if choice == 2 { //revoke certificate
		estimatedGas, err := estimateGas(auth, client, address, "deleteDomain", name)
		if err != nil {
			log.Fatalf("Failed to estimate gas: %v", err)
		}
		auth.GasLimit = estimatedGas // Use the estimated gas limit

		deleteDomain(auth, instance, name)
	} else if choice == 3 { //update certificate
		estimatedGas, err := estimateGas(auth, client, address, "updateDomain", name, util.GetIP(), cert, big.NewInt(expirationTimestamp))
		if err != nil {
			log.Fatalf("Failed to estimate gas: %v", err)
		}
		auth.GasLimit = estimatedGas // Use the estimated gas limit

		updateDomain(auth, instance, name, cert, expirationTimestamp)
	} else if choice == 4 { //get certificate
		ip, certificate, err := getCertificate(auth, instance, name)
		if err != nil {
			log.Fatal(err)
		}
		return ip, certificate, err
	}
	return "", "", nil
}
