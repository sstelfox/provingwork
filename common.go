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
