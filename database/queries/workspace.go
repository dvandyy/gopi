package queries

import (
	"database/sql"

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

/*
**
Return all workspaces for certain user.
**
*/
const GerWorkspacesForUserSQL = `
SELECT workspaces.unique_id, name FROM workspaces JOIN workspace_members
ON workspace_members.workspace_id = workspaces.unique_id JOIN users
ON workspace_members.user_id = users.unique_id
WHERE user_id = (user_id) VALUES ($1)
`

func GetWorkspacesForUser(unique_id string) *sql.Row {
	db := database.GetDatabase()
	if db != nil {
		row := db.Conn.QueryRow(GerWorkspacesForUserSQL, unique_id)
		return row
	}
	return nil
}
