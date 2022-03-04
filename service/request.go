package service

import (
	"bytes"
	"encoding/json"
	"github.com/DenrianWeiss/tracely/model"
	"github.com/bcicen/jstream"
	"log"
	"net/http"
)

func GetTxResult(rpc, txid string) []model.TraceStep {
	result := make([]model.TraceStep, 0)
	req := model.NewRequest(txid)
	reqJson, _ := json.Marshal(req)
	reqStream := bytes.NewReader(reqJson)
	post, err := http.Post(rpc, "application/json", reqStream)
	if err != nil {
		return result
	}
	if post.StatusCode != http.StatusOK {
		log.Println("Failed to fetch")
	}
	resultStream := post.Body
	decoder := jstream.NewDecoder(resultStream, 3)
	for mv := range decoder.Stream() {
		v := mv.Value.(map[string]interface{})
		OpCode := v["op"].(string)
		if OpCodeFocus[OpCode] {
			r := model.TraceStep{}
			r.Depth = int(v["depth"].(float64))
			r.Op = OpCode
			r.Gas = int(v["gas"].(float64))
			r.Stack = v["stack"].([]interface{})
			r.Memory = v["memory"].([]interface{})
			result = append(result, r)
		}

	}
	return result
}
