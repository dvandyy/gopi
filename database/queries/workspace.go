package queries

import (
	"github.com/bodatomas/gopi/database"
)

/*
**
Create new workspace.
**
*/
const NewWorkspaceSQL = `
INSERT INTO workspaces (unique_id, owner, name) VALUES ($1, $2, $3)
`

func NewWorkspace(unique_id string, owner string, name string) error {
	db := database.GetDatabase()
	if db != nil {
		_, err := db.Conn.Exec(NewWorkspaceSQL, unique_id, owner, name)
		if err != nil {
			return err
		}
	}
	return nil
}
