package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"time"

	"crypto/rand"
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
	Data      []byte    `json:"data"`
	Nonce     int64     `json:"nonce"`
	Salt      []byte    `json:"salt"`
	Timestamp time.Time `json:"timestamp"`
}

func (hc *HashCash) Check(zeroes int) bool {
	if hc.ZeroCount() >= zeroes {
		return true
	}

	return false
}

func (hc *HashCash) ContentHash() []byte {
	var buf bytes.Buffer

	buf.Write(hc.Data)
	binary.Write(&buf, binary.BigEndian, hc.Timestamp.Unix())
	buf.Write(hc.Salt)
	binary.Write(&buf, binary.BigEndian, hc.Nonce)

	return buf.Bytes()
}

func (hc *HashCash) FindProof(zeroes int) {
	if hc.Check(zeroes) {
		return
	}

	hc.Nonce = 0
	hc.Timestamp = time.Now()

	hc.Salt = make([]byte, 16)
	rand.Read(hc.Salt)

	for {
		if hc.Check(zeroes) {
			return
		}

		hc.Nonce++
	}
}

func (hc *HashCash) ZeroCount() int {
	digest := sha256.Sum256(hc.ContentHash())
	digestHex := new(big.Int).SetBytes(digest[:])
	return (256 - digestHex.BitLen())
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
