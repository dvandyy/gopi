package queries

import (
	"github.com/bodatomas/gopi/database"
)

/*
**
Create new board.
**
*/
const NewBoardSQL = `
INSERT INTO boards (name, description) VALUES ($1, $2)
`

func NewBoard(name string, description string) {
	db := database.GetDatabase()
	if db != nil {
		_, err := db.Conn.Exec(NewBoardSQL, name, description)
		if err != nil {
			panic(err)
		}
	}
}
