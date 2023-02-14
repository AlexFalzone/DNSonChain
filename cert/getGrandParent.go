package cert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"time"

	"test/util"
)

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
		IPAddresses:           util.GetIP(),
		Subject: pkix.Name{
			CommonName:   "test domain AIA Parent CA",
			SerialNumber: "DA DEFINIRE",
		},
	}

	//open (or create) a file called caAIAKey.pem
	keyOut, err := os.OpenFile("caAIAKey.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Failed to open caAIAKey.pem for writing: %v", err)
		//return
	}
	defer keyOut.Close()

	//convert the private key to PKCS#8 format
	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		log.Fatalf("Unable to marshal private key: %v", err)
	}

	//write the private key to the file
	if err := pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		log.Fatalf("Failed to write data to caAIAKey.pem: %v", err)
	}

	log.Print("wrote caAIAKey.pem\n")

	return template, priv
}
