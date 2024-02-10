package swagger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "github.com/dvandyy/gopi/api/v1/docs"
)

// Make custom config here in the future
func HandleConfiguredSwagger(c *fiber.Ctx) error {
	return swagger.HandlerDefault(c)
}
