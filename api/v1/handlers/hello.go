package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) GetHello(rw http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message": "Hello world!",
	}
	err := json.NewEncoder(rw).Encode(data)
	if err != nil {
		panic(err)
	}
}
