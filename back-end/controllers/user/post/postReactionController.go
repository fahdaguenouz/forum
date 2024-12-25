package post

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"Forum/back-end/models" // Adjust import as necessary
)

type ReactionRequest struct {
	PostID   int    `json:"post_id"`
	Reaction string `json:"reaction"` // 'like' or 'dislike'
}

func PostReactionController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ReactionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || (req.Reaction != "like" && req.Reaction != "dislike") {
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

	// Check if a reaction already exists for this user and post
	var existingReaction string
	err = db.QueryRow("SELECT reaction FROM post_reactions WHERE user_id = ? AND post_id = ?", userID, req.PostID).Scan(&existingReaction)

	if err == sql.ErrNoRows {
		// No existing reaction, insert new one
		_, err = db.Exec("INSERT INTO post_reactions (user_id, post_id, reaction) VALUES (?, ?, ?)", userID, req.PostID, req.Reaction)
	} else if err == nil {
		if existingReaction == req.Reaction {
			// User clicked the same reaction again; delete it
			_, err = db.Exec("DELETE FROM post_reactions WHERE user_id = ? AND post_id = ?", userID, req.PostID)
			if err != nil {
				http.Error(w, "Error removing your reaction", http.StatusInternalServerError)
				return
			}
		} else {
			// Existing reaction found but different; update it
			_, err = db.Exec("UPDATE post_reactions SET reaction = ? WHERE user_id = ? AND post_id = ?", req.Reaction, userID, req.PostID)
			if err != nil {
				http.Error(w, "Error processing your reaction", http.StatusInternalServerError)
				return
			}
		}
	} else {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(w, "Error processing your reaction", http.StatusInternalServerError)
		return
	}

    // Fetch updated likes and dislikes count for this post
	var likesCount int
	var dislikesCount int

	err = db.QueryRow(`
	SELECT 
	    COALESCE(SUM(CASE WHEN reaction = 'like' THEN 1 ELSE 0 END), 0) AS likes,
	    COALESCE(SUM(CASE WHEN reaction = 'dislike' THEN 1 ELSE 0 END), 0) AS dislikes 
	FROM post_reactions 
	WHERE post_id = ?`, req.PostID).Scan(&likesCount, &dislikesCount)

	if err != nil {
	    http.Error(w, "Error fetching counts", http.StatusInternalServerError)
	    return
	}

	response := map[string]int{
	    "likes":     likesCount,
	    "dislikes":  dislikesCount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response) // Send back updated counts as JSON response
}
