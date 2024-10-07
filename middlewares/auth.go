package middlewares

import (
	"btelli-customersupport-app/handlers"
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func Auth(allewedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Step 1- Validate JWT
			claims, err := ValidateJWT(r)
			if err != nil {
				handlers.ApiResponse(w, nil, err.Error(), http.StatusUnauthorized)
				return
			}

			// Step 2- Valite Role
			err = ValidateRole(claims, allewedRoles...)
			if err != nil {
				handlers.ApiResponse(w, nil, err.Error(), http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), "claims", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ValidateJWT(r *http.Request) (*jwt.MapClaims, error) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("no token provided")
	}

	tokenString := strings.Split(authHeader, " ")[1]
	if tokenString == "" {
		return nil, errors.New("invalid token format")
	}

	claims := &jwt.MapClaims{}

	// Validate Token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ValidateRole(claims *jwt.MapClaims, allowedRoles ...string) error {
	userRole := strings.ToLower((*claims)["role"].(string))

	for _, role := range allowedRoles {
		if userRole == strings.ToLower(role) {
			return nil
		}
	}
	return errors.New("forbidden: You have no permission to see this content")
}
