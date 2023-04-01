package util

import (
	"fmt"
	"log"
	"os"
)

// CreateandWriteTemp creates a temporary file and writes the certificate to it.
func CreateandWriteTemp(certString string) error {
	certOut, err := os.Create("temp.pem")
	if err != nil {
		log.Fatal(err)
	}
	defer certOut.Close()

	_, err = certOut.WriteString(certString)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// RemoveTempFile removes the temporary file.
func RemoveTempFile(errPKI error, errMozilla error, err error) {
	if errPKI == nil && errMozilla == nil {
		err := os.Remove("temp.pem")
		if err != nil {
			fmt.Println(err)
		}
	}
}
