package main

import (
	"fmt"
	"net"
	"strconv"

	"test/dns"
	"test/request"
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

	fmt.Println("1. Infura")
	fmt.Println("2. Localhost")
	fmt.Println("(Ctrl+d). Exit")

	var url string
	var choice string
	fmt.Scan(&choice)
	choiceInt, _ := strconv.Atoi(choice)

	// Crea un nuovo conn UDP sulla porta 53
	conn, err := net.ListenPacket("udp", ":53")
	if err != nil {
		fmt.Println("Error listening on DNS port:", err)
		return
	}
	defer conn.Close()

	for {
		check := int8(0)
		hostname, check := dns.HandleRequest(conn)

		if check == 1 {
			fmt.Println("Checking for: ", hostname)
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
		}
	}
}
