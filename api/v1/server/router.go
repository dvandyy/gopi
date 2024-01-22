package server

import (
	"log"

	"github.com/bodatomas/gopi/api/v1/handlers"
	"github.com/bodatomas/gopi/api/v1/swagger"
)

func (s *Server) SetupGetRequests(logger *log.Logger) {
	// Documentation
	docs := s.fiber.Group("/api/v1/docs/")
	docs.Get("/*", swagger.HandleConfiguredSwagger)

	// Api
	api := s.fiber.Group("/api/v1/")
	api.Get("/", handlers.HandleGetHello)
	api.Get("/board/:id", handlers.HandleGetBoardByID)
}

func (s *Server) SetupPostRequests(logger *log.Logger) {
	// User
	user := s.fiber.Group("/user/")
	user.Post("/register", handlers.HandleRegisterUser)

	// Board
	board := s.fiber.Group("/board")
	board.Post("/new", handlers.HandleCreateBoard)
}

func (s *Server) SetupPutRequests(logger *log.Logger) {

}

func (s *Server) SetupDeleteRequests(logger *log.Logger) {

}
