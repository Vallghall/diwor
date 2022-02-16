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

type Sample interface {
}

type Experiment interface {
	SaveResults(userId int, algType string, results results.HashAlgorithmsResults)
}

type Storage struct {
	Authorization
	Experiment
	Sample
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Authorization: NewAuthPostgress(db),
		Experiment:    NewExperimentPostgres(db),
	}
}
