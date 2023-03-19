package controller

import (
	"fmt"
	"net/http"
	"text/template"
)

type LoginDetails struct {
	Email    string
	Password string
}

func Login() {
	tmpl := template.Must(template.ParseFiles("./views/admin/login.html"))

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := LoginDetails{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		if details.Email == "info@marcovaleri.net" && details.Password == "S!lver09" {
			fmt.Println("Login: TRUE")
		} else {
			fmt.Println("Login: FALSE")
		}

		data := PageData{
			PageTitle: "Login Page",
		}

		tmpl.Execute(w, data)
	})
}
