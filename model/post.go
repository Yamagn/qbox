package model

type Post struct {
	ID uint `json:"id"`
	Content string `json:"content"`
	To string `json:"to"`
}