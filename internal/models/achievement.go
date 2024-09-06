package models

type Achievement struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Points    int    `json:"points"`
	Level     int    `json:"level"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
