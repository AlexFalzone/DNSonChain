package cert

import (
	"io/ioutil"
	"log"
	"os"
)

func writeChain() {
	chainOut, err := os.Create("chain.pem")
	if err != nil {
		log.Fatalf("Failed to create chain.pem for writing: %s", err)
	}

	defer chainOut.Close()

	leafCert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		log.Fatalf("Failed to read cert.pem: %v", err)
	}

	_, err = chainOut.Write(leafCert)
	if err != nil {
		log.Fatalf("Failed to write end-entity cert to chain.pem: %v", err)
	}

	caChainOut, err := os.Create("caChain.pem")
	if err != nil {
		log.Fatalf("Failed to open caChain.pem for writing: %v", err)
	}

	defer caChainOut.Close()

	//Riga vuota per separare i certificati
	_, err = chainOut.WriteString("\n\n")
	if err != nil {
		log.Fatalf("Failed to write CA cert padding to chain.pem: %v", err)
	}

	parentToRead := "caCert.pem"

	caCert, err := ioutil.ReadFile(parentToRead)
	if err != nil {
		log.Fatalf("Failed to read %s: %v", parentToRead, err)
	}

	_, err = chainOut.Write(caCert)
	if err != nil {
		log.Fatalf("Failed to write CA cert to chain.pem: %v", err)
	}

	_, err = caChainOut.Write(caCert)
	if err != nil {
		log.Fatalf("Failed to write CA cert to caChain.pem: %v", err)
	}

}
