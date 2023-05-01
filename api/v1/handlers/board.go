package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/bodatomas/gopi/api/v1/models"
	"github.com/labstack/echo/v4"
)

type Board struct {
	logger *log.Logger
}

func NewBoard(logger *log.Logger) *Board {
	return &Board{logger}
}

func (_ *Board) GetBoard(context echo.Context) error {
	data := board
	return context.JSONPretty(http.StatusOK, data, " ")
}

// Dummy data
var board = &models.Board{
	ID:          1,
	Name:        "Test Board",
	Description: "Test",
	CreatedOn:   time.Now().UTC().String(),
}
