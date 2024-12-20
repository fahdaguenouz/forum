package controllers

import (
	errorcont "Forum/back-end/controllers/error"
	utils "Forum/back-end/controllers/utils"
	"database/sql"
	"net/http"
	"time"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorcont.ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		errorcont.ErrorController(w, r, http.StatusNotFound)
		return
	}
	// Check if the user is authenticated
	cookie, err := r.Cookie("session_token")
	if err == nil && cookie.Value != "" {
		// Open the database to verify the session token
		db, err := sql.Open("sqlite3", "./back-end/database/database.db")
		if err != nil {
			errorcont.ErrorController(w, r, http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var userID int
		err = db.QueryRow("SELECT user_id FROM sessions WHERE session_token = ? AND expires_at > ?", cookie.Value, time.Now()).Scan(&userID)
		if err == nil {
			// Redirect to the authenticated home page if the session is valid
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}
	}
	data := ""
	utils.TemplateController(w, r, "/guest/home", data)

}
