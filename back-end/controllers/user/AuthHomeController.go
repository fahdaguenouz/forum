package controllers

import (
	errorcont "Forum/back-end/controllers/error"
	utils "Forum/back-end/controllers/utils"

	"database/sql"
	"fmt"
	"net/http"
	"time"
)

type Comment struct {
	ID        int    `json:"id"`
	PostID    int    `json:"post_id"`
	UserID    int    `json:"user_id"` // Add UserID to associate with the user
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	Username  string `json:"username"` // Add Username for display
}

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
	CreatedAt string    `json:"created_at"`
	Comments  []Comment `json:"comments"`
	Username  string    `json:"username"` // Add Username for display
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}


func AuthHomeController(w http.ResponseWriter, r *http.Request) {
	// Get the session token from the cookie
	cookie, err := r.Cookie("session_token")
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/authentification", http.StatusFound)
		return
	}

	// Check the session token in the database
	db, err := sql.Open("sqlite3", "./back-end/database/database.db")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var userID int
	var expiresAtString string

	// Fetch user ID and expiration time based on session token
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

	if r.Method != "GET" {
		errorcont.ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/home" {
		errorcont.ErrorController(w, r, http.StatusNotFound)
		return
	}

	categories := []Category{}
	rows, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		fmt.Println("categories:", err)
		errorcont.ErrorController(w, r, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			fmt.Println("row scan categories:", err)
			errorcont.ErrorController(w, r, http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}

	posts := []Post{}
	query := `
	SELECT 
		p.id, 
		p.title, 
		p.content, 
		c.name AS category,
        u.username AS username,
        COALESCE(SUM(CASE WHEN pr.reaction = 'like' THEN 1 ELSE 0 END), 0) AS likes,
        COALESCE(SUM(CASE WHEN pr.reaction = 'dislike' THEN 1 ELSE 0 END), 0) AS dislikes,
        p.created_at
	FROM posts p
    LEFT JOIN post_categories pc ON p.id = pc.post_id
    LEFT JOIN categories c ON pc.category_id = c.id
    LEFT JOIN post_reactions pr ON p.id = pr.post_id
    LEFT JOIN users u ON p.user_id = u.id -- Join with users table
	GROUP BY p.id
    ORDER BY p.created_at DESC;
    `
	rowss, err := db.Query(query) // Fetch all posts regardless of user
	if err != nil {
        fmt.Println("select rows: ", err)
        errorcont.ErrorController(w, r, http.StatusInternalServerError)
        return
    }
	defer rowss.Close()
	
	for rowss.Next() {
        var post Post
        if err := rowss.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &post.Username, &post.Likes, &post.Dislikes, &post.CreatedAt); err != nil {
            fmt.Println("row scan posts:", err)
            errorcont.ErrorController(w, r, http.StatusInternalServerError)
            return
        }
	
        comments := []Comment{}
        commentQuery := `
        SELECT 
            id,
            content,
            created_at,
            user_id,
            (SELECT username FROM users WHERE id = user_id) AS username -- Fetch username for each comment
        FROM comments 
        WHERE post_id = ?
        ORDER BY created_at ASC;
        `
		
        commentRows, err := db.Query(commentQuery, post.ID)
        if err != nil {
            fmt.Println("select comments:", err)
            errorcont.ErrorController(w, r, http.StatusInternalServerError)
            return
        }
        defer commentRows.Close()
	
        for commentRows.Next() {
            var comment Comment
            if err := commentRows.Scan(&comment.ID, &comment.Content, &comment.CreatedAt, &comment.UserID,&comment.Username); err != nil { // Include UserID in scan
                fmt.Println("row scan comments:", err)
                errorcont.ErrorController(w, r, http.StatusInternalServerError)
                return
            }
            
            // Fetch username for each comment separately
            var username string
            if err := db.QueryRow("SELECT username FROM users WHERE id = ?", comment.UserID).Scan(&username); err == nil {
                comment.Username = username // Set the username for each comment
            }
			
            comments = append(comments, comment)
        }
	
        post.Comments = comments
        posts = append(posts, post)
    }
	

    // Define data variable to pass to template
	data := struct {
        Posts     []Post
        Categories []Category
    }{
        Posts:     posts,
        Categories: categories,
    }

	utils.TemplateController(w, r, "/user/AuthHome", data) // Ensure data is passed correctly here.
}
