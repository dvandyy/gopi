package middlewares

import (
	"github.com/dvandyy/gopi/api/v1/models"
	"github.com/dvandyy/gopi/config"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// JWT authentication middleware
func AuthMiddleware() func(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.JSON(models.Error{
				Status:  fiber.StatusUnauthorized,
				Message: "Unauthorized",
			})
		},
		SigningKey: jwtware.SigningKey{Key: []byte(config.Get().JWT_Secret)},
	})
}
