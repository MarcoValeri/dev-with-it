package models

import (
	util "devwithit/util"
	"log"
)

var (
	setId       int
	setEmail    string
	setPassword string
)

func UserData(email string) (string, string) {

	// Check if the user exist and gets their email and password
	db := util.Connect()
	rows, err := db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&setId, &setEmail, &setPassword)
		if err != nil {
			log.Fatal(err)
		}
	}

	return setEmail, setPassword

}

func CheckUser(email, password string) bool {

	// Check if the user exist in the db
	db := util.Connect()
	rows, err := db.Query("SELECT * FROM users WHERE email = ? AND password = ?", email, password)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		return true
	}

	return false

}

func UserNew(email, password string) {

	// Add new user into the db
	db := util.Connect()
	insert, err := db.Query("INSERT INTO users (email, password) VALUES (?, ?)", email, password)
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}
