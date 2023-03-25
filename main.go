package main

import (
	controller "devwithit/controllers"
	util "devwithit/util"
	"net/http"
)

type PageData struct {
	PageTitle string
}

func main() {
	// Static files
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// Controllers
	controller.AdminController()
	controller.Home()

	// Database
	util.Connect()

	http.ListenAndServe(":80", nil)
}
