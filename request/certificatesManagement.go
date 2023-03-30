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

	cert.GenerateCert(name)

	certBytes, err := util.ConvertCertToDER("list/certHost.pem")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, _, err = Request(client, choice, name, certBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
}
