package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// for localhost:8080/
func greetingHandler(w http.ResponseWriter, r *http.Request) {
	var greeting = "Hello World!"
	fmt.Fprintln(w, greeting)

}

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

// This handler validates the token and processes the query request.
func queryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_type", "application/json")

	tonkenString := r.Header.Get("Authorization")
	if tonkenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}

	tonkenString = tonkenString[len("Bearer "):]

	err := verifyToken(tonkenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}

	query := AccessData()
	json.NewEncoder(w).Encode(query)

}
