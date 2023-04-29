package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) GetHello(context echo.Context) error {
	data := map[string]string{
		"message": "Hello world!",
	}
  return context.JSON(http.StatusOK, data)
}
