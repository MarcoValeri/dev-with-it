package controller

import (
	"net/http"
	"text/template"
)

type PageData struct {
	PageTitle string
}

func Home() {

	tmpl := template.Must(template.ParseFiles("./views/home.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			PageTitle: "Dev with It",
		}
		tmpl.Execute(w, data)
	})
}
