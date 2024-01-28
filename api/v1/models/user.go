package models

import (
	"github.com/bodatomas/gopi/database/queries"
)

type User struct {
	ID         uint    `json:"id" example:"1"`
	Unique_id  string  `json:"unique_id" example:"u-1706453613063fa2eb4"`
	First_Name *string `json:"first_name" example:"FirstName"`
	Last_Name  *string `json:"last_name" example:"LastName"`
	Email      string  `json:"email" example:"email@email.com"`
	Password   string  `json:"password" example:"hash"`
	Role       *string `json:"role" example:"Tester"`
	Created_at string  `json:"created_at" example:"2024-01-22 17:03:50.283466+00"`
}

type RegisterRequest struct {
	Email    string `json:"email" example:"email@email.com"`
	Password string `json:"password" example:"Password123&"`
}

type RegisterResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Get user with unique id
func GetUserByID(id string) (User, error) {
	var user User

	// Get user from databse
	row := queries.GetUserByID(id)
	err := row.Scan(
		&user.ID,
		&user.Unique_id,
		&user.First_Name,
		&user.Last_Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.Created_at)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// Create new user in database
func CreateNewUser(unique_id string, email string, password string) error {
	return queries.NewUser(unique_id, email, password)
}
