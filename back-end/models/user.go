package models



type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Global variable to hold the currently logged-in user
var CurrentUser *User