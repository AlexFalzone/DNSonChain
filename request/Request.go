package request

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

type EthCall struct {
	To   string `json:"to"`
	Data string `json:"data"`
}

// Struct per rappresentare il JSON RPC request
type JSONRPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

// Struct per rappresentare il JSON RPC response
type JSONRPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	ID      int         `json:"id"`
}

func MakeHTTPRequest(url string, hostname string, params interface{}) (interface{}, error) {

	client, err := rpc.Dial(url)
	if err != nil {
		fmt.Println("Could not connect: ", err)
	}

	call := EthCall{
		To:   "0x2ff4dd4c2c511ed082f28ee5210fe7ab6979e812",
		Data: "0x19ff1d21",
	}

	var response JSONRPCResponse
	err = client.Call(&response.Result, "eth_call", call, "latest")
	if err != nil {
		fmt.Println("Cannot get the call:", err)
		return response.Result, err
	}

	fmt.Println("Response from Infura: ", response.Result)

	return response.Result, nil
}
