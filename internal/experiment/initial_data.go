package experiment

type initialData struct {
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
