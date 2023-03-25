package request

import (
	"crypto/ecdsa"
	"encoding/hex"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
)

// Generate a new ecdsa private key
func GenerateKeyPair() (*ecdsa.PrivateKey, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	return key, nil
}

func ReadOrCreatePrivateKey(keyfile string) (*ecdsa.PrivateKey, error) {
	var privateKey *ecdsa.PrivateKey

	// Check if the keyfile already exists
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
		// Generate a new key
		privateKey, err = GenerateKeyPair()
		if err != nil {
			return nil, err
		}

		// Write the key to file
		err = os.WriteFile(keyfile, []byte(hex.EncodeToString(crypto.FromECDSA(privateKey))), 0600)
		if err != nil {
			return nil, err
		}
	}

	return privateKey, nil
}
