package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

func (hello *Hello) GetHello(context fiber.Ctx) error {
	data := fiber.Map{
		"status":  http.StatusOK,
		"message": "Hello world!",
	}
	return context.JSONP(data)
}
