package router

import (
	"log"

	"github.com/bodatomas/gopi/api/v1/handlers"
	"github.com/bodatomas/gopi/config"
	"github.com/labstack/echo/v4"
)

func InitRouter(logger *log.Logger) *echo.Echo {

	router := echo.New()

	setupGetRequests(router, logger)

	return router
}

func setupGetRequests(router *echo.Echo, logger *log.Logger) {

	basePath := "/api/" + config.Cfg.Version

	// GET routes
	router.GET(basePath+"/", handlers.NewHello(logger).GetHello)
	router.GET(basePath+"/board", handlers.NewBoard(logger).GetBoard)

}
