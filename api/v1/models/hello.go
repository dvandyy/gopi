package models

type HelloResponse struct {
	Status   int    `json:"status"`
	Messsage string `json:"message"`
	UserID   string `json:"user_id"`
}
