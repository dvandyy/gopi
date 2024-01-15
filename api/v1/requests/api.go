package requests

import (
	"log"

	"github.com/bodatomas/gopi/api/v1/handlers"
	"github.com/gofiber/fiber/v3"
)

var endPoint = "/api/v1/"

func SetupApiGetRequests(router *fiber.App, logger *log.Logger) {
	api := router.Group(endPoint)

	// GET routes for "/api" endpoint
	api.Get("/", handlers.NewHello(logger).GetHello)
	api.Get("/board/:id", handlers.NewBoard(logger).GetBoardByID)
}

func SetupApiPostRequests(router *fiber.App, logger *log.Logger) {

}

func SetupApiPutRequests(router *fiber.App, logger *log.Logger) {

}

func SetupApiDeleteRequests(router *fiber.App, logger *log.Logger) {

}
