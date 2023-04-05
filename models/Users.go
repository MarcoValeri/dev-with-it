package models

import (
	util "devwithit/util"
	"log"
	"strconv"
)

type UserAdmin struct {
	Id       string
	Email    string
	Password string
}

var (
	setId       int
	setEmail    string
	setPassword string
)

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

func CheckUserEmail(email string) bool {

	// Check if the user exist in the db
	db := util.Connect()
	rows, err := db.Query("SELECT email FROM users WHERE email = ?", email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		return true
	}

	return false

}

func GetAllUsers() []UserAdmin {

	db := util.Connect()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var Id int
	var Email string
	var Password string
	var userAdminData []UserAdmin

	for rows.Next() {
		err = rows.Scan(&Id, &Email, &Password)
		if err != nil {
			panic(err)
		}
		IdStr := strconv.Itoa(Id)
		userAdminData = append(userAdminData, UserAdmin{Id: IdStr, Email: Email, Password: Password})
	}

	return userAdminData

}

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

func UserNew(email, password string) {

	// Add new user into the db
	db := util.Connect()
	insert, err := db.Query("INSERT INTO users (email, password) VALUES (?, ?)", email, password)
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}
