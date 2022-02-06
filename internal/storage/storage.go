package storage

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/Valghall/diwor/internal/users"
)

type Authorization interface {
	CreateUser(user users.User) (int, error)
	GetUser(username, password string) (users.User, error)
}

type Sample interface {
}

type Experiment interface {
}

type Storage struct {
	Authorization
	Experiment
	Sample
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Authorization: NewAuthPostgress(db),
	}
}
