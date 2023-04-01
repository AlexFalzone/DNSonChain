package request

import (
	"encoding/hex"
	"fmt"
	"test/cert"
)

func CertificatesManagement(choice int, url string) {
	var name string

	switch choice {
	case 1, 3: // Generate or Update certificate
		fmt.Println("Insert the name of the certificate to update:")
	case 2: // Revoke certificate
		fmt.Println("Insert the name of the certificate to revoke:")
	}

	fmt.Scan(&name)

	client, err := DialClient(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	certDER, _, err := cert.GenerateCert(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	certHex := "0x" + hex.EncodeToString(certDER)

	_, _, err = Request(client, choice, name, certHex)
	if err != nil {
		fmt.Println(err)
		return
	}
}
