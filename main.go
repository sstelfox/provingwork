package main

import (
	"fmt"

	"crypto/sha256"
	"math/big"
	"strconv"
	"time"
)

var (
	BuiltAt string
	Version = "Unknown"
)

type Message struct {
	Message   string
	Timestamp int64
	Nonce     int
}

func Prove(msg *Message, zeroes int) {
	nonce := 0
	for {
		nonce++
		digest := sha256.New()
		digest.Write([]byte(msg.Message))
		digest.Write([]byte(strconv.Itoa(int(msg.Timestamp))))
		digest.Write([]byte(strconv.Itoa(nonce)))
		digestResult := digest.Sum(nil)

		digestHex := new(big.Int).SetBytes(digestResult)
		if digestHex.BitLen() == 256-zeroes {
			msg.Nonce = nonce
			return
		}
	}
}

//func Check(message string, zeroes int, nonce int) bool {
//	digest := sha256.Sum256([]byte(message + strconv.Itoa(nonce)))
//	for i := zeroes; i >= 0; i-- {
//		if digest[i] != 0 {
//			return false
//		}
//	}
//	return true
//}

func main() {
	fmt.Printf("Running version %s built on %s\n", Version, BuiltAt)
	m := Message{
		Message:   "testing stuff",
		Timestamp: time.Now().Unix(),
	}

	Prove(&m, 20)

	fmt.Println("None result:", m.Nonce)
}
