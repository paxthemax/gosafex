package account

import (
	"github.com/safex/gosafex/internal/crypto"
)

// KeyPair is a public/private keypair.
type KeyPair struct {
	pub  PublicKey
	priv PrivateKey
}

// KeySet is a complete set of spend and view keypairs.
type KeySet struct {
	view  KeyPair
	spend KeyPair
}

// NewKeyPair constructs a new keypair with the given keys.
func NewKeyPair(pub PublicKey, priv PrivateKey) *KeyPair {
	return &KeyPair{
		pub:  pub,
		priv: priv,
	}
}

// GenerateKeyPair will create a new keypair.
// The implementation relies on system entropy from '/dev/urandom' by default.
func GenerateKeyPair() (*KeyPair, error) {
	pubKey, privKey, err := crypto.GenerateKeys()
	return NewKeyPair(pubKey, privKey), err
}

// KeyPairFromSeed will create a new keypair from a given seed.
func KeyPairFromSeed(seed Seed) *KeyPair {
	pubKey, privKey := crypto.KeysFromSeed(seed)
	return NewKeyPair(pubKey, privKey)
}

// NewKeySet constructs a new keyset with the given keys.
func NewKeySet(view, spend *KeyPair) *KeySet {
	return &KeySet{
		view:  *view,
		spend: *spend,
	}
}

// GenerateKeySet will generate new view and spend keypairs.
//
// NOTE: to preserve the same seed - we generate the private view key from the
// Keccak256 hash of the private spend key.
func GenerateKeySet() (result *KeySet, err error) {
	spend, err := GenerateKeyPair()
	if err != nil {
		return nil, err
	}
	viewSeed := Seed(crypto.Keccak256(spend.priv))
	view := KeyPairFromSeed(viewSeed)
	result = NewKeySet(view, spend)
	return
}

// KeySetFromSeed will generate a key set from a given seed.
//
// NOTE: to preserve the same seed - we generate the private view key from the
// Keccak256 hash of the private spend key.
func KeySetFromSeed(seed Seed) *KeySet {
	spend := KeyPairFromSeed(seed)
	viewSeed := Seed(crypto.Keccak256(spend.priv))
	view := KeyPairFromSeed(viewSeed)
	return NewKeySet(view, spend)
}
