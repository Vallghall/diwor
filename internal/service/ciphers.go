package service

import (
	"math/rand"
	"os"
	"time"
)

const (
	Grasshopper = "Кузнечик"
	DES_GCM     = "DES-GCM"
	AES128_GCM  = "AES128-GCM"

	DES_CFB    = "DES-CFB"
	AES128_CFB = "AES128-CFB"

	RSA = "RSA"
)

var (
	textForCiphering []byte
	letters          = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func init() {
	textForCiphering, _ = os.ReadFile("sample_phrase.txt")
	rand.Seed(time.Now().UnixNano())
}

func generateBytes(n int) ([]byte, time.Duration) {
	start := time.Now()

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return b, time.Since(start)
}
