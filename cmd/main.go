package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bodatomas/gopi/api/v1/app"
	"github.com/bodatomas/gopi/config"
	"github.com/bodatomas/gopi/database"
)

func main() {
	// Logger for one place logging.
	logger := log.New(os.Stdout, "gopi-api", log.LstdFlags)

	// Creating config.
	config.New(logger)

	// Creating main application.
	fiber := app.InitFiberApp(logger)

	// Server setup - config properties are in env file.
	fiber.SetupServer(logger)

	// Creating database.
	database.CreateDatabase()

	// PORT
	port := fmt.Sprintf(":%s", config.Get().Address)

	// Concurrent server run and listen
	go func() {
		logger.Fatal(fiber.App.Listen(port))
	}()
	logger.Println("Server is running!")

	// Graceful shutdown
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	sig := <-sigChan
	logger.Println("Received terminate, graceful shutdown", sig)
	_ = fiber.App.Shutdown()

	// Cleanup
	logger.Println("Running cleanup tasks...")
	// Close database
	database.GetDatabase().Conn.Close()

	logger.Println("Server was successful shutdown")
}
