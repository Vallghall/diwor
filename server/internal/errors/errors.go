package errors

import "errors"

var (
	ErrUsernameAlreadyExists      = errors.New("user with this username already exists")
	ErrEmptyFields                = errors.New("необходимо заполнить все поля")
	ErrLittlePasswordLength       = errors.New("длина пароля должна превышать 6 символов")
	ErrNonAlphabetSymbols         = errors.New("имя должно состоять из символов только либо латинского либо русского алфавита")
	ErrInvalidPasswordCharachters = errors.New("пароль может содержать только символы а-я, А-Я, a-z, A-Z, 0-9")
	ErrInvalidLoginOrPassword     = errors.New("неверный логин или пароль")
	ErrUserCtxNotFound            = errors.New("User context not found")
	ErrInvalidHashAlgorithmInput  = errors.New("введен(-ы) некорректные тип(-ы) алгоритмов хэширования")
	ErrTokenExpired               = errors.New("token has expired")
)
