package models

type Favorite struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	BookID    int    `json:"book_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
