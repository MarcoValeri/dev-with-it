package controller

import (
	adminControllers "devwithit/controllers/adminControllers"
)

func AdminController() {
	adminControllers.Login()
	adminControllers.AdminDashboard()
}
