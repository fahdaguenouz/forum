package auth

import (
	errorcont "Forum/back-end/controllers/error"
	utils "Forum/back-end/controllers/utils"
	"Forum/back-end/models"

	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)



type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LginController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, `{"error": "Method not allowed."}`, http.StatusMethodNotAllowed)
		errorcont.ErrorController(w, r, http.StatusMethodNotAllowed)
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
	var email string // Add email to retrieve it later
	err = db.QueryRow("SELECT id, email, password FROM users WHERE username = ?", req.Username).Scan(&userID, &email, &storedHashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, `{"error": "Username does not exist."}`, http.StatusUnauthorized)
			return
		} else {
			fmt.Println("Query error:", err)
			http.Error(w, `{"error": "Internal server error."}`, http.StatusInternalServerError)
			return
		}
	}

	// Hash the provided password and compare with the stored hash
	inputHashedPassword := HashPassword(req.Password)
	if inputHashedPassword != storedHashedPassword {
		http.Error(w, `{"error": "Incorrect password."}`, http.StatusUnauthorized)
		return
	}

    // Set the global CurrentUser variable with user details
	models.CurrentUser = &models.User{
        ID:       userID,
        Username: req.Username,
        Email:    email,
    }

    // Generate a session token
	sessionToken := GenerateSessionToken()
	expiresAt := time.Now().Add(24 * time.Hour)

    // Store the session token in the database
    _, err = db.Exec("INSERT INTO sessions (user_id, session_token, expires_at) VALUES (?, ?, ?)", userID, sessionToken, expiresAt)
	if err != nil {
        fmt.Println("Session creation error:", err)
        http.Error(w, `{"error": "You are already logged in another device."}`, http.StatusInternalServerError)
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

    // Redirect to aut	henticated home page
	utils.TemplateController(w, r, "/user/AuthHome", nil)
}