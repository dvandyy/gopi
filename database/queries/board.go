package queries

import (
	"database/sql"

	"github.com/bodatomas/gopi/database"
	"github.com/google/uuid"
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
INSERT INTO boards (unique_id, title) VALUES ($1, $2)
`

func NewBoard(title string) error {
	db := database.GetDatabase()
	if db != nil {
		_, err := db.Conn.Exec(NewBoardSQL, uuid.New(), title)
		if err != nil {
			return err
		}
	}
	return nil
}
