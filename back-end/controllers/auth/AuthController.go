package auth

import (
	utils "Forum/back-end/controllers/utils"
	errorcont "Forum/back-end/controllers/error"

	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path == "/login" {
			errorcont.ErrorController(w,r, http.StatusNotFound)
			 return
		 }
	// Check if the user is already logged in
	cookie, err := r.Cookie("session_token")
	if err == nil && cookie.Value != "" {
		// Open SQLite database
		db, err := sql.Open("sqlite3", "./back-end/database/database.db")
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Validate the session token
		var expiresAtString string
		err = db.QueryRow("SELECT expires_at FROM sessions WHERE session_token = ?", cookie.Value).Scan(&expiresAtString)
		if err == nil {
			// Parse the expires_at value
			expiresAt, parseErr := time.Parse("2006-01-02 15:04:05.999999999-07:00", expiresAtString)
			if parseErr == nil && expiresAt.After(time.Now()) {
				// If session is valid, redirect to the home page
				http.Redirect(w, r, "/home", http.StatusFound)
				return
			}
		}
	}
		// Render the authentication page
		utils.TemplateController(w, r, "/auth/Auth", nil)
	case "POST":
		
		// Parse login data
		var req LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil || req.Username == "" || req.Password == "" {
			errorcont.ErrorController(w, r, http.StatusBadRequest)
			return
		}
		
		// Open SQLite database
		db, err := sql.Open("sqlite3", "./back-end/database/database.db")
		if err != nil {
			fmt.Println("row: ",err)
			errorcont.ErrorController(w,r, http.StatusInternalServerError)
			return
		}
		defer db.Close()
		
		// Check credentials
		var userID int

		err1 := db.QueryRow("SELECT id FROM users WHERE username = ? AND password = ?", req.Username, req.Password).Scan(&userID)
		if err1 != nil {
			
			errorcont.ErrorController(w,r, http.StatusUnauthorized)
			return
		}

		// Generate a session token
		sessionToken := generateSessionToken()
		expiresAt := time.Now().Add(24 * time.Hour)

		// Store the session token in the database
		_, err3 := db.Exec("INSERT INTO sessions (user_id, session_token, expires_at) VALUES (?, ?, ?)", userID, sessionToken, expiresAt)
		if err3 != nil {
			fmt.Println(err3)
			errorcont.ErrorController(w,r, http.StatusInternalServerError)
			return
		}

		// Set session cookie
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   sessionToken,
			Expires: expiresAt,
		})

		// Redirect to authenticated home page
		// http.Redirect(w, r, "/home", http.StatusAlreadyReported)
		utils.TemplateController(w, r, "/user/AuthHome", nil)
	default:
		errorcont.ErrorController(w,r, http.StatusMethodNotAllowed)
	}

}

func generateSessionToken() string {
	return "random-session-token" // Replace this with secure token generation
}
