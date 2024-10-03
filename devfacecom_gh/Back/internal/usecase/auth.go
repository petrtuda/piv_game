// File: internal/usecase/auth.go

package usecase

import (
	"myapp/internal/domain"
	"myapp/internal/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthUseCase struct {
	userRepo domain.UserRepository
	config   *utils.Config
}

func NewAuthUseCase(userRepo domain.UserRepository, config *utils.Config) *AuthUseCase {
	return &AuthUseCase{
		userRepo: userRepo,
		config:   config,
	}
}

func (uc *AuthUseCase) Login(username, password string) (string, error) {
	user, err := uc.userRepo.GetByUsername(username)
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", domain.ErrInvalidCredentials
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(uc.config.TokenSecret))
}

func (uc *AuthUseCase) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(uc.config.TokenSecret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), nil
	}

	return "", domain.ErrInvalidCredentials
}
