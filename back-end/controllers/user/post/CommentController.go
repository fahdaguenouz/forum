package post

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"Forum/back-end/models" // Adjust import as necessary
)

type CommentRequest struct {
	PostID  int    `json:"post_id"`
	Content string `json:"content"`
}

func CommentController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CommentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Content == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Check if user is logged in
	if models.CurrentUser == nil {
		http.Error(w, "User not logged in", http.StatusUnauthorized)
		return
	}

	userID := models.CurrentUser.ID // Get current user's ID from global variable

	// Open SQLite database
	db, err := sql.Open("sqlite3", "./back-end/database/database.db")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Insert the new comment into the database
	_, err = db.Exec("INSERT INTO comments (post_id, user_id, content) VALUES (?, ?, ?)", req.PostID, userID, req.Content)
	if err != nil {
		http.Error(w, "Error adding your comment", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"username": models.CurrentUser.Username,
		"content":  req.Content,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response) // Send back the new comment data as JSON response
}
