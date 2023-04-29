package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bodatomas/gopi/api/v1/router"
	"github.com/bodatomas/gopi/env"
)

func main() {
	// Logger for one place logging
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	// Main router Echo
	router := router.InitRouter(logger)

	// Server configuration
	server := &http.Server{
		Handler:      router,
		Addr:         env.GetEnvValue("ADDRESS"),
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
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
