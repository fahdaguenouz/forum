package auth

import (
	utils "Forum/back-end/controllers/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func LginController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, `{"error": "Method not allowed."}`, http.StatusMethodNotAllowed)
		return
	}

	// Parse login data
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Username == "" || req.Password == "" {
		http.Error(w, `{"error": "Both username and password are required."}`, http.StatusBadRequest)
		return
	}

	// Open SQLite database
	db, err := sql.Open("sqlite3", "./back-end/database/database.db")
	if err != nil {
		fmt.Println("Database error:", err)
		http.Error(w, `{"error": "Internal server error."}`, http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Check if username exists
	var storedHashedPassword string
	var userID int
	err = db.QueryRow("SELECT id, password FROM users WHERE username = ?", req.Username).Scan(&userID, &storedHashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			// Username does not exist
			http.Error(w, `{"error": "Username does not exist."}`, http.StatusUnauthorized)
		} else {
			// Other database error
			fmt.Println("Query error:", err)
			http.Error(w, `{"error": "Internal server error."}`, http.StatusInternalServerError)
		}
		return
	}

	// Hash the provided password and compare with the stored hash
	inputHashedPassword := HashPassword(req.Password)
	if inputHashedPassword != storedHashedPassword {
		// Incorrect password
		http.Error(w, `{"error": "Incorrect password."}`, http.StatusUnauthorized)
		return
	}

	// Generate a session token
	sessionToken := GenerateSessionToken()
	expiresAt := time.Now().Add(24 * time.Hour)

	// Store the session token in the database
	_, err = db.Exec("INSERT INTO sessions (user_id, session_token, expires_at) VALUES (?, ?, ?)", userID, sessionToken, expiresAt)
	if err != nil {
		fmt.Println("Session creation error:", err)
		http.Error(w, `{"error": "Internal server error."}`, http.StatusInternalServerError)
		return
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  expiresAt,
		HttpOnly: true,
		Path:     "/",
	})

	// Redirect to authenticated home page
	utils.TemplateController(w, r, "/user/AuthHome", nil)

}
