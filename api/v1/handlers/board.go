package handlers

import (
	"log"
	"net/http"

	"github.com/bodatomas/gopi/api/v1/models"
	"github.com/labstack/echo/v4"
)

type Board struct {
	l *log.Logger
}

func NewBoard(l *log.Logger) *Board {
	return &Board{l}
}

func (b *Board) GetBoard(context echo.Context) error {
	data := models.GetBoard()
  return context.JSON(http.StatusOK, data)
}
