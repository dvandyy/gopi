package queries

import (
	"github.com/bodatomas/gopi/database"
)

func NewBoard(name string, description string) {
	db := database.GetDatabase()
	if db != nil {
		sqlStatement := `INSERT INTO boards (name, description) VALUES ($1, $2)`
		_, err := db.Conn.Exec(sqlStatement, name, description)
		if err != nil {
			panic(err)
		}
	}
}
