package app

import (
	"log"

	"github.com/bodatomas/gopi/api/v1/handlers"
	"github.com/bodatomas/gopi/config"
	"github.com/gofiber/fiber/v3"
)

func SetupGetRequests(router *fiber.App, logger *log.Logger) {

	basePath := "/api/" + config.Cfg.Version

	// GET routes
	router.Get(basePath+"/", handlers.NewHello(logger).GetHello)
	router.Get(basePath+"/board/:id", handlers.NewBoard(logger).GetBoardByID)

}
