package controllers

import (
	"Forum/back-end/models"
	"html/template"
	"net/http"
)

func ErrorController(w http.ResponseWriter, r *http.Request, statusCode int) {
	tmp, err := template.ParseFiles("Front-end/views/error/error.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	errorType := models.Error{
		StatusCode:   statusCode,
		ErrorMessage: http.StatusText(statusCode),
	}
	w.WriteHeader(statusCode)
	// fmt.Println("rrrrrrr")
	if err := tmp.Execute(w, errorType); err != nil {
		// fmt.Println("error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
