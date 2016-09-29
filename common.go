package provingwork

import (
	"crypto/rand"
	"time"
)

var (
	DefaultBitStrength = 22
	DefaultSaltSize    = 16
)

type WorkOptions struct {
	BitStrength int        `json:"-"`
	Extension   []byte     `json:"extension,omitempty"`
	Salt        []byte     `json:"salt,omitempty"`
	Timestamp   *time.Time `json:"timestamp"`
}

func setDefaultWorkOptions(wo *WorkOptions) {
	if wo.Timestamp == nil {
		t := time.Now()
		wo.Timestamp = &t
	}

	if wo.BitStrength == 0 {
		wo.BitStrength = DefaultBitStrength
	}

	if len(wo.Salt) == 0 {
		wo.Salt = make([]byte, DefaultSaltSize)
		rand.Read(wo.Salt)
	}
}
