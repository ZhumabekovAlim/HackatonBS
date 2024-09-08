package models

type Review struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	BookID    int     `json:"book_id"`
	Content   string  `json:"content"`
	Rating    float32 `json:"rating"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
