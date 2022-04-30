package service

import (
	"errors"
	"fmt"
	myerr "gitlab.com/Valghall/diwor/server/internal/errors"
	"gitlab.com/Valghall/diwor/server/internal/storage"
	"gitlab.com/Valghall/diwor/server/internal/users"
	"math/rand"
	"os"
	"regexp"
	"time"

	streebog "github.com/bi-zone/ruwireguard-go/crypto/gost/gost34112012256"
	"github.com/dgrijalva/jwt-go"
)

const (
	AccessTokenTTL  = 5 * time.Minute
	RefreshTokenTTL = 24 * 7 * time.Hour

	NamePatternRU   = `^[а-яА-Я]+$`
	NamePatternEN   = `^[a-zA-Z]+$`
	PasswordPattern = `^[а-яА-Яa-zA-Z0-9]+$`
)

var (
	signingKey = os.Getenv("SIGNING_KEY")
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId  int     `json:"user_id"`
	Marcant float64 `json:"marcant"`
}

type AuthService struct {
	storage storage.Authorization
}

func NewAuthService(storage storage.Authorization) *AuthService {
	return &AuthService{storage: storage}
}

func (as *AuthService) ValidateUserCredentials(user users.User) (bool, error) {
	if user.Name == "" || user.Username == "" || user.Password == "" {
		return false, myerr.ErrEmptyFields
	}

	if len(user.Password) < 6 {
		return false, myerr.ErrLittlePasswordLength
	}

	if ok, err := validateWithEitherOfRegExp(user.Name, NamePatternRU, NamePatternEN); err != nil {
		return false, err
	} else if !ok {
		return false, myerr.ErrNonAlphabetSymbols
	}

	if ok, err := validateWithRegExp(user.Password, PasswordPattern); err != nil {
		return false, err
	} else if !ok {
		return false, myerr.ErrInvalidPasswordCharachters
	}

	return true, nil
}

func (as *AuthService) CreateUser(user users.User) (int, error) {
	if as.storage.LookUpUser(user.Username) {
		return 0, myerr.ErrUsernameAlreadyExists
	}
	user.Password = generatePasswordHash(user.Password)
	return as.storage.CreateUser(user)
}

func (as *AuthService) GenerateTokenPair(username, password string) (string, string, error) {
	user, err := as.storage.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", "", err
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId:  user.Id,
		Marcant: rand.Float64() * rand.Float64() * 1000.0,
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(RefreshTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId:  user.Id,
		Marcant: rand.Float64() * rand.Float64() * 1000.0,
	})

	ATString, err := accessToken.SignedString([]byte(signingKey))
	if err != nil {
		return "", "", err
	}

	RTString, err := refreshToken.SignedString([]byte(signingKey))

	return ATString, RTString, nil
}

func (as *AuthService) RegenerateTokenPair(userId int) (string, string, error) {

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId:  userId,
		Marcant: rand.Float64() * rand.Float64() * 1000.0,
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(RefreshTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId:  userId,
		Marcant: rand.Float64() * rand.Float64() * 1000.0,
	})

	ATString, err := accessToken.SignedString([]byte(signingKey))
	if err != nil {
		return "", "", err
	}

	RTString, err := refreshToken.SignedString([]byte(signingKey))

	return ATString, RTString, nil
}

func generatePasswordHash(password string) string {
	hash := streebog.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func validateWithEitherOfRegExp(str string, patterns ...string) (bool, error) {
	var ok bool
	for _, pattern := range patterns {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return false, err
		}
		ok = re.MatchString(str)
		if ok {
			return ok, nil
		}
	}

	return false, nil
}

func validateWithRegExp(str, pattern string) (bool, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}

	return re.MatchString(str), nil
}

func (as *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New(myerr.ErrTokenExpired.Error())
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (as *AuthService) GetUserById(id int) (users.User, error) {
	user, err := as.storage.GetUserById(id)
	if err != nil {
		return users.User{}, err
	}
	return user, nil
}
