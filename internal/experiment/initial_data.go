package experiment

import "gitlab.com/Valghall/diwor/internal/crypto"

type initialData struct {
	sample
	SampleA *sample
	SampleB *sample
	Mode    string
}

func NewInitialData(algorithmA, algorithmB, mode string) *initialData {
	return &initialData{
		SampleA: createSample(algorithmA),
		SampleB: createSample(algorithmB),
		Mode:    mode,
	}
}

func (i initialData) Encrypt(dst, src []byte) {
	encrypted, _ := crypto.Encrypt(
		src,
		i.Mode,
		i.SampleA.Cipher(),
		i.SampleA.BlockSize(),
	)
	copy(dst, encrypted)
}

func (i initialData) Decrypt(dst, src []byte) {
	decrypted, _ := crypto.Decrypt(
		src,
		i.Mode,
		i.SampleA.Cipher(),
		i.SampleA.BlockSize(),
	)
	copy(dst, decrypted)
}
