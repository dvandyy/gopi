package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dvandyy/gopi/api/v1/server"
	"github.com/dvandyy/gopi/config"
	"github.com/dvandyy/gopi/database"
)

// @title Gopi API
// @version 1.0
// @description REST api
// @securitydefinitions.apikey JWT_TOKEN
// @in header
// @name Authorization

// @contact.name Tomáš Boďa
// @contact.url https://github.com/dvandyy
// @contact.email tomasboda.dev@gmail.com

// @host localhost:4000
// @BasePath /api/v1/
func main() {
	// Logger for one place logging
	logger := log.New(os.Stdout, "gopi-api", log.LstdFlags)

	// Create config
	config.New(logger)

	// Create database
	database.CreateDatabase()

	// Create server
	server := server.NewServer(logger)

	// PORT
	port := fmt.Sprintf(":%s", config.Get().Address)

	// Concurrent server run and listen
	go func() {
		logger.Fatal(server.Start(port))
	}()
	logger.Println("Server is running :)")

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	sig := <-sigChan
	logger.Println("Received terminate, graceful shutdown", sig)
	server.Stop()

	// Cleanup
	logger.Println("Running cleanup tasks...")
	database.GetDatabase().Conn.Close()

	logger.Println("Server was successfully shutdown")
}
