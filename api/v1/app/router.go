package app

import (
	"log"

	"github.com/bodatomas/gopi/api/v1/requests"
	"github.com/gofiber/fiber/v3"
)

func SetupGetRequests(router *fiber.App, logger *log.Logger) {
	requests.SetupApiGetRequests(router, logger)
}

func SetupPostRequests(router *fiber.App, logger *log.Logger) {

}

func SetupPutRequests(router *fiber.App, logger *log.Logger) {

}

func SetupDeleteRequests(router *fiber.App, logger *log.Logger) {

}
