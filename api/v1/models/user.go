package models

import (
	"github.com/bodatomas/gopi/database/queries"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type User struct {
	ID         uint   `json:"id"`
	Unique_id  string `json:"unique_id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	Created_at string `json:"created_at"`
}

// Create new user in database
func CreateNewUser(email string, password string) error {
	return queries.NewUser(email, password)
}
