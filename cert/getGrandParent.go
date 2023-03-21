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
		Subject: pkix.Name{
			CommonName:   name + "Root CA",
			SerialNumber: "DPKI TLS Certificate",
		},
	}

	certificate, err := x509.CreateCertificate(rand.Reader, &template, &template, util.PublicKey(priv), priv)
	if err != nil {
		log.Fatalf("Failed to create Parent certificate: %v", err)
	}

	if _, err := os.Stat("list"); os.IsNotExist(err) {
		err := os.Mkdir("list", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	certOut, err := os.Create("list/certRoot.pem")
	if err != nil {
		log.Fatalf("Failed to open certRoot.pem for writing: %v", err)
	}

	defer certOut.Close()

	//Write the certificate to the file
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certificate}); err != nil {
		log.Fatalf("Failed to write data to certRoot.pem: %v", err)
	}
	log.Print("wrote aia cert\n")

	//open (or create) a file called keyRoot.pem
	keyOut, err := os.OpenFile("list/keyRoot.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Failed to open keyRoot.pem for writing: %v", err)
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
		log.Fatalf("Failed to write data to keyRoot.pem: %v", err)
	}

	log.Print("wrote keyRoot.pem\n")

	return template, priv
}
