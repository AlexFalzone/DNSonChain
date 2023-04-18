package request

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sync"
	"test/domainRegistry"
	"test/util"

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
		return "", "", fmt.Errorf("impossible to get certificate: %v", err)
	}
	return ip, certificate, err
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
		return "", "", fmt.Errorf("impossible to read or create the private key: %v", err)
	}

	// We then need to extract the Ethereum address from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", "", fmt.Errorf("impossible to get the nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", "", fmt.Errorf("impossible to get the gas price: %v", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", "", fmt.Errorf("impossible to get the chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", "", fmt.Errorf("impossible to get the transaction authorizer: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasPrice = gasPrice

	address := common.HexToAddress(contractAdress)
	instance, err := domainRegistry.NewDomainRegistry(address, client)
	if err != nil {
		return "", "", fmt.Errorf("impossible to get the contract instance: %v", err)
	}

	if choice == 1 { //create certificate
		createDomain(auth, instance, name, cert, expirationTimestamp)
	} else if choice == 2 { //revoke certificate
		deleteDomain(auth, instance, name)
	} else if choice == 3 { //update certificate
		updateDomain(auth, instance, name, cert, expirationTimestamp)
	} else if choice == 4 { //get certificate
		ip, certificate, err := getCertificate(auth, instance, name)
		if err != nil {
			return "", "", fmt.Errorf("impossible to get the certificate: %v", err)
		}
		return ip, certificate, err
	}
	return "", "", nil
}
