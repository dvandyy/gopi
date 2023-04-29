package models

import (
	"time"
)

type Board struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedOn   string `json:"createdOn"`
}

func GetBoard() *Board {
	return board
}

// Dummy data
var board = &Board{
	ID:          1,
	Name:        "Test Board",
	Description: "Test",
	CreatedOn:   time.Now().UTC().String(),
}
