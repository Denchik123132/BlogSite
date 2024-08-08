package models

type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"content"`
	Image       string `json:"image"`
}
