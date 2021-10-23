package experimentator

import (
	exp "gitlab.com/Valghall/diwor/internal/experiment"
	ini "gitlab.com/Valghall/diwor/internal/initial_data"
)

func HoldExperiment(algorithmA, algorithmB, modeA, modeB string) (string, string) {
	initialData := ini.NewInitialData(algorithmA, algorithmB, modeA, modeB)
	experiment := exp.NewExperiment(initialData)
	experiment.Start()
	return experiment.EncryptedA(), experiment.EncryptedB()
}
