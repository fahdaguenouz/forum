package auth

import (
	"database/sql"
	"net/http"
	errorcont "Forum/back-end/controllers/error"
)

func LogoutController(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
       errorcont.ErrorController(w,r, http.StatusNotFound)
        return
    }

	cookie, err := r.Cookie("session_token")
	if err != nil {
		// If there's no session token, redirect to the login page
		http.Redirect(w, r, "/authentification", http.StatusFound)
		return
	}

	sessionToken := cookie.Value

	// Open SQLite database
	db, err := sql.Open("sqlite3", "./back-end/database/database.db")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Delete the session from the database
	_, err = db.Exec("DELETE FROM sessions WHERE session_token = ?", sessionToken)
	if err != nil {
		errorcont.ErrorController(w, r, http.StatusInternalServerError)
		return
	}

	// Clear the session token from the cookies
	http.SetCookie(w, &http.Cookie{
		Name:   "session_token",
		Value:  "",
		MaxAge: -1, // Expire the cookie immediately
	})

	// Redirect to the login page
	http.Redirect(w, r, "/authentification", http.StatusFound)
}
