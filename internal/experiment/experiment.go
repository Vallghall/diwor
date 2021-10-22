package experiment

import (
	"os"
)

type experiment struct {
	initialData *initialData
	encrypted   []byte
	decrypted   []byte
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

func (e *experiment) Decrypted() []byte {
	return e.decrypted
}

func (e *experiment) SetEncrypted(encrypted []byte) {
	e.encrypted = encrypted
}

func (e *experiment) SetDecrypted(decrypted []byte) {
	e.decrypted = decrypted
}

func (i *initialData) NewExperiment() *experiment {
	return &experiment{
		initialData: i,
	}
}

func (e *experiment) Start() {
	plaintext, _ := os.ReadFile("./internal/experiment/original_text_sample.txt")
	dst := make([]byte, len(plaintext))
	e.InitialData().Encrypt(dst, plaintext)
	e.SetEncrypted(dst)

	res := make([]byte, len(dst))
	e.InitialData().Decrypt(res, dst)
	e.SetDecrypted(res)
}
