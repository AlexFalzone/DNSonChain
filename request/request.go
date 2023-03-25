package request

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"sync"
	"test/domainRegistry"
	"test/util"
	"time"

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
	contractAdress = "test"
)

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

func createDomain(auth *bind.TransactOpts, instance *domainRegistry.DomainRegistry, domain string, certificate []byte) {
	//FOR NOW
	//insert a 1 day validity
	date := time.Now().Add(time.Duration(time.Now().Day())).Unix()

	tx, err := instance.CreateDomain(auth, domain, util.GetIP(), certificate, big.NewInt(date))
	if err != nil {
		log.Fatalf("Impossibile creare il dominio: %v", err)
	}
	fmt.Printf("Domain created: %s", tx.Hash().Hex())
}

func updateDomain(auth *bind.TransactOpts, instance *domainRegistry.DomainRegistry, domain string, certificate []byte) {
	//FOR NOW
	//insert a 1 day validity
	date := time.Now().Add(time.Duration(time.Now().Day())).Unix()

	tx, err := instance.UpdateDomain(auth, domain, util.GetIP(), certificate, big.NewInt(date))
	if err != nil {
		log.Fatalf("Impossibile aggiornare il dominio: %v", err)
	}
	fmt.Printf("Domain updated: %s", tx.Hash().Hex())
}

func deleteDomain(auth *bind.TransactOpts, instance *domainRegistry.DomainRegistry, domain string) {
	tx, err := instance.DeleteDomain(auth, domain)
	if err != nil {
		log.Fatalf("Impossibile revocare il dominio: %v", err)
	}
	fmt.Printf("Domain revoked: %s", tx.Hash().Hex())
}

func getCertificate(auth *bind.TransactOpts, instance *domainRegistry.DomainRegistry, domain string) (string, []byte, error) {
	ip, certificate, err := instance.GetCertificate(&bind.CallOpts{}, domain)
	if err != nil {
		log.Fatalf("Impossibile recuperare il certificato: %v", err)
	}
	return ip, certificate, err
}

func Request(client *ethclient.Client, choice int, name string, cert []byte) (string, []byte, error) {
	privateKey, err := ReadOrCreatePrivateKey(keyfile)
	if err != nil {
		log.Fatal(err)
	}

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
	auth.GasLimit = uint64(0)  //con 0 si stima automaticamente
	auth.GasPrice = gasPrice

	address := common.HexToAddress(contractAdress)
	instance, err := domainRegistry.NewDomainRegistry(address, client)
	if err != nil {
		log.Fatal(err)
	}

	if choice == 1 { //create certificate
		createDomain(auth, instance, name, cert)
	} else if choice == 2 { //revoke certificate
		deleteDomain(auth, instance, name)
	} else if choice == 3 { //update certificate
		updateDomain(auth, instance, name, cert)
	} else if choice == 4 { //get certificate
		ip, certificate, err := getCertificate(auth, instance, name)
		if err != nil {
			log.Fatal(err)
		}
		return ip, certificate, err
	}
	return "", nil, nil
}
