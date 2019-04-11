package key

import (
	"bytes"
	"encoding/binary"

	"github.com/safex/gosafex/internal/crypto/curve"
	"github.com/safex/gosafex/internal/crypto/hash"
)

func hashToScalar(data ...[]byte) (result *Key) {
	result = new(Key)
	buf := hash.Keccak256(data...)
	copy(result.v[:], buf[:32])
	curve.ScReduce32(&result.v)
	return
}

func idxToVarint(idx uint64) (result []byte) {
	result = make([]byte, 12, 12) // TODO: understand why 12 bytes.
	length := binary.PutUvarint(result, idx)
	return result[:length]
}

func derivationToScalar(outIdx uint64, der *Key) (result *Key) {
	buf := bytes.NewBuffer(der.v[:])
	buf.Write(idxToVarint(outIdx))
	return hashToScalar(buf.Bytes())
}

func hashToEC(data []byte) (result *extendedGE) {
	result = new(extendedGE)
	p1 := new(projectiveGE)
	p2 := new(completedGE)
	keyBuf := new(Key)

	copy(keyBuf.v[:], data[:KeyLength]) // TODO: remove key copying.
	p1.FromBytes(&keyBuf.v)
	curve.GeMul8(p2, p1)

	p2.ToExtended(result)
	return
}

// DeriveKey derives a new private key derivation
// from a given public key and a secret (private key).
// Returns ErrInvalidPrivKey if the given private key (secret)
// is invalid.
// Returns ErrInvalidPubKey if the given public key is invalid.
func DeriveKey(pub, priv *Key) (result *Key, err error) {
	point := new(extendedGE)
	point2 := new(projectiveGE)
	point3 := new(completedGE)
	keyBuf := new(Key)

	if ok := curve.ScCheck(&priv.v); !ok {
		return nil, ErrInvalidPrivKey
	}
	copy(keyBuf.v[:], pub.v[:]) // TODO: remove key copying.
	if ok := point.FromBytes(&keyBuf.v); !ok {
		return nil, ErrInvalidPubKey
	}

	copy(keyBuf.v[:], priv.v[:])
	curve.GeScalarMult(point2, &keyBuf.v, point)
	curve.GeMul8(point3, point2)
	point3.ToProjective(point2)

	point2.ToBytes(&keyBuf.v)
	return keyBuf, nil
}

// DerivationToPrivateKey will compute an ephemereal private key based on
// the key derivation, the given output index and the given private spend.
func DerivationToPrivateKey(idx uint64, base, der *Key) (result *Key) {
	keyBuf := new(Key)
	scalar := derivationToScalar(idx, der)
	copy(keyBuf.v[:], base.v[:])
	curve.ScAdd(&keyBuf.v, &keyBuf.v, &scalar.v)
	return keyBuf
}

// DerivationToPublicKey TODO: comment function
func DerivationToPublicKey(idx uint64, der, base *Key) (result *Key, err error) {
	point1 := new(extendedGE)
	point2 := new(extendedGE)
	point3 := new(cachedGE)
	point4 := new(completedGE)
	point5 := new(projectiveGE)
	keyBuf := new(Key)

	copy(keyBuf.v[:], base.v[:]) // TODO: prevent copying.
	if !point1.FromBytes(&keyBuf.v) {
		return nil, ErrInvalidPubKey
	}

	scalar := derivationToScalar(idx, der)
	curve.GeScalarMultBase(point2, &scalar.v)
	point2.ToCached(point3)
	curve.GeAdd(point4, point1, point3)
	point4.ToProjective(point5)
	point5.ToBytes(&keyBuf.v)
	return keyBuf, nil
}

// GenerateKeyImage will return a key image generated from
// a public/private key pair.
func GenerateKeyImage(pub, priv *Key) (result *Key) {
	result = new(Key)
	proj := new(projectiveGE)

	ext := pub.toECPoint()
	curve.GeScalarMult(proj, &pub.v, ext)
	proj.ToBytes(&result.v)
	return
}
