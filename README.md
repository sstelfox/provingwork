# Proving Work

This is a collection of mechanisms for proving, verifying, and consuming
various [Proof of Work][1] systems.

## HashCash

This follows the official spec for HashCash. There is an oddity that may
surprise users. When base64 encoding the counter value it will always encode a
full 64 bit value where most implementations will only encode the minimum
number of bytes. This won't effect verification or validation, but may effect
other software that attempts to decode the value of the counter.

## StrongWork

This is an implementation of a custom proof of work system. This uses the same
type of data as hashcash, and looks for a hash collision in the same way, but
uses a different encoding mechanism and switches to SHA256 for the digest.

It is primarily designed to exchange it's state through JSON encoding rather
than a colon based mechanism. The resulting value *could* be decoded if so
chosen but the order of fields may change.

## Usage

You can find examples in the examples/ directory. Both versions are pretty
straight forward to use. And breaks down to:

```go
package main

import (
  "encoding/json"
  "fmt"
  "github.com/sstelfox/provingwork"
)

func main() {
	dataToVerify := []byte("Useful Identifier")

	hc := provingwork.NewHashCash(dataToVerify)
	hc.FindProof()
	fmt.Println(hc)

	sw := provingwork.NewStrongWork(dataToVerify)
	sw.FindProof()
	json, _ := json.Marshal(sw)
	fmt.Println(string(json))
}
```

### Advanced Options

Both mechanisms support additional options and customizations. These are all
passed as a second argument to the constructor function. The following shows
how to create the additional options with the default values:

```go
wo := WorkOptions{
	BitStrength: 22,
	Extension:   []byte{},
	Salt:        []byte{},    // 16 random bytes
	Timestamp:   *time.Time,  // Initializes to time.Now()
}
```

As the comments indicate if no Salt is provided it gets initialized to 16
random bytes, and the Timestamp field is the time at initialization. Extension
is officially unused in the spec but can be used to hold additional data at the
user's discretion, without effecting the ability to verify standard usage.

[1]: https://en.wikipedia.org/wiki/Proof-of-work_system
[2]: https://en.wikipedia.org/wiki/Hashcash
