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

	// Hello
	api.Get("/", handlers.HandleGetHello)

	// User
	api.Get("/users/:uid", handlers.HandleGetUserByID)

	// Board
	api.Get("/boards/:uid", handlers.HandleGetBoardByID)
}

func (s *Server) SetupPostRequests(logger *log.Logger) {
	api := s.fiber.Group("/api/v1/")

	// User
	api.Post("/users/register", handlers.HandleRegisterUser)

	// Workspace
	api.Post("/workspace/new", handlers.HandleCreateWorkspace)

	// Team
	api.Post("/team/new", handlers.HandleCreateTeam)

	// Board
	api.Post("/boards/new", handlers.HandleCreateBoard)
}

func (s *Server) SetupPutRequests(logger *log.Logger) {

}

func (s *Server) SetupDeleteRequests(logger *log.Logger) {

}
