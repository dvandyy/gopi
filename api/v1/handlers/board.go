package handlers

import (
	"github.com/bodatomas/gopi/api/v1/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type Board struct {
	logger *log.Logger
}

// NewBoard Constructor
func NewBoard(logger *log.Logger) *Board {
	return &Board{logger}
}

// GetBoardByID Return board with given id
func (_ *Board) GetBoardByID(context echo.Context) error {
	// Get id param from url
	id := context.Param("id")
	// Change id from string to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	// Get board from database
	data, err := models.GetBoardByID(intID)
	if err != nil {
		return context.NoContent(http.StatusNotFound)
	}
	return context.JSONPretty(http.StatusOK, data, " ")
}
