package handlers

import (
	"github.com/dvandyy/gopi/api/v1/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// @Summary      Show a hello message
// @Description  Retun a hello message if everything is ok
// @Tags         Welcome
// @Produce      json
// @Security 	 JWT_TOKEN
// @Success      200  {object}  models.HelloResponse  "Return 'Hello from gopi!'"
// @Router       / [get]
func HandleGetHello(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	return c.JSON(models.HelloResponse{
		Status:   fiber.StatusOK,
		Messsage: "Hello from gopi!",
		UserID:   id,
	})
}
