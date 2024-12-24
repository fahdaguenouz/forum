package controllers

import (
	controllers "Forum/back-end/controllers/error"
	"fmt"
	"net/http"
	"text/template"
)

func TemplateController(w http.ResponseWriter, r *http.Request, temp string, data any) {
	res, err := template.ParseFiles("Front-end/views/" + temp + ".html")
	if err != nil {
		fmt.Println("error parsing")
		controllers.ErrorController(w, r, http.StatusInternalServerError)
		return
	}
	if err = res.Execute(w, data); err != nil {
		fmt.Println("error executing template")
		controllers.ErrorController(w, r, http.StatusInternalServerError)
		return
	}
}
