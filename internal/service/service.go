package service

import (
	"gitlab.com/Valghall/diwor/internal/results"
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

// Result interface represents the data, accumulated during the experiment
type Result interface{}

// Experiment interface encapsulates logic necessary for accumulating resulting data
type Experiment interface {
	ResearchHashingAlgorithm(alg string, har *results.HashAlgorithmsResults) results.HashExpResult
	ResearchCipheringAlgorithm(alg string, car *results.CipherAlgorithmsResults) results.CipherExpResult
	SaveResults(userId int, algType string, results Result)
	GetLastExperimentResults(userId int) results.HashAlgorithmsResults
	GetRecentExperiments(id int) []results.ExperimentDigest
}

type Services struct {
	Authorization
	Experiment
}

func NewServices(storage *storage.Storage) *Services {
	return &Services{
		Authorization: NewAuthService(storage.Authorization),
		Experiment:    NewExperimentService(storage.Experiment),
	}
}
