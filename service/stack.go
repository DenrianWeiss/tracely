package service

import (
	"encoding/hex"
	"strings"
)

func AccessStack(s [][]byte, index int) []byte {
	length := len(s)
	result := s[length-1-index]
	return result
}

func GenerateStack(s []interface{}) [][]byte {
	result := make([][]byte, 0)
	for _, v := range s {
		l := v.(string)
		l = strings.TrimPrefix(l, "0x")
		// Pad stack elements to 32 bytes
		if len(l) < 64 {
			l = strings.Repeat("0", 64-len(l)) + l
		}
		lBytes, _ := hex.DecodeString(l)
		result = append(result, lBytes)
	}
	return result
}
