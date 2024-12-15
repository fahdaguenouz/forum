package auth

import (
	"Forum/back-end/controllers"
	"net/http"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		controllers.ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/register" {
		controllers.ErrorController(w, r, http.StatusNotFound)
		return
	}
	controllers.TemplateController(w, r, "/auth/register", nil)
}
