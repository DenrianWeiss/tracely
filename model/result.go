package model

type TraceStep struct {
	Depth   int           `json:"depth"`
	Gas     int           `json:"gas"`
	GasCost int           `json:"gasCost"`
	Memory  []interface{} `json:"memory"`
	Op      string        `json:"op"`
	Pc      int           `json:"pc"`
	Stack   []interface{} `json:"stack"`
}
