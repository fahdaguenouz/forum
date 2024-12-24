package post

import (
	"database/sql"
	"net/http"
	"time"
)

func AjouterPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the session token from the cookie
	cookie, err := r.Cookie("session_token")
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/authentification", http.StatusFound)
		return
	}

	// Check the session token in the database to retrieve user ID
	db, err := sql.Open("sqlite3", "./back-end/database/database.db")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var userID int
	var expiresAtString string

	err = db.QueryRow("SELECT user_id, expires_at FROM sessions WHERE session_token = ?", cookie.Value).Scan(&userID, &expiresAtString)
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

	// Parse form data
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	// Get selected categories
	categories := r.Form["categories"] // This will be an array of selected category IDs

	// Insert post into database
	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Transaction error", http.StatusInternalServerError)
		return
	}

	var postID int
	err = tx.QueryRow("INSERT INTO posts (title, content, user_id) VALUES (?, ?, ?) RETURNING id", title, content, userID).Scan(&postID)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Error inserting post", http.StatusInternalServerError)
		return
	}

	for _, categoryID := range categories {
		if _, err = tx.Exec("INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)", postID, categoryID); err != nil {
			tx.Rollback()
			http.Error(w, "Error inserting post-category relation", http.StatusInternalServerError)
			return
		}
	}

	if err = tx.Commit(); err != nil {
		http.Error(w, "Transaction commit error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/home", http.StatusFound)
}
