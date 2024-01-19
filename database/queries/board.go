package queries

import (
	"database/sql"

	"github.com/bodatomas/gopi/database"
)

/*
**
Get board with specific ID.
**
*/
const GetBoardByIdSQL = `
SELECT * FROM boards WHERE id = $1;
`

func GetBoardByID(id string) *sql.Row {
	db := database.GetDatabase()
	if db != nil {
		row := db.Conn.QueryRow(GetBoardByIdSQL, id)
		return row
	}
	return nil
}

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
