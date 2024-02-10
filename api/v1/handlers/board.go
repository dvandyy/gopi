package handlers

import (
	"github.com/dvandyy/gopi/api/v1/models"
	"github.com/dvandyy/gopi/utils"
	"github.com/gofiber/fiber/v2"
)

// @Summary      Get board with UUID
// @Description  Return board with unique id
// @Tags         Boards
// @Security 	 JWT_TOKEN
// @Produce      json
// @Param 		 uid path string true "Board unique ID"
// @Success      200  {object}  models.Board
// @Router       /boards/{uid} [get]
func HandleGetBoardByID(c *fiber.Ctx) error {
	// Get id param from url
	id := c.Params("uid")
	// Get board from database
	board, err := models.GetBoardByUID(id)
	if err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusNotFound,
			Message: "Board was not found",
		})
	}
	return c.JSON(board)
}

// @Summary      Create new board
// @Description  Create a new board in database.
// @Tags         Boards
// @Security     JWT_TOKEN
// @Param		 CreateBoardRequest body models.CreateBoardRequest true "Create board with Title"
// @Produce      json
// @Success      200 {object}  models.CreateBoardResponse
// @Router       /boards/new [Post]
func HandleCreateBoard(c *fiber.Ctx) error {
	board := new(models.CreateBoardRequest)

	// Check if input is valid
	if err := c.BodyParser(board); err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
		})
	}

	// Generate unique id with prefix
	uuid := utils.GenerateUniqueID("b-")

	// Store user data in the database
	db_err := models.CreateNewBoard(uuid, board.Team_id, board.Title)
	if db_err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusInternalServerError,
			Message: "Internal server error",
		})
	}

	// Return a success message
	return c.JSON(models.CreateBoardResponse{
		Status:  fiber.StatusAccepted,
		Message: "Board successfully created.",
	})
}
