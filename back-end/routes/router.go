package routes

import (
	guest "Forum/back-end/controllers/guest"
	user "Forum/back-end/controllers/user"
	utils "Forum/back-end/controllers/utils"


	auth "Forum/back-end/controllers/auth"
	"fmt"
	"net/http"
)

func Router() {
	http.HandleFunc("/", guest.HomeController)
	http.HandleFunc("/static/", utils.StaticController)
	http.HandleFunc("/authentification", auth.AuthController)
	http.HandleFunc("/home", user.AuthHomeController)
	http.HandleFunc("/login", auth.AuthController)


	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("Front-end/static"))))
	fmt.Println("Server running on http://localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}

}
