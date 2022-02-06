package service

import (
	"errors"
	"fmt"
	streebog "github.com/bi-zone/ruwireguard-go/crypto/gost/gost34112012256"
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/Valghall/diwor/internal/storage"
	"gitlab.com/Valghall/diwor/internal/users"
	"time"
)

const (
	TokenTTL   = 12 * time.Hour
	signinhKey = "this is my custom Secret key for authentication"
)

var (
	ErrUsernameAlreadyExists = errors.New("user with this username already exists")
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	storage storage.Authorization
}

func NewAuthService(storage storage.Authorization) *AuthService {
	return &AuthService{storage: storage}
}

func (as *AuthService) CreateUser(user users.User) (int, error) {
	if as.storage.LookUpUser(user.Username) {
		return 0, ErrUsernameAlreadyExists
	}
	user.Password = generatePasswordHash(user.Password)
	return as.storage.CreateUser(user)
}

func (as *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := as.storage.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})
	return token.SignedString([]byte(signinhKey))
}

func generatePasswordHash(password string) string {
	hash := streebog.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (as *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signinhKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
