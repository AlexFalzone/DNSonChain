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
	"test/util"
	"time"
)

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
			CommonName:   name + "Intermediate CA",
			SerialNumber: "DPKI TLS Certificate",
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

	//Create a file called keyIntermediate.pem
	keyOut, err := os.OpenFile("list/keyIntermediate.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Failed to open keyIntermediate.pem for writing: %v", err)
	}
	defer keyOut.Close()

	//convert the private key to PKCS8 format
	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		log.Fatalf("Unable to marshal private key: %v", err)
	}

	//Write the private key to the file
	if err := pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		log.Fatalf("Failed to write data to keyIntermediate.pem: %v", err)
	}

	log.Print("wrote keyIntermediate.pem\n")

	return template, priv
}
