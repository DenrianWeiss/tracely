package model

import "github.com/DenrianWeiss/tracely/constants"

type TraceStep struct {
	Depth   int                  `json:"depth"`
	Gas     int                  `json:"gas"`
	GasCost int                  `json:"gasCost"`
	Memory  []interface{}        `json:"memory"`
	Op      constants.OpCodeEnum `json:"op"`
	Pc      int                  `json:"pc"`
	Stack   []interface{}        `json:"stack"`
}
