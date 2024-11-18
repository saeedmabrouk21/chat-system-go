package models

type Message struct {
	ID     int    `json:"id"`
	ChatID int    `json:"chat_id"`
	Body   string `json:"body"`
	Number int    `json:"number"`
}
