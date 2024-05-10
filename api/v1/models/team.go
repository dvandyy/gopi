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
