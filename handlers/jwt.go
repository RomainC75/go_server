package handlers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTData struct {
	jwt.StandardClaims
	CustomClaims map[string]string `json:"custom_claims"`
}

func GenerateJWTAccessToken(userId string, username string) (string, error) {
	secretKey := os.Getenv("JWT_TOKEN")
	fmt.Print("--------", secretKey, "\b")
	fmt.Print("======\n")

	claims := JWTData{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(time.Millisecond * 1000 * 60 * 60)).Unix(),
		},
		CustomClaims: map[string]string{
			"user_id":  userId,
			"username": username,
		},
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenString.SignedString(secretKey)

	return token, err
}
