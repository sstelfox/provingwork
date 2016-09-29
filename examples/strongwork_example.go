package main

import (
	"encoding/json"
	"fmt"

	"github.com/sstelfox/provingwork"
)

func main() {
	hc := provingwork.NewStrongWork([]byte("Just some test data in the string"))
	hc.FindProof()

	json, _ := json.Marshal(hc)
	fmt.Println(string(json))
}
