package main

import (
  "bytes"
  "fmt"
  "time"

  "crypto/sha256"
  "encoding/binary"
  "math/big"
)

var (
  BuiltAt string
  Version = "Unknown"
)

type Message struct {
  Message   string
  Timestamp int64
  Nonce     int64
}

func (m *Message) ContentHash() []byte {
  var buf bytes.Buffer

  buf.WriteString(m.Message)
  binary.Write(&buf, binary.BigEndian, m.Timestamp)
  binary.Write(&buf, binary.BigEndian, m.Nonce)

  return buf.Bytes()
}

func (m *Message) FindProof(zeroes int) {
  // TODO: If Nonce is already set, check to see if it's valid as we may be
  // able to avoid the work.
  m.Nonce = 0

  for {
    digest := sha256.Sum256(m.ContentHash())
    digestHex := new(big.Int).SetBytes(digest[:])
    if digestHex.BitLen() == 256-zeroes {
      return
    }

    m.Nonce++
  }
}

//func Check(message string, zeroes int, nonce int) bool {
//  digest := sha256.Sum256([]byte(message + strconv.Itoa(nonce)))
//  for i := zeroes; i >= 0; i-- {
//    if digest[i] != 0 {
//      return false
//    }
//  }
//  return true
//}

func Verify(msg *Message) int {
  return 1
}

func main() {
  fmt.Printf("Running version %s built on %s\n", Version, BuiltAt)
  m := Message{
    Message:   "testing stuff",
    Timestamp: time.Now().Unix(),
    Nonce: 5,
  }

  m.FindProof(20)

  fmt.Printf("%x\n", sha256.Sum256(m.ContentHash()))
  fmt.Printf("%v\n", m.Nonce)
}
