package experimentator

import "gitlab.com/Valghall/diwor/internal/experiment"

func HoldExperiment(algorithmA, algorithmB, mode string) {
	initialData := experiment.NewInitialData(algorithmA, algorithmB, mode)
	_ = initialData.NewExperiment()
}
