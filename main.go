package main

import (
	"html/template"
	"net/http"
)

type PageData struct {
	PageTitle string
}

func main() {
	// Static files
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// Template
	tmpl := template.Must(template.ParseFiles("./views/home.html"))

	// Handle
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			PageTitle: "Dev With It",
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}
