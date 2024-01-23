package queries

import (
	"database/sql"

	"github.com/bodatomas/gopi/database"
	"github.com/google/uuid"
)

/*
**
Get user with unique ID.
**
*/
const GetUserByIDSQL = `
SELECT * FROM users WHERE unique_id = $1;
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

func NewUser(email string, password string) error {
	db := database.GetDatabase()
	if db != nil {
		_, err := db.Conn.Exec(NewUserSQL, uuid.New(), email, password)
		if err != nil {
			return err
		}
	}
	return nil
}
