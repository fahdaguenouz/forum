package controllers

import (
	"Forum/back-end/models"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func ErrorController(w http.ResponseWriter, r *http.Request, statusCode int) {
	// Only call WriteHeader once
	w.WriteHeader(statusCode)

	// Build the absolute path to the template
	basePath, _ := os.Getwd() // Get the current working directory
	templatePath := filepath.Join(basePath, "Front-end", "views", "error", "error.html")

	// Parse the error template
	tmp, err := template.ParseFiles(templatePath)
	if err != nil {
        // If parsing fails, send an internal server error without calling WriteHeader again
        http.Error(w, err.Error(), http.StatusInternalServerError)
		
        return
    }

	// Create an error structure to pass to the template
	errorType := models.Error{
		StatusCode:   statusCode,
		ErrorMessage: http.StatusText(statusCode),
	}

	// Execute the template with the error data
	if err := tmp.Execute(w, errorType); err != nil {
        // If rendering fails, send an internal server error without calling WriteHeader again
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}