package controllers

import (
	"Forum/models"
	"html/template"
	"net/http"
)

func ErrorController(w http.ResponseWriter, r *http.Request, statusCode int) {
	tmp, err := template.ParseFiles("views/error.html")
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	errorType:=models.Error{
		StatusCode: statusCode,
		ErrorMessage: http.StatusText(statusCode),
	}
	w.WriteHeader(statusCode)
	if err:=tmp.Execute(w,errorType);err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}
}
