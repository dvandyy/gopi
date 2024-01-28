package models

import "github.com/bodatomas/gopi/database/queries"

type Workspace struct {
	ID         uint   `json:"id" example:"1"`
	Unique_id  string `json:"unique_id" example:"u-1706453613063fa2eb4"`
	Name       string `json:"name" example:"My workspace"`
	Owner      string `json:"owner" example:"u-1706453613063fa2eb4"`
	Created_at string `json:"created_at" example:"2024-01-22 17:03:50.283466+00"`
}

type CreateWorkspaceRequest struct {
	Name  string `json:"name" example:"My workspace"`
	Owner string `json:"owner" example:"u-1706453613063fa2eb4"`
}

type CreateWorkspaceResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Create new workspace in database
func CreateNewWorkspace(unique_id string, owner string, name string) error {
	return queries.NewWorkspace(unique_id, owner, name)
}
