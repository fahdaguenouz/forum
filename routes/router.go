package routes

import (
	"Forum/controllers"
	"fmt"
	"net/http"
)

func Router() {
	http.HandleFunc("/", controllers.HomeController)

	fmt.Println("Server running on http://localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}

}
