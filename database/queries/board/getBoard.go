package queries

import (
	"database/sql"

	"github.com/bodatomas/gopi/database"
)

func GetBoardByID(targetId int) *sql.Row {
	db := database.GetDatabase()
	if db != nil {
		sqlStatement := `SELECT * FROM boards WHERE id = $1;`
		row := db.Conn.QueryRow(sqlStatement, targetId)
		return row
	}
	return nil
}
