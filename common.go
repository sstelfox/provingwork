package provingwork

import (
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
	if hc.Timestamp == nil {
		t := time.Now()
		hc.Timestamp = &t
	}

	if hc.BitStrength == 0 {
		hc.BitStrength = DefaultBitStrength
	}

	if len(hc.Salt) == 0 {
		hc.Salt = make([]byte, DefaultSaltSize)
		rand.Read(hc.Salt)
	}
}
