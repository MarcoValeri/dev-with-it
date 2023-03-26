package adminControllers

import (
	"net/http"
	"text/template"

	models "devwithit/models"
)

type NewUserDetails struct {
	Email    string
	Password string
	Submit   string
}

func AdminAddUsers() {
	tmpl := template.Must(template.ParseFiles("./views/admin/addUsers.html", "./views/includes/head.html", "./views/includes/admin-sidebar.html"))

	http.HandleFunc("/admin/add-users", func(w http.ResponseWriter, r *http.Request) {

		// Session: get it
		session, _ := store.Get(r, "cookie-users")

		if session.Values["user-admin-authenticated"] == true {

			// Form Data
			details := NewUserDetails{
				Email:    r.FormValue("email"),
				Password: r.FormValue("password"),
				Submit:   r.FormValue("add-new-user"),
			}

			// If the form has been submitted, process the data
			if details.Submit == "Add new user" {
				models.UserNew(details.Email, details.Password)
				http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
			}

			data := PageData{
				PageTitle: "Admin Add users",
			}

			tmpl.Execute(w, data)

		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}

	})
}
