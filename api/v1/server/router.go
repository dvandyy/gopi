package server

import (
	"log"

	"github.com/bodatomas/gopi/api/v1/handlers"
)

func (s *Server) SetupGetRequests(logger *log.Logger) {
	api := s.fiber.Group("/api/v1/")

	api.Get("/", handlers.HandleGetHello)
	api.Get("/board/:id", handlers.HandleGetBoardByID)
}

func (s *Server) SetupPostRequests(logger *log.Logger) {
	user := s.fiber.Group("/user/")
	user.Post("/register", handlers.HandleRegisterUser)
}

func (s *Server) SetupPutRequests(logger *log.Logger) {

}

func (s *Server) SetupDeleteRequests(logger *log.Logger) {

}
