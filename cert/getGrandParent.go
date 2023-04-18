package cert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"log"
	"math/big"
	"os"
	"time"

	"test/util"
)

// GetGrandParent generates a new grandparent certificate and private key.
// It takes a domain name as input and returns the generated x509.Certificate and the private key.
//
// Parameters:
//   - name: string representing the domain name.
//
// Returns:
//   - x509.Certificate: The generated grandparent certificate.
//   - any: The generated private key.
//   - error: An error, if any occurs during the certificate and private key generation.
func GetGrandParent(name string) (x509.Certificate, any) {
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
		log.Fatalf("Failed to generate grandparent serial number: %v", err)
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
			CommonName:   name + " Root CA",
			SerialNumber: "DNSonChain TLS Certificate",
		},
	}

	_, err = x509.CreateCertificate(rand.Reader, &template, &template, util.PublicKey(priv), priv)
	if err != nil {
		log.Fatalf("Failed to create Parent certificate: %v", err)
	}

	if _, err := os.Stat("list"); os.IsNotExist(err) {
		err := os.Mkdir("list", 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	return template, priv
}
