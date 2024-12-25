package controllers

import (
	errorcont "Forum/back-end/controllers/error"
	post "Forum/back-end/controllers/user/post"
	utils "Forum/back-end/controllers/utils"

	"database/sql"
	"fmt"
	"net/http"
)

func PostController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
			if r.URL.Path == "/ajouter-post" {
				// Fetch categories from the database
				db, err := sql.Open("sqlite3", "./back-end/database/database.db")
				if err != nil {
					fmt.Println("error opening database")
					errorcont.ErrorController(w, r, http.StatusInternalServerError)
					return
				}
				defer db.Close()
		
				categories := []Category{}
				rows, err := db.Query("SELECT id, name FROM categories")
				if err != nil {
					fmt.Println("error quer")
					fmt.Println("Error fetching categories:", err)
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
		
				data := struct {
					Categories []Category // Pass categories to the template
				}{
					Categories: categories,
				}
		
				utils.TemplateController(w, r, "/user/AjouterPost", data) // Pass data to template
			}
	}else if r.Method == "POST"{
		    if r.URL.Path == "/add-post" {
                fmt.Println("add post")
                post.AjouterPost(w,r);
                return
            }else if r.URL.Path == "/reaction"{
				fmt.Println("reaction")
				post.PostReactionController(w, r);
			}else if r.URL.Path == "/add-comment"{
				fmt.Println("add comment")
                post.CommentController(w, r);
                return
			}
	}else{
		errorcont.ErrorController(w, r, http.StatusMethodNotAllowed)
        return
	}

}
