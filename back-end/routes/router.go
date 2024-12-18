package routes

import (
	"Forum/back-end/controllers"
	"Forum/back-end/controllers/auth"
	"fmt"
	"net/http"
)

func Router() {
	http.HandleFunc("/", controllers.HomeController)
	http.HandleFunc("/static/", controllers.StaticController)
	http.HandleFunc("/authentification", auth.AuthController)

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("Front-end/static"))))
	fmt.Println("Server running on http://localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}

}
