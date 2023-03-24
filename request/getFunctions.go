package request

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"math/big"
	"os"
	"sync"
	"test/domainRegistry"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	once sync.Once
)

const (
	keyfile        = "private_key.key"
	contractAdress = "test"
)

func DialClient(client *ethclient.Client, url string) (*ethclient.Client, error) {
	var error error
	once.Do(func() {
		client, error = ethclient.Dial(url)
		if error != nil {
			log.Fatal(error)
		}
	})
	return client, error
}

// Generate a new ecdsa private key
func generateKeyPair() (*ecdsa.PrivateKey, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	return key, nil
}

func readOrCreatePrivateKey(keyfile string) (*ecdsa.PrivateKey, error) {
	var privateKey *ecdsa.PrivateKey

	// Controlla se il file esiste
	if _, err := os.Stat(keyfile); err == nil {
		keyBytes, err := os.ReadFile(keyfile)
		if err != nil {
			return nil, err
		}

		privateKeyBytes, err := hex.DecodeString(string(keyBytes))
		if err != nil {
			return nil, err
		}

		privateKey, err = crypto.ToECDSA(privateKeyBytes)
		if err != nil {
			return nil, err
		}
	} else {
		// Genera una nuova chiave privata
		privateKey, err = generateKeyPair()
		if err != nil {
			return nil, err
		}

		// Salva la chiave privata nel file
		err = os.WriteFile(keyfile, []byte(hex.EncodeToString(crypto.FromECDSA(privateKey))), 0600)
		if err != nil {
			return nil, err
		}
	}

	return privateKey, nil
}

func createDomain(auth *bind.TransactOpts, instance *domainRegistry.DomainRegistry) {
}

func updateDomain(auth *bind.TransactOpts, instance *domainRegistry.DomainRegistry) {
}

func deleteDomain(auth *bind.TransactOpts, instance *domainRegistry.DomainRegistry) {
}

func getCertificate(auth *bind.TransactOpts, instance *domainRegistry.DomainRegistry) {
}

func request(client *ethclient.Client) {
	privateKey, err := generateKeyPair()
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

}
