package cert

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"net"
	"regexp"
	"strings"
	"time"
)

type JSONResultField struct {
	IPAddresses        []net.IP  `json:"ipAddresses"`
	NotAfter           time.Time `json:"notAfter"`
	NotBefore          time.Time `json:"notBefore"`
	PublicKey          string    `json:"publicKey"`
	SerialNumber       *big.Int  `json:"serialNumber"`
	Signature          []byte    `json:"signature"`
	SignatureAlgorithm int64     `json:"signatureAlgorithm"`
	Subject            pkix.Name `json:"subject"`
}

/*
campi certificato: x509
-BasicConstraintsValid 	true
-ExtKeyUsage 			[]x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}
-IPAddresses			(4 o 16 bytes) rispettivamente IPv4 o IPv6
-KeyUsage				x509.KeyUsageDigitalSignature
-NotAfter
-NotBefore
-PublicKey				rsa (2048, 3072, 4096 bytes) ecdsa (256, 384, 521 bytes) ed25519 (256 bytes)
-SerialNumber 			!!MUST be 20 bytes and unique!!
-Signature
-SignatureAlgorithm		SHA512WithRSA
-Subject
*/

/*
	Nella struct che verrà tornata dal smart contract, inserire i campi con i relativi nomi. In questo modo viene più facile prenderli e assegnarli al certificato.
	Controllare i campi notafter e notbefore, non so se sono corretti.
	Controllare il campo SignatureAlgorithm, non so se è corretto.
	I campi del certificato dovrò prenderli manualmente dalla string restituita, assegnare ad ogni stringa il suo campo e poi creare il certificato?
*/

//takes a string in the format of "key:value"
//and returns the value as a JSONResultField
func ExtractJSONResultField(data string) (result JSONResultField, err error) {
	// Initialize an empty JSONResultField struct.
	//result := &JSONResultField{}

	re := regexp.MustCompile(`(\w+):"([^"]+)",`)
	matches := re.FindAllStringSubmatch(data, -1)

	for _, match := range matches {
		key := match[1]
		value := match[2]

		switch key {
		case "ipAddresses":
			ip := net.ParseIP(value)
			if ip == nil {
				return result, fmt.Errorf("invalid IP address: %s", value)
			}
			result.IPAddresses = append(result.IPAddresses, ip)
		case "notAfter":
			t, err := time.Parse("2006-Jan-02", value)
			if err != nil {
				return result, fmt.Errorf("invalid time format: %s", value)
			}
			result.NotAfter = t
		case "notBefore":
			t, err := time.Parse("2006-Jan-02", value)
			if err != nil {
				return result, fmt.Errorf("invalid time format: %s", value)
			}
			result.NotBefore = t
		case "publicKey":
			result.PublicKey = value
		case "serialNumber":
			sn, ok := big.NewInt(0).SetString(value, 10)
			if !ok {
				return result, fmt.Errorf("invalid serial number: %s", value)
			}
			result.SerialNumber = sn
		case "signature":
			result.Signature = []byte(value)
		case "signatureAlgorithm":
			result.SignatureAlgorithm = 10
		case "subject":
			keyValue := strings.Split(value, "=")
			subjectKey := strings.TrimSpace(keyValue[0])
			subjectValue := strings.TrimSpace(keyValue[1])
			fmt.Println("subjectKey: ", subjectKey)
			fmt.Println("subjectValue: ", subjectValue)
			if subjectKey == "CN" {
				result.Subject = pkix.Name{
					CommonName: subjectValue,
				}
			} else {
				return result, fmt.Errorf("invalid subject key: %s", subjectKey)
			}
		default:
			return result, fmt.Errorf("unknown field: %s", key)
		}

	}
	return result, nil

}

func CreateCert(result JSONResultField) (cert *x509.Certificate) {
	template := x509.Certificate{
		IPAddresses:        result.IPAddresses,
		NotAfter:           result.NotAfter,
		NotBefore:          result.NotBefore,
		PublicKey:          result.PublicKey,
		SerialNumber:       result.SerialNumber,
		Signature:          result.Signature,
		SignatureAlgorithm: x509.SHA512WithRSA,
		Subject:            result.Subject,
	}

	x509.CreateCertificate(rand.Reader, &template, &template, template.PublicKey /*add private key*/)
}
