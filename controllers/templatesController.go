package controllers

import (
	"net/http"
	"text/template"
)

func TemplateController(w http.ResponseWriter, r *http.Request, temp string, data any) {
 res,err:=template.ParseFiles("views/"+temp+".html")
 if err!=nil{
	ErrorController(w,r,http.StatusInternalServerError)
	return
 }
 if err=res.Execute(w, data); err != nil {
	ErrorController(w, r, http.StatusInternalServerError)
	return
}
}
