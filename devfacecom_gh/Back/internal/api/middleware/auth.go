// File: internal/api/middleware/auth.go

package middleware

import (
	"context"
	"net/http"
	"myapp/internal/domain"
	"strings"
)

type AuthMiddleware struct {
	authUseCase domain.AuthUseCase
}

func NewAuthMiddleware(authUseCase domain.AuthUseCase) *AuthMiddleware {
	return &AuthMiddleware{authUseCase: authUseCase}
}

func (m *AuthMiddleware) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}

		username, err := m.authUseCase.ValidateToken(tokenParts[1])
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "username", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}