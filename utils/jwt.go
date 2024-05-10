package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT generates a JWT for the user
func GenerateJWT(id string, email string, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
