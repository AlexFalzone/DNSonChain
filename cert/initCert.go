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

/*
	TODO:
	-add the possibility to choose the signature algorithm and the key type
*/

const (
	rsaBits = 2048
)

func GenerateCert(name string /*choice int*/) (x509.Certificate, any) {

	priv, err := rsa.GenerateKey(rand.Reader, rsaBits)
	if err != nil {
		log.Fatal(err)
	}

	//it is necessary for rsa!
	keyUsage := x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment

	notBefore := time.Now()
	//set the validity of the certificate to 1 year
	notAfter := notBefore.Add(365 * 24 * time.Hour)

	//generate a random serial number between 0 and 2^128
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("Failed to generate serial number: %v", err)
	}

	//call getParent to get the parent certificate and the parent private key
	parent, parentPriv := GetParent(name)

	template := x509.Certificate{
		//we set the public key of parent as the public key of the certificate
		//it is necessary for create the certificate from the blockchain
		PublicKey:             util.PublicKey(parentPriv),
		SerialNumber:          serialNumber,
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  false,
		DNSNames:              []string{name},
		Subject: pkix.Name{
			CommonName:   name,
			SerialNumber: "DPKI TLS Certificate",
		},
	}

	//create the certificate using the parent certificate and the parent private key
	certificate, err := x509.CreateCertificate(rand.Reader, &template, &parent, util.PublicKey(priv), parentPriv)
	if err != nil {
		log.Fatal("failed to create certificate: ", err)
	}

	//create a file called cert.pem
	certOut, err := os.Create("list/certHost.pem")
	if err != nil {
		log.Fatal(err)
	}

	defer certOut.Close()

	//write the certificate into cert.pem
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certificate}); err != nil {
		log.Fatal(err)
	}
	log.Print("written certHost.pem\n")

	//create a file called keyHost.pem
	keyOut, err := os.OpenFile("list/keyHost.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Failed to open keyHost.pem for writing: %v", err)
	}

	defer keyOut.Close()

	//convert the private key into PKCS#8 format
	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		log.Fatalf("Unable to marshal private key: %v", err)
	}

	//write the private key into key.pem
	if err := pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		log.Fatalf("Failed to write data to keyHost.pem: %v", err)
	}
	log.Print("wrote keyHost.pem\n")

	writeChain()

	return template, priv
}
