package adminControllers

import (
	"net/http"
	"text/template"
)

func AdminDashboard() {
	tmpl := template.Must(template.ParseFiles("./views/admin/dashboard.html", "./views/includes/head.html", "./views/includes/admin-sidebar.html"))

	http.HandleFunc("/admin/dashboard", func(w http.ResponseWriter, r *http.Request) {

		// Session: get it
		session, _ := store.Get(r, "cookie-users")

		if session.Values["user-admin-authenticated"] == true {

			data := PageAdminData{
				PageTitle: "Admin Dashboard",
			}

			tmpl.Execute(w, data)

		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}

	})
}
