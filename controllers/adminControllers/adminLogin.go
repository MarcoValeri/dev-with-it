package adminControllers

import (
	"devwithit/models"
	util "devwithit/util"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

type LoginDetails struct {
	Email    string
	Password string
}

type PageAdminData struct {
	PageTitle string
	UserError bool
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

		// Check User Login
		if len(details.Email) > 0 && len(details.Password) > 0 {
			_, userPswHashed := models.UserData(details.Email)

			flag := util.CheckPasswordHash(details.Password, userPswHashed)
			if flag {
				// Set session true
				session.Values["user-admin-authenticated"] = true
				session.Save(r, w)
				http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
			} else {
				// Set session false
				session.Values["user-admin-authenticated"] = false
				session.Save(r, w)
			}
		}

		data := PageAdminData{
			PageTitle: "Login Page",
			UserError: false,
		}

		tmpl.Execute(w, data)

	})
}
