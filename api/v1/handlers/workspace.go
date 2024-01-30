package handlers

import (
	"github.com/bodatomas/gopi/api/v1/models"
	"github.com/bodatomas/gopi/utils"
	"github.com/gofiber/fiber/v2"
)

// @Summary      Create new workspace
// @Description  Create a new workspace in database.
// @Tags         Workspaces
// @Security 	 JWT_TOKEN
// @Param		 CreateWorkspaceRequest body models.CreateWorkspaceRequest true "Create new workspace with owner_id and name"
// @Produce      json
// @Success      200 {object}  models.CreateWorkspaceResponse
// @Router       /workspace/new [Post]
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
