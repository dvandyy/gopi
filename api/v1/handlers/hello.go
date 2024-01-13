package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

func (_ *Hello) GetHello(context echo.Context) error {
	data := map[string]string{
		"message": "Hello world!",
	}
	return context.JSONPretty(http.StatusOK, data, " ")
}
