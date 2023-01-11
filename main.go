package main

import (
	"fmt"
	"net"

	"test/dns"
	"test/request"
)

func main() {
	JSONRPCRequest := request.JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "web3_clientVersion",
		Params:  []interface{}{},
		ID:      1,
	}

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
			response, err := request.MakeHTTPRequest(JSONRPCRequest, hostname, "http://localhost")

			if err == nil {
				fmt.Println(response)
			}
		}
	}
}
