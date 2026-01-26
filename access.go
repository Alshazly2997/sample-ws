package main

import (
	"database/sql"
)

// The function responsible for connecting to the DB and read a required data
func AccessData() User {
	var db *sql.DB
	var query User
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
		query.Name = user.Name
		query.Password = user.Password
	}
	return query
}
