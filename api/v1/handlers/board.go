package handlers

import (
	"log"
	"net/http"

	"github.com/bodatomas/gopi/api/v1/models"
)

type Board struct {
	l *log.Logger
}

func NewBoard(l *log.Logger) *Board {
	return &Board{l}
}

func (b *Board) GetBoard(rw http.ResponseWriter, r *http.Request) {
	data := models.GetBoard()
	err := data.ToJSON(rw)
	if err != nil {
		panic(err)
	}
}
