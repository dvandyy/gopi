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

func GetBoardByID(targetId int) *sql.Row {
	db := database.GetDatabase()
	if db != nil {
		row := db.Conn.QueryRow(GetBoardByIdSQL, targetId)
		return row
	}
	return nil
}
