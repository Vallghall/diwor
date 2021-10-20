package experiment

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/bi-zone/ruwireguard-go/crypto/gosthopper"
)

const (
	AES              = "aes"
	Grasshopper      = "grasshopper"
	defaultSampleKey = "abcdabcdacbdacbdabcdabcdacbdacbd"
)

type sample struct {
	algorithm string
	blockSize int
	cipher    cipher.Block
}

func (s sample) Algorithm() string {
	return s.algorithm
}

func (s sample) BlockSize() int {
	return s.blockSize
}

func (s sample) Cipher() cipher.Block {
	return s.cipher
}

func createSample(algorithm string) *sample {
	smpl := new(sample)
	key := []byte(defaultSampleKey)
	switch algorithm {
	case AES:
		block, _ := aes.NewCipher(key)
		smpl = &sample{
			algorithm: AES,
			blockSize: aes.BlockSize,
			cipher:    block,
		}
	case Grasshopper:
		block, _ := gosthopper.NewCipher(key)
		smpl = &sample{
			algorithm: Grasshopper,
			blockSize: gosthopper.BlockSize,
			cipher:    block,
		}
	}
	return smpl
}
