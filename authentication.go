package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// the authentication handler (/auth)
func authHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_type", "application/json")

	var user User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Printf("The user request value %v", user)

	if user.Name == "Mohamed" && user.Password == "1234" {
		tokenString, err := createToken(user.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("No user name found\n")
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
	}
}

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
