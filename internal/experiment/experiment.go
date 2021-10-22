package experiment

import (
	"os"

	ini "gitlab.com/Valghall/diwor/internal/initial_data"
)

func NewExperiment(i *ini.InitialData) *experiment {
	return &experiment{
		initialData: i,
	}
}

func (e *experiment) Start() {
	plaintext := e.getPlainTextFromFile("./internal/experiment/original_text_sample.txt")
	e.makeCryptography(plaintext)
}

func (e *experiment) makeCryptography(plaintext []byte) {
	dstA := make([]byte, len(plaintext))
	dstB := make([]byte, len(plaintext))
	e.InitialData().SampleA.Encrypt(dstA, plaintext)
	e.InitialData().SampleB.Encrypt(dstB, plaintext)
	e.SetEncrypted(dstA, dstB)

	resA := make([]byte, len(dstA))
	resB := make([]byte, len(dstB))
	e.InitialData().SampleA.Decrypt(resA, dstA)
	e.InitialData().SampleB.Decrypt(resB, dstB)
	e.SetDecrypted(resA)
}

func (e experiment) getPlainTextFromFile(filepath string) []byte {
	plaintext, _ := os.ReadFile(filepath)
	return plaintext
}
