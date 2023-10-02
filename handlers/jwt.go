package handlers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTAccessToken(userId string, username string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_TOKEN"))
	fmt.Print("--------", secretKey, "\b")
	fmt.Print("======\n")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
