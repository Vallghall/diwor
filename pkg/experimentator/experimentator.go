package experimentator

import "gitlab.com/Valghall/diwor/internal/experiment"

func HoldExperiment(algorithmA, algorithmB, mode string) []byte {
	initialData := experiment.NewInitialData(algorithmA, algorithmB, mode)
	exp := initialData.NewExperiment()
	exp.Start()
	return exp.Encrypted()
}
