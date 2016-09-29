package main

import (
	"encoding/json"
	"fmt"

	"github.com/sstelfox/provingwork"
)

func main() {
	hc := provingwork.NewHashCash([]byte("testing"))
	hc.FindProof()

	fmt.Printf("%v\n", hc)

	json, err := json.Marshal(hc)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", string(json))
	}
}
