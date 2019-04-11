package key

import "github.com/safex/gosafex/internal/crypto/curve"

func scReduce32(key *Key) {
	curve.ScReduce32(&key.v)
}
