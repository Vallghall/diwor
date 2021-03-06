package service

import (
	"gitlab.com/Valghall/diwor/server/internal/plotconfig"
	"gitlab.com/Valghall/diwor/server/internal/results"
	"gitlab.com/Valghall/diwor/server/internal/storage"
	"gitlab.com/Valghall/diwor/server/internal/users"
)

// Authorization interface encapsulates logic for user registration and authentication
type Authorization interface {
	ValidateUserCredentials(user users.User) (bool, error)
	CreateUser(user users.User) (int, error)
	GetUserById(id int) (users.User, error)
	GenerateTokenPair(username, password string) (string, string, error)
	RegenerateTokenPair(userId int) (string, string, error)
	ParseToken(token string) (int, error)
}

// Result interface represents the data, accumulated during the experiment
type Result interface{}

// Experiment interface encapsulates logic necessary for accumulating resulting data
type Experiment interface {
	ResearchHashingAlgorithm(alg string, conf plotconfig.Config) results.HashExpResult
	ResearchCipheringAlgorithm(alg string, conf plotconfig.Config) results.CipherExpResult
	SaveResults(userId int, algType string, results Result)
	GetLastHashExperimentResults(userId int) results.HashAlgorithmsResults
	GetLastCipherExperimentResults(userId int) results.CipherAlgorithmsResults
	GetRecentExperiments(id int) ([]results.ExperimentDigest, error)
	GetAllUserExperiments(id int) []results.ExperimentDigest
	GetUserExperimentResultBySortedId(alg string, userId, sortedId int) (Result, error)
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
