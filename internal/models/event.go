package models

type Event struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Type      string `json:"type"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
