package main

import (
	"fmt"
	"net"
	"strconv"

	"test/cert"
	"test/dns"
	"test/request"
	"test/util"
)

/*
TO-DO:
-sistemare la funzione HandleRequest in maniera tale che non sia necessario stoppare il servizio DNS
*/

func main() {
	JSONRPCRequest := request.JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_call",
		Params:  []interface{}{map[string]interface{}{"to": "0x2ff4dd4c2c511ed082f28ee5210fe7ab6979e812", "data": "0x19ff1d21"}, "latest"},
		ID:      1,
	}

	// Crea un nuovo conn UDP sulla porta 53
	conn, err := net.ListenPacket("udp", ":53")
	if err != nil {
		fmt.Println("Error listening on DNS port:", err)
		return
	}
	defer conn.Close()

	fmt.Println("1. Infura")
	fmt.Println("2. Localhost")
	fmt.Println("Create certificate")

	var url string
	var choice string
	fmt.Scan(&choice)
	choiceInt, _ := strconv.Atoi(choice)

	input := "ipAddresses:\"192.168.1.1\"," +
		"notAfter:\"2006-Jan-02\"," +
		"notBefore:\"2005-Jan-02\"," +
		"publicKey:\"publicKey\"," +
		"signature:\"MDkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDIgAClxvozFkncQY2R7EFcohBkbGPaq6H2WqBPdDI+ZzC+Ps=\"," +
		"signatureAlgorithm:\"10\""

	for {
		check := int8(0)
		var response, hostname string

		hostname, check = dns.HandleRequest(conn)

		if choiceInt == 3 { //create certificate
			_, _, _ = util.MenuCertificate()
			name := "test123.com"
			_, _ = cert.GenerateCert(name)

		} else if choiceInt == 1 || choiceInt == 2 {

			if check == 1 { //check == 1 -> DNS request of type A with .somet site
				if choiceInt == 1 { //Infura
					url = "https://goerli.infura.io/v3/d777809793694d9dacf5e1f94bfec65a"
				} else if choiceInt == 2 { //Localhost
					url = "http://localhost:8545"
				}

				go func() {
					response, err := request.MakeHTTPRequest(url, hostname, JSONRPCRequest.Params)
					if err != nil {
						fmt.Println(response, err)
					}
				}()

				fmt.Println(response)
				var result cert.JSONResultField
				result, err = cert.ExtractJSONResultField(input)
				if err != nil {
					fmt.Println(err)
					return
				}
				derbytes, err := cert.CreateCert(result, hostname)
				if err != nil {
					fmt.Println(err)
					return
				}

				fmt.Println("derbytes: ", derbytes)

				//INJECT CERTIFICATE
			}
		}
	}
}
