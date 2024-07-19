package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/SolBaa/task-manager/config"
	"github.com/SolBaa/task-manager/internal/auth"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserKey contextKey = "username"

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conf := config.LoadConfig()
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		fmt.Printf("Loaded JWT Secret: %s\n", conf.JWTSecret)

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		fmt.Printf("Token String: %s\n", tokenString)
		claims := &auth.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(conf.JWTSecret), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Invalid token signature", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), UserKey, claims.Username)

		r.Header.Set("username", claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
