package controllers

import "net/http"

func HomeController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" { 
		ErrorController(w,r,http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/"{
		ErrorController(w,r,http.StatusNotFound)
		return
	}
	data:=""
	TemplateController(w,r,"home",data)

}
