package main

import (
	homeController "devwithit/controllers"
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
	homeController.HomeController()

	http.ListenAndServe(":80", nil)
}
