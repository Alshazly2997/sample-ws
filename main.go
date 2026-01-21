package main

import (
	"fmt"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

var secretKey = []byte("test_password")

func main() {

	http.HandleFunc("/", greetingHandler)
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/query", queryHandler)

	err := http.ListenAndServe("localhost:8080", http.DefaultServeMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
