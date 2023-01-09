package main

import (
	"fmt"
	"net"
	"os"

	"example.com/packages/dns"
	"example.com/packages/request"
)

func main() {

	requestJSON := request.JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{"latest", true},
		ID:      1,
	}

	// Apre una connessione sulla porta 53
	conn, err := net.ListenPacket("udp", ":53")
	if err != nil {
		fmt.Println("Errore durante l'apertura della connessione:", err)
		os.Exit(1)
	}
	fmt.Print("ascolto sulla porta 53")
	defer conn.Close()

	for {
		// Ascolta le richieste DNS in arrivo
		dns.HandleRequest(conn)

		request.MakeHTTPRequest(requestJSON, "http://localhost")
	}
}
