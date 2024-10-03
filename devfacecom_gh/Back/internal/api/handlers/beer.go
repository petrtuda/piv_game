// File: internal/api/handlers/beer.go

package handlers

import (
	"encoding/json"
	"net/http"
	"myapp/internal/domain"
	"strconv"
)

type BeerHandler struct {
	beerUseCase domain.BeerUseCase
}

func NewBeerHandler(beerUseCase domain.BeerUseCase) *BeerHandler {
	return &BeerHandler{beerUseCase: beerUseCase}
}

func (h *BeerHandler) AddBeer(w http.ResponseWriter, r *http.Request) {
	var addBeerRequest struct {
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&addBeerRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	username := r.Context().Value("username").(string)

	if err := h.beerUseCase.AddBeer(username, addBeerRequest.Amount); err != nil {
		if err == domain.ErrMinDepositAmount {
			http.Error(w, "Min dep summ 5", http.StatusBadRequest)
		} else {
			http.Error(w, "Failed to add beer", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Beer added successfully"})
}
