package model

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	UserID string  `json:"user"`
}