package initial_data

import (
	sample2 "gitlab.com/Valghall/diwor/internal/sample"
)

type InitialData struct {
	SampleA *sample2.Sample
	SampleB *sample2.Sample
}

type CryptographyData struct {
	EncryptedA string
	EncryptedB string
}

func NewInitialData(algorithmA, algorithmB, modeA, modeB string) *InitialData {
	return &InitialData{
		SampleA: sample2.CreateSample(algorithmA, modeA),
		SampleB: sample2.CreateSample(algorithmB, modeB),
	}
}
