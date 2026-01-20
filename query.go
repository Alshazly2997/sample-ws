package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
)

var db *sql.DB

// this function check the token validity and then process the query
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

	// Open database connection

	db, err := sql.Open("mysql", "root:death notemysql@tcp(127.0.0.1:3306)/usersdb")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//########### to read data from the database ##########
	resluts, err := db.Query("SELECT * FROM my_user")

	if err != nil {
		panic(err.Error())
	}

	var user User
	for resluts.Next() {
		err = resluts.Scan(&user.Name, &user.Password)
		if err != nil {
			panic(err.Error())
		}

		//in real-world scenarois this data must be encrypted
		fmt.Fprint(w, user.Name, user.Password)
	}

}

// the JWT function to verify the token
func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
