package controller

import (
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

type LoginDetails struct {
	Email    string
	Password string
}

// Initialize the session
var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func Login() {

	tmpl := template.Must(template.ParseFiles("./views/admin/login.html"))

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		// Session
		session, _ := store.Get(r, "cookie-name")

		// Session: set the value
		session.Values["authenticated"] = false
		session.Save(r, w)

		details := LoginDetails{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		if details.Email == "info@marcovaleri.net" {
			// Set session true
			session.Values["authenticated"] = true
			session.Save(r, w)
		} else {
			// Set session false
			session.Values["authenticated"] = false
			session.Save(r, w)
		}

		data := PageData{
			PageTitle: "Login Page",
		}

		tmpl.Execute(w, data)

	})
}

func AdminDashboard() {
	tmpl := template.Must(template.ParseFiles("./views/admin/dashboard.html"))

	http.HandleFunc("/admin/dashboard", func(w http.ResponseWriter, r *http.Request) {

		// Session: get it
		session, _ := store.Get(r, "cookie-name")

		if session.Values["authenticated"] == true {

			data := PageData{
				PageTitle: "Admin Dashboard",
			}

			tmpl.Execute(w, data)

		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}

	})
}
