package service

import (
	"crypto/cipher"
	"crypto/rand"
	"io"
)

const (
	Grasshopper = "Кузнечик"
	AES128_GCM  = "AES128-GCM"

	DES_CFB    = "DES-CFB"
	AES128_CFB = "AES128-CFB"

	BF_CFB = "Blowfish-CFB"
	BF_GCM = "Blowfish-GCM"

	AES128_ECB      = "AES128-ECB"
	DES_ECB         = "DES-ECB"
	BF_ECB          = "Blowfish-ECB"
	Grasshopper_ECB = "Кузнечик-ECB"

	RSA = "RSA"
)

func GCMSeal(gcm cipher.AEAD, nonce, text []byte) (res []byte) {
	gcm.Seal(res, nonce, text, nil)
	return
}

func GCMOpen(gcm cipher.AEAD, nonce, src []byte) (res []byte) {
	gcm.Open(res, nonce, src, nil)
	return
}

func CFBSeal(c cipher.Block, bs int, text []byte) (res []byte) {
	res = make([]byte, bs+len(text))
	iv := make([]byte, bs)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(c, iv)
	stream.XORKeyStream(res, text)
	return
}

func CFBOpen(c cipher.Block, bs int, text []byte) (res []byte) {
	res = make([]byte, bs+len(text))
	iv := text[:bs]

	stream := cipher.NewCFBDecrypter(c, iv)
	stream.XORKeyStream(res, text)
	return
}

func ECBSeal(mode cipher.BlockMode, text []byte) (res []byte) {
	res = make([]byte, len(text))
	mode.CryptBlocks(res, text)
	return
}

func ECBOpen(mode cipher.BlockMode, text []byte) (res []byte) {
	res = make([]byte, len(text))
	mode.CryptBlocks(res, text)
	return
}
