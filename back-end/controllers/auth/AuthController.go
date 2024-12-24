package auth

import (
	errorcont "Forum/back-end/controllers/error"
	utils "Forum/back-end/controllers/utils"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	"database/sql"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)



func AuthController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path == "/login" {
			errorcont.ErrorController(w, r, http.StatusNotFound)
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
		if r.URL.Path == "/login" {
			fmt.Println("login")
			LginController(w, r)
			return
		} else if r.URL.Path == "/register" {
			fmt.Println("register")
			RegisterController(w, r)
			return
		}

	default:
		errorcont.ErrorController(w, r, http.StatusMethodNotAllowed)
		return 
	}

}

func GenerateSessionToken() string {
	randomBytes := make([]byte, 16) // 16 bytes = 32 hex characters
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatalf("Failed to generate random session token: %v", err)
	}
	return hex.EncodeToString(randomBytes)
}
func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password)) // Hash the password directly
	return hex.EncodeToString(hash.Sum(nil))
}
