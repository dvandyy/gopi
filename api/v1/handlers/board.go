package handlers

import (
	"net/http"

	"github.com/bodatomas/gopi/api/v1/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary      Return board with unique id
// @Description  Return board with unique id
// @Tags         Boards
// @Produce      json
// @Success      200  {object}  models.Board
// @Router       /board/:uid [get]
func HandleGetBoardByID(c *fiber.Ctx) error {
	// Get id param from url
	id := c.Params("id")
	// Get board from database
	data, err := models.GetBoardByUID(id)
	if err != nil {
		return c.JSON(fiber.Map{"status": http.StatusNotFound})
	}
	return c.JSON(data)
}

// @Summary      Create a new board
// @Description  Create a new board
// @Tags         Boards
// @Produce      json
// @Success      200  {object}  models.Board
// @Router       /board/new [Post]
func HandleCreateBoard(c *fiber.Ctx) error {
	board := new(models.Board)

	// Check if input is valid
	if err := c.BodyParser(board); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error while creating board",
		})
	}

	// Store user data in the database
	db_err := models.CreateNewBoard(board.Title)
	if db_err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error while creating board",
		})
	}

	// Return a success message
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"success": "Board successfully created",
	})
}
