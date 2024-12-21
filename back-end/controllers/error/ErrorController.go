package controllers

import (
	"Forum/back-end/models"
	"html/template"
	"net/http"
)

func ErrorController(w http.ResponseWriter, r *http.Request, statusCode int) {
    w.WriteHeader(statusCode)

    // Build the absolute path to the template
    basePath, _ := os.Getwd() // Get the current working directory
    templatePath := filepath.Join(basePath, "Front-end/views/error/error.html")

    // Parse the error template
    tmp, err := template.ParseFiles(templatePath)
    if err != nil {
        http.Error(w, err.Error()+"fahd", http.StatusInternalServerError)
        return
    }

    errorType := models.Error{
        StatusCode:   statusCode,
        ErrorMessage: http.StatusText(statusCode),
    }

    if err := tmp.Execute(w, errorType); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}    