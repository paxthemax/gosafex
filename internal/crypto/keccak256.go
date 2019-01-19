package crypto

import "github.com/ebfe/keccak"

// KeccakHashLength is the length of the keccak hash in bytes
const KeccakHashLength = 32

// KeccakHash is a keccak digest
type KeccakHash []byte

// Keccak256 returns a keccak256 digest of a sequence of byte slices
func Keccak256(data ...[]byte) (result KeccakHash) {
	h := keccak.New256()
	for _, b := range data {
		h.Write(b)
	}

	return h.Sum(nil)
}
