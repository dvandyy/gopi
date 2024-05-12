package queries

import (
	"database/sql"
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

/*
**
Get teams for user in workspace.
**
*/
const GerTeamsForUserSQL = `
SELECT teams.unique_id, teams.name
FROM teams
JOIN team_members ON teams.unique_id = team_members.team_id
WHERE team_members.user_id = $1
AND teams.workspace_id = $2;
`

func GetTeamsInWorkspaceForUser(user_id string, workspace_id string) *sql.Rows {
	db := database.GetDatabase()
	if db != nil {
		rows, err := db.Conn.Query(GerTeamsForUserSQL, user_id, workspace_id)
		if err != nil {
			return nil
		}
		return rows
	}
	return nil
}

/*
**
Add user to team.
**
*/
const AddUserToTeamSQL = `
INSERT INTO team_members (team_id, user_id) VALUES ($1, $2)
`

func AddUserToTeam(user_id string, team_id string) error {
	db := database.GetDatabase()
	if db != nil {
		_, errMember := db.Conn.Exec(AddUserToTeamSQL, team_id, user_id)
		if errMember != nil {
			return errMember
		}
	}
	return nil
}
