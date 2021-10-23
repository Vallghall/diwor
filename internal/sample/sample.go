package sample

import (
	"crypto/aes"

	"github.com/bi-zone/ruwireguard-go/crypto/gosthopper"
	"gitlab.com/Valghall/diwor/internal/crypto"
)

const (
	AES              = "aes"
	Grasshopper      = "grasshopper"
	defaultSampleKey = "abcdabcdacbdacbdabcdabcdacbdacbd"
)

func (s Sample) Encrypt(dst, src []byte) {
	encrypted, _ := crypto.Encrypt(
		src,
		s.Mode(),
		s.Cipher(),
		s.BlockSize(),
	)
	copy(dst, encrypted)
}

func (s Sample) Decrypt(dst, src []byte) {
	decrypted, _ := crypto.Decrypt(
		src,
		s.Mode(),
		s.Cipher(),
		s.BlockSize(),
	)
	copy(dst, decrypted)
}

func CreateSample(algorithm, mode string) *Sample {
	smpl := new(Sample)
	key := []byte(defaultSampleKey)
	switch algorithm {
	case AES:
		block, _ := aes.NewCipher(key)
		smpl = &Sample{
			algorithm: AES,
			blockSize: aes.BlockSize,
			cipher:    block,
			mode:      mode,
		}
	case Grasshopper:
		block, _ := gosthopper.NewCipher(key)
		smpl = &Sample{
			algorithm: Grasshopper,
			blockSize: gosthopper.BlockSize,
			cipher:    block,
			mode:      mode,
		}
	}
	return smpl
}
