package handlers

import (
	"github.com/bodatomas/gopi/api/v1/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary      Get board with UUID
// @Description  Return board with unique id
// @Tags         Boards
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
// @Param		 CreateBoardRequest body models.CreateBoardRequest true "Create board with Title and Description"
// @Produce      json
// @Success      200 {object}  models.CreateBoardResponse
// @Router       /boards/new [Post]
func HandleCreateBoard(c *fiber.Ctx) error {
	var board models.Board

	// Check if input is valid
	if err := c.BodyParser(board); err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
		})
	}

	// Store user data in the database
	db_err := models.CreateNewBoard(board.Title, board.Description)
	if db_err != nil {
		return c.JSON(models.Error{
			Status:  fiber.StatusInternalServerError,
			Message: "Error while creating board",
		})
	}

	// Return a success message
	response := models.CreateBoardResponse{
		Status:  fiber.StatusAccepted,
		Message: "Board successfully created.",
	}
	return c.JSON(response)
}
