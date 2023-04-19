package cert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"dnsonchain/util"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"time"
)

// GetParent generates a new parent certificate and private key signed by the grandparent certificate.
// It takes a domain name as input and returns the generated x509.Certificate and the private key.
//
// Parameters:
//   - name: string representing the domain name.
//
// Returns:
//   - x509.Certificate: The generated parent certificate.
//   - any: The generated private key.
//   - error: An error, if any occurs during the certificate and private key generation.
func GetParent(name string) (x509.Certificate, any) {
	priv, err := rsa.GenerateKey(rand.Reader, rsaBits)
	if err != nil {
		log.Fatal(err)
	}

	//it is necessary for rsa!
	keyUsage := x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageCertSign

	notBefore := time.Now()
	notAfter := notBefore.Add(365 * 24 * time.Hour)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("Failed to generate parent serial number: %v", err)
	}

	template := x509.Certificate{
		PublicKey:             util.PublicKey(priv),
		SerialNumber:          serialNumber,
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
		Subject: pkix.Name{
			CommonName:   name + " Intermediate CA",
			SerialNumber: "DNSonChain TLS Certificate",
		},
	}

	grandParent, grandParentPriv := GetGrandParent(name)

	certificate, err := x509.CreateCertificate(rand.Reader, &template, &grandParent, util.PublicKey(priv), grandParentPriv)
	if err != nil {
		log.Fatalf("Failed to create Parent certificate: %v", err)
	}

	//Create a file called caCert.pem
	certOut, err := os.Create("list/certIntermediate.pem")
	if err != nil {
		log.Fatalf("Failed to open certIntermediate.pem for writing: %v", err)
	}

	defer certOut.Close()

	//Write the certificate to the file
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certificate}); err != nil {
		log.Fatalf("Failed to write data to certIntermediate.pem: %v", err)
	}
	log.Print("wrote certIntermediate.pem\n")

	return template, priv
}
