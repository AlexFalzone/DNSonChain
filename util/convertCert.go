package util

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
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

func ConvertCertificateToPEM(certificateBytes []byte) (string, error) {
	// Create a PEM block from the certificate bytes
	pemBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certificateBytes,
	}

	// Encode the PEM block as a string
	var pemBytes bytes.Buffer
	err := pem.Encode(&pemBytes, pemBlock)
	if err != nil {
		return "", fmt.Errorf("failed to encode PEM block: %v", err)
	}

	return pemBytes.String(), nil
}
