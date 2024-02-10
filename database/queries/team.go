package queries

import (
	"github.com/dvandyy/gopi/database"
)

/*
**
Create new team.
**
*/
const NewTeamSQL = `
INSERT INTO teams (unique_id, workspace_id, name) VALUES ($1, $2, $3)
`

func NewTeam(unique_id string, workspace_id string, name string) error {
	db := database.GetDatabase()
	if db != nil {
		_, err := db.Conn.Exec(NewTeamSQL, unique_id, workspace_id, name)
		if err != nil {
			return err
		}
	}
	return nil
}
