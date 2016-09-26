package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"time"

	"crypto/sha256"
	"math/big"

	"encoding/binary"
	"encoding/json"
)

var (
	BuiltAt string
	Version = "Unknown"

	printVersion = flag.Bool("v", false, "Display the version and exit")
)

type HashCash struct {
	Data      []byte `json:"data"`
	Timestamp int64  `json:"timestamp"`
	Nonce     int64  `json:"nonce"`
}

func (hc *HashCash) Check(zeroes int) bool {
	digest := sha256.Sum256(hc.ContentHash())

	digestHex := new(big.Int).SetBytes(digest[:])
	if digestHex.BitLen() == 256-zeroes {
		return true
	}

	return false
}

func (hc *HashCash) ContentHash() []byte {
	var buf bytes.Buffer

	buf.Write(hc.Data)
	binary.Write(&buf, binary.BigEndian, hc.Timestamp)
	binary.Write(&buf, binary.BigEndian, hc.Nonce)

	return buf.Bytes()
}

func (hc *HashCash) FindProof(zeroes int) {
	if hc.Check(zeroes) {
		return
	}

	hc.Nonce = 0
	hc.Timestamp = time.Now().Unix()

	for {
		if hc.Check(zeroes) {
			return
		}

		hc.Nonce++
	}
}

func main() {
	flag.Parse()

	if *printVersion {
		fmt.Printf("HashCash version %s built on %s\n", Version, BuiltAt)
		os.Exit(0)
	}

	hc := HashCash{
		Data: []byte("Just some test data in the string"),
	}
	hc.FindProof(20)

	json, _ := json.Marshal(hc)
	fmt.Println(string(json))
}
