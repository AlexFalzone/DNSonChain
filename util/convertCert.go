package util

import (
	"bytes"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"os"
)

// Converte un certificato da PEM a DER
func ConvertCertToDER(certName string) error {
	pemData, err := os.ReadFile(certName)
	if err != nil {
		panic(err)
	}

	// Estrai il blocco PEMsss
	block, _ := pem.Decode(pemData)

	// Parsa il certificato
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic(err)
	}

	// Serializza il certificato in formato DER
	derBytes, err := asn1.Marshal(cert.Raw)
	if err != nil {
		panic(err)
	}

	nameWithoutSuffix, err := removeSuffix(certName)
	if err != nil {
		panic(err)
	}

	// Salva i dati in un file DER
	err = os.WriteFile(nameWithoutSuffix+".der", derBytes, 0644)
	if err != nil {
		panic(err)
	}
	return nil
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
