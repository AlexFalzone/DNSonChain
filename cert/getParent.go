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
		IPAddresses:           util.GetIP(),
		Subject: pkix.Name{
			CommonName:   "test domain CA",
			SerialNumber: "DA DEFINIRE",
		},
	}

	grandParent, grandParentPriv := GetGrandParent(name)

	//Maybe it is not necessary??

	// aiaPubBytes, err := x509.MarshalPKIXPublicKey(util.PublicKey(grandParentPriv))
	// if err != nil {
	// 	log.Fatalf("failed to marshal AIA CA public key: %v", err)
	// }
	// aiaPubHash := sha256.Sum256(aiaPubBytes)
	// aiaPubHashStr := hex.EncodeToString(aiaPubHash[:])

	// aiaBaseURL := "aia.x--nmc.bit/aia"
	// aiaURL := aiaBaseURL + "?domain=" + *&name + "&pubsha256=" + aiaPubHashStr
	// template.IssuingCertificateURL = []string{"http://" + aiaURL}

	certificate, err := x509.CreateCertificate(rand.Reader, &template, &grandParent, util.PublicKey(priv), grandParentPriv)
	if err != nil {
		log.Fatalf("Failed to create Parent certificate: %v", err)
	}

	//Create a file called caCert.pem
	certOut, err := os.Create("caCert.pem")
	if err != nil {
		log.Fatalf("Failed to open caCert.pem for writing: %v", err)
	}

	defer certOut.Close()

	//Write the certificate to the file
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certificate}); err != nil {
		log.Fatalf("Failed to write data to caCert.pem: %v", err)
	}
	log.Print("wrote caCert.pem\n")

	//Create a file called caKey.pem
	keyOut, err := os.OpenFile("caKey.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Failed to open caKey.pem for writing: %v", err)
		//return
	}
	defer keyOut.Close()

	//convert the private key to PKCS8 format
	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		log.Fatalf("Unable to marshal private key: %v", err)
	}

	//Write the private key to the file
	if err := pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		log.Fatalf("Failed to write data to caKey.pem: %v", err)
	}

	log.Print("wrote caKey.pem\n")

	return template, priv
}
