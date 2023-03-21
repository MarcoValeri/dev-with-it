package adminControllers

import (
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

type LoginDetails struct {
	Email    string
	Password string
}

type PageData struct {
	PageTitle string
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
		session, _ := store.Get(r, "cookie-users")

		// Session: set the value
		session.Values["user-admin-authenticated"] = false
		session.Save(r, w)

		details := LoginDetails{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		if details.Email == "info@marcovaleri.net" && details.Password == "S!lver09" {
			// Set session true
			session.Values["user-admin-authenticated"] = true
			session.Save(r, w)
		} else {
			// Set session false
			session.Values["user-admin-authenticated"] = false
			session.Save(r, w)
		}

		data := PageData{
			PageTitle: "Login Page",
		}

		tmpl.Execute(w, data)

	})
}
