package main

import (
	"fmt"
	"encoding/json"

	"github.com/sstelfox/provingwork"
)

func main() {
	hc := provingwork.StrongWork{
		Resource: []byte("Just some test data in the string"),
	}
	hc.FindProof()

	json, _ := json.Marshal(hc)
	fmt.Println(string(json))
}
