package models

import (
	util "devwithit/util"
	"fmt"
	"log"
)

// func CheckUser(email, password string) (*sql.Rows, error) {

// 	// Check if the user exist in the db
// 	db := util.Connect()
// 	check, err := db.Query("SELECT EXISTS(SELECT * FROM users WHERE email = ? AND password = ?)", email, password)
// 	// check, err := db.Query("SELECT EXISTS(SELECT * FROM users WHERE email = 'info@marcovaleri.net' AND password = 'S!lver09')")
// 	return check, err

// }

var (
	setEmail    string
	setPassword string
)

func CheckUser(email, password string) bool {

	// Check hashed password
	passwordHashed, err := util.HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(passwordHashed)

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
