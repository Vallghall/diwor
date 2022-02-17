package results

import "time"

type HashExpResult struct {
	Algorithm string        `json:"algorithm"`
	Duration  time.Duration `json:"duration"`
	Size      int           `json:"size"`
	BlockSize int           `json:"blockSize"`
	Sample    string        `json:"sample"`
}

type HashAlgorithmsResults struct {
	Results    []HashExpResult `json:"results"`
	StartedAt  time.Time       `json:"started_at"`
	FinishedAt time.Time       `json:"finished_at"`
}

func (her HashExpResult) DurationMilliSeconds() int64 {
	return her.Duration.Milliseconds()
}
