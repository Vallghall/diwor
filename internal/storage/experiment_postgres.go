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

func (ep *ExperimentPostgres) GetLastExperimentResults(userId int) (res results.HashAlgorithmsResults) {
	query := fmt.Sprintf(
		`SELECT results, started_at, finished_at
	FROM %s
	WHERE user_id=$1
	ORDER BY started_at DESC
	LIMIT 1;`, experimentsTable)

	row := ep.db.QueryRow(
		query,
		userId,
	)

	err := row.Scan(&res, &res.StartedAt, &res.FinishedAt)
	if err != nil {
		logrus.Error(err)
	}

	return
}

func (ep *ExperimentPostgres) GetRecentExperiments(id, quantity int) (res []results.ExperimentDigest) {
	query := fmt.Sprintf(
		`SELECT
					ROW_NUMBER () OVER (ORDER BY started_at DESC) AS id,
					algorithm_type,
					started_at
				FROM %s
				WHERE user_id=$1
				ORDER BY started_at DESC
				LIMIT $2;`,
		experimentsTable,
	)

	rows, err := ep.db.Query(query, id, quantity)
	if err != nil {
		logrus.Error(err)
	}
	defer rows.Close()

	for rows.Next() {
		var resultSet results.ExperimentDigest
		err = rows.Scan(
			&resultSet.SortedId,
			&resultSet.AlgorithmType,
			&resultSet.StartedAt,
		)
		if err != nil {
			logrus.Error(err)
		}

		res = append(res, resultSet)
	}

	return res
}