package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// the JWT function to generate the token
func createToken(username string) (string, error) {
	tonken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"name": username,
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
		})

	tonkenString, err := tonken.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tonkenString, nil
}
