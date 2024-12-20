package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {
	// Allow only POST requests
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON data
	var req struct {
		Username        string `json:"username"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Username == "" || req.Email == "" || req.Password == "" || req.ConfirmPassword == "" {
		http.Error(w, "Invalid or missing fields", http.StatusBadRequest)
		return
	}

	// Validate password match
	if req.Password != req.ConfirmPassword {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword := HashPassword(req.Password)

	// Open database
	db, err := sql.Open("sqlite3", "./back-end/database/database.db")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Check if the username already exists
	var usernameExists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", req.Username).Scan(&usernameExists)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if usernameExists {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	// Check if the email already exists
	var emailExists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", req.Email).Scan(&emailExists)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if emailExists {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	// Insert new user into the database
	_, err = db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", req.Username, req.Email, hashedPassword)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Generate session token
	sessionToken := GenerateSessionToken()
	expiresAt := time.Now().Add(24 * time.Hour)

	// Store session token in database
	_, err = db.Exec("INSERT INTO sessions (session_token, user_id, expires_at) VALUES (?, (SELECT id FROM users WHERE username = ?), ?)", sessionToken, req.Username, expiresAt)
	if err != nil {
		http.Error(w, "Error creating session", http.StatusInternalServerError)
		return
	}

	// Set session token as a secure cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  expiresAt,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}
