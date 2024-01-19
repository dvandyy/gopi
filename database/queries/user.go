package queries

import (
	"github.com/bodatomas/gopi/database"
	"github.com/google/uuid"
)

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
