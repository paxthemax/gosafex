package key

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/safex/gosafex/internal/crypto/curve"
	"github.com/safex/gosafex/internal/crypto/hash"
	"github.com/safex/gosafex/internal/random"
)

type point = curve.Point

// TODO: move rand generator somewhere appropriate.
var randomGenerator = random.NewGenerator(false, 0)

// KeyLength is the length of ed25519 keys (in bytes).
const KeyLength = curve.PointLength

// SeedLength is the size of the data sequence used as seed.
// Sequence must be compatible with RFC 8032 (private key).
const SeedLength = 32

// Key is the base key type.
type Key struct {
	v point
}

// Seed is a random sequence used a seed for generating keys.
type Seed = [SeedLength]byte

// New will construct a new key with the given data.
func New(data [KeyLength]byte) *Key {
	key := Key{v: data}
	return &key
}

// NewRandomScalar generates a new random key as a scalar point
// on the ed25519 curve.
// The function will make use of the configured random generator.
func NewRandomScalar() (result *Key) {
	result = new(Key)
	seq := randomGenerator.NewSequence()
	curve.ScReduce(&result.v, seq)
	return
}

// NewFromBytes will create a new Key from data bytes.
// Returns an error if sequence length is invalid.
func NewFromBytes(data []byte) (result *Key, err error) {
	if len(data) != KeyLength {
		return nil, ErrKeyLength
	}
	result = new(Key)
	copy(result.v[:], data)
	return
}

// NewFromString will create a new Key
// from its hexadecimal string representation.
func NewFromString(raw string) (result *Key, err error) {
	buf, err := hex.DecodeString(raw)
	if err != nil {
		return nil, err
	}
	return NewFromBytes(buf)
}

// NewFromSeed constructs a private key from a given seed.
// This function is provided for interoperability
// with RFC 8032. RFC 8032's private keys correspond to seeds in this
// package.
func NewFromSeed(seed Seed) (pub, priv *Key) {
	digest := sha512.Sum512(seed[:])
	digest[0] &= 248
	digest[31] &= 127
	digest[31] |= 64

	A := new(extendedGE)
	hashBuf := new(Key)
	copy(hashBuf.v[:], digest[:])
	curve.GeScalarMultBase(A, &hashBuf.v)

	return pub, priv
}

func (key *Key) toECPoint() (result *extendedGE) {
	result = new(extendedGE)
	p1 := new(projectiveGE)
	p2 := new(completedGE)

	hashedKey := New(hash.Keccak256(key.v[:])) // TODO: prevent copying.
	p1.FromBytes(&hashedKey.v)
	curve.GeMul8(p2, p1)
	p2.ToExtended(result)
	return
}

// ToPublic will return the computed public key of
// a private key.
func (key *Key) ToPublic() (result *Key) {
	point := new(extendedGE)
	curve.GeScalarMultBase(point, &key.v)
	result = new(Key)
	point.ToBytes(&result.v)
	return
}

// ValidPublic returns true if the key is a valid public key.
func (key *Key) ValidPublic() bool {
	return new(extendedGE).FromBytes(&key.v)
}

// ValidPrivate returns true if the key is a valid private key.
func (key *Key) ValidPrivate() bool {
	return curve.ScCheck(&key.v)
}

// String implements the Stringer interface.
// Returns a hex string representation of the key.
func (key Key) String() string {
	return fmt.Sprintf("%x", key.v[:])
}

// ToBytes implements ByteMarshaller.
func (key Key) ToBytes() []byte {
	return key.v[:]
}
