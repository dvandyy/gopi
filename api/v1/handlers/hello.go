package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// @Summary      Show a hello message
// @Description  Retun a hello message if everything is ok
// @Tags         Welcome
// @Produce      json
// @Success      200  {string}  string    "Hello world!"
// @Router       / [get]
func HandleGetHello(context *fiber.Ctx) error {
	data := fiber.Map{
		"status":  http.StatusOK,
		"message": "Hello world!",
	}
	return context.JSON(data)
}
