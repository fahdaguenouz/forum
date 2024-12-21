package controllers

import (
	"Forum/back-end/models"
	"html/template"
	"net/http"
)

func ErrorController(w http.ResponseWriter, r *http.Request, statusCode int) {
    tmp, err := template.ParseFiles("Front-end/views/error/error.html")
    if err != nil {
        // Set the status code manually for the error response
        w.WriteHeader(http.StatusInternalServerError)
        // Use http.Error to send the error message in the body
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Set the status code manually before rendering the error page
    w.WriteHeader(statusCode)

    errorType := models.Error{
        StatusCode:   statusCode,
        ErrorMessage: http.StatusText(statusCode),
    }

    // Render the template with the error message
    if err := tmp.Execute(w, errorType); err != nil {
        // If rendering fails, send an internal server error
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}