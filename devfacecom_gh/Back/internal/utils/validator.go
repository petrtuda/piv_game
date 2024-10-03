// File: internal/utils/validator.go

package utils

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func ValidateUsername(username string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]+$", username)
	return matched
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
