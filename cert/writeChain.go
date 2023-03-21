package cert

import (
	"io/ioutil"
	"log"
	"os"
)

func writeChain() {
	chainOut, err := os.Create("list/chain.pem")
	if err != nil {
		log.Fatalf("Failed to create chain.pem for writing: %s", err)
	}

	defer chainOut.Close()

	leafCert, err := ioutil.ReadFile("list/certHost.pem")
	if err != nil {
		log.Fatalf("Failed to read certHost.pem: %v", err)
	}

	_, err = chainOut.Write(leafCert)
	if err != nil {
		log.Fatalf("Failed to write end-entity cert to chain.pem: %v", err)
	}

	//Riga vuota per separare i certificati
	_, err = chainOut.WriteString("\n\n")
	if err != nil {
		log.Fatalf("Failed to write CA cert padding to chain.pem: %v", err)
	}

	parentToRead := "certIntermediate.pem"

	caCert, err := ioutil.ReadFile("list/" + parentToRead)
	if err != nil {
		log.Fatalf("Failed to read %s: %v", parentToRead, err)
	}

	_, err = chainOut.Write(caCert)
	if err != nil {
		log.Fatalf("Failed to write CA cert to chain.pem: %v", err)
	}
}
