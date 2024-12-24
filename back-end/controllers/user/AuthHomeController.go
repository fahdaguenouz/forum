package controllers

import (
	controllers "Forum/back-end/controllers/error"
	utils "Forum/back-end/controllers/utils"
	 errorcont "Forum/back-end/controllers/error"

	"database/sql"
	"fmt"
	"net/http"
	"time"
)

func AuthHomeController(w http.ResponseWriter, r *http.Request) {

	// Get the session token from the cookie
	cookie, err := r.Cookie("session_token")
	fmt.Println("find :", cookie)
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/authentification", http.StatusFound)
		return
	}
	// Check the session token in the database
	db, err := sql.Open("sqlite3", "./back-end/database/database.db")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		errorcont.ErrorController(w,r,http.StatusInternalServerError)

		return
	}
	defer db.Close()

	var expiresAtString string
	err = db.QueryRow("SELECT expires_at FROM sessions WHERE session_token = ?", cookie.Value).Scan(&expiresAtString)
	fmt.Println("err:", err)
	if err != nil {
		http.Redirect(w, r, "/authentification", http.StatusFound)
		return
	}
	// Parse the `expires_at` value to `time.Time`
	expiresAt, err := time.Parse("2006-01-02 15:04:05.999999999-07:00", expiresAtString)
	if err != nil || expiresAt.Before(time.Now()) {
		http.Redirect(w, r, "/authentification", http.StatusFound)
		return
	}

	if r.Method != "GET" {
		controllers.ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/home" {
		controllers.ErrorController(w, r, http.StatusNotFound)
		return
	}
	data := ""

	utils.TemplateController(w, r, "/user/AuthHome", data)

}
