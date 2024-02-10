package queries

import (
	"database/sql"

	"github.com/dvandyy/gopi/database"
)

/*
**
Get user with unique ID.
**
*/
const GetUserByIDSQL = `
SELECT unique_id, first_name, last_name, email, role FROM users WHERE unique_id = $1;
`

func GetUserByID(id string) *sql.Row {
	db := database.GetDatabase()
	if db != nil {
		row := db.Conn.QueryRow(GetUserByIDSQL, id)
		return row
	}
	return nil
}

/*
**
Create new user.
**
*/
const NewUserSQL = `
INSERT INTO users (unique_id, email, password) VALUES ($1, $2, $3)
`

func NewUser(unique_id string, email string, password string) error {
	db := database.GetDatabase()
	if db != nil {
		_, err := db.Conn.Exec(NewUserSQL, unique_id, email, password)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
**
Get user credentials.
**
*/
const GetUserCredentialsSQL = `
SELECT unique_id, password FROM users WHERE email = $1;
`

func GetUserCredentials(email string) *sql.Row {
	db := database.GetDatabase()
	if db != nil {
		row := db.Conn.QueryRow(GetUserCredentialsSQL, email)
		return row
	}
	return nil
}
