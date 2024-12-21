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
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
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
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
func HomeController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorcont.ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		errorcont.ErrorController(w, r, http.StatusNotFound)
		return
	}
	// Check if the user is authenticated
	db, err := sql.Open("sqlite3", "./back-end/database/database.db")
	if err != nil {
		errorcont.ErrorController(w, r, http.StatusInternalServerError)
		return
	}
	defer db.Close()
	cookie, err := r.Cookie("session_token")
	if err == nil && cookie.Value != "" {
		// Open the database to verify the session token
		var userID int
		err = db.QueryRow("SELECT user_id FROM sessions WHERE session_token = ? AND expires_at > ?", cookie.Value, time.Now()).Scan(&userID)
		if err == nil {
			// Redirect to the authenticated home page if the session is valid
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}
	}
	// Fetch all categories
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

	// Fetch posts with their categories and like counts
	posts := []Post{}
	query := `
    SELECT 
    p.id, 
    p.title, 
    p.content, 
    c.name AS category, 
    COALESCE(SUM(CASE WHEN pr.reaction = 'like' THEN 1 ELSE 0 END), 0) AS likes,
    COALESCE(SUM(CASE WHEN pr.reaction = 'dislike' THEN 1 ELSE 0 END), 0) AS dislikes,
    p.created_at
	FROM posts p
	LEFT JOIN post_categories pc ON p.id = pc.post_id
	LEFT JOIN categories c ON pc.category_id = c.id
	LEFT JOIN post_reactions pr ON p.id = pr.post_id
	GROUP BY p.id, c.name
	ORDER BY p.created_at DESC;
	`
	rowss, err := db.Query(query)
	if err != nil {
		fmt.Println("select rows: ", err)
		errorcont.ErrorController(w, r, http.StatusInternalServerError)
		return
	}
	defer rowss.Close()

	for rowss.Next() {
		var post Post
		if err := rowss.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &post.Likes, &post.Dislikes, &post.CreatedAt); err != nil {
			fmt.Println("row scan posts:", err)
			errorcont.ErrorController(w, r, http.StatusInternalServerError)
			return
		}
	
		// Check if the post data looks correct
		fmt.Printf("Post: %+v\n", post)
	
		// Fetch comments for the post
		comments := []Comment{}
		commentQuery := `
			SELECT id, content, created_at 
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
	
		// Debugging log for comment fetching
		fmt.Printf("Fetching comments for post ID %d...\n", post.ID)

		for commentRows.Next() {
			var comment Comment
			if err := commentRows.Scan(&comment.ID, &comment.Content, &comment.CreatedAt); err != nil {
				fmt.Println("row scan comments:", err)
				errorcont.ErrorController(w, r, http.StatusInternalServerError)
				return
			}
			// Log each comment fetched
			fmt.Printf("Fetched comment: %+v\n", comment)
			comments = append(comments, comment)
		}
	
		post.Comments = comments
		posts = append(posts, post)
		fmt.Println("Posts with comments:", posts)
	}

	// Pass posts and categories to the template
	data := struct {
        Posts    []Post
        Categories []Category
    }{
        Posts:     posts,
        Categories: categories,
    }
	utils.TemplateController(w, r, "/guest/home", data)
}