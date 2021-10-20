package experiment

import (
	"os"

	"gitlab.com/Valghall/diwor/internal/crypto"
)

type experiment struct {
	initialData *initialData
	encrypted   []byte
}

func (e *experiment) InitialData() *initialData {
	return e.initialData
}

func (e *experiment) SetInitialData(initialData *initialData) {
	e.initialData = initialData
}

func (e *experiment) Encrypted() []byte {
	return e.encrypted
}

func (e *experiment) SetEncrypted(encrypted []byte) {
	e.encrypted = encrypted
}

func (i *initialData) NewExperiment() *experiment {
	return &experiment{
		initialData: i,
	}
}

func (e *experiment) Start() {
	plaintext, _ := os.ReadFile("./internal/experiment/original_text_sample.txt")
	encrypted, _ := crypto.Encrypt(plaintext,
		e.InitialData().Mode,
		e.InitialData().SampleA.Cipher(),
		e.InitialData().SampleA.BlockSize(),
	)
	e.SetEncrypted(encrypted)
}
