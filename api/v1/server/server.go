package server

import (
	"log"
	"time"

	"github.com/bodatomas/gopi/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	fiber  *fiber.App
	logger *log.Logger
}

// Creating fiber app
func NewServer(l *log.Logger) *Server {
	server := Server{
		fiber:  fiber.New(),
		logger: l,
	}

	// Setup server
	server.setupServer(l)

	// Middlewares
	server.fiber.Use(logger.New())
	server.fiber.Use(recover.New())

	// Requests
	server.SetupGetRequests(l)
	server.SetupPostRequests(l)
	server.SetupPutRequests(l)
	server.SetupDeleteRequests(l)

	return &server
}

func (s *Server) Start(listenAddr string) error {
	return s.fiber.Listen(listenAddr)
}

func (s *Server) Stop() error {
	return s.fiber.Shutdown()
}

// Server configuration
func (s *Server) setupServer(logger *log.Logger) {
	server := s.fiber.Server()
	server.IdleTimeout = config.Get().IdleTimeout * time.Second
	server.WriteTimeout = config.Get().WriteTimeout * time.Second
	server.ReadTimeout = config.Get().ReadTimeout * time.Second
	logger.Println("Server was succesfully configured.")
}
