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
SELECT * FROM boards WHERE unique_id = $1;
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
INSERT INTO boards (unique_id, title, description) VALUES ($1, $2, $3)
`

func NewBoard(unique_id string, title string, description string) error {
	db := database.GetDatabase()
	if db != nil {
		_, err := db.Conn.Exec(NewBoardSQL, unique_id, title, description)
		if err != nil {
			return err
		}
	}
	return nil
}
