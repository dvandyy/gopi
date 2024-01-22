package models

import (
	"github.com/bodatomas/gopi/database/queries"
)

type Board struct {
	ID         int    `json:"id"`
	Unique_id  string `json:"unique_id"`
	Title      string `json:"title"`
	Created_at string `json:"created_at"`
}

// Get board with uuid from database and return as Board struct
func GetBoardByUID(id string) (Board, error) {
	var board Board

	// Get board from databse
	row := queries.GetBoardByID(id)
	err := row.Scan(&board.ID, &board.Unique_id, &board.Title, &board.Created_at)
	if err != nil {
		return Board{}, err
	}

	return board, nil
}

// Create new bord in database
func CreateNewBoard(title string) error {
	return queries.NewBoard(title)
}
