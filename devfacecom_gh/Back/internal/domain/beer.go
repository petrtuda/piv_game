// File: internal/domain/beer.go

package domain

import (
	"errors"
	"time"
)

type BeerDeposit struct {
	Username string
	Amount   float64
	Time     time.Time
}

type BeerUseCase interface {
	AddBeer(username string, amount float64) error
	GetBalance(username string) (float64, error)
}

type BeerRepository interface {
	SaveDeposit(deposit *BeerDeposit) error
	GetBalance(username string) (float64, error)
}

var (
	ErrMinDepositAmount = errors.New("minimum deposit amount not met")
	ErrInvalidAmount    = errors.New("invalid amount")
)