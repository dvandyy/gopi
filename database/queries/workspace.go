package queries

import (
	"database/sql"

	"github.com/dvandyy/gopi/database"
)

/*
**
Create new workspace.
**
*/
const NewWorkspaceSQL = `
INSERT INTO workspaces (unique_id, owner, name) VALUES ($1, $2, $3)
`
const NewWorkspaceMemberSQL = `
INSERT INTO workspace_members (workspace_id, user_id) VALUES ($1, $2)
`

func NewWorkspace(unique_id string, owner string, name string) error {
	db := database.GetDatabase()
	if db != nil {
		// Create new workspace
		_, errWorkspace := db.Conn.Exec(NewWorkspaceSQL, unique_id, owner, name)
		if errWorkspace != nil {
			return errWorkspace
		}
		// Add owner to that workspace
		_, errMember := db.Conn.Exec(NewWorkspaceMemberSQL, unique_id, owner)
		if errMember != nil {
			return errMember
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
SELECT  workspaces.unique_id, name FROM workspaces JOIN workspace_members
ON workspace_members.workspace_id = workspaces.unique_id JOIN users
ON workspace_members.user_id = users.unique_id
WHERE user_id = $1
`

func GetWorkspacesForUser(unique_id string) *sql.Rows {
	db := database.GetDatabase()
	if db != nil {
		rows, err := db.Conn.Query(GerWorkspacesForUserSQL, unique_id)
		if err != nil {
			return nil
		}
		return rows
	}
	return nil
}
