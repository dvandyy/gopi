package handlers

import (
	"github.com/dvandyy/gopi/api/v1/models"
	"github.com/dvandyy/gopi/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// @Summary      Get workspaces for user
// @Description  Return all workspaces for certain user
// @Tags         Workspaces
// @Security 	 JWT_TOKEN
// @Produce      json
// @Success      200 {object}  models.UserWorkspacesResponse
// @Router       /workspaces/user [Get]
func HandleGetUserWorkspaces(c *fiber.Ctx) error {
	// Get user id from JWT token
	user_id := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["id"].(string)

	// Get all workspaces from db for certain user
	workspaces, err := models.GetUserWorkspaces(user_id)
	if err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusInternalServerError,
			Message: "Internal server error",
		})
	}

	// Return response with workspaces
	return c.JSON(models.UserWorkspacesResponse{
		Status:     fiber.StatusOK,
		Workspaces: workspaces,
	})
}

// @Summary      Create new workspace
// @Description  Create a new workspace in database.
// @Tags         Workspaces
// @Security 	 JWT_TOKEN
// @Param		 CreateWorkspaceRequest body models.CreateWorkspaceRequest true "Create new workspace with owner_id and name"
// @Produce      json
// @Success      200 {object}  models.CreateWorkspaceResponse
// @Router       /workspaces/new [Post]
func HandleCreateWorkspace(c *fiber.Ctx) error {
	workspace := new(models.CreateWorkspaceRequest)

	// Check if input is valid
	if err := c.BodyParser(workspace); err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
		})
	}

	// Generate unique id with prefix
	uuid := utils.GenerateUniqueID("w-")

	// Store workspace in the database
	db_err := models.CreateNewWorkspace(uuid, workspace.Owner, workspace.Name)
	if db_err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusInternalServerError,
			Message: "Error while creating workspace",
		})
	}

	// Return a success message
	response := models.CreateWorkspaceResponse{
		Status:  fiber.StatusAccepted,
		Message: "Workspace successfully created.",
	}
	return c.JSON(response)
}

// @Summary      Add user to workspace
// @Description  Add a new user to workspace.
// @Tags         Workspaces
// @Security 	 JWT_TOKEN
// @Param		 AddUserToWorkspaceRequest body models.AddUserToWorkspaceRequest true "Add new user to workspace with user_id and workspace_id"
// @Produce      json
// @Success      200 {object}  models.AddUserToWorkspaceResponse
// @Router       /workspaces/user/add [Post]
func HandleAddUserToWorkspace(c *fiber.Ctx) error {
	req := new(models.AddUserToWorkspaceRequest)

	// Check if input is valid
	if err := c.BodyParser(req); err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
		})
	}

	// Store workspace in the database
	db_err := models.AddUserToWorkspace(req.User_id, req.Workspace_id)
	if db_err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusInternalServerError,
			Message: "Error while adding user to workspace",
		})
	}

	// Return a success message
	return c.JSON(models.AddUserToWorkspaceResponse{
		Status:  fiber.StatusAccepted,
		Message: "User successfully added.",
	})
}
