package plotconfig

type Config struct {
	From int `json:"from"`
	To   int `json:"to"`
	Step int `json:"step"`
}

func NewConfig(to int, step int) *Config {
	return &Config{To: to, Step: step}
}
