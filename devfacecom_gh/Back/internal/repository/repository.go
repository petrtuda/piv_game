/ File: internal/repository/repository.go

package repository

import (
	"myapp/internal/domain"
)

type Repository interface {
	UserRepository() domain.UserRepository
	BeerRepository() domain.BeerRepository
}