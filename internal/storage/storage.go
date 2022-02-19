package storage

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/Valghall/diwor/internal/results"
	"gitlab.com/Valghall/diwor/internal/users"
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
	GetLastExperimentResults(userId int) (res results.HashAlgorithmsResults)
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
