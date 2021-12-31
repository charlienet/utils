package hash

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"github.com/cespare/xxhash/v2"
	"github.com/tjfoc/gmsm/sm3"
)

func Sha1(msg []byte) []byte { return sum(sha1.New, msg) }

func Sha224(msg []byte) []byte { return sum(sha256.New224, msg) }

func Sha256(msg []byte) []byte { return sum(sha256.New, msg) }

func Sha384(msg []byte) []byte { return sum(sha512.New384, msg) }

func Sha512(msg []byte) []byte { return sum(sha512.New, msg) }

func Sm3(msg []byte) []byte { return sum(sm3.New, msg) }

func XXhash(msg []byte) []byte {
	d := xxhash.New()
	_, _ = d.Write(msg)
	return d.Sum(nil)
}

func XXHashUint64(msg []byte) uint64 {
	d := xxhash.New()
	_, _ = d.Write(msg)
	return d.Sum64()
}

func sum(f func() hash.Hash, msg []byte) []byte {
	h := f()

	_, _ = h.Write(msg)
	return h.Sum(nil)
}
