package models

import queries "github.com/bodatomas/gopi/database/queries/board"

type Board struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetBoardByID(id int) (Board, error) {
	var board Board
	row := queries.GetBoardByID(id)
	err := row.Scan(&board.ID, &board.Name, &board.Description)
	if err != nil {
		return Board{}, err
	}
	return board, nil
}
