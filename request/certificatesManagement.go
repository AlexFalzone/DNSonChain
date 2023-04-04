package request

import (
	"fmt"
	"test/cert"
	"test/util"
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

	certPEM, err := util.ConvertCertToPEM(certDER)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, _, err = Request(client, choice, name, certPEM)
	if err != nil {
		fmt.Println(err)
		return
	}
}
