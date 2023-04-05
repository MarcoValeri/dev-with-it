package adminControllers

import (
	"devwithit/models"
	"net/http"
	"text/template"
)

type PageAdminUser struct {
	PageTitle     string
	AllAdminUsers []models.UserAdmin
}

func Users() {
	tmpl := template.Must(template.ParseFiles("./views/admin/users.html", "./views/includes/head.html", "./views/includes/admin-sidebar.html"))

	http.HandleFunc("/admin/users", func(w http.ResponseWriter, r *http.Request) {
		// Session: get it
		session, _ := store.Get(r, "cookie-users")

		if session.Values["user-admin-authenticated"] == true {

			// Get all user data
			allAdminUsers := models.GetAllUsers()

			data := PageAdminUser{
				PageTitle:     "Admin Users",
				AllAdminUsers: allAdminUsers,
			}

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	})
}
