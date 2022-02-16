package storage

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	resulties "gitlab.com/Valghall/diwor/internal/results"
)

type ExperimentPostgres struct {
	db *sqlx.DB
}

func NewExperimentPostgres(db *sqlx.DB) *ExperimentPostgres {
	return &ExperimentPostgres{db: db}
}

type HashExpValues struct {
	Algorithm string `json:"algorithm,"`
	Duration  string `json:"duration"`
	Size      int    `json:"size"`
	BlockSize int    `json:"blockSize"`
	Sample    string `json:"sample"`
}

type HEVArray struct {
	Results []HashExpValues `json:"results"`
}

func NewHEVArray(results resulties.HashAlgorithmsResults) (res HEVArray) {
	for _, result := range results.Results {
		res.Results = append(res.Results, HashExpValues{
			Algorithm: result.(resulties.HashExpResult).Algorithm,
			Duration:  result.(resulties.HashExpResult).Duration.String(),
			Size:      result.(resulties.HashExpResult).Size,
			BlockSize: result.(resulties.HashExpResult).BlockSize,
			Sample:    result.(resulties.HashExpResult).Sample,
		})
	}

	return
}

func (h HEVArray) Value() (driver.Value, error) {
	return json.Marshal(h)
}

func (h *HEVArray) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &h)
}

func (ep *ExperimentPostgres) SaveResults(userId int, algType string, results resulties.HashAlgorithmsResults) {
	hev := NewHEVArray(results)

	query := fmt.Sprintf(
		`INSERT into %s (
		user_id,
		algorithm_type,
		results,
		started_at,
		finished_at) VALUES ($1,$2,$3,$4,$5)`, experimentsTable)
	_, err := ep.db.Query(
		query,
		userId,
		algType,
		hev,
		results.StartedAt,
		results.FinishedAt,
	)

	logrus.Error(err)
}
