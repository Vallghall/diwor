package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gitlab.com/Valghall/diwor/internal/results"
)

type ExperimentPostgres struct {
	db *sqlx.DB
}

func NewExperimentPostgres(db *sqlx.DB) *ExperimentPostgres {
	return &ExperimentPostgres{db: db}
}

func (ep *ExperimentPostgres) SaveHashAlgorithmResults(userId int, algType string, results results.HashAlgorithmsResults) {
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
		results,
		results.StartedAt,
		results.FinishedAt,
	)

	if err != nil {
		logrus.Error(err)
	}
}

func (ep *ExperimentPostgres) SaveCipherAlgorithmResults(userId int, algType string, results results.CipherAlgorithmsResults) {
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
		results,
		results.StartedAt,
		results.FinishedAt,
	)

	if err != nil {
		logrus.Error(err)
	}
}
