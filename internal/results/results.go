package results

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type ExperimentDigest struct {
	SortedId      int       `json:"sorted_id" db:"id"`
	AlgorithmType string    `json:"algorithm_type" db:"algorithm_type"`
	StartedAt     time.Time `json:"started_at" db:"started_at"`
}

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

func (h HashAlgorithmsResults) Value() (driver.Value, error) {
	return json.Marshal(h)
}

func (h *HashAlgorithmsResults) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &h)
}

type CipherExpResult struct {
	Algorithm           string        `json:"algorithm"`
	Type                string        `json:"type"`
	CipheringDuration   time.Duration `json:"ciphering_duration"`
	DecipheringDuration time.Duration `json:"deciphering_duration"`
	KeyLength           int           `json:"key_length"`
	Sample              string        `json:"sample"`
}

type CipherAlgorithmsResults struct {
	Results    []CipherExpResult `json:"results"`
	StartedAt  time.Time         `json:"started_at"`
	FinishedAt time.Time         `json:"finished_at"`
}
