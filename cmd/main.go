package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bodatomas/gopi/api/v1/handlers"
	"github.com/gorilla/mux"
)

func main() {
  // Main router Mux
	router := mux.NewRouter()
  // Logger for one place logging
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

  // Server configuration
  // TODO: implement env variables
	server := &http.Server{
		Handler:      router,
		Addr:         ":4000",
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

  // Get subrouter
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", handlers.NewHello(l).GetHello)

  // Post subrouter
	// postRouter := router.Methods(http.MethodPost).Subrouter()

  // Put subrouter
	// putRouter := router.Methods(http.MethodPut).Subrouter()

  // Delete subrouter
	// deleteRouter := router.Methods(http.MethodDelete).Subrouter()

  // Concurent server run and listen
	go func() {
		l.Fatal(server.ListenAndServe())
	}()
  l.Println("Server is running!")

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	timeContext, cancel := context.WithTimeout(context.Background(), 30* time.Second)
	defer cancel()
	server.Shutdown(timeContext)
}
