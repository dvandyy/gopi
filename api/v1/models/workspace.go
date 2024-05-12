package models

import (
	"github.com/dvandyy/gopi/database/queries"
)

type Workspace struct {
	ID         uint   `json:"id" example:"1"`
	Unique_id  string `json:"unique_id" example:"w-1706453613063fa2eb4"`
	Name       string `json:"name" example:"My workspace"`
	Owner      string `json:"owner" example:"u-1706453613063fa2eb4"`
	Created_at string `json:"created_at" example:"2024-01-22 17:03:50.283466+00"`
}

/*
**
Create workspace
**
*/
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

/*
**
Get workspace
**
*/
type WorkspaceResponse struct {
	Unique_id string `json:"unique_id" example:"w-1706453613063fa2eb4"`
	Name      string `json:"name" example:"My workspace"`
}

type UserWorkspacesResponse struct {
	Status     int                 `json:"status"`
	Workspaces []WorkspaceResponse `json:"workspaces"`
}

func GetUserWorkspaces(user_id string) ([]WorkspaceResponse, error) {
	var workspaces []WorkspaceResponse

	rows := queries.GetWorkspacesForUser(user_id)
	for rows.Next() {
		var workspace WorkspaceResponse
		if err := rows.Scan(&workspace.Unique_id, &workspace.Name); err != nil {
			return nil, err
		}
		workspaces = append(workspaces, workspace)
	}

	return workspaces, nil
}

/*
**
Add User to workspace
**
*/
type AddUserToWorkspaceRequest struct {
	User_id      string `json:"user_id" example:"u-1706453613063fa2eb4"`
	Workspace_id string `json:"workspace_id" example:"w-1706453613063fa2eb4"`
}

type AddUserToWorkspaceResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Adding user to workspace in databse
func AddUserToWorkspace(user_id string, workspace_id string) error {
	return queries.AddUserToWorkspace(user_id, workspace_id)
}
