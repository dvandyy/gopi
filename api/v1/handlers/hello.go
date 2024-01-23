package handlers

import (
	"github.com/bodatomas/gopi/api/v1/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary      Show a hello message
// @Description  Retun a hello message if everything is ok
// @Tags         Welcome
// @Produce      json
// @Success      200  {object}  models.HelloResponse  "Return 'Hello from gopi!'"
// @Router       / [get]
func HandleGetHello(context *fiber.Ctx) error {
	respone := models.HelloResponse{
		Status:   fiber.StatusAccepted,
		Messsage: "Hello from gopi!",
	}
	return context.JSON(respone)
}
