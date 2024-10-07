package utils

import (
	"btelli-customersupport-app/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func GenerateJWT(user models.User) (string, error) {

	claims := jwt.MapClaims{
		"iss":      "www.customersupportapp.com",
		"sub":      fmt.Sprintf("%d", user.ID),
		"aud":      []string{"customersupportapiv1"},
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
		"username": user.UserName,
		"email":    user.Email,
		"role":     user.Role.Name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
