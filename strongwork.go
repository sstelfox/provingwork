package provingwork

import (
	"bytes"
	"time"

	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"encoding/binary"
)

type StrongWork struct {
	Counter  int64  `json:"counter"`
	Resource []byte `json:"resource"`

	*WorkOptions
}

func NewStrongWork(resource []byte, opts ...*WorkOptions) *StrongWork {
	sw := StrongWork{
		Counter: 0,
		Resource: resource,
	}

	if (len(opts) != 0) {
		sw.WorkOptions = opts[0]
	} else {
		sw.WorkOptions = &WorkOptions{}
	}

	if (sw.Timestamp == nil) {
		t := time.Now()
		sw.Timestamp = &t
	}

	if (sw.BitStrength == 0) {
		sw.BitStrength = DefaultBitStrength
	}

	if (len(sw.Salt) == 0) {
		sw.Salt = make([]byte, DefaultSaltSize)
		rand.Read(sw.Salt)
	}

	return &sw
}

func (sw *StrongWork) Check() bool {
	if (sw.ZeroCount() >= sw.BitStrength) {
		return true
	}
	return false
}

func (sw *StrongWork) ContentHash() []byte {
	var buf bytes.Buffer

	buf.Write(sw.Resource)
	buf.Write(sw.Salt)

	ts := sw.Timestamp.Unix()
	binary.Write(&buf, binary.BigEndian, ts)
	binary.Write(&buf, binary.BigEndian, sw.Counter)

	return buf.Bytes()
}

func (hc *StrongWork) FindProof() {
	for {
		if hc.Check() {
			return
		}
		hc.Counter++
	}
}

func (sw *StrongWork) ZeroCount() int {
	digest := sha256.Sum256(sw.ContentHash())
	digestHex := new(big.Int).SetBytes(digest[:])
	return (256 - digestHex.BitLen())
}
