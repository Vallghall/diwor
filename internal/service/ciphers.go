package service

import (
	"math/rand"
	"os"
	"time"
)

const (
	Grasshopper = "Кузнечик"
	DES         = "DES"
	AES         = "AES"
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
