package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"fmt"

	"github.com/bodatomas/gopi/api/v1/router"
	"github.com/bodatomas/gopi/config"
	"github.com/bodatomas/gopi/database"
)

func main() {
	// Logger for one place logging
	logger := log.New(os.Stdout, "gopi-api", log.LstdFlags)
	// Config
	config.New(logger)
	// Main router Echo
	mainRouter := router.InitRouter(logger)
	// Database
	database.CreateDatabase()

	// Server configuration
	server := &http.Server{
		Handler:      mainRouter,
		Addr:         fmt.Sprintf(":%s", config.Cfg.Address ),
		IdleTimeout:  config.Cfg.IdleTimeout * time.Second,
		WriteTimeout: config.Cfg.WriteTimeout * time.Second,
		ReadTimeout:  config.Cfg.ReadTimeout * time.Second,
	}

	// Concurrent server run and listen
	go func() {
		logger.Fatal(server.ListenAndServe())
	}()
	logger.Println("Server is running!")

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Received terminate, graceful shutdown", sig)

	timeContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := server.Shutdown(timeContext)
	if err != nil {
		return
	}
}
