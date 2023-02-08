package cert

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"net"
	"regexp"
	"time"

	"github.com/namecoin/splicesign"
)

//TODO: signature algorithm is hardcoded to RSA, so it's useless to pass it as a parameter

/*
	campi da inserire nella blockchain:
	-IPAddresses
	-NotAfter
	-NotBefore
	-PublicKey PARENT
	-Signature
	-SignatureAlgorithm
*/
type JSONResultField struct {
	IPAddresses        []net.IP  `json:"ipAddresses"`
	NotAfter           time.Time `json:"notAfter"`
	NotBefore          time.Time `json:"notBefore"`
	PublicKey          string    `json:"publicKey"`
	Signature          []byte    `json:"signature"`
	SignatureAlgorithm int64     `json:"signatureAlgorithm"`
}

/*
	Takes a string in the format of ",key:"value,"
	for the SUBJECT field, the value is in the format of ",key=value,".

	For example:
	input := "ipAddresses:\"192.168.1.1\",notAfter:\"2006-Jan-02\",notBefore:\"2005-Jan-02\",publicKey:\"publicKey\",signature:\"signature\",signatureAlgorithm:\"rsa\"
*/
func ExtractJSONResultField(data string) (result JSONResultField, err error) {

	/*
		REGEX EXPLANATION:
		The first capturing group, (\w+), matches one or more word characters (letters, digits or underscores).
		The second capturing group, :"([^"]+)",, matches a string surrounded by double quotes,
		where [^"]+ matches one or more characters that are not double quotes.

		For example, the string ",key:"value"," will be matched by the regex,
		and the first capturing group will match "key" and the second capturing group will match "value".
	*/
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
		case "signature":
			result.Signature = []byte(value)
		case "signatureAlgorithm":
			result.SignatureAlgorithm = 10
		default:
			return result, fmt.Errorf("unknown field: %s", key)
		}
	}
	return result, nil
}

//The function take a name and a JSONResultField and create
//a serial number converting them into hashes and concatenating the first 20 bytes into serialHash
func SerialNumber(name string, result JSONResultField) ([]byte, error) {
	nameHash := sha256.Sum256([]byte(name))

	pubkey64ToString, err := base64.StdEncoding.DecodeString(result.PublicKey)
	if err != nil {
		fmt.Println(err)
	}
	pubkeyHash := sha512.Sum512(pubkey64ToString)

	notBeforeScaledBuf := new(bytes.Buffer)
	err = binary.Write(notBeforeScaledBuf, binary.BigEndian, result.NotBefore)
	if err != nil {
		return nil, fmt.Errorf("binary.Write failed: %s", err)
	}
	notBeforeHash := sha256.Sum256(notBeforeScaledBuf.Bytes())

	notAfterScaledBuf := new(bytes.Buffer)
	err = binary.Write(notAfterScaledBuf, binary.BigEndian, result.NotAfter)
	if err != nil {
		return nil, fmt.Errorf("binary.Write failed: %s", err)
	}
	notAfterHash := sha256.Sum256(notAfterScaledBuf.Bytes())

	serialHash := sha256.New()
	_, err = serialHash.Write(nameHash[:])
	if err != nil {
		return nil, fmt.Errorf("serialHash.Write of nameHash failed: %s", err)
	}
	_, err = serialHash.Write(pubkeyHash[:])
	if err != nil {
		return nil, fmt.Errorf("serialHash.Write of pubkeyHash failed: %s", err)
	}
	_, err = serialHash.Write(notBeforeHash[:])
	if err != nil {
		return nil, fmt.Errorf("serialHash.Write of notBeforeHash failed: %s", err)
	}
	_, err = serialHash.Write(notAfterHash[:])
	if err != nil {
		return nil, fmt.Errorf("serialHash.Write of notAfterHash failed: %s", err)
	}

	// 19 bytes will be less than 2^159, see https://crypto.stackexchange.com/a/260
	return serialHash.Sum(nil)[0:19], nil
}

func CreateCert(result JSONResultField, name string) ([]byte, error) {
	template := x509.Certificate{
		IPAddresses:        result.IPAddresses,
		NotAfter:           result.NotAfter,
		NotBefore:          result.NotBefore,
		PublicKey:          result.PublicKey,
		Signature:          result.Signature,
		SignatureAlgorithm: x509.SHA512WithRSA,

		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  false,
	}

	template.Subject = pkix.Name{
		CommonName:   name,
		SerialNumber: "DA DEFINIRE",
	}

	serialNumberBytes, err := SerialNumber(name, result)
	if err != nil {
		return nil, err
	}
	template.SerialNumber.SetBytes(serialNumberBytes)

	priv := &splicesign.SpliceSigner{
		PublicKey: template.PublicKey,
		Signature: template.Signature,
	}

	certificate, err := x509.CreateCertificate(rand.Reader, &template, &template, template.PublicKey, priv)
	if err != nil {
		return nil, err
	}

	return certificate, err
}