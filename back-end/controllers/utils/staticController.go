package controllers

import (
	controllers "Forum/back-end/controllers/error"
	"fmt"
	"net/http"
	"os"
)

func StaticController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		controllers.ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}

	filePath := "Front-end/static/" + r.URL.Path[len("/static/"):]

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File does not exist:", filePath)
		controllers.ErrorController(w, r, http.StatusNotFound)
		return
	}

	fs := http.Dir("Front-end/static")
	http.StripPrefix("/static/", http.FileServer(fs)).ServeHTTP(w, r)
}
