package service

import "encoding/hex"

func AccessStack(s [][]byte, index int) []byte{
	length := len(s)
	result := s[length - 1 - index]
	return result
}

func GenerateStack(s []interface{}) [][]byte {
	result := make([][]byte, 0)
	for _, v := range s {
		l := v.(string)
		lBytes, _ := hex.DecodeString(l)
		result = append(result, lBytes)
	}
	return result
}
