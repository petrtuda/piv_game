package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"devfacecom/Back/internal/api"
	"devfacecom/Back/internal/config"
	"devfacecom/Back/internal/pkg/logger"
	"devfacecom/Back/internal/repository/csv"
	"devfacecom/Back/internal/usecase"
)

func main() {
	// Initialize logger
	logger := logger.NewLogger()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load configuration", err)
	}

	// Initialize repositories
	userRepo := csv.NewUserRepository("./data/users.json")
	beerRepo := csv.NewBeerRepository("./data/dep_history.csv", "./data/balance_history.csv")

	// Initialize use cases
	authUseCase := usecase.NewAuthUseCase(userRepo)
	beerUseCase := usecase.NewBeerUseCase(beerRepo)

	// Initialize router
	router := api.NewRouter(authUseCase, beerUseCase)

	// Configure server
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	// Start server
	go func() {
		logger.Info("Starting server on port " + cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", err)
	}

	logger.Info("Server exiting")
}
