package auth

import (
	"Forum/back-end/controllers"
	"net/http"
)

func AuthController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		controllers.ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/authentification"{
		controllers.ErrorController(w, r, http.StatusNotFound)
        return
	}
	controllers.TemplateController(w, r, "/auth/Auth", nil)
}
