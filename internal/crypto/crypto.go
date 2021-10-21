package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

const (
	CFB = "cfb"
)

func Encrypt(plaintext []byte, mode string, block cipher.Block, blockSize int) ([]byte, error) {
	encrypted := make([]byte, len(plaintext))
	switch mode {
	case CFB:
		encrypted = encryptCFB(plaintext, block, blockSize)
	default:
		return nil, errors.New("unknown mode")
	}
	return encrypted, nil
}

func Decrypt(plaintext []byte, mode string, block cipher.Block, blockSize int) ([]byte, error) {
	encrypted := make([]byte, len(plaintext))
	switch mode {
	case CFB:
		encrypted = decryptCFB(plaintext, block, blockSize)
	default:
		return nil, errors.New("unknown mode")
	}
	return encrypted, nil
}

func encryptCFB(plaintext []byte, block cipher.Block, blockSize int) []byte {
	ciphertext := make([]byte, blockSize+len(plaintext))
	iv := ciphertext[:blockSize]
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext
}

func decryptCFB(plaintext []byte, block cipher.Block, blockSize int) []byte {
	ciphertext := make([]byte, blockSize+len(plaintext))
	iv := ciphertext[:blockSize]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext
}
