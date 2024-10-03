
// File: internal/usecase/beer.go

package usecase

import (
	"myapp/internal/domain"
	"myapp/internal/utils"
	"time"
)

type BeerUseCase struct {
	beerRepo domain.BeerRepository
	config   *utils.Config
}

func NewBeerUseCase(beerRepo domain.BeerRepository, config *utils.Config) *BeerUseCase {
	return &BeerUseCase{
		beerRepo: beerRepo,
		config:   config,
	}
}

func (uc *BeerUseCase) AddBeer(username string, amount float64) error {
	if amount < uc.config.MinDepositAmount {
		return domain.ErrMinDepositAmount
	}

	deposit := &domain.BeerDeposit{
		Username: username,
		Amount:   amount,
		Time:     time.Now(),
	}

	return uc.beerRepo.SaveDeposit(deposit)
}

func (uc *BeerUseCase) GetBalance(username string) (float64, error) {
	return uc.beerRepo.GetBalance(username)
}

// File: internal/utils/csv.go

package utils

import (
	"encoding/csv"
	"os"
)

func ReadCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func WriteCSV(filename string, records [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.WriteAll(records)
}

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