package main

import (
	"crypto/sha1"
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
	fmt.Printf("%x\n", sha1.Sum([]byte(hc.String())))
}
