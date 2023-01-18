package router

import (
	"log"
	"net/http"

	"github.com/bodatomas/gopi/api/v1/handlers"
	"github.com/gorilla/mux"
)

func InitRouter(l *log.Logger) *mux.Router {

	router := mux.NewRouter()

  // Get subrouter
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", handlers.NewHello(l).GetHello)
	getRouter.HandleFunc("/board", handlers.NewBoard(l).GetBoard)

  // Post subrouter
	// postRouter := router.Methods(http.MethodPost).Subrouter()

  // Put subrouter
	// putRouter := router.Methods(http.MethodPut).Subrouter()

  // Delete subrouter
	// deleteRouter := router.Methods(http.MethodDelete).Subrouter()

  return router
}

