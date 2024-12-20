package controllers

import (
	controllers "Forum/back-end/controllers/error"
	"net/http"
	"text/template"
)

func TemplateController(w http.ResponseWriter, r *http.Request, temp string, data any) {
 res,err:=template.ParseFiles("Front-end/views/"+temp+".html")
 if err!=nil{
	controllers.ErrorController(w,r,http.StatusInternalServerError)
	return
 }
 if err=res.Execute(w, data); err != nil {
	controllers.ErrorController(w, r, http.StatusInternalServerError)
	return
}
}
