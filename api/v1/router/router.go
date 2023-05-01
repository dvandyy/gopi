package router

import (
	"log"

	"github.com/bodatomas/gopi/api/v1/handlers"
	"github.com/labstack/echo/v4"
)

func InitRouter(logger *log.Logger) *echo.Echo {

	router := echo.New()

  setupGetRequests(router, logger);

	return router
}

func setupGetRequests(router *echo.Echo, logger *log.Logger) {

  // GET routes
  router.GET("/", handlers.NewHello(logger).GetHello)
  router.GET("/board", handlers.NewBoard(logger).GetBoard)

}
