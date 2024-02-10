package handlers

import (
	"github.com/dvandyy/gopi/api/v1/models"
	"github.com/dvandyy/gopi/utils"
	"github.com/gofiber/fiber/v2"
)

// @Summary      Create new team
// @Description  Create a new team in database.
// @Tags         Teams
// @Security 	 JWT_TOKEN
// @Param		 CreateTeamRequest body models.CreateTeamRequest true "Create new team with workspace_id and name"
// @Produce      json
// @Success      200 {object}  models.CreateTeamResponse
// @Router       /team/new [Post]
func HandleCreateTeam(c *fiber.Ctx) error {
	team := new(models.CreateTeamRequest)

	// Check if input is valid
	if err := c.BodyParser(team); err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
		})
	}

	// Generate unique id with prefix
	uuid := utils.GenerateUniqueID("t-")

	// Store team in the database
	db_err := models.CreateNewTeam(uuid, team.Workspace_id, team.Name)
	if db_err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusInternalServerError,
			Message: "Error while creating team",
		})
	}

	// Return a success message
	response := models.CreateTeamResponse{
		Status:  fiber.StatusAccepted,
		Message: "Team successfully created.",
	}
	return c.JSON(response)
}
