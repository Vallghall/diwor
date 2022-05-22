package service

import (
	"crypto/cipher"
	"crypto/rand"
	"io"
)

const (
	Grasshopper_GCM = "Кузнечик-GCM"
	Grasshopper_ECB = "Кузнечик-ECB"
	Grasshopper_OFB = "Кузнечик-OFB"
	Grasshopper_CFB = "Кузнечик-CFB"
	Grasshopper_CTR = "Кузнечик-CTR"

	AES128_GCM = "AES128-GCM"
	AES128_ECB = "AES128-ECB"
	AES128_CFB = "AES128-CFB"
	AES128_OFB = "AES128-OFB"
	AES128_CTR = "AES128-CTR"

	MGM_GCM = "MGM-GCM"
	MGM_ECB = "Магма-ECB"
	MGM_CFB = "Магма-CFB"
	MGM_OFB = "Магма-OFB"
	MGM_CTR = "Магма-CTR"

	DES_CFB = "DES-CFB"
	DES_ECB = "DES-ECB"
	DES_OFB = "DES-OFB"
	DES_CTR = "DES-CTR"

	BF_CFB = "Blowfish-CFB"
	BF_OFB = "Blowfish-OFB"
	BF_ECB = "Blowfish-ECB"
	BF_CTR = "Blowfish-CTR"

	RSA = "RSA"
	EG  = "Эль-Гамаль"
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

func OFBSealOpen(c cipher.Block, bs int, text []byte) (res []byte) {
	res = make([]byte, bs+len(text))
	iv := make([]byte, bs)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewOFB(c, iv)
	stream.XORKeyStream(res, text)
	return
}

func CTRSeal(c cipher.Block, bs int, text []byte) (res []byte) {
	res = make([]byte, bs+len(text))
	iv := make([]byte, bs)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCTR(c, iv)
	stream.XORKeyStream(res, text)
	return
}

func CTROpen(c cipher.Block, bs int, text []byte) (res []byte) {
	res = make([]byte, bs+len(text))
	iv := text[:bs]

	stream := cipher.NewCTR(c, iv)
	stream.XORKeyStream(res, text)
	return
}
