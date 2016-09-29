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

	json, _ := json.Marshal(hc)
	fmt.Printf("%v\n", string(json))
}
