package handlers

import (
	"net/http"

	"github.com/bodatomas/gopi/api/v1/models"
	"github.com/gofiber/fiber/v2"
)

// GetBoardByID Return board with given id
// @Summary      Return board with given id
// @Description  Return board with given id
// @Tags         Boards
// @Produce      json
// @Success      200  {object}  models.Board
// @Router       /board/:id [get]
func HandleGetBoardByID(c *fiber.Ctx) error {
	// Get id param from url
	id := c.Params("id")
	// Get board from database
	data, err := models.GetBoardByID(id)
	if err != nil {
		return c.JSON(fiber.Map{"status": http.StatusNotFound})
	}
	return c.JSON(data)
}
