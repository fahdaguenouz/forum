package controllers

import ("net/http"
		errorcont "Forum/back-end/controllers/error"
		utils "Forum/back-end/controllers/utils"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" { 
		errorcont.ErrorController(w,r,http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/"{
		errorcont.ErrorController(w,r,http.StatusNotFound)
		return
	}
	data:=""
	utils.TemplateController(w,r,"/guest/home",data)

}
