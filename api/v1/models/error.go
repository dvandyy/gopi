package models

type Error struct {
	Status  int    `json:"error"`
	Message string `json:"message"`
}
