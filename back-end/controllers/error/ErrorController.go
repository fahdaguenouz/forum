package controllers

import (
	"Forum/back-end/models"
	"html/template"
	"net/http"
)

func ErrorController(w http.ResponseWriter, r *http.Request, statusCode int) {
    // Set the status code manually before rendering the error page
    w.WriteHeader(statusCode)

    // Parse the error template
    tmp, err := template.ParseFiles("Front-end/views/error/error.html")
    if err != nil {
        // If parsing fails, send an internal server error without calling WriteHeader again
        http.Error(w, err.Error()+"fahd", http.StatusInternalServerError)
        return
    }

    // Define the error type to pass to the template
    errorType := models.Error{
        StatusCode:   statusCode,
        ErrorMessage: http.StatusText(statusCode),
    }

    // Render the template with the error message
    if err := tmp.Execute(w, errorType); err != nil {
        // If rendering fails, send an internal server error without calling WriteHeader again
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}