// Author: James Mallon <jamesmallondev@gmail.com>
// routers package -
package routers

import (
	. "GoAuthorization/controllers"
	. "GoAuthorization/middlewares"
	"net/http"
)

// LoadRoutes function -
func LoadAdminRoutes() {
	am := AuthMiddleware()
	aclm := ACLMiddleware()

	http.HandleFunc("/dashboard", am.CheckLogged(AdminController().Dashboard))
	http.HandleFunc("/dashboard/users", am.CheckLogged(aclm.Authorized("superuser", AdminController().Users)))
}
