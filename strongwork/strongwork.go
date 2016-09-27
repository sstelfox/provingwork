package strongwork

import (
	"bytes"
	"time"

	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"encoding/binary"
)

type StrongWork struct {
	Data      []byte    `json:"data"`
	Nonce     int64     `json:"nonce"`
	Salt      []byte    `json:"salt"`
	Timestamp time.Time `json:"timestamp"`
}

func (sw *StrongWork) Check(zeroes int) bool {
	if sw.ZeroCount() >= zeroes {
		return true
	}
	return false
}

func (sw *StrongWork) ContentHash() []byte {
	var buf bytes.Buffer

	buf.Write(sw.Data)
	buf.Write(sw.Salt)

	binary.Write(&buf, binary.BigEndian, sw.Timestamp.Unix())
	binary.Write(&buf, binary.BigEndian, sw.Nonce)

	return buf.Bytes()
}

func (sw *StrongWork) FindProof(zeroes int) {
	if sw.Check(zeroes) {
		return
	}

	sw.Nonce = 0
	sw.Timestamp = time.Now()

	sw.Salt = make([]byte, 16)
	rand.Read(sw.Salt)

	for {
		if sw.Check(zeroes) {
			return
		}
		sw.Nonce++
	}
}

func (sw *StrongWork) ZeroCount() int {
	digest := sha256.Sum256(sw.ContentHash())
	digestHex := new(big.Int).SetBytes(digest[:])
	return (256 - digestHex.BitLen())
}
