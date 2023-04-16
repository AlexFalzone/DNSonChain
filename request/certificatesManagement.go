package request

import (
	"fmt"
	"test/cert"
	"test/util"
)

// CertificatesManagement is a helper function for managing certificates.
// It takes an integer representing the desired action (1: generate, 2: revoke, 3: update),
// and the Ethereum node URL as input. Based on the action, it calls the appropriate function
// in the domain registry contract.
//
// Parameters:
//   - choice: Integer representing the desired action (1: generate, 2: revoke, 3: update).
//   - url: The Ethereum node URL as a string.
func CertificatesManagement(choice int, url string) {
	var name string
	var expirationDate string

	switch choice {
	case 1, 3: // Generate or Update certificate
		fmt.Println("Insert the name of the certificate to update:")
		fmt.Scan(&name)
		fmt.Println("Insert the expiration date in the format yyyy-mm-dd:")
		fmt.Scan(&expirationDate)
	case 2: // Revoke certificate
		fmt.Println("Insert the name of the certificate to revoke:")
		fmt.Scan(&name)
	}

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

	expirationTimestamp, err := util.ParseExpirationDate(expirationDate)
	if err != nil {
		fmt.Println("Error parsing expiration date:", err)
		return
	}

	_, _, err = Request(client, choice, name, certPEM, expirationTimestamp)
	if err != nil {
		fmt.Println(err)
		return
	}
}
