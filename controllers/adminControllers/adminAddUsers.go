package adminControllers

import (
	"log"
	"net/http"
	"text/template"

	models "devwithit/models"
	util "devwithit/util"
)

type NewUserDetails struct {
	Email    string
	Password string
	Submit   string
}

var userError bool = false

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
				// Add the user if the email is not already in the db
				flag := models.CheckUserEmail(details.Email)
				if !flag {
					passwordHashed, err := util.HashPassword(details.Password)
					if err != nil {
						log.Fatal(err)
					}
					models.UserNew(details.Email, passwordHashed)
					http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
				} else {
					userError = true
				}
			}

			data := PageAdminData{
				PageTitle: "Admin Add users",
				UserError: userError,
			}

			tmpl.Execute(w, data)

		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}

	})
}
