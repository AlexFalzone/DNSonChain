package request

import (
	"crypto/ecdsa"
	"encoding/hex"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
)

// GenerateKeyPair generates a new ECDSA private key.
//
// Returns:
//   - *ecdsa.PrivateKey: A pointer to the generated ECDSA private key.
//   - error: An error if the operation fails, nil otherwise.
func GenerateKeyPair() (*ecdsa.PrivateKey, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	return key, nil
}

// ReadOrCreatePrivateKey reads an existing ECDSA private key from a file or generates a new one if the file does not exist.
//
// Parameters:
//   - keyfile: string representing the path to the private key file.
//
// Returns:
//   - *ecdsa.PrivateKey: A pointer to the ECDSA private key read from the file or newly generated.
//   - error: An error if the operation fails, nil otherwise.
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
