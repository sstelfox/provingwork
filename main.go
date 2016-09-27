package main

import (
	"flag"
	"fmt"
	"os"
	"encoding/json"

	"github.com/sstelfox/go-provingwork/strongwork"
)

var (
	BuiltAt string
	Version = "Unknown"

	printVersion = flag.Bool("v", false, "Display the version and exit")
)

func main() {
	flag.Parse()

	if *printVersion {
		fmt.Printf("HashCash version %s built on %s\n", Version, BuiltAt)
		os.Exit(0)
	}

	hc := strongwork.StrongWork{
		Data: []byte("Just some test data in the string"),
	}

	hc.FindProof(20)

	json, _ := json.Marshal(hc)
	fmt.Println(string(json))
}
