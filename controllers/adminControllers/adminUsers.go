package adminControllers

import (
	"net/http"
	"text/template"
)

func Users() {
	tmpl := template.Must(template.ParseFiles("./views/admin/users.html", "./views/includes/head.html", "./views/includes/admin-sidebar.html"))

	http.HandleFunc("/admin/users", func(w http.ResponseWriter, r *http.Request) {
		// Session: get it
		session, _ := store.Get(r, "cookie-users")

		if session.Values["user-admin-authenticated"] == true {

			data := PageData{
				PageTitle: "Admin Users",
			}

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	})
}
