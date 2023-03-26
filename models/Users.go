package models

import (
	util "devwithit/util"
)

func UserNew(email, password string) {

	// Add new user into the db
	db := util.Connect()
	insert, err := db.Query("INSERT INTO users (email, password) VALUES (?, ?)", email, password)
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}
