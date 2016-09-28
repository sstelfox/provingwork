package main

import (
	"fmt"

	"github.com/sstelfox/provingwork"
)

func main() {
	hc := provingwork.NewHashCash([]byte("testing"))
	hc.FindProof()

	fmt.Printf("%v\n", hc)
}
