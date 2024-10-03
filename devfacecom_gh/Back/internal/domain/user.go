// File: internal/domain/user.go

package domain

import (
	"errors"
)

type User struct {
	Username string
	Password string // В реальном приложении здесь должен быть хеш пароля
}

type AuthUseCase interface {
	Login(username, password string) (string, error)
	ValidateToken(token string) (string, error)
}

type UserRepository interface {
	GetByUsername(username string) (*User, error)
	Save(user *User) error
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
)