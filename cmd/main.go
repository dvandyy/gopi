package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bodatomas/gopi/api/v1/router"
	"github.com/bodatomas/gopi/config"
)

func main() {
	// Logger for one place logging
	logger := log.New(os.Stdout, "gopi-api", log.LstdFlags)
	// Config
	cfg := config.New(logger)
	// Main router Echo
	router := router.InitRouter(logger)

	// Server configuration
	server := &http.Server{
		Handler:      router,
		Addr:         cfg.Address,
		IdleTimeout:  cfg.IdleTimeout * time.Second,
		WriteTimeout: cfg.WriteTimeout * time.Second,
		ReadTimeout:  cfg.ReadTimeout * time.Second,
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
	logger.Println("Recieved terminate, graceful shutdown", sig)

	timeContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(timeContext)
}
