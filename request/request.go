package request

import (
	"encoding/hex"
	"errors"
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

func MakeHTTPRequest(url string, hostname string, params interface{}) (string, error) {

	client, err := rpc.Dial(url)
	if err != nil {
		fmt.Println("Could not connect: ", err)
	}

	call := EthCall{
		//for testing
		To:   "0x2ff4dd4c2c511ed082f28ee5210fe7ab6979e812",
		Data: "0x19ff1d21",
	}

	var response JSONRPCResponse
	err = client.Call(&response.Result, "eth_call", call, "latest")
	if err != nil {
		fmt.Println("Error calling eth_call: ", err)
		return "", err
	}

	hexData, ok := response.Result.(string)
	if !ok {
		fmt.Println("Error converting response to string")
		return "", errors.New("error converting response to string")
	}

	//hexData parte dalla posizione 2, quindi elimina i primi 2 caratteri (0x)
	bytes, err := hex.DecodeString(hexData[2:])
	if err != nil {
		fmt.Println("Error decoding hex string")
		return "", err
	}

	str := string(bytes)

	return str, nil
}
