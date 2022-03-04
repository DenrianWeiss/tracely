package model

type Request struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Method  string `json:"method"`
	Params  []interface{} `json:"params"`
}

func NewRequest(txId string) Request {
	p := Request{
		Jsonrpc: "2.0",
		ID:      0,
		Method:  "debug_traceTransaction",
		Params:  nil,
	}
	p.Params = make([]interface{}, 0)
	p.Params = append(p.Params, txId)
	p.Params = append(p.Params, map[string]bool{
		"disableStorage": true,
	})
	return p
}
