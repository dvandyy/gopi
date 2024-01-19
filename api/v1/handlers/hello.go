package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func HandleGetHello(context *fiber.Ctx) error {
	data := fiber.Map{
		"status":  http.StatusOK,
		"message": "Hello world!",
	}
	return context.JSON(data)
}
