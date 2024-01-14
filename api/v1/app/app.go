package app

import (
	"log"
	"time"

	"github.com/bodatomas/gopi/config"
	"github.com/gofiber/fiber/v3"
)

type FiberApp struct {
	App *fiber.App
}

// Creating fiber app
func InitFiberApp(logger *log.Logger) *FiberApp {
	app := FiberApp{App: fiber.New()}
	//Setup get request (may be temporary)
	SetupGetRequests(app.App, logger)
	return &app
}

// Server configuration
func (fiber *FiberApp) SetupServer(logger *log.Logger) {
	server := fiber.App.Server()
	server.IdleTimeout = config.Cfg.IdleTimeout * time.Second
	server.WriteTimeout = config.Cfg.WriteTimeout * time.Second
	server.ReadTimeout = config.Cfg.ReadTimeout * time.Second
}
