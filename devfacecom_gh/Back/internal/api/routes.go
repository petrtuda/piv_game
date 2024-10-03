// File: internal/api/routes.go

package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"myapp/internal/api/handlers"
	"myapp/internal/api/middleware"
	"myapp/internal/domain"
)

func NewRouter(authUseCase domain.AuthUseCase, beerUseCase domain.BeerUseCase) http.Handler {
	router := mux.NewRouter()

	authHandler := handlers.NewAuthHandler(authUseCase)
	beerHandler := handlers.NewBeerHandler(beerUseCase)
	authMiddleware := middleware.NewAuthMiddleware(authUseCase)

	router.HandleFunc("/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/beer/add", authMiddleware.Authenticate(beerHandler.AddBeer)).Methods("POST")

	return router
}