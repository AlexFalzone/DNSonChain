package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Struct per rappresentare il JSON RPC request
type JSONRPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

// Struct per rappresentare il JSON RPC response
type JSONRPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	ID      int         `json:"id"`
}

func MakeHTTPRequest(request JSONRPCRequest, hostname string, externalClient string) (interface{}, error) {
	// Crea un client HTTP
	client := &http.Client{}

	// Crea il corpo del request
	requestBody, err := json.Marshal(request)
	if err != nil {
		return JSONRPCResponse{}, err
	}

	// Crea il request HTTP
	req, err := http.NewRequest("POST", externalClient+":8545", bytes.NewBuffer(requestBody))
	if err != nil {
		return JSONRPCResponse{}, err
	}

	// Imposta l'header "Content-Type"
	req.Header.Set("Content-Type", "application/json")

	// Esegue il request
	resp, err := client.Do(req)
	if err != nil {
		return JSONRPCResponse{}, err
	}
	defer resp.Body.Close()

	// Legge il corpo del response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return JSONRPCResponse{}, err
	}

	// Decodifica il corpo del response come JSON
	var response JSONRPCResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return JSONRPCResponse{}, err
	}

	return response.Result, nil
}
