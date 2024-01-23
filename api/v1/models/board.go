package models

import (
	"github.com/bodatomas/gopi/database/queries"
)

type Board struct {
	ID          int    `json:"id" example:"1"`
	Unique_id   string `json:"unique_id" example:"926e7309-12e4-4c50-824c-33737fb45f8a"`
	Title       string `json:"title" example:"Title"`
	Description string `json:"description" example:"Board description"`
	Created_at  string `json:"created_at" example:"2024-01-22 17:03:50.283466+00"`
}

type CreateBoardRequest struct {
	Title       string `json:"title" example:"My title"`
	Description string `json:"description" example:"My description"`
}

type CreateBoardResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Get board with uuid from database and return as Board struct
func GetBoardByUID(id string) (Board, error) {
	var board Board

	// Get board from databse
	row := queries.GetBoardByID(id)
	err := row.Scan(&board.ID, &board.Unique_id, &board.Title, &board.Description, &board.Created_at)
	if err != nil {
		return Board{}, err
	}

	return board, nil
}

// Create new bord in database
func CreateNewBoard(title string, description string) error {
	return queries.NewBoard(title, description)
}
