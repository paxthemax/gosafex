package curve

import (
	"crypto/sha512"
	"fmt"

	"github.com/safex/gosafex/internal/random"
)

// TODO: move rand generator somewhere appropriate.
var randomGenerator = random.NewGenerator(false, 0)

// KeyLength is the length of ed25519 keys (in bytes).
const KeyLength = 32

// SeedLength is the size of the data sequence used as seed.
// Sequence must be compatible with RFC 8032 (private key).
const SeedLength = 32

// Key is the base key type.
type Key [KeyLength]byte

// Seed is a random sequence used a seed for generating keys.
type Seed = [SeedLength]byte

// NewRandomScalar generates a new random key as a scalar point on the
// ed25519 curve.
// The function will make use of the system random generator.
func NewRandomScalar() (result *Key) {
	result = new(Key)
	seq := randomGenerator.NewSequence()
	ScReduce(result, seq)
	return
}

// NewKeyFromSeed calculates a private key from a given seed.
// This function is provided for interoperability
// with RFC 8032. RFC 8032's private keys correspond to seeds in this
// package.
func NewKeyFromSeed(seed Seed) (pub, priv *Key) {
	digest := sha512.Sum512(seed[:])
	digest[0] &= 248
	digest[31] &= 127
	digest[31] |= 64

	A := new(ExtendedGroupElement)
	hashBuf := new(Key)
	copy(hashBuf[:], digest[:])
	GeScalarMultBase(A, hashBuf)

	return pub, priv
}

// String implements the Stringer interface.
// Returns a hex string representation of the key.
func (key Key) String() string {
	return fmt.Sprintf("%x", key[:])
}
