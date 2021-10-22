package sample

import "crypto/cipher"

type Sample struct {
	algorithm string
	blockSize int
	cipher    cipher.Block
	mode      string
}

func (s Sample) Algorithm() string {
	return s.algorithm
}

func (s Sample) BlockSize() int {
	return s.blockSize
}

func (s Sample) Cipher() cipher.Block {
	return s.cipher
}

func (s Sample) Mode() string {
	return s.mode
}
