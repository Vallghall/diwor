package service

import (
	"gitlab.com/Valghall/diwor/internal/storage"
	"gitlab.com/Valghall/diwor/internal/users"
)

type Authorization interface {
	CreateUser(user users.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Sample interface {
}

type Experiment interface {
}

type Service struct {
	Authorization
	Experiment
	Sample
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Authorization: NewAuthService(storage.Authorization),
	}
}
