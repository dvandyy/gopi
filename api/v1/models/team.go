package models

import "github.com/dvandyy/gopi/database/queries"

type Team struct {
	ID           uint    `json:"id" example:"1"`
	Unique_id    string  `json:"unique_id" example:"u-1706453613063fa2eb4"`
	Name         *string `json:"name" example:"My team"`
	Workspace_id string  `json:"workspace_id" example:"w-1706453613063fa2eb4"`
	Created_at   string  `json:"created_at" example:"2024-01-22 17:03:50.283466+00"`
}

/*
**
Create team
**
*/
type CreateTeamRequest struct {
	Name         string `json:"name" example:"My team"`
	Workspace_id string `json:"workspace_id" example:"w-1706453613063fa2eb4"`
}

type CreateTeamResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Create new team in database
func CreateNewTeam(unique_id string, workspace_id string, name string) error {
	return queries.NewTeam(unique_id, workspace_id, name)
}

/*
**
Get team
**
*/
type TeamResponse struct {
	Unique_id string `json:"unique_id" example:"t-1706453613063fa2eb4"`
	Name      string `json:"name" example:"Team name"`
}

type UserTeamsResponse struct {
	Status int            `json:"status"`
	Teams  []TeamResponse `json:"teams"`
}

// Get all teams in workspace where user is assigned
func GetTeamsForUserInWorkspace(user_id string, workspace_id string) ([]TeamResponse, error) {
	var teams []TeamResponse

	rows := queries.GetTeamsInWorkspaceForUser(user_id, workspace_id)
	for rows.Next() {
		var team TeamResponse
		if err := rows.Scan(&team.Unique_id, &team.Name); err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	return teams, nil
}

/*
**
Add User to team
**
*/
type AddUserToTeamRequest struct {
	User_id string `json:"user_id" example:"u-1706453613063fa2eb4"`
	Team_id string `json:"team_id" example:"t-1706453613063fa2eb4"`
}

type AddUserToTeamResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Adding user to team in database
func AddUserToTeam(user_id string, team_id string) error {
	return queries.AddUserToTeam(user_id, team_id)
}
