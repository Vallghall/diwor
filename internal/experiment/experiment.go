package experiment

type experiment struct {
	InitialData *initialData
}

func (i *initialData) NewExperiment() *experiment {
	return &experiment{
		InitialData: i,
	}
}

/*
func (e *experiment) Start() {
	plaintext, _ := os.ReadFile("./internal/experiment/original_text_sample.txt")
	crypto.Encrypt(plaintext, e.InitialData.Mode, e.InitialData.SampleA.Algorithm(), e.InitialData.SampleA.BlockSize())
}
*/
