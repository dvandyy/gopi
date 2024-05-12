package handlers

import (
	"github.com/dvandyy/gopi/api/v1/models"
	"github.com/dvandyy/gopi/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// @Summary      Get teams for user
// @Description  Return all teams for certain user in workspace
// @Tags         Teams
// @Security 	 JWT_TOKEN
// @Produce      json
// @Param		 wid path string true "Workspace unique ID"
// @Success      200 {object}  models.UserTeamsResponse
// @Router       /teams/user/{wid} [Get]
func HandleGetUserTeamsInWorkspace(c *fiber.Ctx) error {
	// Get user id from JWT token
	user_id := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["id"].(string)

	// Get id param from url
	workspace_id := c.Params("wid")

	// Get all teams from db for certain user in workspace
	teams, err := models.GetTeamsForUserInWorkspace(user_id, workspace_id)
	if err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusInternalServerError,
			Message: "Internal server error",
		})
	}

	// Return response with teams
	return c.JSON(models.UserTeamsResponse{
		Status: fiber.StatusOK,
		Teams:  teams,
	})
}

// @Summary      Create new team
// @Description  Create a new team in database.
// @Tags         Teams
// @Security 	 JWT_TOKEN
// @Param		 CreateTeamRequest body models.CreateTeamRequest true "Create new team with workspace_id and name"
// @Produce      json
// @Success      200 {object}  models.CreateTeamResponse
// @Router       /teams/new [Post]
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

// @Summary      Add user to team
// @Description  Add a new user to team.
// @Tags         Teams
// @Security 	 JWT_TOKEN
// @Param		 AddUserToTeamRequest body models.AddUserToTeamRequest true "Add new user to team with user_id and workspace_id"
// @Produce      json
// @Success      200 {object}  models.AddUserToTeamResponse
// @Router       /teams/user/add [Post]
func HandleAddUserToTeam(c *fiber.Ctx) error {
	req := new(models.AddUserToTeamRequest)

	// Check if input is valid
	if err := c.BodyParser(req); err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
		})
	}

	// Add user to team in the database
	db_err := models.AddUserToTeam(req.User_id, req.Team_id)
	if db_err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusInternalServerError,
			Message: "Error while adding user to team",
		})
	}

	// Return a success message
	return c.JSON(models.AddUserToTeamResponse{
		Status:  fiber.StatusAccepted,
		Message: "User successfully added.",
	})
}
