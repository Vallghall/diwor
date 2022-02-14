package service

import (
	"gitlab.com/Valghall/diwor/internal/storage"
	"gitlab.com/Valghall/diwor/internal/users"
)

// Authorization interface encapsulates logic for user registration and authentication
type Authorization interface {
	ValidateUserCredentials(user users.User) (bool, error)
	CreateUser(user users.User) (int, error)
	GetUserById(id int) (users.User, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

// InitialData interface represents the data, received from the user's input
type InitialData interface {
	AlgorithmType() string
}

// Result interface represents the data, accumulated during the experiment
type Result interface {
	Duration() int64
}

// Experiment interface encapsulates logic necessary for accumulating resulting data
type Experiment interface {
	Hold(ini InitialData) Result
}

type Services struct {
	Authorization
	Experiment
	InitialData
}

func NewServices(storage *storage.Storage) *Services {
	return &Services{
		Authorization: NewAuthService(storage.Authorization),
	}
}
