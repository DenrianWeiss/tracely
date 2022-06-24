package main

import (
	"flag"
	"fmt"
	"github.com/DenrianWeiss/tracely/service"
	"log"
)

func main() {
	r := flag.String("r", "http://127.0.0.1:8545", "Specify RPC http address")
	t := flag.String("t", "", "TxID")
	sep := flag.String("sep", "", "Separate payload")
	flag.Parse()
	if sep != nil && *sep != "" {
		service.SeperatePayload(sep)
		return
	}
	log.Println("query")
	if t == nil || r == nil || *t == "" {
		fmt.Printf("Usage: tracely -r rpc -t txID\n" +
			"example: tracely -r http://127.0.0.1:8545 -t 0x7b3cb706b2da1356daaad40860948800870b2df1201a42be9cb274b2eca61053\n" +
			"Other commands:\n" +
			"    -sep [payload]: Seperate payload\n" +
			"    -byte [signature]: query 4byte.directory [todo]\n")
		return
	}
	log.Printf("rpc: %s, txid %s", *r, *t)
	res := service.GetTxResult(*r, *t)
	for _, b := range res {
		service.PrintStep(b)
	}
}
