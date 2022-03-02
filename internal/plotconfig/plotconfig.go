package plotconfig

type Config struct {
	NumMeasurements int `json:"num_measurements"`
	From            int `json:"from"`
	To              int `json:"to"`
	Step            int `json:"step"`
}

func NewConfig(numMeasurements int, to int, step int) *Config {
	return &Config{NumMeasurements: numMeasurements, To: to, Step: step}
}
