package service

import (
	hash2 "hash"
)

const (
	Streebog256 = "Streebog-256"
	Streebog512 = "Streebog-512"
	SHA224      = "SHA-224"
	SHA256      = "SHA-256"
	SHA384      = "SHA-384"
	SHA512      = "SHA-512"
	SHA3_224    = "SHA3-224"
	SHA3_256    = "SHA3-256"
	SHA3_384    = "SHA3-384"
	SHA3_512    = "SHA3-512"

	RIPEMD128 = "RIPEMD-128"
	RIPEMD160 = "RIPEMD-160"
	RIPEMD256 = "RIPEMD-256"
	RIPEMD320 = "RIPEMD-320"
	MD5       = "MD5"
	BCRYPT    = "bcrypt"
)

func hash(h hash2.Hash, text []byte) (res []byte) {
	h.Write(text)
	res = h.Sum(nil)
	return
}
