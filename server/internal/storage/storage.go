package storage

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/Valghall/diwor/server/internal/results"
	"gitlab.com/Valghall/diwor/server/internal/users"
)

type Authorization interface {
	CreateUser(user users.User) (int, error)
	GetUser(username, password string) (users.User, error)
	GetUserById(id int) (users.User, error)
	LookUpUser(username string) bool
}

type Experiment interface {
	SaveHashAlgorithmResults(userId int, algType string, results results.HashAlgorithmsResults)
	SaveCipherAlgorithmResults(userId int, algType string, results results.CipherAlgorithmsResults)
	GetLastHashExperimentResults(userId int) (res results.HashAlgorithmsResults)
	GetLastCipherExperimentResults(userId int) (res results.CipherAlgorithmsResults)
	GetRecentExperiments(id, quantity int) ([]results.ExperimentDigest, error)
	GetAllUserExperiments(id int) []results.ExperimentDigest
	GetUserHashExperimentResultBySortedId(userId, sortedId int) (results.HashAlgorithmsResults, error)
	GetUserCipherExperimentResultBySortedId(userId, sortedId int) (results.CipherAlgorithmsResults, error)
}

type Storage struct {
	Authorization
	Experiment
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Authorization: NewAuthPostgress(db),
		Experiment:    NewExperimentPostgres(db),
	}
}
