package server

import (
	"log"

	"github.com/dvandyy/gopi/api/v1/handlers"
	"github.com/dvandyy/gopi/api/v1/middlewares"
	"github.com/dvandyy/gopi/api/v1/swagger"
)

func (s *Server) SetupGetRequests(logger *log.Logger) {
	// Documentation
	docs := s.fiber.Group("/api/v1/docs/")
	docs.Get("/*", swagger.HandleConfiguredSwagger)

	// Api
	api := s.fiber.Group("/api/v1/")

	// Hello
	api.Get("/", middlewares.AuthMiddleware(), handlers.HandleGetHello)

	// User
	api.Get("/users/:uid", middlewares.AuthMiddleware(), handlers.HandleGetUserByID)

	// Workspace
	api.Get("/workspaces/user", middlewares.AuthMiddleware(), handlers.HandleGetUserWorkspaces)

	// Team
	api.Get("/teams/user/:wid", middlewares.AuthMiddleware(), handlers.HandleGetUserTeamsInWorkspace)

	// Board
	api.Get("/boards/:uid", middlewares.AuthMiddleware(), handlers.HandleGetBoardByID)

}

func (s *Server) SetupPostRequests(logger *log.Logger) {
	api := s.fiber.Group("/api/v1/")

	// Register / Login - unauthorized routes
	api.Post("/users/register", handlers.HandleRegisterUser)
	api.Post("/users/login", handlers.HandleLoginUser)

	// Workspace
	api.Post("/workspaces/new", middlewares.AuthMiddleware(), handlers.HandleCreateWorkspace)
	api.Post("/workspaces/user/add", middlewares.AuthMiddleware(), handlers.HandleAddUserToWorkspace)

	// Team
	api.Post("/teams/new", middlewares.AuthMiddleware(), handlers.HandleCreateTeam)
	api.Post("/teams/user/add", middlewares.AuthMiddleware(), handlers.HandleAddUserToTeam)

	// Board
	api.Post("/boards/new", middlewares.AuthMiddleware(), handlers.HandleCreateBoard)
}

func (s *Server) SetupPutRequests(logger *log.Logger) {

}

func (s *Server) SetupDeleteRequests(logger *log.Logger) {

}
