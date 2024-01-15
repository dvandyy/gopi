package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/bodatomas/gopi/api/v1/models"
	"github.com/gofiber/fiber/v3"
)

type Board struct {
	logger *log.Logger
}

// NewBoard Constructor
func NewBoard(logger *log.Logger) *Board {
	return &Board{logger}
}

// GetBoardByID Return board with given id
func (board *Board) GetBoardByID(context fiber.Ctx) error {
	// Get id param from url
	id := context.Params("id")
	// Change id from string to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	// Get board from database
	data, err := models.GetBoardByID(intID)
	if err != nil {
		return context.JSON(fiber.Map{"status": http.StatusNotFound})
	}
	return context.JSON(data)
}
