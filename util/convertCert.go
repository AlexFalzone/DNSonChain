package util

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

// Convert a certificate from PEM to DER format
func ConvertCertToDER(certFile string) ([]byte, error) {
	pemData, err := os.ReadFile(certFile)
	if err != nil {
		fmt.Printf("Impossibile leggere il file del certificato PEM: %v\n", err)
		return nil, err
	}

	// Decode the PEM data
	block, _ := pem.Decode(pemData)
	if block == nil {
		fmt.Println("Impossibile decodificare il certificato PEM")
		return nil, err
	}

	if block.Type != "CERTIFICATE" {
		fmt.Printf("Il blocco PEM non contiene un certificato: %v\n", err)
		return nil, err
	}

	// create a x509 certificate from the PEM data
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic(err)
	}

	return cert.Raw, nil
}

func ConvertCertToPEM(certBytes []byte) (string, error) {
	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		return "", errors.New("failed to parse certificate")
	}

	// Create a PEM block from the certificate bytes
	pemBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	}

	// Encode the PEM block as a string
	pemBytes := pem.EncodeToMemory(pemBlock)
	if pemBytes == nil {
		return "", errors.New("failed to encode PEM")
	}

	return string(pemBytes), nil
}
